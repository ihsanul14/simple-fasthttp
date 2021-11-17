package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

const database = "MySQL"

func ConnectMySQL() (*gorm.DB, error) {
	var (
		err      error
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASS")
		dbname   = os.Getenv("DB_NAME")
	)
	msqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbname)

	Db, err = gorm.Open(mysql.Open(msqlInfo), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	if err != nil {
		log.Println("Connection Database Error ", err.Error())
	} else {
		log.Println(database, "is Connected")
	}
	Db = Db.Debug()
	return Db, err
}
