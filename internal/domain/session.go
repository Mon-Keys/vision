package domain

import "time"

type Session struct {
	Cookie     string
	ID         string
	Expiration time.Time
}

type SessionRepository interface {
	GetSessionByCookie(cookie string) (*Session, error)
	NewSessionCookie(session *Session) error
	DeleteSessionCookie(cookie string) error
}

type SessionUsecase interface {
	Create(session Session) error
	Logout(session Session) error
	GetSessionByCookie(cookie string) (*Session, error)
}
