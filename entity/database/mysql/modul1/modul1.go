package repository

import (
	"fmt"
	repository "simple-fasthttp/entity/database/mysql"
	models "simple-fasthttp/entity/model"

	"github.com/valyala/fasthttp"
	"gorm.io/gorm"
)

type Repo struct {
	Dbconn *gorm.DB
}

const createQuery = "INSERT INTO %s (nama, nomor, created_at) VALUES (?,?,NOW())"
const updateQuery = "UPDATE %s SET nama = ?, nomor = ?, updated_at = NOW() WHERE id = ?"
const deleteQuery = "DELETE FROM %s WHERE id = ?"
const table = "testing"

func NewRepository(dbconn *gorm.DB) repository.Repository {
	return &Repo{dbconn}
}
func (r Repo) GetData(ctx *fasthttp.RequestCtx, param *models.Request) ([]*models.Response, error) {
	var (
		res []*models.Response
		err error
	)
	query := r.Dbconn.Table(table)
	if param.Id != 0 {
		query = query.Where("id = ?", param.Id)
	}
	err = query.Scan(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r Repo) CreateData(ctx *fasthttp.RequestCtx, param *models.Request) (err error) {
	err = r.Dbconn.Exec(fmt.Sprintf(createQuery, table), param.Nama, param.Nomor).Error
	if err != nil {
		return err
	}
	return err
}

func (r Repo) UpdateData(ctx *fasthttp.RequestCtx, param *models.Request) (err error) {
	query := r.Dbconn.Exec(fmt.Sprintf(updateQuery, table), param.Nama, param.Nomor, param.Id)
	err = query.Error
	if err != nil {
		return err
	}
	return err
}

func (r Repo) DeleteData(ctx *fasthttp.RequestCtx, param *models.Request) (err error) {
	query := r.Dbconn.Exec(fmt.Sprintf(deleteQuery, table), param.Id)
	err = query.Error
	if err != nil {
		return err
	}
	return err
}
