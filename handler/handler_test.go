package handler_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	mock_expopush "github.com/wheatandcat/PeperomiaBackend/backend/client/expo_push/mocks"
	mock_timegen "github.com/wheatandcat/PeperomiaBackend/backend/client/timegen/mocks"
	mock_uuidgen "github.com/wheatandcat/PeperomiaBackend/backend/client/uuidgen/mocks"
	handler "github.com/wheatandcat/PeperomiaBackend/backend/handler"
	repository "github.com/wheatandcat/PeperomiaBackend/backend/repository"
)

func JsonEncode(v interface{}) string {
	b, err := json.Marshal(v)

	if err != nil {
		panic(err)
	}

	return string(b)
}

func NewRequest(body string) *http.Request {
	r, _ := http.NewRequest(http.MethodPost, "/", strings.NewReader(body))
	r.Header.Set("Content-Type", "application/json")
	return r
}

func Execute(hf gin.HandlerFunc, req *http.Request) *httptest.ResponseRecorder {
	res := httptest.NewRecorder()
	c, r := gin.CreateTestContext(res)

	r.Use(func(gc *gin.Context) {
		gc.Set("firebaseUID", "test")
		gc.Next()
	})

	r.POST("/", hf)

	c.Request = req

	r.ServeHTTP(res, c.Request)

	return res
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
