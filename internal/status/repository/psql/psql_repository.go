package psql

import (
	"github.com/jackc/pgx"
	"github.com/patrickmn/go-cache"
	"github.com/perlinleo/vision/internal/domain"
)

type statusPsql struct {
	Conn  *pgx.ConnPool
	Cache *cache.Cache
}

func NewStatusPSQLRepository(ConnectionPool *pgx.ConnPool, Cache *cache.Cache) domain.StatusRepository {
	return &statusPsql{
		ConnectionPool,
		Cache,
	}
}

func (r statusPsql) GetUsersAmount() (uint32, error) {
	var usersAmount uint32

	row := r.Conn.QueryRow(CountUsersPSQLQuerry)
	err := row.Scan(&usersAmount)

	if err != nil {
		return 0, err
	}

	return usersAmount, nil
}
