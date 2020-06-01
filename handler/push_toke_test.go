package handler_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/wheatandcat/PeperomiaBackend/backend/domain"
	mock_domain "github.com/wheatandcat/PeperomiaBackend/backend/domain/mocks"
	handler "github.com/wheatandcat/PeperomiaBackend/backend/handler"
)

func TestCreatePushToken(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mock := mock_domain.NewMockPushTokenRepository(ctrl)

	i := domain.PushTokenRecord{
		ID:    "sample-uuid-string",
		UID:   "test",
		Token: "test",
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
					Token: "test",
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
