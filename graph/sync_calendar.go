package graph

import (
	"context"
	"time"

	"github.com/wheatandcat/PeperomiaBackend/domain"
	"github.com/wheatandcat/PeperomiaBackend/graph/model"
)

// SyncCalendar アイテムを同期する
func (g *Graph) SyncCalendar(ctx context.Context, calendars model.SyncCalendars) (bool, error) {
	h := g.Handler
	uid := g.UID
	loc := GetLoadLocation()

	if err := h.App.CalendarRepository.DeleteByUID(ctx, h.FirestoreClient, uid); err != nil {
		return false, nil
	}

	for _, c := range calendars.Calendars {
		date, err := time.ParseInLocation("2006-01-02T15:04:05", c.Date, loc)
		if err != nil {
			return false, err
		}
		cr := &domain.CalendarRecord{
			ID:   h.Client.UUID.Get(),
			UID:  uid,
			Date: &date,
		}
		err = h.App.CalendarRepository.Create(ctx, h.FirestoreClient, *cr)
		if err != nil {
			return false, err
		}
		item := domain.ItemRecord{
			ID:    h.Client.UUID.Get(),
			UID:   uid,
			Title: c.Item.Title,
			Kind:  c.Item.Kind,
		}
		itemKey := domain.ItemKey{
			UID:  uid,
			Date: &date,
		}
		err = h.App.ItemRepository.Create(ctx, h.FirestoreClient, item, itemKey)
		if err != nil {
			return false, err
		}

		for _, i := range c.Item.ItemDetails {
			date, err := time.ParseInLocation("2006-01-02T15:04:05", c.Date, loc)
			if err != nil {
				return false, err
			}

			itemDetail := domain.ItemDetailRecord{
				ID:       h.Client.UUID.Get(),
				UID:      uid,
				Title:    i.Title,
				Kind:     i.Kind,
				Place:    i.Place,
				URL:      i.URL,
				Memo:     i.Memo,
				Priority: i.Priority,
			}

			itemDetailKey := domain.ItemDetailKey{
				UID:          uid,
				Date:         &date,
				ItemID:       item.ID,
				ItemDetailID: itemDetail.ID,
			}

			err = h.App.ItemDetailRepository.Create(ctx, h.FirestoreClient, itemDetail, itemDetailKey)
			if err != nil {
				return false, err
			}
		}
	}

	return true, nil
}
