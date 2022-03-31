package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/patrickmn/go-cache"
	"github.com/perlinleo/vision/internal/domain"
)

var ctx = context.Background()

type sessionRedis struct {
	rdb   *redis.Client
	Cache *cache.Cache
}

func NewSessionRedisRepository(rc *redis.Client, Cache *cache.Cache) domain.SessionRepository {
	return &sessionRedis{
		rc,
		Cache,
	}
}

func (rr sessionRedis) GetSessionByCookie(cookie string) (*domain.Session, error) {
	id, err := rr.rdb.Get(ctx, cookie).Result()
	if err != nil {
		return nil, err
	}

	result := &domain.Session{
		Cookie: cookie,
		ID:     id,
	}

	return result, nil
}

func (rr sessionRedis) NewSessionCookie(session *domain.Session) error {
	err := rr.rdb.Set(ctx, session.Cookie, session.ID, 0).Err()

	if err != nil {
		return err
	}
	return nil
}

func (rr sessionRedis) DeleteSessionCookie(cookie string) error {
	err := rr.rdb.Del(ctx, cookie).Err()
	if err != nil {
		return err
	}

	return nil
}
