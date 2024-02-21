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

func (uc *usecase) Save(param *domain.Profile) error {
	err := uc.Repo.Save(param)
	if err != nil {
		logrus.Error(err)
	}

	return err
}

func (uc *usecase) SaveRadcheck(param *domain.Radgroupcheck) error {
	err := uc.Repo.SaveRadcheck(param)
	if err != nil {
		logrus.Error(err)
	}

	return err
}

func (uc *usecase) SaveRadreply(param *domain.Radgroupreply) error {
	err := uc.Repo.SaveRadreply(param)
	if err != nil {
		logrus.Error(err)
	}

	return err
}

func (uc *usecase) DeleteRadcheck(id uint) error {
	err := uc.Repo.DeleteRadcheck(id)
	if err != nil {
		logrus.Error(err)
	}

	return err
}

func (uc *usecase) DeleteRadreply(id uint) error {
	err := uc.Repo.DeleteRadreply(id)
	if err != nil {
		logrus.Error(err)
	}

	return err
}

func (uc *usecase) Delete(username string) error {
	err := uc.Repo.Delete(username)
	if err != nil {
		logrus.Error(err)
	}

	return err
}
