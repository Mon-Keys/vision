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

func (ap *accountPsql) ChangeUserRole(userID int, newRoleID int) error {
	query := `UPDATE account
				SET account_role_id = $1
				WHERE account_id = $2 returning account_id
	`
	var passIDCheck int32
	err := ap.Conn.QueryRow(query, newRoleID, userID).Scan(&passIDCheck)
	if err != nil {
		return err
	}
	return nil
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

func (ap *accountPsql) FindAccountByUserID(userID int) (*domain.Account, error) {
	account := new(domain.Account)

	query := "SELECT * FROM account WHERE account_user_id= $1"

	err := ap.Conn.QueryRow(query, &userID).Scan(&account.ID, &account.RoleID, &account.UserID, &account.Fullname)

	if err != nil {
		return nil, err
	}

	return account, nil
}
