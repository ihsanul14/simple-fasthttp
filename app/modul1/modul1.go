package modul1

import (
	"simple-fasthttp/app/middleware"
	fe "simple-fasthttp/framework/error"
	"simple-fasthttp/framework/infra"
	"simple-fasthttp/framework/logger"
	"simple-fasthttp/framework/validator"
	usecase "simple-fasthttp/usecase/modul1"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Validator validator.IValidator
	Logger    logger.ILogger
}

type IHandler interface {
	GetDataHandler(*fiber.Ctx) error
	CreateHandler(*fiber.Ctx) error
	UpdateHandler(*fiber.Ctx) error
	DeleteHandler(*fiber.Ctx) error
	ErrorHandler(*fiber.Ctx, *fe.Error) error
}

func NewApplication(router *fiber.App, r *infra.Infra) {
	r.Validator.AddRules(getRule, usecase.Request{})
	r.Validator.AddRules(createRule, usecase.CreateRequest{})
	r.Validator.AddRules(updateRule, usecase.UpdateRequest{})
	r.Validator.AddRules(deleteRule, usecase.DeleteRequest{})

	u := &Handler{
		Validator: r.Validator,
		Logger:    r.Logger,
	}

	router.Use(middleware.UseCaseMiddleware(r))

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
		IUsecase := c.Locals("IUsecase").(usecase.IUsecase)
		result, err := IUsecase.GetData(ctx, param)
		if err == nil {
			result.Code = 200
			result.Message = "success retrieve data"
			return c.Status(200).JSON(result)
		} else {
			return u.ErrorHandler(c, err)
		}
	} else {
		return c.Status(400).JSON(usecase.ResponseAll{Code: 400, Message: err.Error()})
	}
}

func (u *Handler) CreateHandler(c *fiber.Ctx) error {
	param := new(usecase.CreateRequest)
	ctx := c.Context()
	if err := c.BodyParser(param); err != nil {
		return c.Status(400).JSON(usecase.ResponseAll{Code: 400, Message: err.Error()})
	}
	err := u.Validator.Check(param)
	if err == nil {
		IUsecase := c.Locals("IUsecase").(usecase.IUsecase)
		err := IUsecase.CreateData(ctx, param)
		if err == nil {
			return c.Status(200).JSON(usecase.ResponseAll{
				Code:    200,
				Message: "success create data",
			})
		} else {
			return u.ErrorHandler(c, err)
		}
	} else {
		return c.Status(400).JSON(usecase.ResponseAll{Code: 400, Message: err.Error()})
	}
}

func (u *Handler) UpdateHandler(c *fiber.Ctx) error {
	param := new(usecase.UpdateRequest)
	ctx := c.Context()
	if err := c.BodyParser(param); err != nil {
		return c.Status(400).JSON(usecase.ResponseAll{Code: 400, Message: err.Error()})
	}
	err := u.Validator.Check(param)
	if err == nil {
		IUsecase := c.Locals("IUsecase").(usecase.IUsecase)
		err := IUsecase.UpdateData(ctx, param)
		if err == nil {
			return c.Status(200).JSON(
				usecase.ResponseAll{
					Code:    200,
					Message: "success update data",
				},
			)
		} else {
			return u.ErrorHandler(c, err)
		}
	} else {
		return c.Status(400).JSON(usecase.ResponseAll{Code: 400, Message: err.Error()})
	}
}

func (u *Handler) DeleteHandler(c *fiber.Ctx) error {
	param := new(usecase.DeleteRequest)
	ctx := c.Context()
	if err := c.BodyParser(param); err != nil {
		return c.Status(400).JSON(usecase.ResponseAll{Code: 400, Message: err.Error()})
	}
	err := u.Validator.Check(param)
	if err == nil {
		IUsecase := c.Locals("IUsecase").(usecase.IUsecase)
		err := IUsecase.DeleteData(ctx, param)
		if err == nil {
			return c.Status(200).JSON(usecase.ResponseAll{
				Code:    200,
				Message: "success create data",
			})
		} else {
			return u.ErrorHandler(c, err)
		}
	} else {
		return c.Status(400).JSON(usecase.ResponseAll{Code: 400, Message: err.Error()})
	}
}

func (u *Handler) ErrorHandler(c *fiber.Ctx, err *fe.Error) error {
	u.Logger.Error(err.Message)

	var result = usecase.ResponseAll{
		Code:    err.Code,
		Message: err.Message,
	}

	return c.Status(err.Code).JSON(result)
}
