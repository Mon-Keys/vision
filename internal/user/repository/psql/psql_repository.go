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

func (m *userPsql) Create(forum *domain.User) error {
	return nil
}
func (m *userPsql) FindByNickname(nickname string) (*domain.User, error) {
	return nil, nil
}
func (m *userPsql) Find(nickname string, email string) ([]domain.User, error) {
	return nil, nil
}
func (m *userPsql) Update(user *domain.User) (*domain.User, error) {
	return nil, nil
}
