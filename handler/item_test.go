package handler_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/wheatandcat/PeperomiaBackend/domain"
	mock_domain "github.com/wheatandcat/PeperomiaBackend/domain/mocks"
	handler "github.com/wheatandcat/PeperomiaBackend/handler"
)

func TestCreateItem(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mock := mock_domain.NewMockItemRepository(ctrl)
	i := domain.ItemRecord{
		ID:    "sample-uuid-string",
		UID:   "test",
		Title: "test",
		Kind:  "test",
	}

	date, _ := time.Parse("2006-01-02", "2019-01-01")
	key := domain.ItemKey{
		UID:  "test",
		Date: &date,
	}

	mock.EXPECT().Create(gomock.Any(), gomock.Any(), i, key).Return(nil)

	h := NewTestHandler(ctx)
	h.App.ItemRepository = mock

	tests := []struct {
		name       string
		request    handler.CreateItemRequest
		statusCode int
	}{
		{
			name: "ok",
			request: handler.CreateItemRequest{
				Item: handler.CreateItem{
					Title: "test",
					Kind:  "test",
				},
				Date: &date,
			},
			statusCode: http.StatusCreated,
		},
	}

	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			res := Execute(h.CreateItem, NewRequest(JsonEncode(td.request)))
			assert.Equal(t, td.statusCode, res.Code)
		})
	}
}

func TestUpdateItem(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mock := mock_domain.NewMockItemRepository(ctrl)
	i := domain.ItemRecord{
		ID:    "test",
		UID:   "test",
		Title: "test",
		Kind:  "test",
	}

	date, _ := time.Parse("2006-01-02", "2019-01-01")
	key := domain.ItemKey{
		UID:  "test",
		Date: &date,
	}

	mock.EXPECT().FindByDoc(gomock.Any(), gomock.Any(), "test", "test").Return(i, nil)
	mock.EXPECT().Update(gomock.Any(), gomock.Any(), i, key).Return(nil)

	h := NewTestHandler(ctx)
	h.App.ItemRepository = mock

	tests := []struct {
		name       string
		request    handler.UpdateItemRequest
		statusCode int
	}{
		{
			name: "ok",
			request: handler.UpdateItemRequest{
				Item: handler.UpdateItem{
					ID:    "test",
					Title: "test",
					Kind:  "test",
				},
				Date: &date,
			},
			statusCode: http.StatusOK,
		},
	}

	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			res := Execute(h.UpdateItem, NewRequest(JsonEncode(td.request)))
			assert.Equal(t, td.statusCode, res.Code)
		})
	}
}
