package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/wheatandcat/PeperomiaBackend/graph/generated"
	"github.com/wheatandcat/PeperomiaBackend/graph/model"
)

func (r *queryResolver) Item(ctx context.Context, id string) (*model.Item, error) {
	h := r.Handler
	item := &model.Item{}

	i, err := h.App.ItemRepository.FindByPublicAndID(ctx, h.FirestoreClient, id)
	if err != nil {
		return item, err
	}

	c, _ := h.App.CalendarRepository.FindByItemID(ctx, h.FirestoreClient, id)

	ids, _ := h.App.ItemDetailRepository.FindByItemID(ctx, h.FirestoreClient, id)

	item = i.ToModel()
	item.Calendar = c.ToModel()

	for _, id := range ids {
		item.ItemDetails = append(item.ItemDetails, id.ToModel())
	}

	return item, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }