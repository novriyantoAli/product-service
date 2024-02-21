package usecase

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/m1/go-generate-password/generator"
	"github.com/novriyantoAli/product-service/domain"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type usecase struct {
	Repo      domain.VoucherRepository
	RProduct  domain.ProductRepository
	RRadcheck domain.RadcheckRepository
	RRadacct  domain.RadacctRepository
}

func NewUsecase(repo domain.VoucherRepository, rproduct domain.ProductRepository, rradcheck domain.RadcheckRepository, rradacct domain.RadacctRepository) domain.VoucherUsecase {
	return &usecase{Repo: repo, RProduct: rproduct, RRadcheck: rradcheck, RRadacct: rradacct}
}

func (uc *usecase) Find(param *domain.Voucher) (res []domain.Voucher, err error) {
	res, err = uc.Repo.Find(param)
	if err != nil {
		logrus.Error(err)
	}

	return
}
func (uc *usecase) Create(param *domain.CreateVoucherRequest) error {
	// first search id product if valid or not
	products, err := uc.RProduct.Find(&domain.Product{ID: param.IDProduct})
	if err != nil {
		logrus.Error(err)
		return err
	}

	if len(products) != 1 {
		logrus.Error("anomali behaviour product result")
		return errors.New("anomali behaviour product result")
	}

	// cari kode batch yang belum terpakai
	var timestamp string
	for {
		timestamp = strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
		_, err = uc.Repo.FirstBatch(timestamp)
		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			// found uiq
			break
		}
	}

	// konfigurasi tipe
	configUsername := generator.Config{
		Length:                     param.LengthUsername,
		IncludeSymbols:             false,
		IncludeNumbers:             param.Number,
		IncludeLowercaseLetters:    param.LowerCase,
		IncludeUppercaseLetters:    param.UpperCase,
		ExcludeSimilarCharacters:   true,
		ExcludeAmbiguousCharacters: true,
	}

	gUsername, err := generator.New(&configUsername)
	if err != nil {
		logrus.Error(err)
		return err
	}

	// sekarang mencari username yang valid
	usernames := make([]string, 0)
	for {
		unm, err := gUsername.Generate()
		if err != nil {
			logrus.Error(err)
			return err
		}
		// cek username pada voucher
		_, err = uc.Repo.FirstUsername(*unm)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			logrus.Error(err)
			return err
		}

		// cek username pada radcheck
		_, err = uc.RRadcheck.First(*unm)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			logrus.Error(err)
			return err
		}

		// cek username pada radacct
		_, err = uc.RRadacct.FirstUsername(*unm)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			logrus.Error(err)
			return err
		}
		usernames = append(usernames, *unm)
		if param.Sequentially {
			break
		}
		if len(usernames) == int(param.Amount) {
			break
		}
	}

	configPassword := generator.Config{
		Length:                     param.LengthPassword,
		IncludeSymbols:             false,
		IncludeNumbers:             param.Number,
		IncludeLowercaseLetters:    param.LowerCase,
		IncludeUppercaseLetters:    param.UpperCase,
		ExcludeSimilarCharacters:   true,
		ExcludeAmbiguousCharacters: true,
	}
	gPassword, err := generator.New(&configPassword)
	if err != nil {
		logrus.Error(err)
		return err
	}

	radchecks := make([]domain.Radcheck, 0)
	vouchers := make([]domain.Voucher, 0)
	if param.Sequentially {
		for i := 0; i < int(param.Amount); i++ {
			// membuat username dan password untuk radcheck
			uspas := domain.Radcheck{
				Username:  (usernames[0] + fmt.Sprint((i + 1))),
				Attribute: viper.GetString("setting.dictionary.password"),
				OP:        ":=",
				Value:     (usernames[0] + fmt.Sprint((i + 1))),
			}
			if !param.UPSame {
				pss, err := gPassword.Generate()
				if err != nil {
					logrus.Error(err)
					return err
				}
				uspas.Value = *pss
			}
			radchecks = append(radchecks, uspas)

			// membuat profil untuk radcheck
			profile := domain.Radcheck{
				Username:  (usernames[0] + fmt.Sprint((i + 1))),
				Attribute: viper.GetString("setting.dictionary.profile"),
				OP:        ":=",
				Value:     products[0].Profile,
			}
			radchecks = append(radchecks, profile)

			if param.Active {
				myDate := time.Now().Add(time.Hour * 1)
				if products[0].ValUnit == "MONTH" {
					myDate = time.Now().AddDate(0, int(products[0].ValVal), 0)
				} else if products[0].ValUnit == "DAY" {
					myDate = time.Now().AddDate(0, 0, int(products[0].ValVal))
				} else if products[0].ValUnit == "HOUR" {
					myDate = time.Now().Add(time.Hour * time.Duration(int(products[0].ValVal)))
				}
				// langsung aktifkan
				exp := domain.Radcheck{
					Username:  (usernames[0] + fmt.Sprint((i + 1))),
					Attribute: viper.GetString("setting.dictionary.expire"),
					OP:        ":=",
					Value:     myDate.Format(viper.GetString("setting.timeLayoutFormat")),
				}
				radchecks = append(radchecks, exp)
			}
			voucher := domain.Voucher{
				IDProduct: products[0].ID,
				Username:  (usernames[0] + fmt.Sprint((i + 1))),
				Batch:     timestamp,
			}
			vouchers = append(vouchers, voucher)
		}
	} else {
		for _, el := range usernames {
			// membuat username dan password untuk radcheck
			uspas := domain.Radcheck{
				Username:  el,
				Attribute: viper.GetString("setting.dictionary.password"),
				OP:        ":=",
				Value:     el,
			}
			if !param.UPSame {
				pss, err := gPassword.Generate()
				if err != nil {
					logrus.Error(err)
					return err
				}
				uspas.Value = *pss
			}
			radchecks = append(radchecks, uspas)

			// membuat profil untuk radcheck
			profile := domain.Radcheck{
				Username:  el,
				Attribute: viper.GetString("setting.dictionary.profile"),
				OP:        ":=",
				Value:     products[0].Profile,
			}
			radchecks = append(radchecks, profile)

			if param.Active {
				myDate := time.Now().Add(time.Hour * 1)
				if products[0].ValUnit == "MONTH" {
					myDate = time.Now().AddDate(0, int(products[0].ValVal), 0)
				} else if products[0].ValUnit == "DAY" {
					myDate = time.Now().AddDate(0, 0, int(products[0].ValVal))
				} else if products[0].ValUnit == "HOUR" {
					myDate = time.Now().Add(time.Hour * time.Duration(int(products[0].ValVal)))
				}
				// langsung aktifkan
				exp := domain.Radcheck{
					Username:  el,
					Attribute: viper.GetString("setting.dictionary.expire"),
					OP:        ":=",
					Value:     myDate.Format(viper.GetString("setting.timeLayoutFormat")),
				}
				radchecks = append(radchecks, exp)
			}
			voucher := domain.Voucher{
				IDProduct: products[0].ID,
				Username:  el,
				Batch:     timestamp,
			}
			vouchers = append(vouchers, voucher)
		}
	}

	err = uc.Repo.Create(vouchers, radchecks)
	if err != nil {
		logrus.Error(err)
	}

	return err
}
