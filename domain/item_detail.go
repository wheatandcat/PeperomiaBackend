package domain

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/wheatandcat/PeperomiaBackend/graph/model"
)

// ItemDetailRecord is itemDetail data
type ItemDetailRecord struct {
	ID          string `json:"id" firestore:"id" binding:"required"`
	UID         string `json:"uid" firestore:"uid"`
	ItemID      string `json:"itemId" firestore:"itemId" binding:"required"`
	Title       string `json:"title" firestore:"title" binding:"required"`
	Kind        string `json:"kind" firestore:"kind" binding:"required"`
	MoveMinutes int    `json:"moveMinutes" firestore:"moveMinutes"`
	Place       string `json:"place" firestore:"place"`
	URL         string `json:"url" firestore:"url"`
	Memo        string `json:"memo" firestore:"memo"`
	Priority    int    `json:"priority" firestore:"priority"`
}

// ItemDetailKey is item_detail key
type ItemDetailKey struct {
	UID    string
	Date   *time.Time
	ItemID string
}

// ItemDetailRepository is repository interface
type ItemDetailRepository interface {
	Create(ctx context.Context, f *firestore.Client, i ItemDetailRecord, key ItemDetailKey) error
	Update(ctx context.Context, f *firestore.Client, i ItemDetailRecord, key ItemDetailKey) error
	Delete(ctx context.Context, f *firestore.Client, i ItemDetailRecord, key ItemDetailKey) error
	Get(ctx context.Context, f *firestore.Client, i ItemDetailRecord, key ItemDetailKey) (ItemDetailRecord, error)
	FindByItemID(ctx context.Context, f *firestore.Client, itemID string) ([]ItemDetailRecord, error)
}

// ToModel Modelに変換する
func (r *ItemDetailRecord) ToModel() *model.ItemDetail {
	item := &model.ItemDetail{
		ID:          r.ID,
		ItemID:      r.ItemID,
		Title:       r.Title,
		Kind:        r.Kind,
		MoveMinutes: r.MoveMinutes,
		Place:       r.Place,
		URL:         r.URL,
		Memo:        r.Memo,
		Priority:    r.Priority,
	}

	return item
}
