package domain

import "github.com/novriyantoAli/product-service/model"

type Radcheck struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Username  string `json:"username"`
	Attribute string `json:"attribute"`
	OP        string `json:"op"`
	Value     string `json:"value"`
}

func (Radcheck) TableName() string {
	return "radcheck"
}

type RadcheckRepository interface {
	CreateBatch(params []Radcheck) (err error)
	Save(param Radcheck) (err error)
}

type RadcheckUsecase interface {
	CreateBatch(params *model.RadcheckList) error
}
