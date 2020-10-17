package graph_test

import (
	"context"

	firebase "firebase.google.com/go"
	mock_expopush "github.com/wheatandcat/PeperomiaBackend/client/expo_push/mocks"
	mock_timegen "github.com/wheatandcat/PeperomiaBackend/client/timegen/mocks"
	mock_uuidgen "github.com/wheatandcat/PeperomiaBackend/client/uuidgen/mocks"
	graph "github.com/wheatandcat/PeperomiaBackend/graph"
	"github.com/wheatandcat/PeperomiaBackend/graph/generated"
	handler "github.com/wheatandcat/PeperomiaBackend/handler"
	"github.com/wheatandcat/PeperomiaBackend/repository"
)

func NewResolver(ctx context.Context) generated.Config {
	h := NewTestHandler(ctx)
	r := graph.Resolver{
		Handler: &h,
	}

	return generated.Config{
		Resolvers: &r,
	}
}

func NewTestHandler(ctx context.Context) handler.Handler {
	config := &firebase.Config{ProjectID: "my-test-project"}
	f, _ := firebase.NewApp(ctx, config)

	app := &handler.Application{
		ItemRepository:       repository.NewItemRepository(),
		ItemDetailRepository: repository.NewItemDetailRepository(),
		CalendarRepository:   repository.NewCalendarRepository(),
		PushTokenRepository:  repository.NewPushTokenRepository(),
		UserRepository:       repository.NewUserRepository(),
	}

	client := &handler.Client{
		UUID:     &mock_uuidgen.UUID{},
		ExpoPush: &mock_expopush.ExpoPushClient{},
		Time:     &mock_timegen.Time{},
	}

	return handler.Handler{
		FirebaseApp:     f,
		FirestoreClient: nil,
		App:             app,
		Client:          client,
	}
}
