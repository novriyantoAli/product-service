package usecase

import (
	"github.com/novriyantoAli/product-service/domain"
	"github.com/novriyantoAli/product-service/model"
	"github.com/sirupsen/logrus"
)

type usecase struct {
	Repo domain.ProductRepository
}

func NewUsecase(repo domain.ProductRepository) domain.ProductUsecase {
	return &usecase{Repo: repo}
}

func (uc *usecase) Delete(param uint) error {
	err := uc.Repo.Delete(param)
	if err != nil {
		logrus.Error(err)
	}

	return err
}

func (uc *usecase) Save(param *domain.Product) error {
	err := uc.Repo.Save(param)
	if err != nil {
		logrus.Error(err)
	}

	return err
}

func (uc *usecase) Create(param *domain.Product) error {
	err := uc.Repo.Create(param)
	if err != nil {
		logrus.Error(err)
	}

	return err
}

func (uc *usecase) FindProduct(param *domain.Product) (res []domain.Product, err error) {
	res, err = uc.Repo.Find(param)
	if err != nil {
		logrus.Error(err)
	}
	return
}

func (uc *usecase) Find(param *model.Product) (*model.ProductList, error) {
	// convert from proto object to local struct
	pb := new(model.ProductList)
	pb.List = make([]*model.Product, 0)

	prod := domain.Product{
		ID:      uint(param.ID),
		Name:    param.Name,
		ValUnit: param.ValUnit,
		ValVal:  uint(param.ValVal),
		Price:   uint(param.ValVal),
		Profile: param.Profile,
	}

	products, err := uc.Repo.Find(&prod)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	// res.List = make([]*model.Product, 0)

	for _, v := range products {
		prd := model.Product{
			ID:      uint64(v.ID),
			Name:    v.Name,
			ValUnit: v.ValUnit,
			ValVal:  uint64(v.ValVal),
			Price:   uint64(v.Price),
			Profile: v.Profile,
		}
		pb.List = append(pb.List, &prd)
	}

	return pb, err
}
