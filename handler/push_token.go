package handler

import (
	"context"
	"net/http"
	"time"

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

const location = "Asia/Tokyo"

// SendCalendarPushNotifications カレンダーに設定されている
func (h *Handler) SendCalendarPushNotifications(gc *gin.Context) {
	loc, _ := time.LoadLocation(location)

	ctx := context.Background()
	dateQuery := gc.Query("date")
	date := time.Now().In(loc)

	if dateQuery != "" {
		d, err := time.Parse("2006-01-02T15:04:05", dateQuery)
		if err != nil {
			gc.JSON(http.StatusInternalServerError, err)
			return
		}
		date = d
	}

	today := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, loc)
	pts, err := h.App.PushTokenRepository.FindAll(ctx, h.FirestoreClient)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, err)
		return
	}

	cs, err := h.App.CalendarRepository.FindByDate(ctx, h.FirestoreClient, &today)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, err)
		return
	}

	res := []string{}

	for _, c := range cs {
		for _, pt := range pts {
			if pt.UID == c.UID {
				req := expopush.SendRequest{
					Body:  "本文",
					Title: "タイトル",
					Data:  map[string]string{"urlSchema": "schedule/" + c.ItemID},
					Token: pt.Token,
				}
				err := h.Client.ExpoPush.Send(req)
				if err == nil {
					res = append(res, c.UID)
				}

			}
		}
	}

	gc.JSON(http.StatusOK, res)
}
