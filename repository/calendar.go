package repository

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/wheatandcat/PeperomiaBackend/backend/domain"
)

// CalendarRepository is repository for calendars
type CalendarRepository struct {
}

// NewCalendarRepository is Create new CalendarRepository
func NewCalendarRepository() domain.CalendarRepository {
	return &CalendarRepository{}
}

func getCalendarDocID(uID string, itemID string, calendarID string) string {
	doc := uID + "_" + calendarID + "_" + itemID
	return doc
}

// Create カレンダーを作成する
func (re *CalendarRepository) Create(ctx context.Context, f *firestore.Client, i domain.CalendarRecord) error {
	idDoc := getCalendarDocID(i.UID, i.ItemID, i.ID)

	_, err := f.Collection("calendars").Doc(idDoc).Set(ctx, i)

	return err
}

// Update カレンダーを更新する
func (re *CalendarRepository) Update(ctx context.Context, f *firestore.Client, i domain.CalendarRecord) error {
	idDoc := getCalendarDocID(i.UID, i.ItemID, i.ID)

	_, err := f.Collection("calendars").Doc(idDoc).Set(ctx, i)

	return err
}

// FindByDate 日付から取得する
func (re *CalendarRepository) FindByDate(ctx context.Context, f *firestore.Client, date *time.Time) ([]domain.CalendarRecord, error) {
	var items []domain.CalendarRecord

	matchItem := f.Collection("calendars").Where("date", "==", date).Documents(ctx)
	docs, err := matchItem.GetAll()
	if err != nil {
		return items, err
	}

	for _, doc := range docs {
		var item domain.CalendarRecord
		doc.DataTo(&item)
		items = append(items, item)
	}

	return items, nil
}

// Delete カレンダーを削除する
func (re *CalendarRepository) Delete(ctx context.Context, f *firestore.Client, i domain.CalendarRecord) error {
	idDoc := getCalendarDocID(i.UID, i.ItemID, i.ID)

	_, err := f.Collection("calendars").Doc(idDoc).Delete(ctx)

	return err
}

// DeleteByUID ユーザーIDから削除する
func (re *CalendarRepository) DeleteByUID(ctx context.Context, f *firestore.Client, uid string) error {
	matchItem := f.Collection("calendars").Where("uid", "==", uid).Documents(ctx)
	docs, err := matchItem.GetAll()
	if err != nil {
		return err
	}

	for _, doc := range docs {
		if _, err := doc.Ref.Delete(ctx); err != nil {
			return err
		}
	}

	return nil
}

// DeleteByItemID ItemIDから削除する
func (re *CalendarRepository) DeleteByItemID(ctx context.Context, f *firestore.Client, itemID string) error {
	matchItem := f.Collection("calendars").Where("itemId", "==", itemID).Documents(ctx)
	docs, err := matchItem.GetAll()
	if err != nil {
		return err
	}

	for _, doc := range docs {
		if _, err := doc.Ref.Delete(ctx); err != nil {
			return err
		}
	}

	return nil
}
