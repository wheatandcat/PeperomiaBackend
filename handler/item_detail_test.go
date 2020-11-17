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

func TestCreateItemDetail(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mock := mock_domain.NewMockItemDetailRepository(ctrl)
	i := domain.ItemDetailRecord{
		ID:       "sample-uuid-string",
		UID:      "test",
		Title:    "test",
		Kind:     "test",
		Memo:     "test",
		Place:    "test",
		URL:      "test",
		Priority: 0,
	}

	date, _ := time.Parse("2006-01-02", "2019-01-01")
	key := domain.ItemDetailKey{
		UID:          "test",
		Date:         &date,
		ItemID:       "test",
		ItemDetailID: i.ID,
	}

	mock.EXPECT().Create(gomock.Any(), gomock.Any(), i, key).Return(nil)

	h := NewTestHandler(ctx)
	h.App.ItemDetailRepository = mock

	tests := []struct {
		name       string
		request    handler.CreateItemDetailRequest
		statusCode int
	}{
		{
			name: "ok",
			request: handler.CreateItemDetailRequest{
				ItemDetail: handler.CreateItemDetail{
					ItemID:   "test",
					Title:    "test",
					Kind:     "test",
					Memo:     "test",
					Place:    "test",
					URL:      "test",
					Priority: 0,
				},
				Date: &date,
			},
			statusCode: http.StatusCreated,
		},
	}

	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			res := Execute(h.CreateItemDetail, NewRequest(JsonEncode(td.request)))
			assert.Equal(t, td.statusCode, res.Code)
		})
	}
}

func TestUpdateItemDetail(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mock := mock_domain.NewMockItemDetailRepository(ctrl)
	i := domain.ItemDetailRecord{
		ID:       "test",
		UID:      "test",
		Title:    "test",
		Kind:     "test",
		Memo:     "test",
		URL:      "test",
		Place:    "test",
		Priority: 0,
	}

	date, _ := time.Parse("2006-01-02", "2019-01-01")
	key := domain.ItemDetailKey{
		UID:          "test",
		Date:         &date,
		ItemID:       "test",
		ItemDetailID: i.ID,
	}

	mock.EXPECT().Update(gomock.Any(), gomock.Any(), i, key).Return(nil)

	h := NewTestHandler(ctx)
	h.App.ItemDetailRepository = mock

	tests := []struct {
		name       string
		request    handler.UpdateItemDetailRequest
		statusCode int
	}{
		{
			name: "ok",
			request: handler.UpdateItemDetailRequest{
				ItemDetail: handler.UpdateItemDetail{
					ID:       "test",
					ItemID:   "test",
					Title:    "test",
					Kind:     "test",
					Memo:     "test",
					URL:      "test",
					Place:    "test",
					Priority: 0,
				},
				Date: &date,
			},
			statusCode: http.StatusOK,
		},
	}

	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			res := Execute(h.UpdateItemDetail, NewRequest(JsonEncode(td.request)))
			assert.Equal(t, td.statusCode, res.Code)
		})
	}
}

func TestDeleteItemDetail(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mock := mock_domain.NewMockItemDetailRepository(ctrl)
	i := domain.ItemDetailRecord{
		ID:  "test",
		UID: "test",
	}

	date, _ := time.Parse("2006-01-02", "2019-01-01")
	key := domain.ItemDetailKey{
		UID:          "test",
		Date:         &date,
		ItemID:       "test",
		ItemDetailID: i.ID,
	}

	mock.EXPECT().Delete(gomock.Any(), gomock.Any(), key).Return(nil)

	h := NewTestHandler(ctx)
	h.App.ItemDetailRepository = mock

	tests := []struct {
		name       string
		request    handler.DeleteItemDetailRequest
		statusCode int
	}{
		{
			name: "ok",
			request: handler.DeleteItemDetailRequest{
				ItemDetail: handler.DeleteItemDetail{
					ID:     "test",
					ItemID: "test",
				},
				Date: &date,
			},
			statusCode: http.StatusOK,
		},
	}

	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			res := Execute(h.DeleteItemDetail, NewRequest(JsonEncode(td.request)))
			assert.Equal(t, td.statusCode, res.Code)
		})
	}
}
