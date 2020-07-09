package handler_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/wheatandcat/PeperomiaBackend/domain"
	mock_domain "github.com/wheatandcat/PeperomiaBackend/domain/mocks"
	handler "github.com/wheatandcat/PeperomiaBackend/handler"
)

func TestCreatePushToken(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mock := mock_domain.NewMockPushTokenRepository(ctrl)

	i := domain.PushTokenRecord{
		ID:       "sample-uuid-string",
		UID:      "test",
		Token:    "test-Token",
		DeviceID: "test-DeviceID",
	}

	mock.EXPECT().Create(gomock.Any(), gomock.Any(), i).Return(nil)

	h := NewTestHandler(ctx)
	h.App.PushTokenRepository = mock

	tests := []struct {
		name       string
		request    handler.CreatePushTokenRequest
		statusCode int
	}{
		{
			name: "ok",
			request: handler.CreatePushTokenRequest{
				PushToken: handler.CreatePushToken{
					Token:    "test-Token",
					DeviceID: "test-DeviceID",
				},
			},
			statusCode: http.StatusCreated,
		},
	}

	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			res := Execute(h.CreatePushToken, NewRequest(JsonEncode(td.request)))
			assert.Equal(t, td.statusCode, res.Code)
		})
	}
}

func TestSentPushNotifications(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mock := mock_domain.NewMockPushTokenRepository(ctrl)

	i := []domain.PushTokenRecord{{
		ID:       "",
		UID:      "",
		Token:    "",
		DeviceID: "",
	}}

	mock.EXPECT().FindByUID(gomock.Any(), gomock.Any(), "test").Return(i, nil)

	h := NewTestHandler(ctx)
	h.App.PushTokenRepository = mock

	tests := []struct {
		name       string
		request    handler.SentPushNotificationsRequest
		statusCode int
	}{
		{
			name: "ok",
			request: handler.SentPushNotificationsRequest{
				Title:     "test",
				UID:       "test",
				Body:      "test",
				URLScheme: "test",
			},
			statusCode: http.StatusOK,
		},
	}

	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			res := Execute(h.SentPushNotifications, NewRequest(JsonEncode(td.request)))
			assert.Equal(t, td.statusCode, res.Code)
		})
	}
}

const location = "Asia/Tokyo"

func TestSendCalendarPushNotifications(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mockCalendar := mock_domain.NewMockCalendarRepository(ctrl)
	mockPushToken := mock_domain.NewMockPushTokenRepository(ctrl)
	mockItem := mock_domain.NewMockItemRepository(ctrl)

	date := TimeNow()

	pts := []domain.PushTokenRecord{{
		ID:       "",
		UID:      "test",
		Token:    "",
		DeviceID: "",
	}}

	cs := []domain.CalendarRecord{{
		ID:     "",
		UID:    "test",
		ItemID: "",
	}}

	i := domain.ItemRecord{
		ID:    "",
		UID:   "test",
		Title: "",
	}

	today := Day(date)

	mockPushToken.EXPECT().FindAll(gomock.Any(), gomock.Any()).Return(pts, nil)
	mockCalendar.EXPECT().FindByDate(gomock.Any(), gomock.Any(), &today).Return(cs, nil)
	mockItem.EXPECT().FindByDoc(gomock.Any(), gomock.Any(), "test", "").Return(i, nil)

	h := NewTestHandler(ctx)
	h.App.PushTokenRepository = mockPushToken
	h.App.CalendarRepository = mockCalendar
	h.App.ItemRepository = mockItem

	tests := []struct {
		name       string
		statusCode int
	}{
		{
			name:       "ok",
			statusCode: http.StatusOK,
		},
	}

	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			res := Execute(h.SendCalendarPushNotifications, NewRequest(JsonEncode(nil)))
			assert.Equal(t, td.statusCode, res.Code)
		})
	}
}
