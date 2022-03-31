package domain

type Status struct {
	UsersAmount uint32
}

type StatusRepository interface {
	GetUsersAmount() (uint32, error)
}

type StatusUsecase interface {
	Status() (*Status, error)
}
