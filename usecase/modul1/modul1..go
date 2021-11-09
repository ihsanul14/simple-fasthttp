package modul1

import (
	"fmt"
	"log"
	mdl "simple-fasthttp/models/modul1"
	repo "simple-fasthttp/repository"
	"simple-fasthttp/usecase"

	_ "github.com/go-sql-driver/mysql"
	"github.com/valyala/fasthttp"
)

type UsecaseModul struct {
	Repo repo.Repository
}

func NewUsecase(u repo.Repository) usecase.Usecase {
	return &UsecaseModul{
		Repo: u,
	}
}
func (u UsecaseModul) ShowData(ctx *fasthttp.RequestCtx, param *mdl.Request) (*mdl.ResponseAll, error) {
	res, err := u.Repo.GetData(ctx, param)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return res, err
}

func (u *UsecaseModul) CreateData(ctx *fasthttp.RequestCtx, param *mdl.Request) error {
	err := u.Repo.CreateData(ctx, param)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return err
}

func (u *UsecaseModul) UpdateData(ctx *fasthttp.RequestCtx, param *mdl.Request) error {
	err := u.Repo.UpdateData(ctx, param)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return err
}

func (u *UsecaseModul) DeleteData(ctx *fasthttp.RequestCtx, param *mdl.Request) error {
	err := u.Repo.DeleteData(ctx, param)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return err
}
