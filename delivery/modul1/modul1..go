package modul1

import (
	"fmt"
	mdl "simple-fasthttp/models/modul1"
	usecase "simple-fasthttp/usecase/modul1"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Usecase usecase.UsecaseModul
}

const modul = "modul1"

func Router(router *fiber.App) {
	var u Handler
	// router.Get("/", handler.Get)
	router.Post("api/data", u.GetDataHandler)
	// router.Post("api/data/add", handler.CreateHandler)
	// router.Put("api/data/update", handler.UpdateHandler)
	// router.Delete("api/data/delete", handler.DeleteHandler)
}

func (u Handler) GetDataHandler(c *fiber.Ctx) error {
	param := new(mdl.Request)
	var err error

	ctx := c.Context()
	if err := c.BodyParser(param); err != nil {
		return err
	}
	if param.Id >= 0 {
		result, err := u.Usecase.ShowData(ctx, param)
		if err == nil {
			result.Code = 200
			result.Message = "success retrieve data"
			c.JSON(result)
		} else {
			result.Code = 500
			result.Message = "internal server error"
			c.JSON(result)
		}
	} else {
		c.JSON(mdl.ResponseAll{Code: 400, Message: fmt.Sprintf("kode %s must be a string", modul)})
	}
	return err
}
