package modul1

import (
	"fmt"

	usecase "simple-fasthttp/usecase/modul1"

	"github.com/gofiber/fiber/v2"
)

const modul = "modul1"

type Handler struct {
	Usecase usecase.IUsecase
}

type IHandler interface {
	Routes(*fiber.App)
	GetDataHandler(*fiber.Ctx) error
	CreateHandler(*fiber.Ctx) error
	UpdateHandler(*fiber.Ctx) error
	DeleteHandler(*fiber.Ctx) error
}

func NewApplication(u usecase.IUsecase) IHandler {
	return &Handler{Usecase: u}
}

func (u *Handler) Routes(router *fiber.App) {
	router.Post("api/data", u.GetDataHandler)
	router.Post("api/data/add", u.CreateHandler)
	router.Put("api/data/update", u.UpdateHandler)
	router.Delete("api/data/delete", u.DeleteHandler)
}

func (u *Handler) GetDataHandler(c *fiber.Ctx) error {
	param := new(usecase.Request)
	var err error

	ctx := c.Context()
	if err := c.BodyParser(param); err != nil {
		return err
	}
	if param.Id >= 0 {
		result, err := u.Usecase.GetData(ctx, param)
		if err == nil {
			result.Code = 200
			result.Message = "success retrieve data"
			c.JSON(result)
		} else {
			result.Code = 500
			result.Message = err.Error()
			c.JSON(result)
		}
	} else {
		c.JSON(usecase.ResponseAll{Code: 400, Message: fmt.Sprintf("kode %s must be a string", modul)})
	}
	return err
}

func (u *Handler) CreateHandler(c *fiber.Ctx) error {
	param := new(usecase.Request)
	var err error

	ctx := c.Context()
	if err := c.BodyParser(param); err != nil {
		return err
	}
	if param.Id >= 0 {
		result, err := u.Usecase.GetData(ctx, param)
		if err == nil {
			result.Code = 200
			result.Message = "success retrieve data"
			c.JSON(result)
		} else {
			result.Code = 500
			result.Message = err.Error()
			c.JSON(result)
		}
	} else {
		c.JSON(usecase.ResponseAll{Code: 400, Message: fmt.Sprintf("kode %s must be a string", modul)})
	}
	return err
}

func (u *Handler) UpdateHandler(c *fiber.Ctx) error {
	param := new(usecase.Request)
	var err error

	ctx := c.Context()
	if err := c.BodyParser(param); err != nil {
		return err
	}
	if param.Id >= 0 {
		result, err := u.Usecase.GetData(ctx, param)
		if err == nil {
			result.Code = 200
			result.Message = "success retrieve data"
			c.JSON(result)
		} else {
			result.Code = 500
			result.Message = err.Error()
			c.JSON(result)
		}
	} else {
		c.JSON(usecase.ResponseAll{Code: 400, Message: fmt.Sprintf("kode %s must be a string", modul)})
	}
	return err
}

func (u *Handler) DeleteHandler(c *fiber.Ctx) error {
	param := new(usecase.Request)
	var err error

	ctx := c.Context()
	if err := c.BodyParser(param); err != nil {
		return err
	}
	if param.Id >= 0 {
		result, err := u.Usecase.GetData(ctx, param)
		if err == nil {
			result.Code = 200
			result.Message = "success retrieve data"
			c.JSON(result)
		} else {
			result.Code = 500
			result.Message = err.Error()
			c.JSON(result)
		}
	} else {
		c.JSON(usecase.ResponseAll{Code: 400, Message: fmt.Sprintf("kode %s must be a string", modul)})
	}
	return err
}
