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
	const location = "Asia/Tokyo"
	loc, _ := time.LoadLocation(location)

	doc := date.In(loc).Format("2006-01-02")
	return doc
}

func calendarCollectionRef(f *firestore.Client, uid string) *firestore.CollectionRef {
	return f.Collection("version/1/users/" + uid + "/calendars")
}

func calendarCollection(f *firestore.Client, i domain.CalendarRecord) *firestore.DocumentRef {
	idDoc := getCalendarDocID(i.Date)

	return calendarCollectionRef(f, i.UID).Doc(idDoc)
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

	matchItem := f.CollectionGroup("calendars").Where("date", "==", date).OrderBy("id", firestore.Asc).Documents(ctx)
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

// FindBetweenDateAndUID 開始日から終了日まで取得する
func (re *CalendarRepository) FindBetweenDateAndUID(ctx context.Context, f *firestore.Client, uid string, startDate *time.Time, endDate *time.Time) ([]domain.CalendarRecord, error) {
	var items []domain.CalendarRecord

	matchItem := calendarCollectionRef(f, uid).Where("date", ">=", startDate).Where("date", "<=", endDate).OrderBy("date", firestore.Asc).Documents(ctx)
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

// FindByDateAndUID 日付から取得する
func (re *CalendarRepository) FindByDateAndUID(ctx context.Context, f *firestore.Client, uid string, date *time.Time) (domain.CalendarRecord, error) {
	cr := domain.CalendarRecord{
		UID:  uid,
		Date: date,
	}

	doc, err := calendarCollection(f, cr).Get(ctx)
	if err != nil {
		return cr, err
	}

	doc.DataTo(&cr)

	docItem, err := GetItemDoc(ctx, doc)
	if err != nil {
		return cr, err
	}
	docItem.DataTo(&cr.Item)
	ids, err := GetItemDetailsByDocument(ctx, docItem)
	if err != nil {
		return cr, err
	}

	cr.Item.ItemDetails = ids

	return cr, nil
}

// Delete カレンダーを削除する
func (re *CalendarRepository) Delete(ctx context.Context, f *firestore.Client, i domain.CalendarRecord) error {
	_, err := calendarCollection(f, i).Delete(ctx)

	return err
}

// DeleteByDateAndUID カレンダーを削除する
func (re *CalendarRepository) DeleteByDateAndUID(ctx context.Context, f *firestore.Client, uid string, date *time.Time) error {
	cr := domain.CalendarRecord{
		UID:  uid,
		Date: date,
	}

	doc, err := calendarCollection(f, cr).Get(ctx)
	if err != nil {
		return err
	}

	if err := DeleteItemDoc(ctx, doc); err != nil {
		return err
	}

	if _, err := doc.Ref.Delete(ctx); err != nil {
		return err
	}

	return nil
}

// DeleteByUID ユーザーIDから削除する
func (re *CalendarRepository) DeleteByUID(ctx context.Context, f *firestore.Client, uid string) error {

	matchItem := calendarCollectionRef(f, uid).Documents(ctx)
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
