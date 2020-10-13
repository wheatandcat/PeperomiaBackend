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

func TestCreateCalendar(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mock1 := mock_domain.NewMockCalendarRepository(ctrl)
	loc := graph.GetLoadLocation()

	date, _ := time.ParseInLocation("2006-01-02", "2019-01-01", loc)

	ca := domain.CalendarRecord{
		ID:   "sample-uuid-string",
		UID:  "test",
		Date: &date,
	}

	mock1.EXPECT().Create(gomock.Any(), gomock.Any(), ca).Return(nil)

	mock2 := mock_domain.NewMockItemRepository(ctrl)
	i := domain.ItemRecord{
		ID:    "sample-uuid-string",
		UID:   "test",
		Title: "test",
		Kind:  "test",
	}

	key := domain.ItemKey{
		UID:  "test",
		Date: &date,
	}

	mock2.EXPECT().Create(gomock.Any(), gomock.Any(), i, key).Return(nil)

	h := NewTestHandler(ctx)
	h.App.CalendarRepository = mock1
	h.App.ItemRepository = mock2

	g := graph.NewGraph(&h, "test")

	item := model.NewItem{
		Title: "test",
		Kind:  "test",
	}
	cm := model.NewCalendar{
		Date: "2019-01-01T00:00:00",
		Item: &item,
	}

	t.Run("カレンダーを作成する", func(t *testing.T) {
		r, _ := g.CreateCalendar(ctx, cm)
		assert.Equal(t, r.ID, "sample-uuid-string")
	})

}
