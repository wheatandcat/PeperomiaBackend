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

func TestSyncCalendar(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mock1 := mock_domain.NewMockCalendarRepository(ctrl)
	mock2 := mock_domain.NewMockItemRepository(ctrl)
	mock3 := mock_domain.NewMockItemDetailRepository(ctrl)

	loc := graph.GetLoadLocation()

	mock1.EXPECT().DeleteByUID(gomock.Any(), gomock.Any(), "test").Return(nil)

	date, _ := time.ParseInLocation("2006-01-02", "2019-01-01", loc)
	ca := domain.CalendarRecord{
		ID:   "sample-uuid-string",
		UID:  "test",
		Date: &date,
	}

	mock1.EXPECT().Create(gomock.Any(), gomock.Any(), ca).Return(nil)

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

	idr := domain.ItemDetailRecord{
		ID:       "sample-uuid-string",
		UID:      "test",
		Title:    "test",
		Kind:     "test",
		Memo:     "test",
		URL:      "test",
		Place:    "test",
		Priority: 1,
	}
	itemKey := domain.ItemDetailKey{
		UID:          "test",
		Date:         &date,
		ItemID:       "sample-uuid-string",
		ItemDetailID: idr.ID,
	}

	mock3.EXPECT().Create(gomock.Any(), gomock.Any(), idr, itemKey).Return(nil)

	h := NewTestHandler(ctx)
	h.App.CalendarRepository = mock1
	h.App.ItemRepository = mock2
	h.App.ItemDetailRepository = mock3

	g := graph.NewGraph(&h, "test")

	sid := model.SyncItemDetail{
		ID:       "sample-uuid-string",
		Title:    "test",
		Kind:     "test",
		Memo:     "test",
		URL:      "test",
		Place:    "test",
		Priority: 1,
	}

	sids := []*model.SyncItemDetail{&sid}

	si := model.SyncItem{
		ID:          "sample-uuid-string",
		Title:       "test",
		Kind:        "test",
		ItemDetails: sids,
	}

	sc := model.SyncCalendar{
		ID:   "sample-uuid-string",
		Date: "2019-01-01T00:00:00",
		Item: &si,
	}

	scs := []*model.SyncCalendar{&sc}

	mscs := model.SyncCalendars{
		Calendars: scs,
	}

	tests := []struct {
		name   string
		param  model.SyncCalendars
		result bool
	}{
		{
			name:   "カレンダーを同期する",
			param:  mscs,
			result: true,
		},
	}

	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			r, _ := g.SyncCalendar(ctx, td.param)
			diff := cmp.Diff(r, td.result)
			if diff != "" {
				t.Errorf("differs: (-got +want)\n%s", diff)
			} else {
				assert.Equal(t, diff, "")
			}
		})
	}

}
