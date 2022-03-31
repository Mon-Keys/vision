package vision

import (
	"errors"

	"github.com/go-redis/redis"
)

func NewRedisDataBase(addr string, password string, db int) (redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	if client == nil {
		return *client, errors.New("Can`t connect to redis")
	}
	return *client, nil
}
