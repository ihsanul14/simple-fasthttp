package modul1

import (
	"fmt"
	"net/http/httptest"
	"testing"

	fe "simple-fasthttp/framework/error"
	"simple-fasthttp/framework/infra"
	"simple-fasthttp/framework/logger"
	mocks "simple-fasthttp/framework/mocks/usecase/modul1"
	"simple-fasthttp/framework/validator"
	usecase "simple-fasthttp/usecase/modul1"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func initRouter() *fiber.App {
	router := fiber.New()
	return router
}

func TestNewApplication(t *testing.T) {
	router := initRouter()
	NewApplication(router, &infra.Infra{
		Validator: validator.NewValidator(),
	})
}

func TestGetData(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockUsecase := mocks.NewMockIUsecase(mockController)
	handler := &Handler{
		Usecase:   mockUsecase,
		Validator: validator.NewValidator(),
		Logger:    logger.NewLogger(),
	}
	assert.NotNil(t, handler)

	router := initRouter()
	router.Get("api/data", handler.GetDataHandler)
	t.Run("200", func(t *testing.T) {
		mockUsecase.EXPECT().GetData(gomock.Any(), gomock.Any()).Return(&usecase.ResponseAll{
			Code:    200,
			Message: "success",
		}, nil)

		res := httptest.NewRequest("GET", "/api/data", nil)
		resp, err := router.Test(res, 1)
		assert.Nil(t, err)
		assert.Equal(t, resp.StatusCode, 200)
	})

	t.Run("500", func(t *testing.T) {
		mockUsecase.EXPECT().GetData(gomock.Any(), gomock.Any()).Return(nil, fe.NewError(500, fmt.Errorf("error")))
		res := httptest.NewRequest("GET", "/api/data", nil)
		resp, err := router.Test(res, 1)
		assert.Nil(t, err)
		assert.Equal(t, resp.StatusCode, 500)
	})

	t.Run("400", func(t *testing.T) {
		mockUsecase.EXPECT().GetData(gomock.Any(), gomock.Any()).Return(nil, fe.NewError(400, fmt.Errorf("id must be greater than 0")))
		res := httptest.NewRequest("GET", "/api/data?id=-1", nil)
		resp, err := router.Test(res, 1)
		assert.Nil(t, err)
		assert.Equal(t, resp.StatusCode, 400)
	})
}

func TestValidator(t *testing.T) {
	v := validator.NewValidator()
	t.Run("Success", func(t *testing.T) {
		err := v.Check(&usecase.Request{
			Id: "",
		})
		assert.Nil(t, err)
	})
}
