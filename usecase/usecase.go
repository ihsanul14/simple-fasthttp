package usecase

import (
	models "simple-fasthttp/models/modul1"

	"github.com/valyala/fasthttp"
)

type Usecase interface {
	ReadData(ctx *fasthttp.RequestCtx, request *models.Request) (res *models.ResponseAll, err error)
	AlterData(ctx *fasthttp.RequestCtx, request *models.Request) error
}
