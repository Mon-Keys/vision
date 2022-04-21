package usecase

import (
	"github.com/perlinleo/vision/internal/domain"
	cookie "github.com/perlinleo/vision/internal/pkg/cookie_generator"
)

type sessionUsecase struct {
	sessionRepository domain.SessionRepository
	userRepository    domain.UserRepository
}

func NewSessionUsecase(r domain.SessionRepository, ur domain.UserRepository) domain.SessionUsecase {
	return &sessionUsecase{
		sessionRepository: r,
		userRepository:    ur,
	}
}

func (su sessionUsecase) Login(session domain.LoginCredentials) (*domain.Session, error) {
	// su.sessionRepository.NewSessionCookie()
	userData, err := su.userRepository.Find(session.Email)
	if err != nil {
		return nil, domain.ErrorCantFindUserWithEmail
	}
	if userData.Password != session.Password {
		return nil, domain.ErrorWrongPassword
	}

	userSession := cookie.CreateAuthSessionUUID(string(userData.ID))

	return userSession, nil
}
func (su sessionUsecase) Logout(session domain.Session) error {
	return nil
}

func (su sessionUsecase) GetSessionByCookie(cookie string) (*domain.Session, error) {
	return nil, nil
}
