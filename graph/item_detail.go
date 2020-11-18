package graph

import (
	"context"
	"time"

	"github.com/wheatandcat/PeperomiaBackend/domain"
	"github.com/wheatandcat/PeperomiaBackend/graph/model"
)

// CreateItemDetail アイテム詳細を作成
func (g *Graph) CreateItemDetail(ctx context.Context, itemDetail model.NewItemDetail) (*model.ItemDetail, error) {
	h := g.Handler
	uid := g.UID
	loc := GetLoadLocation()

	date, err := time.ParseInLocation("2006-01-02T15:04:05", itemDetail.Date, loc)
	if err != nil {
		return nil, err
	}

	item := domain.ItemDetailRecord{
		ID:       h.Client.UUID.Get(),
		UID:      uid,
		Title:    itemDetail.Title,
		Kind:     itemDetail.Kind,
		Place:    itemDetail.Place,
		URL:      itemDetail.URL,
		Memo:     itemDetail.Memo,
		Priority: itemDetail.Priority,
	}

	itemKey := domain.ItemDetailKey{
		UID:          uid,
		Date:         &date,
		ItemID:       itemDetail.ItemID,
		ItemDetailID: item.ID,
	}

	err = h.App.ItemDetailRepository.Create(ctx, h.FirestoreClient, item, itemKey)
	if err != nil {
		return nil, err
	}

	result := item.ToModel()

	return result, nil
}

// UpdateItemDetail アイテム詳細を更新
func (g *Graph) UpdateItemDetail(ctx context.Context, itemDetail model.UpdateItemDetail) (*model.ItemDetail, error) {
	h := g.Handler
	uid := g.UID
	loc := GetLoadLocation()

	date, err := time.ParseInLocation("2006-01-02T15:04:05", itemDetail.Date, loc)
	if err != nil {
		return nil, err
	}

	idr := domain.ItemDetailRecord{
		ID:       itemDetail.ID,
		UID:      uid,
		Title:    itemDetail.Title,
		Kind:     itemDetail.Kind,
		Place:    itemDetail.Place,
		URL:      itemDetail.URL,
		Memo:     itemDetail.Memo,
		Priority: itemDetail.Priority,
	}

	idrKey := domain.ItemDetailKey{
		UID:          uid,
		Date:         &date,
		ItemID:       itemDetail.ItemID,
		ItemDetailID: idr.ID,
	}

	if err = h.App.ItemDetailRepository.Update(ctx, h.FirestoreClient, idr, idrKey); err != nil {
		return nil, err
	}

	result := idr.ToModel()

	if idr.Priority == 1 {
		// Priority:1の場合はitemも更新する
		item := domain.ItemRecord{
			ID:    itemDetail.ItemID,
			UID:   uid,
			Title: itemDetail.Title,
			Kind:  itemDetail.Kind,
		}
		itemKey := domain.ItemKey{
			UID:  uid,
			Date: &date,
		}

		if err = h.App.ItemRepository.Update(ctx, h.FirestoreClient, item, itemKey); err != nil {
			return nil, err
		}
	}

	return result, nil
}

// DeleteItemDetail アイテム詳細を削除
func (g *Graph) DeleteItemDetail(ctx context.Context, itemDetail model.DeleteItemDetail) (*model.ItemDetail, error) {
	h := g.Handler
	uid := g.UID
	loc := GetLoadLocation()

	date, err := time.ParseInLocation("2006-01-02T15:04:05", itemDetail.Date, loc)
	if err != nil {
		return nil, err
	}

	idrKey := domain.ItemDetailKey{
		UID:          uid,
		Date:         &date,
		ItemID:       itemDetail.ItemID,
		ItemDetailID: itemDetail.ID,
	}

	if err = h.App.ItemDetailRepository.Delete(ctx, h.FirestoreClient, idrKey); err != nil {
		return nil, err
	}

	result := &model.ItemDetail{}

	return result, nil
}

// GetItemDetail アイテム詳細を取得する
func (g *Graph) GetItemDetail(ctx context.Context, date string, itemID string, itemDetailID string) (*model.ItemDetail, error) {
	h := g.Handler
	uid := g.UID
	loc := GetLoadLocation()

	d, err := time.ParseInLocation("2006-01-02T15:04:05", date, loc)
	if err != nil {
		return nil, err
	}

	itemKey := domain.ItemDetailKey{
		UID:          uid,
		Date:         &d,
		ItemID:       itemID,
		ItemDetailID: itemDetailID,
	}

	item, err := h.App.ItemDetailRepository.Get(ctx, h.FirestoreClient, itemKey)
	if err != nil {
		return nil, err
	}

	result := item.ToModel()

	return result, nil
}
