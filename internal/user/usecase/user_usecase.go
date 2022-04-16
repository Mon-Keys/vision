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

func (u userUsecase) SignUpUser(user *domain.NewUserWithoutAccount) ([]domain.User, error) {
	newUser := new(domain.NewUser)
	newUser.Email = user.Email
	newUser.Password = user.Password
	u.userRepository.CreateUser(newUser)
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
