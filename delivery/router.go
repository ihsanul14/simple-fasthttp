package delivery

import (
	mod1_router "simple-fasthttp/delivery/modul1"
	"simple-fasthttp/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

//InitRouter is used for initiate router
func InitRouter(usecase usecase.Usecase) *fiber.App {
	router := fiber.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "*",
		AllowHeaders:     "Origin, X-Requested-With, Content-Type, Accept, Authorization, Access-Control-Allow-Headers, Accept-Encoding, X-CSRF-Token",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true,
	}))

	mod1_router.Router(router, usecase)
	return router
}
