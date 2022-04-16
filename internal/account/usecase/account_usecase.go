package usecase

import "github.com/perlinleo/vision/internal/domain"

type accountUsecase struct {
	accountRepository domain.AccountRepository
}

func NewUserUsecase(r domain.AccountRepository) domain.AccountUsecase {
	return &accountUsecase{
		accountRepository: r,
	}
}

func (au accountUsecase) CreateAccount(na *domain.NewAccount) error {
	return nil
}
