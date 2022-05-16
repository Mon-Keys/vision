package redis

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/patrickmn/go-cache"
	"github.com/perlinleo/vision/internal/domain"
)

var ctx = context.Background()

type sessionRedis struct {
	rdbUser    *redis.Client
	rdbAccount *redis.Client
	Cache      *cache.Cache
}

func NewSessionRedisRepository(rcUser *redis.Client, rcAccount *redis.Client, Cache *cache.Cache) domain.SessionRepository {
	return &sessionRedis{
		rcUser,
		rcAccount,
		Cache,
	}
}

func (rr sessionRedis) GetUserSessionByCookie(cookie string) (*domain.UserSession, error) {
	id, err := rr.rdbUser.Get(ctx, cookie).Result()
	if err != nil {
		return nil, err
	}

	fmt.Println(id)

	decodedID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	result := &domain.UserSession{
		Cookie: cookie,
		UserID: int32(decodedID),
	}

	return result, nil
}

func (rr sessionRedis) GetAccountSessionByCookie(cookie string) (*domain.AccountSession, error) {
	id, err := rr.rdbUser.Get(ctx, cookie).Result()
	if err != nil {
		return nil, err
	}

	fmt.Println(id)

	decodedID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	result := &domain.AccountSession{
		Cookie:    cookie,
		AccountID: int32(decodedID),
	}

	return result, nil
}

func (rr sessionRedis) NewUserSessionCookie(session *domain.UserSession) error {
	err := rr.rdbUser.Set(ctx, session.Cookie, session.UserID, 0).Err()

	if err != nil {
		return err
	}
	return nil
}

func (rr sessionRedis) NewAccountSessionCookie(session *domain.AccountSession) error {
	err := rr.rdbUser.Set(ctx, session.Cookie, session.AccountID, 0).Err()

	if err != nil {
		return err
	}
	return nil
}

func (rr sessionRedis) DeleteUserSessionCookie(cookie string) error {
	err := rr.rdbUser.Del(ctx, cookie).Err()
	if err != nil {
		return err
	}

	return nil
}

func (rr sessionRedis) DeleteAccountSessionCookie(cookie string) error {
	err := rr.rdbAccount.Del(ctx, cookie).Err()
	if err != nil {
		return err
	}

	return nil
}
