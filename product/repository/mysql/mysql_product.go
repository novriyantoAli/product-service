package mysql

import (
	"github.com/novriyantoAli/product-service/domain"
	"gorm.io/gorm"
)

type mysqlClient struct {
	DB *gorm.DB
}

func NewMysqlClient(db *gorm.DB) domain.ProductRepository {
	return &mysqlClient{DB: db}
}

func (m *mysqlClient) Find(product *domain.Product) (res []domain.Product, err error) {
	err = m.DB.Find(&res, product).Error
	return
}

func (m *mysqlClient) Save(product *domain.Product) error {
	return m.DB.Save(product).Error
}

func (m *mysqlClient) Create(product *domain.Product) error {
	return m.DB.Create(product).Error
}

func (m *mysqlClient) Delete(id uint) error {
	return m.DB.Delete(&domain.Product{}, id).Error
}
