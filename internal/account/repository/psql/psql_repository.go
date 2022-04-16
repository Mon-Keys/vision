package psql

import (
	"github.com/jackc/pgx"
	"github.com/patrickmn/go-cache"
	"github.com/perlinleo/vision/internal/domain"
)

type accountPsql struct {
	Conn  *pgx.ConnPool
	Cache *cache.Cache
}

func NewUserPSQLRepository(ConnectionPool *pgx.ConnPool, Cache *cache.Cache) domain.AccountRepository {
	return &accountPsql{
		ConnectionPool,
		Cache,
	}
}

func (ap *accountPsql) CreateAccount(na *domain.NewAccount) error {
	return nil
}

func (ap *accountPsql) Find(fn string) ([]domain.Account, error) {
	return nil, nil
}
