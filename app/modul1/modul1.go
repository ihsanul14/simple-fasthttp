package modul1

import (
	fe "simple-fasthttp/framework/error"
	"simple-fasthttp/framework/infra"
	"simple-fasthttp/framework/logger"
	"simple-fasthttp/framework/validator"
	usecase "simple-fasthttp/usecase/modul1"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Usecase   usecase.IUsecase
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
	r.Validator.AddRules(createRule, usecase.CreateRequest{})
	r.Validator.AddRules(updateRule, usecase.UpdateRequest{})
	r.Validator.AddRules(deleteRule, usecase.DeleteRequest{})

	u := &Handler{
		Usecase:   usecase.NewUsecase(r),
		Validator: r.Validator,
		Logger:    r.Logger,
	}

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
	result, err := u.Usecase.GetData(ctx, param)
	if err == nil {
		result.Code = 200
		result.Message = "success retrieve data"
		return c.Status(200).JSON(result)
	} else {
		return u.ErrorHandler(c, err)
	}
}

func (u *Handler) CreateHandler(c *fiber.Ctx) error {
	param := new(usecase.CreateRequest)
	ctx := c.Context()
	if err := c.BodyParser(param); err != nil {
		return u.ErrorHandler(c, fe.NewError(400, err))
	}
	err := u.Validator.Check(param)
	if err == nil {
		err := u.Usecase.CreateData(ctx, param)
		if err == nil {
			return c.Status(200).JSON(usecase.ResponseAll{
				Code:    200,
				Message: "success create data",
			})
		} else {
			return u.ErrorHandler(c, err)
		}
	} else {
		return u.ErrorHandler(c, fe.NewError(400, err))
	}
}

func (u *Handler) UpdateHandler(c *fiber.Ctx) error {
	param := new(usecase.UpdateRequest)
	ctx := c.Context()
	if err := c.BodyParser(param); err != nil {
		return u.ErrorHandler(c, fe.NewError(400, err))
	}
	err := u.Validator.Check(param)
	if err == nil {
		err := u.Usecase.UpdateData(ctx, param)
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
		return u.ErrorHandler(c, fe.NewError(400, err))
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
		err := u.Usecase.DeleteData(ctx, param)
		if err == nil {
			return c.Status(200).JSON(usecase.ResponseAll{
				Code:    200,
				Message: "success create data",
			})
		} else {
			return u.ErrorHandler(c, err)
		}
	} else {
		return u.ErrorHandler(c, fe.NewError(400, err))
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
