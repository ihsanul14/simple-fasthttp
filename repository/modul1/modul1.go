package repository

import (
	"context"
	"fmt"
	models "simple-fasthttp/models/modul1"
	"simple-fasthttp/repository"

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
	return &Repo{
		dbconn,
	}
}
func (r Repo) GetDataQuery(ctx context.Context, param *models.Request) (models.ResponseAll, error) {
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

func (r Repo) CreateDataQuery(ctx context.Context, param models.Request) (err error) {
	err = r.Dbconn.Exec(fmt.Sprintf(createQuery, table), param.Nama, param.Nomor).Error
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return err
}

func (r Repo) UpdateDataQuery(ctx context.Context, param models.Request) (err error) {
	query := r.Dbconn.Exec(fmt.Sprintf(updateQuery, table), param.Nama, param.Nomor, param.Id)
	err = query.Error
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return err
}

func (r Repo) DeleteDataQuery(ctx context.Context, param models.Request) (err error) {
	query := r.Dbconn.Exec(fmt.Sprintf(deleteQuery, table), param.Id)
	err = query.Error
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return err
}
