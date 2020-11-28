package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/wheatandcat/PeperomiaBackend/graph/generated"
	"github.com/wheatandcat/PeperomiaBackend/graph/model"
)

func (r *mutationResolver) CreateCalendar(ctx context.Context, calendar model.NewCalendar) (*model.Calendar, error) {
	uid, err := GetSelfUID(ctx)
	if err != nil {
		return nil, err
	}

	g := NewGraph(r.Handler, uid)

	result, err := g.CreateCalendar(ctx, calendar)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *mutationResolver) DeleteCalendar(ctx context.Context, calendar model.DeleteCalendar) (*model.Calendar, error) {
	uid, err := GetSelfUID(ctx)
	if err != nil {
		return nil, err
	}

	g := NewGraph(r.Handler, uid)

	result, err := g.DeleteCalendar(ctx, calendar.Date)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *mutationResolver) UpdateItemDetail(ctx context.Context, itemDetail model.UpdateItemDetail) (*model.ItemDetail, error) {
	uid, err := GetSelfUID(ctx)
	if err != nil {
		return nil, err
	}

	g := NewGraph(r.Handler, uid)

	result, err := g.UpdateItemDetail(ctx, itemDetail)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *mutationResolver) CreateItemDetail(ctx context.Context, itemDetail model.NewItemDetail) (*model.ItemDetail, error) {
	uid, err := GetSelfUID(ctx)
	if err != nil {
		return nil, err
	}

	g := NewGraph(r.Handler, uid)

	result, err := g.CreateItemDetail(ctx, itemDetail)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *mutationResolver) DeleteItemDetail(ctx context.Context, itemDetail model.DeleteItemDetail) (*model.ItemDetail, error) {
	uid, err := GetSelfUID(ctx)
	if err != nil {
		return nil, err
	}

	g := NewGraph(r.Handler, uid)

	result, err := g.DeleteItemDetail(ctx, itemDetail)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *mutationResolver) UpdateMainItemDetail(ctx context.Context, itemDetail model.UpdateMainItemDetail) (*model.ItemDetail, error) {
	uid, err := GetSelfUID(ctx)
	if err != nil {
		return nil, err
	}

	g := NewGraph(r.Handler, uid)

	result, err := g.UpdateMainItem(ctx, itemDetail)
	if err != nil {
		return nil, err
	}

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
	uid, err := GetSelfUID(ctx)
	if err != nil {
		return nil, err
	}

	g := NewGraph(r.Handler, uid)

	result, err := g.GetCalendars(ctx, startDate, endDate)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *queryResolver) Calendar(ctx context.Context, date string) (*model.Calendar, error) {
	uid, err := GetSelfUID(ctx)
	if err != nil {
		return nil, err
	}

	g := NewGraph(r.Handler, uid)

	result, err := g.GetCalendar(ctx, date)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *queryResolver) ItemDetail(ctx context.Context, date string, itemID string, itemDetailID string) (*model.ItemDetail, error) {
	uid, err := GetSelfUID(ctx)
	if err != nil {
		return nil, err
	}

	g := NewGraph(r.Handler, uid)

	result, err := g.GetItemDetail(ctx, date, itemID, itemDetailID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *queryResolver) SuggestionTitle(ctx context.Context, text string) ([]string, error) {
	uid, err := GetSelfUID(ctx)
	if err != nil {
		return nil, err
	}

	if text == "" {
		return nil, nil
	}

	g := NewGraph(r.Handler, uid)

	result, err := g.GetSuggestionText(ctx, text)
	if err != nil {
		return nil, err
	}

	return result, nil
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
