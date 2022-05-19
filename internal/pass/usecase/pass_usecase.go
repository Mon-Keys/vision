package usecase

import (
	"fmt"

	"github.com/perlinleo/vision/internal/domain"
	"github.com/perlinleo/vision/internal/pkg/timecode_generator"
)

type passUsecase struct {
	passRepostiory domain.PassRepository
}

func NewPassUsecase(p domain.PassRepository) domain.PassUsecase {
	return &passUsecase{
		passRepostiory: p,
	}
}

func (p passUsecase) CheckPass(data string) (*domain.CheckResult, error) {
	dataDecoded, err := timecode_generator.Decode(data, 10)
	if err != nil {
		return nil, err
	}
	checkRes, err := p.passRepostiory.CheckPassByData(dataDecoded)

	if err != nil {
		checkRes := new(domain.CheckResult)
		checkRes.Access = false
		checkRes.Error = "Not found"
		return checkRes, err
	}
	checkRes.Access = true
	passage := new(domain.AddPassage)
	passage.Exit = false
	passage.PassID = checkRes.PassID
	passage.Status = 1
	err = p.passRepostiory.AddPassage(*passage)
	return checkRes, nil
}

func (p passUsecase) GetUserPasses(accountID int32) ([]domain.Pass, error) {
	passes, err := p.passRepostiory.PassesByAccountID(accountID)
	for index, _ := range passes {
		fmt.Println(passes[index].SecureData)
		passes[index].SecureData, err = timecode_generator.Encode(passes[index].SecureData, 10)
		fmt.Println(passes[index].SecureData)
		decodedVal, err := timecode_generator.Decode(passes[index].SecureData, 10)
		fmt.Println(decodedVal)
		if err != nil {
			return nil, err
		}
	}
	if err != nil {
		return nil, err
	}
	return passes, nil
}

func (p passUsecase) AllPassages() ([]domain.Passage, error) {
	return p.passRepostiory.AllPassages()
}
