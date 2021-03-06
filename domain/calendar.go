package domain

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/wheatandcat/PeperomiaBackend/graph/model"
)

// CalendarRecord is Calendar data
type CalendarRecord struct {
	ID     string      `json:"id" firestore:"id" binding:"required"`
	UID    string      `json:"uid" firestore:"uid"`
	Public bool        `json:"public" firestore:"public"`
	Date   *time.Time  `json:"date" firestore:"date" binding:"required"`
	Item   *ItemRecord `json:"item" firestore:"item"`
}

// CalendarRepository is repository interface
type CalendarRepository interface {
	Create(ctx context.Context, f *firestore.Client, i CalendarRecord) error
	Update(ctx context.Context, f *firestore.Client, i CalendarRecord) error
	Delete(ctx context.Context, f *firestore.Client, i CalendarRecord) error
	DeleteByUID(ctx context.Context, f *firestore.Client, uid string) error
	DeleteByDateAndUID(ctx context.Context, f *firestore.Client, uid string, date *time.Time) error
	FindBetweenDateAndUID(ctx context.Context, f *firestore.Client, uid string, startDate *time.Time, endDate *time.Time) ([]CalendarRecord, error)
	FindByDateAndUID(ctx context.Context, f *firestore.Client, uid string, date *time.Time) (CalendarRecord, error)
	FindByDate(ctx context.Context, f *firestore.Client, date *time.Time) ([]CalendarRecord, error)
	FindByPublicAndID(ctx context.Context, f *firestore.Client, id string) (CalendarRecord, error)
}

// ToShareItemModel Modelに変換する
func (r *CalendarRecord) ToShareItemModel() *model.ShareItem {
	const location = "Asia/Tokyo"
	loc, _ := time.LoadLocation(location)

	item := &model.ShareItem{
		ID:   r.ID,
		Date: r.Date.In(loc).Format("2006-01-02 15:04:05"),
		Item: r.Item.ToModel(),
	}

	return item
}

// ToModel Modelに変換する
func (r *CalendarRecord) ToModel() *model.Calendar {
	const location = "Asia/Tokyo"
	loc, _ := time.LoadLocation(location)

	item := &model.Calendar{
		ID:     r.ID,
		Date:   r.Date.In(loc).Format("2006-01-02 15:04:05"),
		Public: r.Public,
	}
	if r.Item != nil {
		item.Item = r.Item.ToModel()
	}

	return item
}
