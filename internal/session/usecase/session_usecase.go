package usecase

import (
	"github.com/perlinleo/vision/internal/domain"
	cookie "github.com/perlinleo/vision/internal/pkg/cookie_generator"
)

type sessionUsecase struct {
	sessionRepository domain.SessionRepository
	userRepository    domain.UserRepository
	accountRepository domain.AccountRepository
}

func NewSessionUsecase(r domain.SessionRepository, ur domain.UserRepository, ar domain.AccountRepository) domain.SessionUsecase {
	return &sessionUsecase{
		sessionRepository: r,
		userRepository:    ur,
		accountRepository: ar,
	}
}

func (su sessionUsecase) Login(session domain.LoginCredentials) (*domain.UserSession, *domain.AccountSession, error) {
	// su.sessionRepository.NewSessionCookie()
	userData, err := su.userRepository.Find(session.Email)
	if err != nil {
		return nil, nil, domain.ErrorCantFindUserWithEmail
	}

	accountData, err := su.accountRepository.FindAccountByUserID(int(userData.ID))

	if userData.Password != session.Password {
		return nil, nil, domain.ErrorWrongPassword
	}

	accountSession := cookie.CreateAccountSessionUUID(accountData.ID)
	su.sessionRepository.NewAccountSessionCookie(accountSession)

	userSession := cookie.CreateAuthSessionUUID(userData.ID)
	su.sessionRepository.NewUserSessionCookie(userSession)

	return userSession, accountSession, nil
}

func (su sessionUsecase) Logout(accountCookie string, userCookie string) error {
	err := su.sessionRepository.DeleteAccountSessionCookie(accountCookie)
	if err != nil {
		return err
	}
	err = su.sessionRepository.DeleteUserSessionCookie(userCookie)
	if err != nil {
		return err
	}
	return nil
}

func (su sessionUsecase) GetUserSessionByCookie(cookie string) (*domain.UserSession, error) {
	session := new(domain.UserSession)
	session, err := su.sessionRepository.GetUserSessionByCookie(cookie)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (su sessionUsecase) GetAccountSessionByCookie(cookie string) (*domain.AccountSession, error) {
	session := new(domain.AccountSession)
	session, err := su.sessionRepository.GetAccountSessionByCookie(cookie)
	if err != nil {
		return nil, err
	}
	return session, nil
}
