package domain

import (
	"time"
)

type User struct {
	ID       int32     `json:"id,omitempty"`
	Email    string    `json:"email"`
	Created  time.Time `json:"created"`
	Login    string    `json:"login"`
	Password string    `json:"password"`
}

type NewUser struct {
	Email    string    `json:"email"`
	Created  time.Time `json:"created"`
	Login    string    `json:"login"`
	Password string    `json:"password"`
}

type UserRepository interface {
	Create(forum *User) error
	FindByNickname(nickname string) (*User, error)
	Find(nickname string, email string) ([]User, error)
	Update(user *User) (*User, error)
}

type UserUsecase interface {
	CreateUser(user *User) ([]User, error)
	DuplicateUser(user *User) ([]User, error)
	Find(nickname string) (*User, error)
	Update(user *User) (*User, error)
}
