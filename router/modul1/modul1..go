package modul1

import (
	handler "simple-fasthttp/handler/modul1"

	"github.com/gofiber/fiber/v2"
)

func Router(router *fiber.App) {
	router.Get("/", handler.Get)
	router.Post("api/data", handler.DataHandler)
	router.Post("api/data/add", handler.CreateHandler)
	router.Put("api/data/update", handler.UpdateHandler)
	router.Delete("api/data/delete", handler.DeleteHandler)
}
