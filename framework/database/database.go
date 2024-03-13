package database

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	MySQL *gorm.DB
}

func ConnectMySQL() (*Database, error) {
	msqlInfo := os.Getenv("MYSQL_DIALECTOR")
	Db, err := gorm.Open(mysql.Open(msqlInfo), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		return nil, err
	}

	return &Database{MySQL: Db}, err
}
