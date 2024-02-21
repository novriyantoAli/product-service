package mysql

import (
	"github.com/novriyantoAli/product-service/domain"
	"gorm.io/gorm"
)

type mysql struct {
	DB *gorm.DB
}

func NewMysqlClient(db *gorm.DB) domain.RadacctRepository {
	return &mysql{DB: db}
}

func (m *mysql) FirstUsername(param string) (res domain.Radacct, err error) {
	err = m.DB.First(&res, "username = ?", param).Error
	return
}
