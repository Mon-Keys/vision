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

func (u userUsecase) FindUserAccountByID(userID int32) (*domain.User, *domain.Account, error) {
	user, err := u.userRepository.FindUserByID(userID)
	if err != nil {
		return nil, nil, err
	}

	account, err := u.accountRepository.FindAccountByUserID(int(userID))

	if err != nil {
		return nil, nil, err
	}

	return user, account, nil
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

func (u userUsecase) All() ([]domain.UserAccountFull, error) {
	users, err := u.userRepository.All()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u userUsecase) FindAllByName(name string) ([]domain.UserAccountFull, error) {
	users, err := u.userRepository.FindAllByName(name)
	if err != nil {
		return nil, err
	}

	return users, nil
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
