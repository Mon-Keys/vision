package status

import "github.com/perlinleo/vision/internal/domain"

type statusUsecase struct {
	statusRepository domain.StatusRepository
}

func NewStatusUsecase(r domain.StatusRepository) domain.StatusUsecase {
	return &statusUsecase{
		statusRepository: r,
	}
}

func (u statusUsecase) Status() (*domain.Status, error) {
	usersAmount, err := u.statusRepository.GetUsersAmount()
	if err != nil {
		return nil, err
	}

	currentStatus := &domain.Status{
		UsersAmount: usersAmount,
	}

	return currentStatus, nil
}
