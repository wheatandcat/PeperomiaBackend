package graph_test

import (
	"context"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/wheatandcat/PeperomiaBackend/domain"
	mock_domain "github.com/wheatandcat/PeperomiaBackend/domain/mocks"
	graph "github.com/wheatandcat/PeperomiaBackend/graph"
	"github.com/wheatandcat/PeperomiaBackend/graph/model"
	"gopkg.in/go-playground/assert.v1"
)

func TestCreateItemDetail(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mock1 := mock_domain.NewMockItemDetailRepository(ctrl)
	loc := graph.GetLoadLocation()

	date, _ := time.ParseInLocation("2006-01-02", "2019-01-01", loc)

	idr := domain.ItemDetailRecord{
		ID:       "sample-uuid-string",
		UID:      "test",
		Title:    "Title",
		Kind:     "Kind",
		Place:    "Place",
		URL:      "URL",
		Memo:     "Memo",
		Priority: 1,
	}
	itemKey := domain.ItemDetailKey{
		UID:          "test",
		Date:         &date,
		ItemID:       "ItemID",
		ItemDetailID: idr.ID,
	}

	mock1.EXPECT().Create(gomock.Any(), gomock.Any(), idr, itemKey).Return(nil)

	h := NewTestHandler(ctx)
	h.App.ItemDetailRepository = mock1

	g := graph.NewGraph(&h, "test")

	nid := model.NewItemDetail{
		Date:     "2019-01-01T00:00:00",
		ItemID:   "ItemID",
		Title:    "Title",
		Kind:     "Kind",
		Place:    "Place",
		URL:      "URL",
		Memo:     "Memo",
		Priority: 1,
	}

	tests := []struct {
		name   string
		param  model.NewItemDetail
		result *model.ItemDetail
	}{
		{
			name:  "アイテム詳細を作成",
			param: nid,
			result: &model.ItemDetail{
				ID:       "sample-uuid-string",
				Title:    "Title",
				Kind:     "Kind",
				Place:    "Place",
				URL:      "URL",
				Memo:     "Memo",
				Priority: 1,
			},
		},
	}

	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			r, _ := g.CreateItemDetail(ctx, td.param)
			diff := cmp.Diff(r, td.result)
			if diff != "" {
				t.Errorf("differs: (-got +want)\n%s", diff)
			} else {
				assert.Equal(t, diff, "")
			}
		})
	}

}

func TestGetItemDetail(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mock1 := mock_domain.NewMockItemDetailRepository(ctrl)
	loc := graph.GetLoadLocation()

	date, _ := time.ParseInLocation("2006-01-02T15:04:05", "2019-01-01T00:00:00", loc)

	idr := domain.ItemDetailRecord{
		ID:       "itemDetailID",
		UID:      "test",
		Title:    "Title",
		Kind:     "Kind",
		Place:    "Place",
		URL:      "URL",
		Memo:     "Memo",
		Priority: 1,
	}

	itemKey := domain.ItemDetailKey{
		UID:          "test",
		Date:         &date,
		ItemID:       "ItemID",
		ItemDetailID: idr.ID,
	}

	mock1.EXPECT().Get(gomock.Any(), gomock.Any(), itemKey).Return(idr, nil)

	h := NewTestHandler(ctx)
	h.App.ItemDetailRepository = mock1

	g := graph.NewGraph(&h, "test")

	mid := &model.ItemDetail{
		ID:       "itemDetailID",
		Title:    "Title",
		Kind:     "Kind",
		Place:    "Place",
		URL:      "URL",
		Memo:     "Memo",
		Priority: 1,
	}

	type paramType struct {
		date         string
		itemID       string
		itemDetailID string
	}

	tests := []struct {
		name   string
		param  paramType
		result *model.ItemDetail
	}{
		{
			name: "アイテム詳細を取得",
			param: paramType{
				date:         "2019-01-01T00:00:00",
				itemID:       "ItemID",
				itemDetailID: "itemDetailID",
			},
			result: mid,
		},
	}

	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			r, _ := g.GetItemDetail(ctx, td.param.date, td.param.itemID, td.param.itemDetailID)
			diff := cmp.Diff(r, td.result)
			if diff != "" {
				t.Errorf("differs: (-got +want)\n%s", diff)
			} else {
				assert.Equal(t, diff, "")
			}
		})
	}
}

