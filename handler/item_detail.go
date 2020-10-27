package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wheatandcat/PeperomiaBackend/domain"
)

// CreateItemDetailRequest is CreateItemDetail request
type CreateItemDetailRequest struct {
	ItemDetail CreateItemDetail `json:"itemDetail" binding:"required"`
	Date       *time.Time       `json:"date" binding:"required"`
}

// CreateItemDetail is CreateItemDetail request
type CreateItemDetail struct {
	ItemID   string `json:"itemID" binding:"required"`
	Title    string `json:"title" binding:"required"`
	Kind     string `json:"kind" binding:"required"`
	Memo     string `json:"memo"`
	URL      string `json:"url"`
	Place    string `json:"place"`
	Priority int    `json:"priority"`
}

// UpdateItemDetailRequest is UpdateItemDetail request
type UpdateItemDetailRequest struct {
	ItemDetail UpdateItemDetail `json:"itemDetail" binding:"required"`
	Date       *time.Time       `json:"date" binding:"required"`
}

// UpdateItemDetail is UpdateItemDetail request
type UpdateItemDetail struct {
	ID       string `json:"id" binding:"required"`
	ItemID   string `json:"itemID" binding:"required"`
	Title    string `json:"title" binding:"required"`
	Kind     string `json:"kind" binding:"required"`
	Memo     string `json:"memo"`
	URL      string `json:"url"`
	Place    string `json:"place"`
	Priority int    `json:"priority"`
}

// DeleteItemDetailRequest is DeleteItemDetail request
type DeleteItemDetailRequest struct {
	ItemDetail DeleteItemDetail `json:"itemDetail" binding:"required"`
	Date       *time.Time       `json:"date" binding:"required"`
}

// DeleteItemDetail is DeleteItemDetail request
type DeleteItemDetail struct {
	ID     string `json:"id" binding:"required"`
	ItemID string `json:"itemID" binding:"required"`
}

// CreateItemDetail 予定の詳細を作成する
func (h *Handler) CreateItemDetail(gc *gin.Context) {
	ctx := context.Background()
	req := &CreateItemDetailRequest{}
	if err := gc.Bind(req); err != nil {
		NewErrorResponse(err).Render(gc)
		return
	}

	uid, err := GetSelfUID(gc)
	if err != nil {
		NewErrorResponse(err).Render(gc)
		return
	}

	item := domain.ItemDetailRecord{
		ID:       h.Client.UUID.Get(),
		Title:    req.ItemDetail.Title,
		Kind:     req.ItemDetail.Kind,
		UID:      uid,
		Memo:     req.ItemDetail.Memo,
		URL:      req.ItemDetail.URL,
		Place:    req.ItemDetail.Place,
		Priority: req.ItemDetail.Priority,
	}

	key := domain.ItemDetailKey{
		UID:    uid,
		Date:   req.Date,
		ItemID: req.ItemDetail.ItemID,
	}

	if err := h.App.ItemDetailRepository.Create(ctx, h.FirestoreClient, item, key); err != nil {
		NewErrorResponse(err).Render(gc)
		return
	}

	gc.JSON(http.StatusCreated, item)
}

// UpdateItemDetail 予定の詳細を更新する
func (h *Handler) UpdateItemDetail(gc *gin.Context) {
	ctx := context.Background()
	req := &UpdateItemDetailRequest{}
	if err := gc.Bind(req); err != nil {
		NewErrorResponse(err).Render(gc)
		return
	}

	uid, err := GetSelfUID(gc)
	if err != nil {
		NewErrorResponse(err).Render(gc)
		return
	}

	item := domain.ItemDetailRecord{
		ID:       req.ItemDetail.ID,
		Title:    req.ItemDetail.Title,
		Kind:     req.ItemDetail.Kind,
		UID:      uid,
		Memo:     req.ItemDetail.Memo,
		URL:      req.ItemDetail.URL,
		Place:    req.ItemDetail.Place,
		Priority: req.ItemDetail.Priority,
	}

	key := domain.ItemDetailKey{
		UID:    uid,
		Date:   req.Date,
		ItemID: req.ItemDetail.ItemID,
	}

	if err := h.App.ItemDetailRepository.Update(ctx, h.FirestoreClient, item, key); err != nil {
		NewErrorResponse(err).Render(gc)
		return
	}

	gc.JSON(http.StatusOK, nil)
}

// DeleteItemDetail 予定の詳細を削除する
func (h *Handler) DeleteItemDetail(gc *gin.Context) {
	ctx := context.Background()
	req := &DeleteItemDetailRequest{}
	if err := gc.Bind(req); err != nil {
		NewErrorResponse(err).Render(gc)
		return
	}

	uid, err := GetSelfUID(gc)
	if err != nil {
		NewErrorResponse(err).Render(gc)
		return
	}

	item := domain.ItemDetailRecord{
		ID:  req.ItemDetail.ID,
		UID: uid,
	}

	key := domain.ItemDetailKey{
		UID:    uid,
		Date:   req.Date,
		ItemID: req.ItemDetail.ItemID,
	}

	if err := h.App.ItemDetailRepository.Delete(ctx, h.FirestoreClient, item, key); err != nil {
		NewErrorResponse(err).Render(gc)
		return
	}

	gc.JSON(http.StatusOK, nil)
}
