package usecase

import "github.com/perlinleo/vision/internal/domain"

type userUsecase struct {
	userRepository    domain.UserRepository
	accountRepository domain.AccountRepository
}

func NewUserUsecase(r domain.UserRepository, ar domain.AccountRepository) domain.UserUsecase {
	return &userUsecase{
		userRepository:    r,
		accountRepository: ar,
	}
}

func (u userUsecase) SignUpUser(user *domain.NewUserWithoutAccount) error {
	newUser := new(domain.NewUser)
	newUser.Email = user.Email
	newUser.Password = user.Password
	userID, err := u.userRepository.CreateUser(newUser)

	if err != nil {
		// такой пользователь уже существует
		return domain.ErrorUserConflict
	}

	newAccount := new(domain.NewAccount)

	newAccount.Fullname = user.LastName + " " + user.FirstName
	newAccount.UserID = userID

	newAccount.RoleID = defaultNewAccountRoleID

	err = u.accountRepository.CreateAccount(newAccount)

	if err != nil {
		return err
	}

	return nil
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
