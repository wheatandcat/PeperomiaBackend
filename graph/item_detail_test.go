package graph_test

import (
	"context"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
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
		ID:          "sample-uuid-string",
		UID:         "test",
		ItemID:      "ItemID",
		Title:       "Title",
		Kind:        "Kind",
		MoveMinutes: 0,
		Place:       "Place",
		URL:         "URL",
		Memo:        "Memo",
		Priority:    1,
	}
	itemKey := domain.ItemDetailKey{
		UID:    "test",
		Date:   &date,
		ItemID: "ItemID",
	}

	mock1.EXPECT().Create(gomock.Any(), gomock.Any(), idr, itemKey).Return(nil)

	h := NewTestHandler(ctx)
	h.App.ItemDetailRepository = mock1

	g := graph.NewGraph(&h, "test")

	nid := model.NewItemDetail{
		Date:        "2019-01-01T00:00:00",
		ItemID:      "ItemID",
		Title:       "Title",
		Kind:        "Kind",
		MoveMinutes: 0,
		Place:       "Place",
		URL:         "URL",
		Memo:        "Memo",
		Priority:    1,
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
				ID: "sample-uuid-string",
			},
		},
	}

	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			r, _ := g.CreateItemDetail(ctx, td.param)
			assert.Equal(t, td.result.ID, r.ID)
		})
	}

}
