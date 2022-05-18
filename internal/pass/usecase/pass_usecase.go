package usecase

import "github.com/perlinleo/vision/internal/domain"

type passUsecase struct {
	passRepostiory domain.PassRepository
}

func NewPassUsecase(p domain.PassRepository) domain.PassUsecase {
	return &passUsecase{
		passRepostiory: p,
	}
}

func (p passUsecase) GetUserPasses(accountID int32) ([]domain.Pass, error) {
	passes, err := p.passRepostiory.PassesByAccountID(accountID)
	if err != nil {
		return nil, err
	}
	return passes, nil
}
