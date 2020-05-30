package handler_test

import (
	"context"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/wheatandcat/PeperomiaBackend/backend/domain"
	mock_domain "github.com/wheatandcat/PeperomiaBackend/backend/domain/mocks"
	handler "github.com/wheatandcat/PeperomiaBackend/backend/handler"
)

func TestCreateCalendar(t *testing.T) {
	gin.SetMode(gin.TestMode)
	log.Print("001")

	ctrl := gomock.NewController(t)
	log.Print("002")
	defer ctrl.Finish()
	ctx := context.Background()

	mock := mock_domain.NewMockCalendarRepository(ctrl)
	log.Print("003")
	date, _ := time.Parse("2006-01-02", "2019-01-01")
	log.Print("004")

	i := domain.CalendarRecord{
		ID:     "sample-uuid-string",
		ItemID: "test",
		UID:    "test",
		Date:   &date,
	}

	log.Print("005")
	mock.EXPECT().Create(gomock.Any(), gomock.Any(), i).Return(nil)

	log.Print("006")
	h := NewTestHandler(ctx)
	h.App.CalendarRepository = mock
	log.Print("0061")

	tests := []struct {
		name       string
		request    handler.CreateCalendarRequest
		statusCode int
	}{
		{
			name: "ok",
			request: handler.CreateCalendarRequest{
				Calendar: handler.CreateCalendar{
					ItemID: "test",
					Date:   &date,
				},
			},
			statusCode: http.StatusCreated,
		},
	}

	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			log.Print("007")
			res := Execute(h.CreateCalendar, NewRequest(JsonEncode(td.request)))
			assert.Equal(t, td.statusCode, res.Code)
		})
	}
}

func TestUpdateCalendar(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()
	date, _ := time.Parse("2006-01-02", "2019-01-01")

	mock := mock_domain.NewMockCalendarRepository(ctrl)
	i := domain.CalendarRecord{
		ID:     "test",
		ItemID: "test",
		UID:    "test",
		Date:   &date,
	}

	mock.EXPECT().Update(gomock.Any(), gomock.Any(), i).Return(nil)

	h := NewTestHandler(ctx)
	h.App.CalendarRepository = mock

	tests := []struct {
		name       string
		request    handler.UpdateCalendarRequest
		statusCode int
	}{
		{
			name: "ok",
			request: handler.UpdateCalendarRequest{
				Calendar: handler.UpdateCalendar{
					ID:     "test",
					ItemID: "test",
					Date:   &date,
				},
			},
			statusCode: http.StatusOK,
		},
	}

	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			res := Execute(h.UpdateCalendar, NewRequest(JsonEncode(td.request)))
			assert.Equal(t, td.statusCode, res.Code)
		})
	}
}

func TestDeleteCalendar(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mock := mock_domain.NewMockCalendarRepository(ctrl)
	i := domain.CalendarRecord{
		ID:     "test",
		ItemID: "test",
		UID:    "test",
	}

	mock.EXPECT().Delete(gomock.Any(), gomock.Any(), i).Return(nil)

	h := NewTestHandler(ctx)
	h.App.CalendarRepository = mock

	tests := []struct {
		name       string
		request    handler.DeleteCalendarRequest
		statusCode int
	}{
		{
			name: "ok",
			request: handler.DeleteCalendarRequest{
				Calendar: handler.DeleteCalendar{
					ID:     "test",
					ItemID: "test",
				},
			},
			statusCode: http.StatusOK,
		},
	}

	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			res := Execute(h.DeleteCalendar, NewRequest(JsonEncode(td.request)))
			assert.Equal(t, td.statusCode, res.Code)
		})
	}
}
