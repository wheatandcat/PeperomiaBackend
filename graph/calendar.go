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
