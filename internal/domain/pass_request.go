package domain

type NewPassRequest struct {
	Comment string  `json:"comment"`
	IssueTo Account `json:"account"`
}

type PassRequest struct {
	Comment  string  `json:"comment"`
	IssueTo  Account `json:"account"`
	Approved bool    `json:"approved"`
}
