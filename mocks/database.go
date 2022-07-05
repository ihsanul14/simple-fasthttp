package mocks

import (
	models "simple-fasthttp/models/modul1"
	repo "simple-fasthttp/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository repo.Repository
	person     *models.Request
}
