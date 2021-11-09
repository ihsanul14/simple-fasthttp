package modul1

import (
	"log"
	mdl "simple-fasthttp/models/modul1"
	repo "simple-fasthttp/repository"

	_ "github.com/go-sql-driver/mysql"
	"github.com/valyala/fasthttp"
)

type Usecase struct {
	Repo repo.Repository
}

func NewUsecase(u repo.Repository) Usecase {
	return &Usecase{Repo: u}
}
func (u Usecase) ShowData(ctx *fasthttp.RequestCtx, param *mdl.Request) (*mdl.ResponseAll, error) {
	res, err := u.Repo.GetData(ctx, param)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return res, err
}

// func (u *Usecase)CreateData(param mdl.Request) (err error) {
// 	db := database.GetDB().Debug()
// 	err = db.Exec(fmt.Sprintf(createQuery, table), param.Nama, param.Nomor).Error
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return err
// 	}
// 	return err
// }

// func (u *Usecase)UpdateData(param mdl.Request) (err error) {
// 	db := database.GetDB().Debug()
// 	query := db.Exec(fmt.Sprintf(updateQuery, table), param.Nama, param.Nomor, param.Id)
// 	err = query.Error
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return err
// 	}
// 	return err
// }

// func (u *Usecase)DeleteData(param mdl.Request) (err error) {
// 	db := database.GetDB().Debug()
// 	query := db.Exec(fmt.Sprintf(deleteQuery, table), param.Id)
// 	err = query.Error
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return err
// 	}
// 	return err
// }
