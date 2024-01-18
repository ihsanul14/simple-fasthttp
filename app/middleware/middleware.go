package middleware

import (
	"simple-fasthttp/framework/infra"
	uc "simple-fasthttp/usecase/modul1"

	"github.com/gofiber/fiber/v2"
)

const IUsecase = "IUsecase"

func UseCaseMiddleware(r *infra.Infra) fiber.Handler {
	return func(c *fiber.Ctx) error {
		usecase := uc.NewUsecase(r)
		c.Locals(IUsecase, usecase)
		return c.Next()
	}
}
