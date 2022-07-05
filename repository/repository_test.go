package repository

import (
	"log"
	"regexp"
	"testing"

	model "simple-fasthttp/models/modul1"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	gb, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("error connect to database")
	}
	return gb, mock, err
}

func TestGetData(t *testing.T) {
	gb, mock, _ := InitDB()
	r := NewRepository(gb)
	var ctx *fasthttp.RequestCtx
	req := &model.Request{}
	t.Run("Test Get Data - Pass", func(t *testing.T) {
		req.Id = 1
		rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
		query := "SELECT * FROM `testing` WHERE id = ?"
		mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)
		anArticle, err := r.GetData(ctx, req)
		assert.NoError(t, err)
		assert.NotNil(t, anArticle)
	})
	t.Run("Test Get Data - Failed", func(t *testing.T) {
		req.Id = 0
		rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
		query := "SELECT * FROM `testing` WHERE id = ?"
		mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)
		_, err := r.GetData(ctx, req)
		assert.NotNil(t, err)
	})
}

func TestCreateData(t *testing.T) {
	gb, mock, _ := InitDB()
	r := NewRepository(gb)
	var ctx *fasthttp.RequestCtx
	req := &model.Request{}
	t.Run("Test Insert Data - Pass", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO testing (nama, nomor, created_at)")).WithArgs(req.Nama, req.Nomor).WillReturnResult(sqlmock.NewResult(1, 1))
		err := r.CreateData(ctx, req)
		assert.Nil(t, err)
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expections: %s", err)
		}
		assert.Nil(t, err)
	})
	t.Run("Test Insert Data - Fail", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO testing (nama, nomor, created)")).WithArgs(req.Nama, req.Nomor).WillReturnResult(sqlmock.NewResult(1, 1))
		err := r.CreateData(ctx, req)
		assert.NotNil(t, err)
	})
}

func TestUpdateData(t *testing.T) {
	gb, mock, _ := InitDB()
	r := NewRepository(gb)
	var ctx *fasthttp.RequestCtx
	req := &model.Request{}
	req.Id = 1
	t.Run("Test Insert Data - Pass", func(t *testing.T) {
		mock.ExpectExec("UPDATE testing").WillReturnResult(sqlmock.NewResult(1, 1))
		err := r.UpdateData(ctx, req)
		assert.Nil(t, err)
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expections: %s", err)
		}
		assert.Nil(t, err)
	})
	t.Run("Test Insert Data - Fail", func(t *testing.T) {
		mock.ExpectExec("UPDATES testing").WillReturnResult(sqlmock.NewResult(1, 1))
		err := r.UpdateData(ctx, req)
		assert.NotNil(t, err)
	})
}

func TestDeleteData(t *testing.T) {
	gb, mock, _ := InitDB()
	r := NewRepository(gb)
	var ctx *fasthttp.RequestCtx
	req := &model.Request{}
	req.Id = 1
	t.Run("Test Delete Data - Pass", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta("DELETE")).WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))
		err := r.DeleteData(ctx, req)
		assert.Nil(t, err)
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expections: %s", err)
		}
		assert.Nil(t, err)
	})
	t.Run("Test Delete Data - Fail", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta("DELETES")).WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))
		err := r.DeleteData(ctx, req)
		assert.NotNil(t, err)
	})
}
