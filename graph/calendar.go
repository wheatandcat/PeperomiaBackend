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

	itemDetail := domain.ItemDetailRecord{
		ID:       h.Client.UUID.Get(),
		UID:      uid,
		Title:    calendar.Item.Title,
		Kind:     calendar.Item.Kind,
		Place:    calendar.Item.Place,
		URL:      calendar.Item.URL,
		Memo:     calendar.Item.Memo,
		Priority: 1,
	}

	itemDetailKey := domain.ItemDetailKey{
		UID:          uid,
		Date:         &date,
		ItemID:       item.ID,
		ItemDetailID: itemDetail.ID,
	}

	err = h.App.ItemDetailRepository.Create(ctx, h.FirestoreClient, itemDetail, itemDetailKey)
	if err != nil {
		return nil, err
	}

	cr.Item = &item
	cr.Item.ItemDetails = append(cr.Item.ItemDetails, &itemDetail)

	result := cr.ToModel()

	return result, nil
}

// UpdateMainItem メインアイテムに更新する
func (g *Graph) UpdateMainItem(ctx context.Context, umi model.UpdateMainItemDetail) (*model.ItemDetail, error) {
	h := g.Handler
	uid := g.UID
	loc := GetLoadLocation()

	d, err := time.ParseInLocation("2006-01-02T15:04:05", umi.Date, loc)
	if err != nil {
		return nil, err
	}

	cr, err := h.App.CalendarRepository.FindByDateAndUID(ctx, h.FirestoreClient, uid, &d)
	if err != nil {
		return nil, err
	}

	idrKey := domain.ItemDetailKey{
		UID:          uid,
		Date:         &d,
		ItemID:       umi.ItemID,
		ItemDetailID: umi.ID,
	}

	idr, err := h.App.ItemDetailRepository.Get(ctx, h.FirestoreClient, idrKey)
	if err != nil {
		return nil, err
	}

	for _, itemDetail := range cr.Item.ItemDetails {
		if itemDetail.Priority == 1 {
			// メインのアイテムのPriorityを更新
			itemDetail.Priority = idr.Priority

			idrKey1 := domain.ItemDetailKey{
				UID:          uid,
				Date:         &d,
				ItemID:       umi.ItemID,
				ItemDetailID: itemDetail.ID,
			}
			if err = h.App.ItemDetailRepository.Update(ctx, h.FirestoreClient, *itemDetail, idrKey1); err != nil {
				return nil, err
			}
		}
		if itemDetail.Priority == idr.Priority {
			// 選択したアイテムをメインに更新
			itemDetail.Priority = 1

			idrKey1 := domain.ItemDetailKey{
				UID:          uid,
				Date:         &d,
				ItemID:       umi.ItemID,
				ItemDetailID: itemDetail.ID,
			}
			if err = h.App.ItemDetailRepository.Update(ctx, h.FirestoreClient, *itemDetail, idrKey1); err != nil {
				return nil, err
			}
		}

	}

	item := *cr.Item
	item.Title = idr.Title
	item.Kind = idr.Kind
	itemKey := domain.ItemKey{
		UID:  uid,
		Date: &d,
	}
	if err = h.App.ItemRepository.Update(ctx, h.FirestoreClient, item, itemKey); err != nil {
		return nil, err
	}

	idr.Priority = 1
	result := idr.ToModel()

	return result, nil
}

// UpdateCalendarPublic カレンダーの公開/非公開を更新する
func (g *Graph) UpdateCalendarPublic(ctx context.Context, ucp model.UpdateCalendarPublic) (*model.Calendar, error) {
	h := g.Handler
	uid := g.UID
	loc := GetLoadLocation()

	d, err := time.ParseInLocation("2006-01-02T15:04:05", ucp.Date, loc)
	if err != nil {
		return nil, err
	}

	cr, err := h.App.CalendarRepository.FindByDateAndUID(ctx, h.FirestoreClient, uid, &d)
	if err != nil {
		return nil, err
	}

	cr.Public = ucp.Public

	if err := h.App.CalendarRepository.Update(ctx, h.FirestoreClient, cr); err != nil {
		return nil, err
	}

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

// DeleteCalendar カレンダーを削除する
func (g *Graph) DeleteCalendar(ctx context.Context, date string) (*model.Calendar, error) {
	h := g.Handler
	uid := g.UID
	loc := GetLoadLocation()

	d, err := time.ParseInLocation("2006-01-02T15:04:05", date, loc)
	if err != nil {
		return nil, err
	}

	err = h.App.CalendarRepository.DeleteByDateAndUID(ctx, h.FirestoreClient, uid, &d)
	if err != nil {
		return nil, err
	}

	result := &model.Calendar{}

	return result, nil
}
