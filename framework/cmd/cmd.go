package cmd

import (
	"log"
	"os"
	modul1Handler "simple-fasthttp/app/modul1"
	repo "simple-fasthttp/entity/database/mysql/modul1"
	"simple-fasthttp/framework/database"
	fr "simple-fasthttp/framework/router"
	u "simple-fasthttp/usecase/modul1"

	"github.com/gofiber/fiber/v2"
	"github.com/subosito/gotenv"
)

type Cmd struct {
	Router *fiber.App
}

func Run() {
	if err := gotenv.Load(); err != nil {
		log.Fatal(err)
	}
	mysqlConn, err := database.ConnectMySQL()
	if err != nil {
		log.Fatal(err)
	}
	repository := repo.NewRepository(mysqlConn)
	usecase := u.NewUsecase(repository)
	handler := modul1Handler.NewApplication(usecase)
	router := fr.Router{
		Modul1Handler: handler,
	}
	cmd := &Cmd{
		Router: router.NewRouter(),
	}
	cmd.Router.Listen(":" + os.Getenv("PORT"))
}
