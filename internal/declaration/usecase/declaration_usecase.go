package usecase

import (
	"errors"
	"fmt"
	"time"

	"github.com/perlinleo/vision/internal/domain"
	"github.com/perlinleo/vision/internal/pkg/secure_data_generator"
)

type declarationUsecase struct {
	declarationRepository domain.DeclarationRepository
	passRepostiory        domain.PassRepository
}

func NewDeclarationUsecase(r domain.DeclarationRepository, p domain.PassRepository) domain.DeclarationUsecase {
	return &declarationUsecase{
		declarationRepository: r,
		passRepostiory:        p,
	}
}

func (u declarationUsecase) CreateRoleDeclaration(declaration domain.AskRoleDeclaration) error {
	return u.declarationRepository.CreateRoleDeclaration(declaration)
}
func (u declarationUsecase) CreatePassDeclaration(declaration domain.AskPass, userID int32) error {
	newInactivePass := new(domain.Pass)
	newInactivePass.AccountID = userID
	newInactivePass.DynamicQR = true
	newInactivePass.ExpirationDate = declaration.PassExpirationDate
	newInactivePass.Name = declaration.PassName
	newInactivePass.SecureData = secure_data_generator.RandStringBytesMaskImprSrcSB(1024)
	newInactivePass.IssueDate = time.Now()
	newInactivePass.Active = false
	err, newPassID := u.passRepostiory.CreatePass(*newInactivePass)
	if err != nil {
		return err
	}
	newPassDeclaration := new(domain.AskPassDeclaration)
	newPassDeclaration.Comment = declaration.Comment
	newPassDeclaration.CreatorID = userID
	newPassDeclaration.NewPassID = newPassID
	newPassDeclaration.PassExpirationDate = declaration.PassExpirationDate
	err = u.declarationRepository.CreatePassDeclaration(*newPassDeclaration)
	if err != nil {
		return err
	}

	return nil
}
func (u declarationUsecase) CreateTimeDeclaration(declaration domain.AskTimeDeclaration) error {
	return nil
}
func (u declarationUsecase) AllDeclarations() ([]domain.DeclarationCommon, error) {
	declarations, err := u.declarationRepository.AllDeclarations()
	if err != nil {
		return nil, err
	}
	return declarations, nil
}

func (u declarationUsecase) AllDeclarationsByID(accountID int32) ([]domain.DeclarationCommon, error) {
	declarations, err := u.declarationRepository.AllDeclarationsByAccountID(accountID)
	if err != nil {
		return nil, err
	}
	return declarations, nil
}

func (u declarationUsecase) AcceptDeclaration(declaration domain.DeclarationCommon) error {
	var err error
	switch declaration.DeclarationType {
	case 0:
		// pass
		dec, err := u.declarationRepository.PassRequestDeclarationByID(declaration.DeclarationInnerID)
		if err != nil {
			return err
		}
		if dec.Denied {
			return errors.New("Declaration already denied")
		}
		if dec.Approved {
			return errors.New("Declaration already approved")
		} else {
			err = u.passRepostiory.ActivatePass(dec.NewPassID)
			if err != nil {
				return err
			}
			err = u.declarationRepository.AcceptPassDeclaration(declaration.DeclarationInnerID)
			if err != nil {
				return err
			}
		}
	case 1:
		err = u.declarationRepository.AcceptTimeDeclaration(declaration.DeclarationInnerID)
	case 2:
		err = u.declarationRepository.AcceptRoleDeclaration(declaration.DeclarationInnerID)
	default:
		fmt.Println("Unknown declaration type")
	}
	if err != nil {
		return err
	}

	return nil
}

func (u declarationUsecase) DenyDeclaration(declaration domain.DeclarationCommon) error {
	var err error
	switch declaration.DeclarationType {
	case 0:
		// pass
		dec, err := u.declarationRepository.PassRequestDeclarationByID(declaration.DeclarationInnerID)
		if err != nil {
			return err
		}
		if dec.Approved {
			return errors.New("Declaration already approved")
		}
		if dec.Denied {
			return errors.New("Declaration already denied")
		} else {
			err = u.passRepostiory.DisablePass(dec.NewPassID)
			if err != nil {
				return err
			}
			err = u.declarationRepository.DenyPassDeclaration(declaration.DeclarationInnerID)
			if err != nil {
				return err
			}
		}
	case 1:
		err = u.declarationRepository.DenyTimeDeclaration(declaration.DeclarationInnerID)
	case 2:
		err = u.declarationRepository.DenyRoleDeclaration(declaration.DeclarationInnerID)
	default:
		fmt.Println("Unknown declaration type")
	}
	if err != nil {
		return err
	}

	return nil
}
