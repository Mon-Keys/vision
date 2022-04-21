package domain

type NewAccount struct {
	Fullname string `json:"Fullname"`
	UserID   int32
	RoleID   int32
}

type Account struct {
	ID       int32  `json:"id,omitempty"`
	Fullname string `json:"Fullname"`
	RoleID   int32
	UserID   int32
}

type AccountRepository interface {
	CreateAccount(na *NewAccount) error
	Find(fn string) ([]Account, error)
}

type AccountUsecase interface {
	CreateAccount(na *NewAccount) error
}
