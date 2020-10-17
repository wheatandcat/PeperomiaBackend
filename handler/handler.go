package handler

import (
	"context"
	"errors"
	"net/http"
	"time"

	expopush "github.com/wheatandcat/PeperomiaBackend/client/expo_push"
	"github.com/wheatandcat/PeperomiaBackend/client/timegen"
	"github.com/wheatandcat/PeperomiaBackend/client/uuidgen"
	repository "github.com/wheatandcat/PeperomiaBackend/repository"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/wheatandcat/PeperomiaBackend/domain"
)

// Role 役割
const (
	RoleApp        = "app"
	RoleAdmin      = "admin"
	RoleCron       = "cron"
	RoleGraphql    = "graphql"
	RoleAppGraphql = "app/graphql"
)

// Application is app interface
type Application struct {
	ItemRepository       domain.ItemRepository
	ItemDetailRepository domain.ItemDetailRepository
	CalendarRepository   domain.CalendarRepository
	PushTokenRepository  domain.PushTokenRepository
	UserRepository       domain.UserRepository
}

// Client is Client type
type Client struct {
	UUID     uuidgen.UUIDGenerator
	ExpoPush expopush.ExpoPushClientGenerator
	Time     timegen.TimeGenerator
}

// Handler is Handler type
type Handler struct {
	FirebaseApp     *firebase.App
	FirestoreClient *firestore.Client
	App             *Application
	Client          *Client
}

// ErrorResponse is Error Response
type ErrorResponse struct {
	StatusCode int    `json:"-"`
	ErrorCode  string `json:"code"`
	Message    string `json:"message"`
	Error      error  `json:"-"`
}

// NewApplication アプリケーションを作成する
func newApplication() *Application {
	return &Application{
		ItemRepository:       repository.NewItemRepository(),
		ItemDetailRepository: repository.NewItemDetailRepository(),
		CalendarRepository:   repository.NewCalendarRepository(),
		PushTokenRepository:  repository.NewPushTokenRepository(),
		UserRepository:       repository.NewUserRepository(),
	}
}

// NewHandler is Create Handler
func NewHandler(ctx context.Context, f *firebase.App) (*Handler, error) {
	fc, err := f.Firestore(ctx)
	if err != nil {
		h := &Handler{}
		return h, nil
	}

	epc, err := expopush.NewExpoPushClient()
	if err != nil {
		h := &Handler{}
		return h, nil
	}

	client := &Client{
		UUID:     &uuidgen.UUID{},
		ExpoPush: epc,
		Time:     &timegen.Time{},
	}

	app := newApplication()

	return &Handler{
		FirebaseApp:     f,
		FirestoreClient: fc,
		App:             app,
		Client:          client,
	}, nil
}

// GetSelfUID 自身のUIDを取得する
func GetSelfUID(gc *gin.Context) (string, error) {
	fuid, ok := gc.Get("firebaseUID")
	if ok {
		uid, ok := fuid.(string)
		if ok {
			return uid, nil
		}
	}

	return "", errors.New("not uid")
}

func getRole(gc *gin.Context) string {
	r, ok := gc.Get("role")
	if ok {
		role, ok := r.(string)
		if ok {
			return role
		}
	}

	return ""
}

// GetSelfAmazonUID 自身のAmazonUIDを取得する
func GetSelfAmazonUID(gc *gin.Context) (string, error) {
	fuid, ok := gc.Get("amazonUID")
	if ok {
		uid, ok := fuid.(string)
		if ok {
			return uid, nil
		}
	}

	return "", errors.New("not uid")
}

// NewErrorResponse エラーレスポンス作成する
func NewErrorResponse(err error) *ErrorResponse {

	e := &ErrorResponse{
		ErrorCode:  getErrorCode(),
		StatusCode: getStatusCode(),
		Error:      err,
	}

	if err != nil {
		e.Message = err.Error()
	}

	return e
}

// Render 書き込み
func (e *ErrorResponse) Render(gc *gin.Context) {
	if hub := sentrygin.GetHubFromContext(gc); hub != nil {
		hub.WithScope(func(scope *sentry.Scope) {
			r := getRole(gc)
			if r == RoleApp || r == RoleGraphql || r == RoleAdmin {
				// アプリの場合はユーザーIDを追加
				uid, _ := GetSelfUID(gc)
				scope.SetUser(sentry.User{ID: uid})
			}
			hub.Scope().SetTag("role", r)
			hub.CaptureException(e.Error)
		})
	}

	gc.JSON(e.StatusCode, e)
}

func getErrorCode() string {
	return domain.ErrorCodeDefault
}

func getStatusCode() int {
	return http.StatusInternalServerError
}

const location = "Asia/Tokyo"

// TimeNow 現在日時を取得する
func TimeNow() time.Time {
	loc, _ := time.LoadLocation(location)
	date := time.Now().In(loc)
	return date
}

// Day 日付データを取得する
func Day(date time.Time) time.Time {
	loc, _ := time.LoadLocation(location)
	day := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, loc)

	return day
}
