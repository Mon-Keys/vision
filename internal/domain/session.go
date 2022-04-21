package domain

type Session struct {
	Cookie     string
	ID         string
	Expiration int32
}

type LoginCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SessionRepository interface {
	GetSessionByCookie(cookie string) (*Session, error)
	NewSessionCookie(session *Session) error
	DeleteSessionCookie(cookie string) error
}

type SessionUsecase interface {
	Login(session LoginCredentials) (*Session, error)
	Logout(session Session) error
	GetSessionByCookie(cookie string) (*Session, error)
}
