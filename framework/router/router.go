package router

import (
	modul1_handler "simple-fasthttp/app/modul1"

	"simple-fasthttp/framework/infra"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type IRouter interface {
	Routes() *fiber.App
	Listen(string) error
}

type Router struct {
	Infra  *infra.Infra
	Router *fiber.App
}

func NewRouter(infra *infra.Infra) IRouter {
	return &Router{Infra: infra}
}

func (fr *Router) Routes() *fiber.App {
	router := fiber.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "*",
		AllowHeaders:     "Origin, X-Requested-With, Content-Type, Accept, Authorization, Access-Control-Allow-Headers, Accept-Encoding, X-CSRF-Token",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true,
	}))

	modul1_handler.NewApplication(router, fr.Infra)
	return router
}

func (fr *Router) Listen(port string) error {
	return fr.Routes().Listen(port)
}
