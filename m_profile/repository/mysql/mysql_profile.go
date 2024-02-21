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

func (db *mysqlClient) Save(param *domain.Profile) error {
	return db.DB.Create(param).Error
}

func (db *mysqlClient) SaveRadcheck(param *domain.Radgroupcheck) error {
	return db.DB.Save(param).Error
}

func (db *mysqlClient) SaveRadreply(param *domain.Radgroupreply) error {
	return db.DB.Save(param).Error
}

func (db *mysqlClient) DeleteRadcheck(id uint) error {
	return db.DB.Delete(&domain.Radgroupcheck{}, id).Error
}

func (db *mysqlClient) DeleteRadreply(id uint) error {
	return db.DB.Delete(&domain.Radgroupreply{}, id).Error
}

func (db *mysqlClient) Delete(username string) error {
	return db.DB.Where("username = ?", username).Delete(&domain.Profile{}).Error
}
