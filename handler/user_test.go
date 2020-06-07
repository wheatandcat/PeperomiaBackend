package handler_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/wheatandcat/PeperomiaBackend/backend/domain"
	mock_domain "github.com/wheatandcat/PeperomiaBackend/backend/domain/mocks"
)

func TestCreateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mock := mock_domain.NewMockUserRepository(ctrl)

	u := domain.UserRecord{
		UID:       "test",
		CreatedAt: time.Date(2020, 1, 1, 00, 00, 00, 0, time.Local),
		UpdatedAt: time.Date(2020, 1, 1, 00, 00, 00, 0, time.Local),
	}

	mock.EXPECT().ExistsByUID(gomock.Any(), gomock.Any(), u.UID).Return(false, nil)
	mock.EXPECT().Create(gomock.Any(), gomock.Any(), u).Return(nil)

	h := NewTestHandler(ctx)
	h.App.UserRepository = mock

	tests := []struct {
		name       string
		statusCode int
	}{
		{
			name:       "ok",
			statusCode: http.StatusCreated,
		},
	}

	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			res := Execute(h.CreateUser, NewRequest(JsonEncode(nil)))
			assert.Equal(t, td.statusCode, res.Code)
		})
	}
}
