package repository

import (
	models "simple-fasthttp/entity/model"

	fe "simple-fasthttp/framework/error"

	"github.com/valyala/fasthttp"
)

type Repository interface {
	GetData(*fasthttp.RequestCtx, *models.Request) ([]*models.Response, *fe.Error)
	CreateData(*fasthttp.RequestCtx, *models.Request) *fe.Error
	UpdateData(*fasthttp.RequestCtx, *models.Request) *fe.Error
	DeleteData(*fasthttp.RequestCtx, *models.Request) *fe.Error
}
