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

type CheckResult struct {
	PassID         int32     `json:"pass_id"`
	ExpirationDate time.Time `json:"expiration_date"`
	IssueDate      time.Time `json:"issue_date"`
	UserRoleID     int32     `json:"user_role_id"`
	PassName       string    `json:"pass_name"`
	UserEmail      string    `json:"user_email"`
	UserCreated    time.Time `json:"user_created"`
	UserName       string    `json:"owner_name"`
	Access         bool      `json:"access"`
	Error          string    `json:"error"`
}

type PassRepository interface {
	CreatePass(pass Pass) (error, int32)
	ActivatePass(passID int32) error
	DisablePass(passID int32) error
	DeletePass(passID int32) error
	PassesByAccountID(accountID int32) ([]Pass, error)
	CheckPassByData(data string) (*CheckResult, error)
	FindPassByID(passID int32) (*Pass, error)
	UpdatePassTime(data time.Time, passID int32) error
	AddPassage(passage AddPassage) error
	AllPassages() ([]Passage, error)
}

type PassUsecase interface {
	GetUserPasses(accountID int32) ([]Pass, error)
	CheckPass(data string) (*CheckResult, error)
	AllPassages() ([]Passage, error)
}
