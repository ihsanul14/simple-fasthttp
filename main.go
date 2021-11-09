package main

import (
	"os"
	"simple-fasthttp/database"
	"simple-fasthttp/delivery"
	repo "simple-fasthttp/repository/modul1"
	use "simple-fasthttp/usecase/modul1"

	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	mysqlConn, err := database.ConnectMySQL()
	if err != nil {
		return
	}
	repository := repo.NewRepository(mysqlConn)
	usecase := use.NewUsecase(repository)
	router := delivery.InitRouter(usecase)

	router.Listen(":" + os.Getenv("PORT"))
}
