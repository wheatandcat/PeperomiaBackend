package domain

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/wheatandcat/PeperomiaBackend/graph/model"
)

// ItemRecord is item data
type ItemRecord struct {
	ID          string              `json:"id" firestore:"id" binding:"required"`
	UID         string              `json:"uid" firestore:"uid"`
	Title       string              `json:"title" firestore:"title" binding:"required"`
	Kind        string              `json:"kind" firestore:"kind" binding:"required"`
	CreatedAt   time.Time           `json:"-" firestore:"createdAt"`
	UpdatedAt   time.Time           `json:"-" firestore:"updatedAt"`
	ItemDetails []*ItemDetailRecord `json:"itemDetails" firestore:"itemDetails"`
}

// ItemKey is item key
type ItemKey struct {
	UID  string
	Date *time.Time
}

// ItemRepository is repository interface
type ItemRepository interface {
	Create(ctx context.Context, f *firestore.Client, i ItemRecord, key ItemKey) error
	Update(ctx context.Context, f *firestore.Client, i ItemRecord, key ItemKey) error
	FindByDoc(ctx context.Context, f *firestore.Client, uid string, itemID string) (ItemRecord, error)
	FindByUID(ctx context.Context, f *firestore.Client, uid string) ([]ItemRecord, error)
}

// ToModel Modelに変換する
func (r *ItemRecord) ToModel() *model.Item {
	var ids []*model.ItemDetail

	for _, id := range r.ItemDetails {
		ids = append(ids, id.ToModel())
	}

	item := &model.Item{
		ID:          r.ID,
		Kind:        r.Kind,
		Title:       r.Title,
		ItemDetails: ids,
	}

	return item
}
