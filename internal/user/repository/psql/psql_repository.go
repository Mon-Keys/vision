package psql

import (
	"github.com/jackc/pgx"
	"github.com/patrickmn/go-cache"
	"github.com/perlinleo/vision/internal/domain"
)

type userPsql struct {
	Conn  *pgx.ConnPool
	Cache *cache.Cache
}

func NewUserPSQLRepository(ConnectionPool *pgx.ConnPool, Cache *cache.Cache) domain.UserRepository {
	return &userPsql{
		ConnectionPool,
		Cache,
	}
}

func (m *userPsql) CreateUser(user *domain.NewUser) error {
	query := "INSERT INTO users (user_password_hash, user_email) " +
		"VALUES ($1, $2)"
	_, err := m.Conn.Exec(
		query, user.Password, user.Email)
	return err
}
func (m *userPsql) FindByNickname(nickname string) (*domain.User, error) {
	return nil, nil
}
func (m *userPsql) Find(email string) (*domain.User, error) {
	query := "SELECT * FROM users WHERE user_email=$1"
	userData := new(domain.User)

	err := m.Conn.QueryRow(query, email).Scan(&userData.ID, &userData.Password, &userData.Email, &userData.Created)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
func (m *userPsql) Update(user *domain.User) (*domain.User, error) {
	return nil, nil
}
