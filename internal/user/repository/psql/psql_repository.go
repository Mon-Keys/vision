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

func (m *userPsql) All() ([]domain.UserAccountFull, error) {
	var users []domain.UserAccountFull

	query := `SELECT user_id, user_email, user_created_date,
	 account_id, account_role_id,  
	 account_fullname from users 
	 JOIN account ON users.user_id=account.account_user_id`

	rows, err := m.Conn.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		obj := domain.UserAccountFull{}
		err := rows.Scan(&obj.UserID, &obj.Email, &obj.Created, &obj.AccountID, &obj.UserRoleID, &obj.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, obj)
	}
	return users, nil
}

func (m *userPsql) FindAllByName(name string) ([]domain.UserAccountFull, error) {
	var users []domain.UserAccountFull

	query := `SELECT user_id, user_email, user_created_date,
	 account_id, account_role_id,  
	 account_fullname from users 
	 JOIN account ON users.user_id=account.account_user_id
	 where account_fullname LIKE $1`

	rows, err := m.Conn.Query(query, "%"+name+"%")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		obj := domain.UserAccountFull{}
		err := rows.Scan(&obj.UserID, &obj.Email, &obj.Created, &obj.AccountID, &obj.UserRoleID, &obj.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, obj)
	}
	return users, nil
}

func (m *userPsql) FindUserByID(userID int32) (*domain.User, error) {
	query := "SELECT * from users where user_id=$1"
	userData := new(domain.User)

	err := m.Conn.QueryRow(query, userID).Scan(&userData.ID, &userData.Password, &userData.Email, &userData.Created)
	if err != nil {
		return nil, err
	}

	return userData, nil

}

func (m *userPsql) CreateUser(user *domain.NewUser) (int32, error) {
	query := "INSERT INTO users (user_password_hash, user_email) " +
		"VALUES ($1, $2) RETURNING user_id"

	var newUserID int32

	err := m.Conn.QueryRow(
		query, user.Password, user.Email).Scan(&newUserID)

	return newUserID, err
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

	return userData, nil
}
func (m *userPsql) Update(user *domain.User) (*domain.User, error) {
	return nil, nil
}
