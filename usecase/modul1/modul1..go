package modul1

import (
	repo "simple-fasthttp/entity/database/mysql"
	model "simple-fasthttp/entity/model"
	fe "simple-fasthttp/framework/error"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/valyala/fasthttp"
)

type Request struct {
	Id string `json:"id" xml:"id" form:"id"`
}

type CreateRequest struct {
	Id    int    `json:"id" xml:"id" form:"id"`
	Nama  string `json:"nama" xml:"nama" form:"nama"`
	Nomor int    `json:"nomor" xml:"nomor" form:"nomor"`
}

type UpdateRequest struct {
	Id    int    `json:"id" xml:"id" form:"id"`
	Nama  string `json:"nama" xml:"nama" form:"nama"`
	Nomor int    `json:"nomor" xml:"nomor" form:"nomor"`
}

type DeleteRequest struct {
	Id int `json:"id" xml:"id" form:"id"`
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
	Data    []*model.Response `json:"data,omitempty"`
}

type Usecase struct {
	Repo repo.Repository
}

type IUsecase interface {
	GetData(*fasthttp.RequestCtx, *Request) (*ResponseAll, *fe.Error)
	CreateData(*fasthttp.RequestCtx, *CreateRequest) *fe.Error
	UpdateData(*fasthttp.RequestCtx, *UpdateRequest) *fe.Error
	DeleteData(*fasthttp.RequestCtx, *DeleteRequest) *fe.Error
}

func NewUsecase(u repo.Repository) IUsecase {
	return &Usecase{Repo: u}
}
func (u Usecase) GetData(ctx *fasthttp.RequestCtx, param *Request) (*ResponseAll, *fe.Error) {
	var id int
	if param.Id != "" {
		i, errv := strconv.Atoi(param.Id)
		if errv != nil {
			return nil, fe.NewError(400, errv)
		}
		id = i
	}

	data := &model.Request{
		Id: id,
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

func (u *Usecase) CreateData(ctx *fasthttp.RequestCtx, param *CreateRequest) *fe.Error {
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

func (u *Usecase) UpdateData(ctx *fasthttp.RequestCtx, param *UpdateRequest) *fe.Error {
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

func (u *Usecase) DeleteData(ctx *fasthttp.RequestCtx, param *DeleteRequest) *fe.Error {
	req := &model.Request{
		Id: param.Id,
	}
	err := u.Repo.DeleteData(ctx, req)
	if err != nil {
		return err
	}
	return nil
}
