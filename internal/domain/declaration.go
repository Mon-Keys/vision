package domain

import "time"

type AskRole struct {
	RoleID  int32  `json:"roleID"`
	Comment string `json:"comment"`
}

type AskRoleDeclaration struct {
	CreatorID   int32  `json:"account_id"`
	RoleID      int32  `json:"roleID"`
	Comment     string `json:"comment"`
	CurrentRole int32  `json:"currentRoleID"`
}

type AskPass struct {
	Comment            string    `json:"comment"`
	PassExpirationDate time.Time `json:"expirationDate"`
	PassName           string    `json:"name"`
}

type AskTimeDeclaration struct {
	Comment      string    `json:"comment"`
	CreatorID    int32     `json:"creatorID"`
	PassID       int32     `json:"passID"`
	TimeExtended time.Time `json:"time"`
}

type AskTime struct {
	Comment      string    `json:"comment"`
	PassID       int32     `json:"passID"`
	TimeExtended time.Time `json:"time"`
}

type AskPassDeclarationPass struct {
	Declaration AskPassDeclaration `json:"declaration"`
	Pass        Pass               `json:"pass"`
}

type AskTimeDeclarationPass struct {
	Declaration AskTimeDeclaration `json:"declaration"`
	Pass        Pass               `json:"pass"`
}

type AskPassDeclaration struct {
	Comment            string    `json:"comment"`
	CreatorID          int32     `json:"creatorID"`
	NewPassID          int32     `json:"passID"`
	PassExpirationDate time.Time `json:"expirationDate"`
	Approved           bool      `json:"approved"`
	Denied             bool      `json:"denied"`
}

type DeclarationCommon struct {
	DeclarationType    int32     `json:"type"`
	CreatedDate        time.Time `json:"created"`
	AuthorName         string    `json:"creator"`
	DeclarationInnerID int32     `json:"innerID"`
	AuthorID           int32     `json:"creatorID"`
	Accepted           bool      `json:"accepted"`
	Denied             bool      `json:"denied"`
}

type DeclarationRepository interface {
	CreateRoleDeclaration(declaration AskRoleDeclaration) error
	CreatePassDeclaration(declaration AskPassDeclaration) error
	CreateTimeDeclaration(declaration AskTimeDeclaration) error
	AcceptRoleDeclaration(RoleDeclarationID int32) error
	AcceptPassDeclaration(PassDeclarationID int32) error
	AcceptTimeDeclaration(TimeDeclarationID int32) error
	DenyRoleDeclaration(RoleDeclarationID int32) error
	DenyPassDeclaration(PassDeclarationID int32) error
	DenyTimeDeclaration(TimeDeclarationID int32) error
	PassRequestDeclarationByID(AskPassDeclarationID int32) (*AskPassDeclaration, error)
	AllDeclarations() ([]DeclarationCommon, error)
	AllDeclarationsByAccountID(accountID int32) ([]DeclarationCommon, error)
	PassRequestDeclarationsAll() ([]DeclarationCommon, error)
	RoleRequestDeclarationsAll() ([]DeclarationCommon, error)
	TimeRequestDeclarationsAll() ([]DeclarationCommon, error)
	PassRequestDeclarationsByAccountID(accountID int32) ([]DeclarationCommon, error)
	RoleRequestDeclarationsByAccountID(accountID int32) ([]DeclarationCommon, error)
	TimeRequestDeclarationsByAccountID(accountID int32) ([]DeclarationCommon, error)
	RoleDeclarationByID(id int32) (*AskRoleDeclaration, error)
	TimeDeclarationByID(id int32) (*AskTimeDeclaration, error)
}

type DeclarationUsecase interface {
	CreateRoleDeclaration(declaration AskRoleDeclaration) error
	CreatePassDeclaration(declaration AskPass, userID int32) error
	CreateTimeDeclaration(declaration AskTimeDeclaration) error
	AllDeclarations() ([]DeclarationCommon, error)
	AllDeclarationsByID(accountID int32) ([]DeclarationCommon, error)
	AcceptDeclaration(DeclarationCommon) error
	DenyDeclaration(DeclarationCommon) error
	RoleDeclarationByID(id int32) (*AskRoleDeclaration, error)
	PassDeclarationByID(id int32) (*AskPassDeclaration, *Pass, error)
	TimeDeclarationByID(id int32) (*AskTimeDeclaration, *Pass, error)
}
