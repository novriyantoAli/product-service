package domain

import "time"

type Voucher struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	IDProduct uint      `json:"id_product"`
	Username  string    `json:"username" gorm:"index"`
	Batch     string    `json:"batch"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateVoucherRequest struct {
	IDReseller     string `json:"id_reseller"`
	IDProduct      uint   `json:"id_product" validate:"required"`
	LengthUsername uint   `json:"l_uname" validate:"required"` // panjang username
	LengthPassword uint   `json:"l_psswd" validate:"required"` // panjang password
	Number         bool   `json:"number"`                      // angka
	LowerCase      bool   `json:"lower_case"`                  // huruf besar
	UpperCase      bool   `json:"upper_case"`                  // huruf kecil
	Amount         uint   `json:"amount" validate:"required"`
	UPSame         bool   `json:"up_same"`      // username dan password sama
	Active         bool   `json:"active"`       // fungsi mengijinkan pembuatan langsung aktiv
	Sequentially   bool   `json:"sequentially"` // berurutan
}

type VoucherRepository interface {
	Find(param *Voucher) (res []Voucher, err error)
	FirstBatch(param string) (res Voucher, err error)
	FirstUsername(param string) (res Voucher, err error)
	Create(param []Voucher, param1 []Radcheck) error
}

type VoucherUsecase interface {
	Find(param *Voucher) (res []Voucher, err error)
	Create(param *CreateVoucherRequest) error
}
