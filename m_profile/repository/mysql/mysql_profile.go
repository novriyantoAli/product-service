package mysql

import (
	"github.com/novriyantoAli/product-service/domain"
	"gorm.io/gorm"
)

type mysqlClient struct {
	DB *gorm.DB
}

func NewMysqlClient(db *gorm.DB) domain.ProfileRepository {
	return &mysqlClient{DB: db}
}

func (db *mysqlClient) Find(param *domain.Profile) (res []domain.Profile, err error) {
	err = db.DB.Find(&res, param).Error
	return
}

func (db *mysqlClient) FindRadcheck(param *domain.Radgroupcheck) (res []domain.Radgroupcheck, err error) {
	err = db.DB.Find(&res, param).Error
	return
}

func (db *mysqlClient) FindRadreply(param *domain.Radgroupreply) (res []domain.Radgroupreply, err error) {
	err = db.DB.Find(&res, param).Error
	return
}
