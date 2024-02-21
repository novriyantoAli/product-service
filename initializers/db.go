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
	GORM.AutoMigrate(&domain.Product{}, &domain.Radcheck{}, &domain.Voucher{})
}

func Trigger() {
	GORM.Exec(`
	CREATE TRIGGER irtgr AFTER INSERT ON radacct FOR EACH ROW

	BEGIN

	SET @expiration = (SELECT COUNT(*) FROM radcheck WHERE username = New.username AND attribute = 'Expiration');

	IF (@expiration = 0) THEN
		SET @validity_value = (SELECT products.val_val FROM vouchers INNER JOIN products ON products.id = vouchers.id_product WHERE vouchers.username = New.username ORDER BY vouchers.id DESC LIMIT 1);
		SET @validity_unit = (SELECT products.val_unit FROM vouchers INNER JOIN products ON products.id = vouchers.id_product WHERE vouchers.username = New.username ORDER BY vouchers.id DESC LIMIT 1);

		IF (@validity_unit = 'HOUR') THEN
			INSERT INTO radcheck(username, attribute, op, value) VALUES(New.username, "Expiration", ":=", DATE_FORMAT((NOW() + INTERVAL @validity_value HOUR), "%d %b %Y %H:%I:%S"));

		ELSEIF (@validity_unit = 'DAY') THEN
			INSERT INTO radcheck(username, attribute, op, value) VALUES(New.username, "Expiration", ":=", DATE_FORMAT((NOW() + INTERVAL @validity_value DAY), "%d %b %Y %H:%I:%S"));

		ELSEIF (@validity_unit = 'MONTH') THEN
			INSERT INTO radcheck(username, attribute, op, value) VALUES(New.username, "Expiration", ":=", DATE_FORMAT((NOW() + INTERVAL @validity_value MONTH), "%d %b %Y %H:%I:%S"));

		ELSEIF (@validity_unit = 'YEAR') THEN
			INSERT INTO radcheck(username, attribute, op, value) VALUES(New.username, "Expiration", ":=", DATE_FORMAT((NOW() + INTERVAL @validity_unit YEAR), "%d %b %Y %H:%I:%S"));

		END IF;

	END IF;
	END;
	`)
}
