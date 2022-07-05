package usecase

// import (
// 	"context"
// 	"testing"
// 	"time"

// 	mocks "simple-fasthttp/mocks"
// 	models "simple-fasthttp/models/modul1"
// 	domain "simple-fasthttp/repository"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// func TestReadData(t *testing.T) {
// 	mockArticleRepo := new(mocks.Repository)
// 	mockArticle := models.Request{
// 		Id: 1,
// 	}

// 	mockListArtilce := make([]domain.Repository, 0)
// 	mockListArtilce = append(mockListArtilce, mockArticle)

// 	t.Run("success", func(t *testing.T) {
// 		mockArticleRepo.On("Fetch", mock.Anything, mock.AnythingOfType("string"),
// 			mock.AnythingOfType("int64")).Return(mockListArtilce, "next-cursor", nil).Once()
// 		u := ucase.NewArticleUsecase(mockArticleRepo, mockAuthorrepo, time.Second*2)
// 		num := int64(1)
// 		cursor := "12"
// 		list, nextCursor, err := u.Fetch(context.TODO(), cursor, num)
// 		cursorExpected := "next-cursor"
// 		assert.Equal(t, cursorExpected, nextCursor)
// 		assert.NotEmpty(t, nextCursor)
// 		assert.NoError(t, err)
// 		assert.Len(t, list, len(mockListArtilce))

// 		mockArticleRepo.AssertExpectations(t)
// 		mockAuthorrepo.AssertExpectations(t)
// 	})
// }
