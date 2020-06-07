package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wheatandcat/PeperomiaBackend/backend/domain"
)

// CreateUser ユーザーを作成する
func (h *Handler) CreateUser(gc *gin.Context) {
	ctx := context.Background()
	uid, err := GetSelfUID(gc)
	if err != nil {
		NewErrorResponse(err).Render(gc)
		return
	}

	exists, err := h.App.UserRepository.ExistsByUID(ctx, h.FirestoreClient, uid)
	if err != nil {
		NewErrorResponse(err).Render(gc)
		return
	}

	if exists {
		// 既に作成済み
		gc.JSON(http.StatusOK, nil)
		return
	}

	u := domain.UserRecord{
		UID:       uid,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := h.App.UserRepository.Create(ctx, h.FirestoreClient, u); err != nil {
		NewErrorResponse(err).Render(gc)
		return
	}

	gc.JSON(http.StatusCreated, nil)
}
