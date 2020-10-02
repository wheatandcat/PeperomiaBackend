package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	domain "github.com/wheatandcat/PeperomiaBackend/domain"
)

// SyncItemsRequest is SyncItemsRequest request
type SyncItemsRequest struct {
	Items       []domain.ItemRecord       `json:"items" binding:"required"`
	ItemDetails []domain.ItemDetailRecord `json:"itemDetails" binding:"required"`
	Calendars   []domain.CalendarRecord   `json:"calendars"`
}

// CreateItemRequest is CreateItem request
type CreateItemRequest struct {
	Item CreateItem `json:"item" binding:"required"`
	Date *time.Time `json:"date" binding:"required"`
}

// CreateItem is CreateItem request
type CreateItem struct {
	Title string `json:"title" binding:"required"`
	Kind  string `json:"kind" binding:"required"`
}

// UpdateItemRequest is UpdateItem request
type UpdateItemRequest struct {
	Item UpdateItem `json:"item" binding:"required"`
	Date *time.Time `json:"date" binding:"required"`
}

// UpdateItem is UpdateItem request
type UpdateItem struct {
	ID    string `json:"id" binding:"required"`
	Title string `json:"title" binding:"required"`
	Kind  string `json:"kind" binding:"required"`
}

// CreateItem 予定を作成する
func (h *Handler) CreateItem(gc *gin.Context) {
	ctx := context.Background()
	req := &CreateItemRequest{}
	if err := gc.Bind(req); err != nil {
		NewErrorResponse(err).Render(gc)
		return
	}

	uid, err := GetSelfUID(gc)
	if err != nil {
		NewErrorResponse(err).Render(gc)
		return
	}

	if err != nil {
		NewErrorResponse(err).Render(gc)
		return
	}

	item := domain.ItemRecord{
		ID:    h.Client.UUID.Get(),
		Title: req.Item.Title,
		Kind:  req.Item.Kind,
		UID:   uid,
	}

	key := domain.ItemKey{
		UID:  uid,
		Date: req.Date,
	}

	if err := h.App.ItemRepository.Create(ctx, h.FirestoreClient, item, key); err != nil {
		NewErrorResponse(err).Render(gc)
		return
	}

	gc.JSON(http.StatusCreated, item)
}

// UpdateItem 予定を更新する
func (h *Handler) UpdateItem(gc *gin.Context) {
	ctx := context.Background()
	req := &UpdateItemRequest{}
	if err := gc.Bind(req); err != nil {
		NewErrorResponse(err).Render(gc)
		return
	}

	uid, err := GetSelfUID(gc)
	if err != nil {
		NewErrorResponse(err).Render(gc)
		return
	}

	item, err := h.App.ItemRepository.FindByDoc(ctx, h.FirestoreClient, uid, req.Item.ID)
	if err != nil {
		NewErrorResponse(err).Render(gc)
		return
	}

	item.Title = req.Item.Title
	item.Kind = req.Item.Kind

	key := domain.ItemKey{
		UID:  uid,
		Date: req.Date,
	}

	if err := h.App.ItemRepository.Update(ctx, h.FirestoreClient, item, key); err != nil {
		NewErrorResponse(err).Render(gc)
		return
	}

	gc.JSON(http.StatusOK, nil)
}

// SyncItems アイテムを同期させる
func (h *Handler) SyncItems(gc *gin.Context) {
	gc.JSON(http.StatusCreated, nil)
}
