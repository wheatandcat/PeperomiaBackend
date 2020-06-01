package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wheatandcat/PeperomiaBackend/backend/domain"
)

// CreatePushTokenRequest is CreatePushToke request
type CreatePushTokenRequest struct {
	PushToken CreatePushToken `json:"pushToken" binding:"required"`
}

// CreatePushToken is CreatePushToken request
type CreatePushToken struct {
	Token string `json:"token" binding:"required"`
}

// CreatePushToken Expo Push通知トークンを作成する
func (h *Handler) CreatePushToken(gc *gin.Context) {
	ctx := context.Background()
	req := &CreatePushTokenRequest{}
	if err := gc.Bind(req); err != nil {
		gc.JSON(http.StatusInternalServerError, err)
		return
	}

	uid, err := GetSelfUID(gc)
	if err != nil {
		NewErrorResponse(err).Render(gc)
		return
	}

	i := domain.PushTokenRecord{
		ID:    h.Client.UUID.Get(),
		UID:   uid,
		Token: req.PushToken.Token,
	}

	if err := h.App.PushTokenRepository.Create(ctx, h.FirestoreClient, i); err != nil {
		NewErrorResponse(err).Render(gc)
		return
	}

	gc.JSON(http.StatusCreated, nil)
}
