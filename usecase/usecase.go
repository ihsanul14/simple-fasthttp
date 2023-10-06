package usecase

import (
	"fmt"
	"log"
	mdl "simple-fasthttp/models/modul1"
	repo "simple-fasthttp/repository"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-lib/metrics/prometheus"
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
	tracer := opentracing.GlobalTracer()
	metricsFactory := prometheus.New()
	tracer, closer, err := config.Configuration{
		ServiceName: "simple-fasthttp",
	}.NewTracer(
		config.Metrics(metricsFactory),
	)
	defer closer.Close()

	span := tracer.StartSpan("read data")
	res, err := u.Repo.GetData(ctx, param)
	if err != nil {
		log.Println(err.Error())
		return res, err
	}
	span.Finish()
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
