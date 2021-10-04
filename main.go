package main

import (
	"os"
	"simple-fasthttp/database"
	"simple-fasthttp/router"

	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	database.ConnectMySQL()
	router := router.InitRouter()
	router.Listen(":" + os.Getenv("PORT"))
}
