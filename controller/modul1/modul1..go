package jadwaldonor

import (
	"fmt"
	"simple-fasthttp/database"
	mdl "simple-fasthttp/models/modul1"

	_ "github.com/go-sql-driver/mysql"
)

const createQuery = "INSERT INTO %s (nama, nomor, created_at) VALUES (?,?,NOW())"
const updateQuery = "UPDATE %s SET nama = ?, nomor = ?, updated_at = NOW() WHERE id = ?"
const deleteQuery = "DELETE FROM %s WHERE id = ?"
const table = "testing"

func GetData(param mdl.Request) (mdl.ResponseAll, error) {
	var (
		result []mdl.Response
		res    mdl.ResponseAll
		err    error
	)
	db := database.GetDB().Debug()
	query := db.Table(table)
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

func CreateData(param mdl.Request) (err error) {
	db := database.GetDB().Debug()
	err = db.Exec(fmt.Sprintf(createQuery, table), param.Nama, param.Nomor).Error
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return err
}

func UpdateData(param mdl.Request) (err error) {
	db := database.GetDB().Debug()
	query := db.Exec(fmt.Sprintf(updateQuery, table), param.Nama, param.Nomor, param.Id)
	err = query.Error
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return err
}

func DeleteData(param mdl.Request) (err error) {
	db := database.GetDB().Debug()
	query := db.Exec(fmt.Sprintf(deleteQuery, table), param.Id)
	err = query.Error
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return err
}
