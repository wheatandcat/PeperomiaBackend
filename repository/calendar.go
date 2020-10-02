package repository

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/wheatandcat/PeperomiaBackend/domain"
)

// CalendarRepository is repository for calendars
type CalendarRepository struct {
}

// NewCalendarRepository is Create new CalendarRepository
func NewCalendarRepository() domain.CalendarRepository {
	return &CalendarRepository{}
}

func getCalendarDocID(date *time.Time) string {
	doc := date.Format("2006-01-02")
	return doc
}

func calendarCollection(f *firestore.Client, i domain.CalendarRecord) *firestore.DocumentRef {
	idDoc := getCalendarDocID(i.Date)

	return f.Collection("version/1/" + i.UID + "/calendars").Doc(idDoc)
}

// Create カレンダーを作成する
func (re *CalendarRepository) Create(ctx context.Context, f *firestore.Client, i domain.CalendarRecord) error {
	_, err := calendarCollection(f, i).Set(ctx, i)

	return err
}

// Update カレンダーを更新する
func (re *CalendarRepository) Update(ctx context.Context, f *firestore.Client, i domain.CalendarRecord) error {
	_, err := calendarCollection(f, i).Set(ctx, i)

	return err
}

// FindByPublicAndID IDかつPublicから取得する
func (re *CalendarRepository) FindByPublicAndID(ctx context.Context, f *firestore.Client, id string) (domain.CalendarRecord, error) {
	var item domain.CalendarRecord
	matchItem := f.CollectionGroup("calendars").Where("id", "==", id).Where("public", "==", true).Limit(1).Documents(ctx)
	docs, err := matchItem.GetAll()
	if err != nil {
		return item, err
	}

	doc := docs[0]

	doc.DataTo(&item)
	docItem, err := GetItemDoc(ctx, doc)
	if err != nil {
		return item, err
	}

	docItem.DataTo(&item.Item)

	ids, err := GetItemDetailsByDocument(ctx, docItem)
	if err != nil {
		return item, err
	}

	item.Item.ItemDetails = ids

	return item, nil
}

// FindByDate 日付から取得する
func (re *CalendarRepository) FindByDate(ctx context.Context, f *firestore.Client, date *time.Time) ([]domain.CalendarRecord, error) {
	var items []domain.CalendarRecord

	matchItem := f.CollectionGroup("calendars").Where("date", "==", date).Documents(ctx)
	docs, err := matchItem.GetAll()
	if err != nil {
		return items, err
	}

	for _, doc := range docs {
		var item domain.CalendarRecord
		doc.DataTo(&item)
		docItem, err := GetItemDoc(ctx, doc)
		if err != nil {
			return items, err
		}
		docItem.DataTo(&item.Item)

		items = append(items, item)
	}

	return items, nil
}

// Delete カレンダーを削除する
func (re *CalendarRepository) Delete(ctx context.Context, f *firestore.Client, i domain.CalendarRecord) error {
	_, err := calendarCollection(f, i).Delete(ctx)

	return err
}

// DeleteByUID ユーザーIDから削除する
func (re *CalendarRepository) DeleteByUID(ctx context.Context, f *firestore.Client, uid string) error {
	matchItem := f.Collection("version/1/" + uid + "/calendars").Documents(ctx)
	docs, err := matchItem.GetAll()
	if err != nil {
		return err
	}

	for _, doc := range docs {
		err := DeleteItemDoc(ctx, doc)
		if err != nil {
			return err
		}

		if _, err := doc.Ref.Delete(ctx); err != nil {
			return err
		}
	}

	return nil
}
