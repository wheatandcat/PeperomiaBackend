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

	mock3 := mock_domain.NewMockItemDetailRepository(ctrl)
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

	item := model.NewItem{
		Title: "test",
		Kind:  "test",
		Memo:  "test",
		URL:   "test",
		Place: "test",
	}
	cm := model.NewCalendar{
		Date: "2019-01-01T00:00:00",
		Item: &item,
	}

	tests := []struct {
		name   string
		param  model.NewCalendar
		result *model.Calendar
	}{
		{
			name:  "カレンダー作成",
			param: cm,
			result: &model.Calendar{
				ID:   "sample-uuid-string",
				Date: "2019-01-01 00:00:00",
				Item: &model.Item{ID: "sample-uuid-string", Title: "test", Kind: "test", ItemDetails: []*model.ItemDetail{
					{
						ID:       "sample-uuid-string",
						Title:    "test",
						Kind:     "test",
						Memo:     "test",
						URL:      "test",
						Place:    "test",
						Priority: 1,
					},
				},
				},
			},
		},
	}

	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			r, _ := g.CreateCalendar(ctx, td.param)
			diff := cmp.Diff(r, td.result)
			if diff != "" {
				t.Errorf("differs: (-got +want)\n%s", diff)
			} else {
				assert.Equal(t, diff, "")
			}
		})
	}
}

func TestUpdateCalendarPublic(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mock1 := mock_domain.NewMockCalendarRepository(ctrl)
	loc := graph.GetLoadLocation()

	date, _ := time.ParseInLocation("2006-01-02T15:04:05", "2019-01-01T00:00:00", loc)

	cr := domain.CalendarRecord{
		ID:     "uuid-string",
		UID:    "test",
		Date:   &date,
		Public: true,
	}

	mock1.EXPECT().FindByDateAndUID(gomock.Any(), gomock.Any(), "test", &date).Return(cr, nil)
	mock1.EXPECT().Update(gomock.Any(), gomock.Any(), cr).Return(nil)

	h := NewTestHandler(ctx)
	h.App.CalendarRepository = mock1

	g := graph.NewGraph(&h, "test")

	cm := &model.Calendar{
		ID:     "uuid-string",
		Date:   "2019-01-01 00:00:00",
		Public: true,
	}

	tests := []struct {
		name   string
		param  model.UpdateCalendarPublic
		result *model.Calendar
	}{
		{
			name: "カレンダーを公開に更新する",
			param: model.UpdateCalendarPublic{
				Date:   "2019-01-01T00:00:00",
				Public: true,
			},
			result: cm,
		},
	}

	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			r, _ := g.UpdateCalendarPublic(ctx, td.param)
			diff := cmp.Diff(r, td.result)
			if diff != "" {
				t.Errorf("differs: (-got +want)\n%s", diff)
			} else {
				assert.Equal(t, diff, "")
			}
		})
	}
}

func TestGetCalendars(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mock1 := mock_domain.NewMockCalendarRepository(ctrl)
	loc := graph.GetLoadLocation()

	startDate, _ := time.ParseInLocation("2006-01-02T15:04:05", "2019-01-01T00:00:00", loc)
	endDate, _ := time.ParseInLocation("2006-01-02T15:04:05", "2019-01-02T00:00:00", loc)

	crs := []domain.CalendarRecord{{
		ID:   "uuid-string",
		Date: &startDate,
		Item: &domain.ItemRecord{
			ID: "uuid-string",
		},
	}}

	mock1.EXPECT().FindBetweenDateAndUID(gomock.Any(), gomock.Any(), "test", &startDate, &endDate).Return(crs, nil)

	h := NewTestHandler(ctx)
	h.App.CalendarRepository = mock1

	g := graph.NewGraph(&h, "test")

	cms := []*model.Calendar{{
		ID:   "uuid-string",
		Date: "2019-01-01 00:00:00",
		Item: &model.Item{ID: "uuid-string"},
	}}

	type paramType struct {
		startDate string
		endDate   string
	}

	tests := []struct {
		name   string
		param  paramType
		result []*model.Calendar
	}{
		{
			name: "期間でカレンダーを取得",
			param: paramType{
				startDate: "2019-01-01T00:00:00",
				endDate:   "2019-01-02T00:00:00",
			},
			result: cms,
		},
	}

	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			r, _ := g.GetCalendars(ctx, td.param.startDate, td.param.endDate)
			diff := cmp.Diff(r, td.result)
			if diff != "" {
				t.Errorf("differs: (-got +want)\n%s", diff)
			} else {
				assert.Equal(t, diff, "")
			}
		})
	}
}

func TestGetCalendar(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mock1 := mock_domain.NewMockCalendarRepository(ctrl)
	loc := graph.GetLoadLocation()

	date, _ := time.ParseInLocation("2006-01-02T15:04:05", "2019-01-01T00:00:00", loc)

	cr := domain.CalendarRecord{
		ID:   "uuid-string",
		Date: &date,
		Item: &domain.ItemRecord{
			ID: "uuid-string",
		},
	}

	mock1.EXPECT().FindByDateAndUID(gomock.Any(), gomock.Any(), "test", &date).Return(cr, nil)

	h := NewTestHandler(ctx)
	h.App.CalendarRepository = mock1

	g := graph.NewGraph(&h, "test")

	cm := &model.Calendar{
		ID:   "uuid-string",
		Date: "2019-01-01 00:00:00",
		Item: &model.Item{ID: "uuid-string"},
	}

	type paramType struct {
		date string
	}

	tests := []struct {
		name   string
		param  paramType
		result *model.Calendar
	}{
		{
			name: "カレンダーを取得",
			param: paramType{
				date: "2019-01-01T00:00:00",
			},
			result: cm,
		},
	}

	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			r, _ := g.GetCalendar(ctx, td.param.date)
			diff := cmp.Diff(r, td.result)
			if diff != "" {
				t.Errorf("differs: (-got +want)\n%s", diff)
			} else {
				assert.Equal(t, diff, "")
			}
		})
	}
}

func TestDeleteCalendar(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mock1 := mock_domain.NewMockCalendarRepository(ctrl)
	loc := graph.GetLoadLocation()

	date, _ := time.ParseInLocation("2006-01-02T15:04:05", "2019-01-01T00:00:00", loc)

	mock1.EXPECT().DeleteByDateAndUID(gomock.Any(), gomock.Any(), "test", &date).Return(nil)

	h := NewTestHandler(ctx)
	h.App.CalendarRepository = mock1

	g := graph.NewGraph(&h, "test")

	cm := &model.Calendar{}

	type paramType struct {
		date string
	}

	tests := []struct {
		name   string
		param  paramType
		result *model.Calendar
	}{
		{
			name: "カレンダーを削除",
			param: paramType{
				date: "2019-01-01T00:00:00",
			},
			result: cm,
		},
	}

	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			r, _ := g.DeleteCalendar(ctx, td.param.date)
			diff := cmp.Diff(r, td.result)
			if diff != "" {
				t.Errorf("differs: (-got +want)\n%s", diff)
			} else {
				assert.Equal(t, diff, "")
			}
		})
	}
}
