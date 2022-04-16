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
	CreateUser(nuna *NewUser) error
	FindByNickname(nickname string) (*User, error)
	Find(email string) (*User, error)
	Update(user *User) (*User, error)
}

type UserUsecase interface {
	SignUpUser(nuna *NewUserWithoutAccount) ([]User, error)
	DuplicateUser(user *User) ([]User, error)
	Find(nickname string) (*User, error)
	Update(user *User) (*User, error)
}
