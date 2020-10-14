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
		ID:          h.Client.UUID.Get(),
		UID:         uid,
		ItemID:      itemDetail.ItemID,
		Title:       itemDetail.Title,
		Kind:        itemDetail.Kind,
		MoveMinutes: itemDetail.MoveMinutes,
		Place:       itemDetail.Place,
		URL:         itemDetail.URL,
		Memo:        itemDetail.Memo,
		Priority:    itemDetail.Priority,
	}

	itemKey := domain.ItemDetailKey{
		UID:    uid,
		Date:   &date,
		ItemID: itemDetail.ItemID,
	}

	err = h.App.ItemDetailRepository.Create(ctx, h.FirestoreClient, item, itemKey)
	if err != nil {
		return nil, err
	}

	result := item.ToModel()

	return result, nil
}
