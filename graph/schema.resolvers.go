package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/wheatandcat/PeperomiaBackend/graph/generated"
	"github.com/wheatandcat/PeperomiaBackend/graph/model"
)

func (r *queryResolver) ShareItem(ctx context.Context, id string) (*model.ShareItem, error) {
	h := r.Handler
	item := &model.ShareItem{}

	if isPublic(ctx) {
		return item, fmt.Errorf("not public")
	}

	i, err := h.App.CalendarRepository.FindByPublicAndID(ctx, h.FirestoreClient, id)
	if err != nil {
		return item, err
	}

	item = i.ToShareItemModel()

	return item, nil
}

func (r *queryResolver) User(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Calendars(ctx context.Context, startDate string, endDate string) ([]*model.Calendar, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Item(ctx context.Context, itemID string) (*model.Item, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ItemDetail(ctx context.Context, itemDetailID string) (*model.ItemDetail, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ExpoPushToken(ctx context.Context) (*model.ExpoPushToken, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
