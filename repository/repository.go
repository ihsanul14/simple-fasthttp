package repository

import (
	models "simple-fasthttp/models/modul1"

	"github.com/valyala/fasthttp"
)

type Repository interface {
	GetData(ctx *fasthttp.RequestCtx, request *models.Request) (res *models.ResponseAll, err error)
	CreateData(ctx *fasthttp.RequestCtx, request *models.Request) error
	UpdateData(ctx *fasthttp.RequestCtx, request *models.Request) error
	DeleteData(ctx *fasthttp.RequestCtx, request *models.Request) error
}
