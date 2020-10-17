package graph

import (
	"context"
	"time"

	"github.com/wheatandcat/PeperomiaBackend/domain"
	"github.com/wheatandcat/PeperomiaBackend/graph/model"
)

// CreateCalendar カレンダーを作成
func (g *Graph) CreateCalendar(ctx context.Context, calendar model.NewCalendar) (*model.Calendar, error) {
	h := g.Handler
	uid := g.UID
	loc := GetLoadLocation()

	date, err := time.ParseInLocation("2006-01-02T15:04:05", calendar.Date, loc)
	if err != nil {
		return nil, err
	}

	cr := &domain.CalendarRecord{
		ID:   h.Client.UUID.Get(),
		UID:  uid,
		Date: &date,
	}

	err = h.App.CalendarRepository.Create(ctx, h.FirestoreClient, *cr)
	if err != nil {
		return nil, err
	}
	item := domain.ItemRecord{
		ID:    h.Client.UUID.Get(),
		UID:   uid,
		Title: calendar.Item.Title,
		Kind:  calendar.Item.Kind,
	}
	itemKey := domain.ItemKey{
		UID:  uid,
		Date: &date,
	}

	err = h.App.ItemRepository.Create(ctx, h.FirestoreClient, item, itemKey)
	if err != nil {
		return nil, err
	}
	cr.Item = &item

	result := cr.ToModel()

	return result, nil
}

// GetCalendars カレンダーリストを取得する
func (g *Graph) GetCalendars(ctx context.Context, startDate string, endDate string) ([]*model.Calendar, error) {
	h := g.Handler
	uid := g.UID
	loc := GetLoadLocation()

	item := []*model.Calendar{}
	sd, err := time.ParseInLocation("2006-01-02T15:04:05", startDate, loc)
	if err != nil {
		return item, err
	}
	ed, err := time.ParseInLocation("2006-01-02T15:04:05", endDate, loc)
	if err != nil {
		return item, err
	}

	crs, err := h.App.CalendarRepository.FindBetweenDateAndUID(ctx, h.FirestoreClient, uid, &sd, &ed)
	if err != nil {
		return item, err
	}

	for _, cr := range crs {
		item = append(item, cr.ToModel())
	}

	return item, nil
}

// GetCalendar カレンダーを取得する
func (g *Graph) GetCalendar(ctx context.Context, date string) (*model.Calendar, error) {
	h := g.Handler
	uid := g.UID
	loc := GetLoadLocation()

	d, err := time.ParseInLocation("2006-01-02T15:04:05", date, loc)
	if err != nil {
		return nil, err
	}

	cr, err := h.App.CalendarRepository.FindByDateAndUID(ctx, h.FirestoreClient, uid, &d)
	if err != nil {
		return nil, err
	}

	item := cr.ToModel()

	return item, nil
}
