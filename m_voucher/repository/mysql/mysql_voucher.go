package mysql

import (
	"github.com/novriyantoAli/product-service/domain"
	"gorm.io/gorm"
)

type mysql struct {
	DB *gorm.DB
}

func NewMysqlClient(db *gorm.DB) domain.VoucherRepository {
	return &mysql{DB: db}
}

func (m *mysql) Find(param *domain.Voucher) (res []domain.Voucher, err error) {
	err = m.DB.Find(&res, param).Error
	return
}

func (m *mysql) FirstBatch(param string) (res domain.Voucher, err error) {
	err = m.DB.First(&res, "batch = ?", param).Error
	return
}

func (m *mysql) FirstUsername(param string) (res domain.Voucher, err error) {
	err = m.DB.First(&res, "username = ?", param).Error
	return
}

func (m *mysql) Create(param []domain.Voucher, param2 []domain.Radcheck) error {
	tx := m.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	// save voucher
	if err := tx.Create(param).Error; err != nil {
		tx.Rollback()
		return err
	}

	// save account
	if err := tx.Create(param2).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
