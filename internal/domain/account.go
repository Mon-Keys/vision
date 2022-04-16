package domain

type NewAccount struct {
	Fullname string `json:"Fullname"`
}

type Account struct {
	ID       int32  `json:"id,omitempty"`
	Fullname string `json:"Fullname"`
}

type AccountRepository interface {
	CreateAccount(na *NewAccount) error
	Find(fn string) ([]Account, error)
}

type AccountUsecase interface {
	CreateAccount(na *NewAccount) error
}
