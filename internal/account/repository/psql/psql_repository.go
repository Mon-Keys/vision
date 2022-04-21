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

func NewAccountPSQLRepository(ConnectionPool *pgx.ConnPool, Cache *cache.Cache) domain.AccountRepository {
	return &accountPsql{
		ConnectionPool,
		Cache,
	}
}

// Создает аккаунт пользователя в базе данных PostgreSQL
func (ap *accountPsql) CreateAccount(na *domain.NewAccount) error {
	query := "INSERT INTO account (account_role_id, account_user_id, account_fullname) " +
		"VALUES ($1, $2, $3)"
	_, err := ap.Conn.Exec(
		query, na.RoleID, na.UserID, na.Fullname)

	return err
}

func (ap *accountPsql) Find(fn string) ([]domain.Account, error) {
	return nil, nil
}
