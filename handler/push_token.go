package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	expopush "github.com/wheatandcat/PeperomiaBackend/backend/client/expo_push"
	"github.com/wheatandcat/PeperomiaBackend/backend/domain"
)

// CreatePushTokenRequest is CreatePushToke request
type CreatePushTokenRequest struct {
	PushToken CreatePushToken `json:"pushToken" binding:"required"`
}

// CreatePushToken is CreatePushToken request
type CreatePushToken struct {
	Token    string `json:"token" binding:"required"`
	DeviceID string `json:"deviceId" binding:"required"`
}

// SentPushNotificationsRequest is SentPushNotifications request
type SentPushNotificationsRequest struct {
	UID       string `json:"uid" binding:"required"`
	Title     string `json:"title" binding:"required"`
	Body      string `json:"body" binding:"required"`
	URLScheme string `json:"urlScheme"`
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
		ID:       h.Client.UUID.Get(),
		UID:      uid,
		Token:    req.PushToken.Token,
		DeviceID: req.PushToken.DeviceID,
	}

	if err := h.App.PushTokenRepository.Create(ctx, h.FirestoreClient, i); err != nil {
		NewErrorResponse(err).Render(gc)
		return
	}

	gc.JSON(http.StatusCreated, nil)
}

// SentPushNotifications Expo Push通知を送信する（テスト用）
func (h *Handler) SentPushNotifications(gc *gin.Context) {
	ctx := context.Background()
	req := &SentPushNotificationsRequest{}
	if err := gc.Bind(req); err != nil {
		gc.JSON(http.StatusInternalServerError, err)
		return
	}

	pts, err := h.App.PushTokenRepository.FindByUID(ctx, h.FirestoreClient, req.UID)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, err)
		return
	}

	for _, pt := range pts {
		req := expopush.SendRequest{
			Body:  req.Body,
			Title: req.Title,
			Data:  map[string]string{"urlSchema": req.URLScheme},
			Token: pt.Token,
		}

		if err != h.Client.ExpoPush.Send(req) {
			gc.JSON(http.StatusInternalServerError, err)
			return
		}
	}

	gc.JSON(http.StatusOK, nil)
}
