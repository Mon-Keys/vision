package usecase

import (
	"errors"

	"github.com/perlinleo/vision/internal/domain"
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

func (su sessionUsecase) Login(session domain.LoginCredentials) error {
	// su.sessionRepository.NewSessionCookie()
	userData, err := su.userRepository.Find(session.Email)
	if err != nil {
		return err
	}
	if userData.Password != session.Password {
		return errors.New("Wrong")
	}

	return nil
}
func (su sessionUsecase) Logout(session domain.Session) error {
	return nil
}

func (su sessionUsecase) GetSessionByCookie(cookie string) (*domain.Session, error) {
	return nil, nil
}
