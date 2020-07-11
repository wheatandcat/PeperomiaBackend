package domain

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/wheatandcat/PeperomiaBackend/graph/model"
)

// CalendarRecord is Calendar data
type CalendarRecord struct {
	ID     string     `json:"id" firestore:"id" binding:"required"`
	UID    string     `json:"uid" firestore:"uid"`
	ItemID string     `json:"itemId" firestore:"itemId" binding:"required"`
	Date   *time.Time `json:"date" firestore:"date" binding:"required"`
}

// CalendarRepository is repository interface
type CalendarRepository interface {
	Create(ctx context.Context, f *firestore.Client, i CalendarRecord) error
	Update(ctx context.Context, f *firestore.Client, i CalendarRecord) error
	Delete(ctx context.Context, f *firestore.Client, i CalendarRecord) error
	DeleteByUID(ctx context.Context, f *firestore.Client, uid string) error
	DeleteByItemID(ctx context.Context, f *firestore.Client, itemID string) error
	FindByDate(ctx context.Context, f *firestore.Client, date *time.Time) ([]CalendarRecord, error)
	FindByItemID(ctx context.Context, f *firestore.Client, itemID string) (CalendarRecord, error)
}

// ToModel Modelに変換する
func (r *CalendarRecord) ToModel() *model.Calendar {
	const location = "Asia/Tokyo"
	loc, _ := time.LoadLocation(location)

	item := &model.Calendar{
		ID:     r.ID,
		ItemID: r.ItemID,
		Date:   r.Date.In(loc).Format("2006-01-02 15:04:05"),
	}

	return item
}
