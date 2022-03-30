package usecase

import "github.com/perlinleo/vision/internal/domain"

type userUsecase struct {
	userRepository domain.UserRepository
}

func NewUserUsecase(r domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepository: r,
	}
}

func (u userUsecase) CreateUser(user *domain.User) ([]domain.User, error) {
	return nil, nil
}

func (u userUsecase) DuplicateUser(user *domain.User) ([]domain.User, error) {
	return nil, nil
}
func (u userUsecase) Find(nickname string) (*domain.User, error) {
	return nil, nil
}
func (u userUsecase) Update(user *domain.User) (*domain.User, error) {
	return nil, nil
}
