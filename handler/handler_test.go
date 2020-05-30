package handler_test

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
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
	log.Print("0062")

	res := httptest.NewRecorder()

	log.Print("0063")

	c, r := gin.CreateTestContext(res)

	log.Print("0064")

	r.Use(func(gc *gin.Context) {
		gc.Set("firebaseUID", "test")
		gc.Next()
	})

	log.Print("0065")

	r.POST("/", hf)

	log.Print("0066")

	c.Request = req

	log.Print("0067")

	r.ServeHTTP(res, c.Request)

	log.Print("0068")

	return res
}

func NewTestHandler(ctx context.Context) handler.Handler {
	f, _ := firebase.NewApp(ctx, nil)
	fc, _ := f.Firestore(ctx)

	app := &handler.Application{
		ItemRepository:       repository.NewItemRepository(),
		ItemDetailRepository: repository.NewItemDetailRepository(),
		CalendarRepository:   repository.NewCalendarRepository(),
	}

	client := &handler.Client{
		UUID: &mock_uuidgen.UUID{},
	}

	return handler.Handler{
		FirebaseApp:     f,
		FirestoreClient: fc,
		App:             app,
		Client:          client,
	}
}
