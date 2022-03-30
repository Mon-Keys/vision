package domain

type NewAccount struct {
	Fullname string `json:"Fullname"`
}

type Account struct {
	ID       int32  `json:"id,omitempty"`
	Fullname string `json:"Fullname"`
}
