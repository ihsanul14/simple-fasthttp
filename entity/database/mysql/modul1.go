package repository

import (
	models "simple-fasthttp/entity/model"

	"github.com/valyala/fasthttp"
)

type Repository interface {
	GetData(*fasthttp.RequestCtx, *models.Request) ([]*models.Response, error)
	CreateData(*fasthttp.RequestCtx, *models.Request) error
	UpdateData(*fasthttp.RequestCtx, *models.Request) error
	DeleteData(*fasthttp.RequestCtx, *models.Request) error
}
