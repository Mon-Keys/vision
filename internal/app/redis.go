package vision

import (
	"errors"

	"github.com/go-redis/redis/v8"
)

func NewRedisDataBase(addr string, password string, db int) (redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	if client == nil {
		return *client, errors.New("Can`t connect to redis")
	}
	return *client, nil
}
