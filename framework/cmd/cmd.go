package cmd

import (
	"os"
	"simple-fasthttp/framework/infra"
	fr "simple-fasthttp/framework/router"
)

func Run() {
	infra := infra.Setup()
	router := fr.NewRouter(infra)
	if err := router.Listen(":" + os.Getenv("PORT")); err != nil {
		infra.Logger.Fatal(err.Error())
	}
}
