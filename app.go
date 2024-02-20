package main

import (
	"io"
	"log"
	"net"
	"os"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/novriyantoAli/product-service/initializers"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	_productHandlerGRPC "github.com/novriyantoAli/product-service/product/delivery/grpc"
	_productHandler "github.com/novriyantoAli/product-service/product/delivery/http"
	_productRepository "github.com/novriyantoAli/product-service/product/repository/mysql"
	_productUsecase "github.com/novriyantoAli/product-service/product/usecase"

	_radcheckHandlerGRPC "github.com/novriyantoAli/product-service/m_radcheck/delivery/grpc"
	_radcheckRepository "github.com/novriyantoAli/product-service/m_radcheck/repository/mysql"
	_radcheckUsecase "github.com/novriyantoAli/product-service/m_radcheck/usecase"

	_profileHandler "github.com/novriyantoAli/product-service/m_profile/delivery/http"
	_profileRepository "github.com/novriyantoAli/product-service/m_profile/repository/mysql"
	_profileUsecase "github.com/novriyantoAli/product-service/m_profile/usecase"
)

type customValidator struct {
	validator *validator.Validate
}

func (cv *customValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	logrus.SetReportCaller(true)

	initializers.LoadConfig()

	initializers.ConnectDB()

	initializers.Migrations()
}

func main() {
	// initialize
	f, err := os.OpenFile(viper.GetString(`log`), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	defer f.Close()

	wrt := io.MultiWriter(os.Stdout, f)

	logrus.SetOutput(wrt)

	srv := grpc.NewServer()
	e := echo.New()

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"OPTIONS", "GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "X-CSRF-Token", "application/json"},
		Debug:          true,
	})

	e.Validator = &customValidator{validator: validator.New()}

	e.Use(echo.WrapMiddleware(corsMiddleware.Handler))

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	productRepo := _productRepository.NewMysqlClient(initializers.GORM)
	productUCase := _productUsecase.NewUsecase(productRepo)
	_productHandlerGRPC.NewHandler(srv, productUCase)
	_productHandler.NewHandler(e, productUCase)

	radcheckRepo := _radcheckRepository.NewMysqlClient(initializers.GORM)
	radcheckUCase := _radcheckUsecase.NewUsecase(radcheckRepo)
	_radcheckHandlerGRPC.NewHandler(srv, radcheckUCase)

	profileRepo := _profileRepository.NewMysqlClient(initializers.GORM)
	profileUCase := _profileUsecase.NewUsecase(profileRepo)
	_profileHandler.NewHandler(e, profileUCase)

	go func() {
		l, err := net.Listen("tcp", viper.GetString(`server.grpc`))
		if err != nil {
			logrus.Fatalln(err)
		}
		logrus.Infof("start serve on port %s", viper.GetString(`server.grpc`))

		logrus.Fatal(srv.Serve(l))
	}()

	logrus.Fatal(e.Start(viper.GetString(`server.address`)))
}
