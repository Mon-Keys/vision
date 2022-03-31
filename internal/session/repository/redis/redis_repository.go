package redis

import (
	"github.com/jackc/pgx"
	"github.com/patrickmn/go-cache"
)

type sessionRedis struct {
	Client *pgx.ConnPool
	Cache  *cache.Cache
}
