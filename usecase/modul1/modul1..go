package modul1

import (
	repo "simple-fasthttp/entity/database/mysql"
	model "simple-fasthttp/entity/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/valyala/fasthttp"
)

type Request struct {
	Id         int    `json:"id" xml:"id" form:"id"`
	Nama       string `json:"nama" xml:"nama" form:"nama"`
	Nomor      int    `json:"nomor" xml:"nomor" form:"nomor"`
	Created_at string `json:"created_at" xml:"created_at" form:"created_at"`
	Updated_at string `json:"updated_at" xml:"updated_at" form:"upadated_at"`
}

type Response struct {
	Id         int    `json:"id"`
	Nama       string `json:"nama"`
	Nomor      int    `json:"nomor"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

type ResponseAll struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    []*model.Response `json:"data"`
}

type Usecase struct {
	Repo repo.Repository
}

type IUsecase interface {
	GetData(*fasthttp.RequestCtx, *Request) (*ResponseAll, error)
	CreateData(*fasthttp.RequestCtx, *Request) error
	UpdateData(*fasthttp.RequestCtx, *Request) error
	DeleteData(*fasthttp.RequestCtx, *Request) error
}

func NewUsecase(u repo.Repository) IUsecase {
	return &Usecase{Repo: u}
}
func (u Usecase) GetData(ctx *fasthttp.RequestCtx, param *Request) (*ResponseAll, error) {
	data := &model.Request{
		Id:    param.Id,
		Nama:  param.Nama,
		Nomor: param.Nomor,
	}
	res, err := u.Repo.GetData(ctx, data)
	if err != nil {
		return nil, err
	}

	return &ResponseAll{
		Code:    200,
		Message: "success",
		Data:    res,
	}, nil
}

func (u *Usecase) CreateData(ctx *fasthttp.RequestCtx, param *Request) error {
	req := &model.Request{
		Id:    param.Id,
		Nama:  param.Nama,
		Nomor: param.Nomor,
	}
	err := u.Repo.CreateData(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (u *Usecase) UpdateData(ctx *fasthttp.RequestCtx, param *Request) error {
	req := &model.Request{
		Id:    param.Id,
		Nama:  param.Nama,
		Nomor: param.Nomor,
	}
	err := u.Repo.UpdateData(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (u *Usecase) DeleteData(ctx *fasthttp.RequestCtx, param *Request) error {
	req := &model.Request{
		Id: param.Id,
	}
	err := u.Repo.DeleteData(ctx, req)
	if err != nil {
		return err
	}
	return nil
}
