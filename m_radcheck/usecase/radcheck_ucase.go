package usecase

import (
	"github.com/novriyantoAli/product-service/domain"
	"github.com/novriyantoAli/product-service/model"
	"gorm.io/gorm"
)

type usecase struct {
	Repo domain.RadcheckRepository
}

func NewUsecase(repo domain.RadcheckRepository) domain.RadcheckUsecase {
	return &usecase{Repo: repo}
}

func (uc *usecase) CreateBatch(params *model.RadcheckList) error {
	// first convert
	err := gorm.ErrEmptySlice

	if len(params.List) > 0 {
		radchecks := make([]domain.Radcheck, 0)
		for _, v := range params.List {
			rc := domain.Radcheck{
				Username:  v.Username,
				Attribute: v.Attribute,
				OP:        v.OP,
				Value:     v.Value,
			}

			radchecks = append(radchecks, rc)
		}

		err = uc.Repo.CreateBatch(radchecks)
	}

	return err
}
