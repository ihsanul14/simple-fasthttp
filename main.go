package main

import (
	"os"
	"simple-fasthttp/database"
	"simple-fasthttp/delivery"
	repo "simple-fasthttp/repository"
	use "simple-fasthttp/usecase"

	"github.com/gofiber/fiber/v2/middleware/logger"
	// "github.com/sirupsen/logrus"

	"github.com/subosito/gotenv"
)

// type StandardLogger struct {
// 	*logrus.Logger
// }

func main() {
	gotenv.Load()
	// var baseLogger = logrus.New()

	// var standardLogger = &StandardLogger{baseLogger}

	// standardLogger.Formatter = &logrus.JSONFormatter{}

	mysqlConn, err := database.ConnectMySQL()
	if err != nil {
		return
	}
	repository := repo.NewRepository(mysqlConn)
	usecase := use.NewUsecase(repository)
	router := delivery.InitRouter(usecase)
	router.Use(logger.New())

	router.Listen(":" + os.Getenv("PORT"))
}
