package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/wheatandcat/PeperomiaBackend/domain"
	"github.com/wheatandcat/PeperomiaBackend/graph/generated"
	"github.com/wheatandcat/PeperomiaBackend/graph/model"
)

func (r *mutationResolver) CreateCalendar(ctx context.Context, calendar model.NewCalendar) (*model.Calendar, error) {
	uid, err := GetSelfUID(ctx)
	if err != nil {
		return nil, err
	}

	h := r.Handler
	loc, _ := time.LoadLocation(location)
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

func (r *mutationResolver) CreateItemDetail(ctx context.Context, itemDetail model.NewItemDetail) (*model.ItemDetail, error) {
	uid, err := GetSelfUID(ctx)
	if err != nil {
		return nil, err
	}

	h := r.Handler
	loc, _ := time.LoadLocation(location)
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

func (r *queryResolver) ShareItem(ctx context.Context, id string) (*model.ShareItem, error) {
	item := &model.ShareItem{}

	if !isPublic(ctx) {
		return item, fmt.Errorf("public")
	}

	h := r.Handler
	i, err := h.App.CalendarRepository.FindByPublicAndID(ctx, h.FirestoreClient, id)
	if err != nil {
		return item, err
	}

	item = i.ToShareItemModel()

	return item, nil
}

func (r *queryResolver) User(ctx context.Context) (*model.User, error) {
	user := &model.User{}

	uid, err := GetSelfUID(ctx)
	if err != nil {
		return user, err
	}

	h := r.Handler
	u, err := h.App.UserRepository.FindByUID(ctx, h.FirestoreClient, uid)
	if err != nil {
		return user, err
	}

	user = u.ToModel()

	return user, nil
}

func (r *queryResolver) Calendars(ctx context.Context, startDate string, endDate string) ([]*model.Calendar, error) {
	item := []*model.Calendar{}

	uid, err := GetSelfUID(ctx)
	if err != nil {
		return item, err
	}

	loc, _ := time.LoadLocation(location)
	sd, err := time.ParseInLocation("2006-01-02T15:04:05", startDate, loc)
	if err != nil {
		return item, err
	}
	ed, err := time.ParseInLocation("2006-01-02T15:04:05", endDate, loc)
	if err != nil {
		return item, err
	}

	h := r.Handler
	crs, err := h.App.CalendarRepository.FindBetweenDateAndUID(ctx, h.FirestoreClient, uid, &sd, &ed)
	if err != nil {
		return item, err
	}

	for _, cr := range crs {
		item = append(item, cr.ToModel())
	}

	return item, nil
}

func (r *queryResolver) Calendar(ctx context.Context, date string) (*model.Calendar, error) {
	var item *model.Calendar

	uid, err := GetSelfUID(ctx)
	if err != nil {
		return item, err
	}

	loc, _ := time.LoadLocation(location)
	d, err := time.ParseInLocation("2006-01-02T15:04:05", date, loc)
	if err != nil {
		return item, err
	}

	h := r.Handler
	cr, err := h.App.CalendarRepository.FindByDateAndUID(ctx, h.FirestoreClient, uid, &d)
	if err != nil {
		return item, err
	}

	item = cr.ToModel()

	return item, nil
}

func (r *queryResolver) ItemDetail(ctx context.Context, date string, itemDetailID string) (*model.ItemDetail, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
const location = "Asia/Tokyo"

func (r *queryResolver) Item(ctx context.Context, date string, itemID string) (*model.Item, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *queryResolver) ExpoPushToken(ctx context.Context) (*model.ExpoPushToken, error) {
	panic(fmt.Errorf("not implemented"))
}
