package mysql

import (
	"github.com/novriyantoAli/product-service/domain"
	"gorm.io/gorm"
)

type mysqlClient struct {
	DB *gorm.DB
}

func NewMysqlClient(db *gorm.DB) domain.RadcheckRepository {
	return &mysqlClient{DB: db}
}

func (m *mysqlClient) First(param string) (res domain.Radcheck, err error) {
	err = m.DB.First(&res, "username = ?", param).Error
	return
}

func (m *mysqlClient) CreateBatch(param []domain.Radcheck) (err error) {
	err = m.DB.Create(param).Error
	return
}

func (m *mysqlClient) Save(param domain.Radcheck) (err error) {
	err = m.DB.Save(param).Error
	return
}
