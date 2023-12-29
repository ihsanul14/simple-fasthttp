package router

import (
	modul1_usecase "simple-fasthttp/app/modul1"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Router struct {
	Modul1Handler modul1_usecase.IHandler
}

func (r *Router) NewRouter() *fiber.App {
	router := fiber.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "*",
		AllowHeaders:     "Origin, X-Requested-With, Content-Type, Accept, Authorization, Access-Control-Allow-Headers, Accept-Encoding, X-CSRF-Token",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true,
	}))

	r.Modul1Handler.Routes(router)
	return router
}
