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

	if isPublic(ctx) {
		return user, fmt.Errorf("not public")
	}

	h := r.Handler
	uid, err := GetSelfUID(ctx)
	if err != nil {
		return user, err
	}

	u, err := h.App.UserRepository.FindByUID(ctx, h.FirestoreClient, uid)
	if err != nil {
		return user, err
	}

	user = u.ToModel()

	return user, nil
}

func (r *queryResolver) Calendars(ctx context.Context, startDate string, endDate string) ([]*model.Calendar, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Calendar(ctx context.Context, date string) (*model.Calendar, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Item(ctx context.Context, date string, itemID string) (*model.Item, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ItemDetail(ctx context.Context, date string, itemDetailID string) (*model.ItemDetail, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) ExpoPushToken(ctx context.Context) (*model.ExpoPushToken, error) {
	panic(fmt.Errorf("not implemented"))
}
