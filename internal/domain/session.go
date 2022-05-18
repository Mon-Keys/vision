package domain

type UserSession struct {
	Cookie     string
	UserID     int32
	Expiration int32
}

type AccountSession struct {
	Cookie     string
	AccountID  int32
	Expiration int32
}

type LoginCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SessionRepository interface {
	GetUserSessionByCookie(cookie string) (*UserSession, error)
	NewUserSessionCookie(session *UserSession) error
	DeleteUserSessionCookie(cookie string) error
	GetAccountSessionByCookie(cookie string) (*AccountSession, error)
	NewAccountSessionCookie(session *AccountSession) error
	DeleteAccountSessionCookie(cookie string) error
}

type SessionUsecase interface {
	Login(session LoginCredentials) (*UserSession, *AccountSession, error)
	Logout(accountCookie string, userCookie string) error
	GetUserSessionByCookie(cookie string) (*UserSession, error)
	GetAccountSessionByCookie(cookie string) (*AccountSession, error)
}