func TestUpdateItemDetail(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mock1 := mock_domain.NewMockItemDetailRepository(ctrl)
	mock2 := mock_domain.NewMockItemRepository(ctrl)
	loc := graph.GetLoadLocation()

	date, _ := time.ParseInLocation("2006-01-02", "2019-01-01", loc)

	idr := domain.ItemDetailRecord{
		ID:       "sample-uuid-string",
		UID:      "test",
		Title:    "Title",
		Kind:     "Kind",
		Place:    "Place",
		URL:      "URL",
		Memo:     "Memo",
		Priority: 1,
	}
	idrKey := domain.ItemDetailKey{
		UID:          "test",
		Date:         &date,
		ItemID:       "ItemID",
		ItemDetailID: idr.ID,
	}

	mock1.EXPECT().Update(gomock.Any(), gomock.Any(), idr, idrKey).Return(nil)

	item := domain.ItemRecord{
		ID:    "ItemID",
		UID:   "test",
		Title: "Title",
		Kind:  "Kind",
	}
	itemKey := domain.ItemKey{
		UID:  "test",
		Date: &date,
	}

	mock2.EXPECT().Update(gomock.Any(), gomock.Any(), item, itemKey).Return(nil)

	h := NewTestHandler(ctx)
	h.App.ItemDetailRepository = mock1
	h.App.ItemRepository = mock2

	g := graph.NewGraph(&h, "test")

	nid := model.UpdateItemDetail{
		ID:       "sample-uuid-string",
		Date:     "2019-01-01T00:00:00",
		ItemID:   "ItemID",
		Title:    "Title",
		Kind:     "Kind",
		Place:    "Place",
		URL:      "URL",
		Memo:     "Memo",
		Priority: 1,
	}

	tests := []struct {
		name   string
		param  model.UpdateItemDetail
		result *model.ItemDetail
	}{
		{
			name:  "アイテム詳細を更新",
			param: nid,
			result: &model.ItemDetail{
				ID:       "sample-uuid-string",
				Title:    "Title",
				Kind:     "Kind",
				Place:    "Place",
				URL:      "URL",
				Memo:     "Memo",
				Priority: 1,
			},
		},
	}

	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			r, _ := g.UpdateItemDetail(ctx, td.param)
			diff := cmp.Diff(r, td.result)
			if diff != "" {
				t.Errorf("differs: (-got +want)\n%s", diff)
			} else {
				assert.Equal(t, diff, "")
			}
		})
	}

}

func TestDeleteItemDetail(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mock1 := mock_domain.NewMockItemDetailRepository(ctrl)
	loc := graph.GetLoadLocation()

	date, _ := time.ParseInLocation("2006-01-02", "2019-01-01", loc)

	idr := domain.ItemDetailRecord{
		ID:       "sample-uuid-string",
		UID:      "test",
		Title:    "Title",
		Kind:     "Kind",
		Place:    "Place",
		URL:      "URL",
		Memo:     "Memo",
		Priority: 1,
	}
	idrKey := domain.ItemDetailKey{
		UID:          "test",
		Date:         &date,
		ItemID:       "ItemID",
		ItemDetailID: idr.ID,
	}

	mock1.EXPECT().Delete(gomock.Any(), gomock.Any(), idrKey).Return(nil)

	h := NewTestHandler(ctx)
	h.App.ItemDetailRepository = mock1

	g := graph.NewGraph(&h, "test")

	nid := model.DeleteItemDetail{
		ID:     "sample-uuid-string",
		Date:   "2019-01-01T00:00:00",
		ItemID: "ItemID",
	}

	tests := []struct {
		name   string
		param  model.DeleteItemDetail
		result *model.ItemDetail
	}{
		{
			name:  "アイテム詳細を削除",
			param: nid,
			result: &model.ItemDetail{
				ID:       "",
				Title:    "",
				Kind:     "",
				Place:    "",
				URL:      "",
				Memo:     "",
				Priority: 0,
			},
		},
	}

	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			r, _ := g.DeleteItemDetail(ctx, td.param)
			diff := cmp.Diff(r, td.result)
			if diff != "" {
				t.Errorf("differs: (-got +want)\n%s", diff)
			} else {
				assert.Equal(t, diff, "")
			}
		})
	}

}
