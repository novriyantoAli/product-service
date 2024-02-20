package initializers

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"

	"github.com/novriyantoAli/product-service/domain"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *sql.DB

var GORM *gorm.DB

func ConnectDB() {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add(`parseTime`, "1")
	val.Add(`loc`, "Asia/Makassar")

	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())

	DB, err := sql.Open(`mysql`, dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	GORM, err = gorm.Open(mysql.New(mysql.Config{
		Conn: DB,
	}), &gorm.Config{
		TranslateError: true,
	})

	if err != nil {
		logrus.Fatalln(err)
	}
}

// func ExecTrigger(command string) {

// }

func Migrations() {
	GORM.AutoMigrate(&domain.Product{}, &domain.Radcheck{})
}
