package infra

import (
	"simple-fasthttp/framework/database"
	fd "simple-fasthttp/framework/database"
	fl "simple-fasthttp/framework/logger"
	fv "simple-fasthttp/framework/validator"

	"github.com/subosito/gotenv"
)

type Infra struct {
	Logger    fl.ILogger
	Validator fv.IValidator
	Database  *fd.Database
}

func Setup() *Infra {
	logger := fl.NewLogger()
	if err := gotenv.Load(); err != nil {
		logger.Fatal(err.Error())
	}
	db, err := database.ConnectMySQL()
	if err != nil {
		logger.Fatal(err.Error())
	}
	validator := fv.NewValidator()
	return &Infra{
		Logger:    logger,
		Validator: validator,
		Database:  db,
	}
}
