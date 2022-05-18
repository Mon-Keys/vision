package timecode_generator

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"time"
)

func currentCipher(time_period int64) (cipher.Block, error) {
	now := time.Now()
	sec := now.Unix()
	floor := (sec / time_period)
	hex := fmt.Sprintf("%x", floor)
	key := sha256.Sum256([]byte(hex))
	c, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}
	return c, nil
}

func Encode(msg string, time_period int64) (string, error) {
	if len(msg) < 16 {
		return "", errors.New("too short")
	}
	resBytes := make([]byte, 16)
	c, err := currentCipher(time_period)
	if err != nil {
		return "", err
	}
	c.Encrypt(resBytes, []byte(msg))
	res := hex.EncodeToString(resBytes)
	return res, nil
}

func Decode(msg string, time_period int64) (string, error) {
	if len(msg) < 16 {
		return "", errors.New("too short")
	}
	resBytes := make([]byte, 16)
	c, err := currentCipher(time_period)
	if err != nil {
		return "", err
	}
	str, err := hex.DecodeString(msg)
	c.Decrypt(resBytes, []byte(str))

	res := string(resBytes)
	return res, nil
}

func main() {
	c, err := Encode(`abcdabcdabcdabcd1`, 1)
	fmt.Println()
	if err != nil {
		log.Printf(err.Error())
	}
	// time.Sleep(2000000000)
	b, err := Decode(c, 1)

	fmt.Println(b)
}
