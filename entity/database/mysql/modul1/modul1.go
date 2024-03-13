package repository

import (
	repository "simple-fasthttp/entity/database/mysql"
	models "simple-fasthttp/entity/model"
	fe "simple-fasthttp/framework/error"
	"simple-fasthttp/framework/infra"

	"github.com/valyala/fasthttp"
	"gorm.io/gorm"
)

const table = "testing"

type Repo struct {
	Dbconn *gorm.DB
}

func NewRepository(i *infra.Infra) repository.Repository {
	return &Repo{i.Database.MySQL}
}
func (r Repo) GetData(ctx *fasthttp.RequestCtx, param *models.Request) ([]*models.Response, *fe.Error) {
	var res []*models.Response

	query := r.Dbconn.Table(table)
	if param.Id != 0 {
		query = query.Where("id = ?", param.Id)
	}
	err := query.Find(&res).Error
	if err != nil {
		return nil, fe.NewError(500, err)
	}
	return res, nil
}

func (r Repo) CreateData(ctx *fasthttp.RequestCtx, param *models.Request) *fe.Error {
	err := r.Dbconn.Omit("updated_at").Create(&param).Error
	if err != nil {
		return fe.NewError(500, err)
	}
	return nil
}

func (r Repo) UpdateData(ctx *fasthttp.RequestCtx, param *models.Request) *fe.Error {
	err := r.Dbconn.Omit("created_at").Updates(&param).Error
	if err != nil {
		return fe.NewError(500, err)
	}
	return nil
}

func (r Repo) DeleteData(ctx *fasthttp.RequestCtx, param *models.Request) *fe.Error {
	query := r.Dbconn.Delete(&param)
	err := query.Error
	if err != nil {
		return fe.NewError(500, err)
	}
	return nil
}
