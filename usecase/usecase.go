package usecase

import (
	"fmt"
	"log"
	mdl "simple-fasthttp/models/modul1"
	repo "simple-fasthttp/repository"

	"github.com/valyala/fasthttp"
)

type UsecaseModul struct {
	Repo repo.Repository
}

type Usecase interface {
	ReadData(ctx *fasthttp.RequestCtx, request *mdl.Request) (res mdl.ResponseAll, err error)
	CreateData(ctx *fasthttp.RequestCtx, request *mdl.Request) error
	UpdateData(ctx *fasthttp.RequestCtx, request *mdl.Request) error
	DeleteData(ctx *fasthttp.RequestCtx, request *mdl.Request) error
}

func NewUsecase(u repo.Repository) Usecase {
	return &UsecaseModul{u}
}
func (u UsecaseModul) ReadData(ctx *fasthttp.RequestCtx, param *mdl.Request) (mdl.ResponseAll, error) {
	res, err := u.Repo.GetData(ctx, param)
	if err != nil {
		log.Println(err.Error())
		return res, err
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
