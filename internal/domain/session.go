package domain

type Session struct {
	Cookie string
	ID     uint64
}

type SessionRepository interface {
	GetSessionByCookie(cookie string) (Session, error)
	NewSessionCookie(cookie string, id uint64) error
	DeleteSessionCookie(cookie string) error
}

type SessionUsecase interface {
	Login(session Session) error
	Logout(session Session) error
	GetIDByCookie(session Session) error
}
