package domain

import (
	"time"
)

type User struct {
	ID       int32     `json:"id,omitempty"`
	Email    string    `json:"email"`
	Created  time.Time `json:"created"`
	Password string    `json:"password"`
}

type UserAccount struct {
	Name       string    `json:"name"`
	Created    time.Time `json:"created"`
	UserRoleID int32     `json:"RoleID"`
}

type UserAccountFull struct {
	UserID     int32     `json:"userID"`
	Email      string    `json:"email"`
	Name       string    `json:"name"`
	Created    time.Time `json:"created"`
	UserRoleID int32     `json:"RoleID"`
	AccountID  int32     `json:"accountID"`
}

type NewUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type NewUserWithoutAccount struct {
	FirstName string `json:firstName`
	LastName  string `json:lastName`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UserRepository interface {
	CreateUser(nuna *NewUser) (int32, error)
	FindByNickname(nickname string) (*User, error)
	Find(email string) (*User, error)
	Update(user *User) (*User, error)
	FindUserByID(userID int32) (*User, error)
	All() ([]UserAccountFull, error)
	FindAllByName(name string) ([]UserAccountFull, error)
}

type UserUsecase interface {
	FindUserAccountByID(userID int32) (*User, *Account, error)
	SignUpUser(nuna *NewUserWithoutAccount) error
	DuplicateUser(user *User) ([]User, error)
	Find(nickname string) (*User, error)
	Update(user *User) (*User, error)
	All() ([]UserAccountFull, error)
	FindAllByName(name string) ([]UserAccountFull, error)
}
