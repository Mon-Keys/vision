package domain

import "time"

type Pass struct {
	PassID         int32     `json:"id"`
	AccountID      int32     `json:"account_id"`
	IssueDate      time.Time `json:"issueDate"`
	ExpirationDate time.Time `json:"dueDate"`
	SecureData     string    `json:"secure_data"`
	Name           string    `json:"pass_name"`
	DynamicQR      bool      `json:"is_dynamic"`
	Active         bool      `json:"is_active"`
	Disabled       bool      `json:"is_disabled"`
}

type PassRepository interface {
	CreatePass(pass Pass) (error, int32)
	ActivatePass(passID int32) error
	DisablePass(passID int32) error
	DeletePass(passID int32) error
	PassesByAccountID(accountID int32) ([]Pass, error)
}

type PassUsecase interface {
	GetUserPasses(accountID int32) ([]Pass, error)
}
