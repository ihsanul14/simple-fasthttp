package modul1

import (
	"simple-fasthttp/framework/validator"
	usecase "simple-fasthttp/usecase/modul1"

	"github.com/gofiber/fiber/v2"
)

const modul = "modul1"

type Handler struct {
	Usecase   usecase.IUsecase
	Validator validator.IValidator
}

type IHandler interface {
	Routes(*fiber.App)
	GetDataHandler(*fiber.Ctx) error
	CreateHandler(*fiber.Ctx) error
	UpdateHandler(*fiber.Ctx) error
	DeleteHandler(*fiber.Ctx) error
}

func NewApplication(u usecase.IUsecase, v validator.IValidator) IHandler {
	v.AddRules(getRule, usecase.Request{})
	v.AddRules(createRule, usecase.CreateRequest{})
	v.AddRules(updateRule, usecase.UpdateRequest{})
	v.AddRules(deleteRule, usecase.DeleteRequest{})
	return &Handler{Usecase: u, Validator: v}
}

func (u *Handler) Routes(router *fiber.App) {
	router.Get("api/data", u.GetDataHandler)
	router.Post("api/data/add", u.CreateHandler)
	router.Put("api/data/update", u.UpdateHandler)
	router.Delete("api/data/delete", u.DeleteHandler)
}

func (u *Handler) GetDataHandler(c *fiber.Ctx) error {
	ctx := c.Context()

	param := &usecase.Request{
		Id: c.Query("id"),
	}

	err := u.Validator.Check(param)
	if err == nil {
		result, err := u.Usecase.GetData(ctx, param)
		if err == nil {
			result.Code = 200
			result.Message = "success retrieve data"
			return c.Status(200).JSON(result)
		} else {
			result.Code = 500
			result.Message = err.Error()
			return c.Status(500).JSON(result)
		}
	} else {
		return c.Status(400).JSON(usecase.ResponseAll{Code: 400, Message: err.Error()})
	}
}

func (u *Handler) CreateHandler(c *fiber.Ctx) error {
	param := new(usecase.Request)
	ctx := c.Context()
	if err := c.BodyParser(param); err != nil {
		return c.Status(400).JSON(usecase.ResponseAll{Code: 400, Message: err.Error()})
	}
	err := u.Validator.Check(param)
	if err == nil {
		result, err := u.Usecase.GetData(ctx, param)
		if err == nil {
			result.Code = 200
			result.Message = "success retrieve data"
			return c.Status(200).JSON(result)
		} else {
			result.Code = 500
			result.Message = err.Error()
			return c.Status(500).JSON(result)
		}
	} else {
		return c.Status(400).JSON(usecase.ResponseAll{Code: 400, Message: err.Error()})
	}
}

func (u *Handler) UpdateHandler(c *fiber.Ctx) error {
	param := new(usecase.Request)
	ctx := c.Context()
	if err := c.BodyParser(param); err != nil {
		return c.Status(400).JSON(usecase.ResponseAll{Code: 400, Message: err.Error()})
	}
	err := u.Validator.Check(param)
	if err == nil {
		result, err := u.Usecase.GetData(ctx, param)
		if err == nil {
			result.Code = 200
			result.Message = "success retrieve data"
			return c.Status(200).JSON(result)
		} else {
			result.Code = 500
			result.Message = err.Error()
			return c.Status(500).JSON(result)
		}
	} else {
		return c.Status(400).JSON(usecase.ResponseAll{Code: 400, Message: err.Error()})
	}
}

func (u *Handler) DeleteHandler(c *fiber.Ctx) error {
	param := new(usecase.Request)
	ctx := c.Context()
	if err := c.BodyParser(param); err != nil {
		return c.Status(400).JSON(usecase.ResponseAll{Code: 400, Message: err.Error()})
	}
	err := u.Validator.Check(param)
	if err == nil {
		result, err := u.Usecase.GetData(ctx, param)
		if err == nil {
			result.Code = 200
			result.Message = "success retrieve data"
			return c.Status(200).JSON(result)
		} else {
			result.Code = 500
			result.Message = err.Error()
			return c.Status(500).JSON(result)
		}
	} else {
		return c.Status(400).JSON(usecase.ResponseAll{Code: 400, Message: err.Error()})
	}
}
