package usecase

import "github.com/perlinleo/vision/internal/domain"

type sessionUsecase struct {
	sessionRepository domain.SessionRepository
}

func NewSessionUsecase(r domain.SessionRepository) domain.SessionUsecase {
	return &sessionUsecase{
		sessionRepository: r,
	}
}

func (su sessionUsecase) Create(session Session) error {
	su.sessionRepository.NewSessionCookie()
}
func (su sessionUsecase) Logout(session Session) error {

}

func (su sessionUsecase) GetSessionByCookie(cookie string) (*Session, error) {

}
