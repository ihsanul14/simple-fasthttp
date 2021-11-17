package repository

import (
	"fmt"
	models "simple-fasthttp/models/modul1"

	"github.com/valyala/fasthttp"
	"gorm.io/gorm"
)

const createQuery = "INSERT INTO %s (nama, nomor, created_at) VALUES (?,?,NOW())"
const updateQuery = "UPDATE %s SET nama = ?, nomor = ?, updated_at = NOW() WHERE id = ?"
const deleteQuery = "DELETE FROM %s WHERE id = ?"
const table = "testing"

type Repo struct {
	Dbconn *gorm.DB
}

type Repository interface {
	GetData(ctx *fasthttp.RequestCtx, request *models.Request) (res models.ResponseAll, err error)
	CreateData(ctx *fasthttp.RequestCtx, request *models.Request) error
	UpdateData(ctx *fasthttp.RequestCtx, request *models.Request) error
	DeleteData(ctx *fasthttp.RequestCtx, request *models.Request) error
}

func NewRepository(dbconn *gorm.DB) Repository {
	return &Repo{dbconn}
}
func (r Repo) GetData(ctx *fasthttp.RequestCtx, param *models.Request) (models.ResponseAll, error) {
	var (
		result []models.Response
		res    models.ResponseAll
		err    error
	)
	query := r.Dbconn.Table(table)
	if param.Id != 0 {
		query = query.Where("id = ?", param.Id)
	}
	err = query.Scan(&result).Error
	if err != nil {
		fmt.Println(err.Error())
		return res, err
	}
	res.Data = result
	return res, err
}

func (r Repo) CreateData(ctx *fasthttp.RequestCtx, param *models.Request) (err error) {
	err = r.Dbconn.Exec(fmt.Sprintf(createQuery, table), param.Nama, param.Nomor).Error
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return err
}

func (r Repo) UpdateData(ctx *fasthttp.RequestCtx, param *models.Request) (err error) {
	query := r.Dbconn.Exec(fmt.Sprintf(updateQuery, table), param.Nama, param.Nomor, param.Id)
	err = query.Error
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return err
}

func (r Repo) DeleteData(ctx *fasthttp.RequestCtx, param *models.Request) (err error) {
	query := r.Dbconn.Exec(fmt.Sprintf(deleteQuery, table), param.Id)
	err = query.Error
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return err
}
