package usecase

import (
	"github.com/novriyantoAli/product-service/domain"
	"github.com/sirupsen/logrus"
)

type usecase struct {
	Repo domain.ProfileRepository
}

func NewUsecase(repo domain.ProfileRepository) domain.ProfileUsecase {
	return &usecase{Repo: repo}
}

func (uc *usecase) Find(param *domain.Profile) (res []domain.Profile, err error) {
	res, err = uc.Repo.Find(param)
	if err != nil {
		logrus.Error(err)
	}

	return
}

func (uc *usecase) FindRadcheck(param *domain.Radgroupcheck) (res []domain.Radgroupcheck, err error) {
	res, err = uc.Repo.FindRadcheck(param)
	if err != nil {
		logrus.Error(err)
	}

	return
}

func (uc *usecase) FindRadreply(param *domain.Radgroupreply) (res []domain.Radgroupreply, err error) {
	res, err = uc.Repo.FindRadreply(param)
	if err != nil {
		logrus.Error(err)
	}

	return
}
