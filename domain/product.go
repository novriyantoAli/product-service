package domain

import (
	"github.com/novriyantoAli/product-service/model"
)

type Product struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Name    string `json:"name" validate:"required"`
	ValUnit string `json:"val_unit" validate:"required" gorm:"default:HOUR"`
	ValVal  uint   `json:"val_val" validate:"required"`
	Price   uint   `json:"price" validate:"required"`
	Profile string `json:"profile" validate:"required"`
}

type ProductRepository interface {
	Find(product *Product) (res []Product, err error)
	Save(product *Product) error
	Create(product *Product) error
	Delete(id uint) error
}

type ProductUsecase interface {
	Find(param *model.Product) (res *model.ProductList, err error)
	FindProduct(param *Product) (res []Product, err error)
	Save(param *Product) error
	Create(param *Product) error
	Delete(param uint) error
}
