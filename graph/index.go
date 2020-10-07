package graph

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wheatandcat/PeperomiaBackend/domain"
	"github.com/wheatandcat/PeperomiaBackend/handler"
)

// Graph Graph struct
type Graph struct {
	Handler *handler.Handler
	UID     string
}

// NewGraph Graphを作成
func NewGraph(h *handler.Handler, uid string) *Graph {
	return &Graph{
		Handler: h,
		UID:     uid,
	}
}

// GetLoadLocation ロケーションを取得する
func GetLoadLocation() *time.Location {
	const location = "Asia/Tokyo"
	loc, _ := time.LoadLocation(location)
	return loc
}

// GetSelfUID 自身のUIDを取得する
func GetSelfUID(ctx context.Context) (string, error) {
	if isPublic(ctx) {
		return "", fmt.Errorf("not public")
	}

	gc, err := ginContextFromContext(ctx)
	if err != nil {
		return "", err
	}

	fuid, ok := gc.Get("firebaseUID")
	if ok {
		uid, ok := fuid.(string)
		if ok {
			return uid, nil
		}
	}

	return "", errors.New("not uid")
}

func ginContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value(domain.GinContextKey)
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}

func getPublic(gc *gin.Context) bool {
	r, ok := gc.Get("public")
	if ok {
		role, ok := r.(bool)
		if ok {
			return role
		}
	}

	return false
}

// IsPublic Publicか判定する
func isPublic(ctx context.Context) bool {
	gc, err := ginContextFromContext(ctx)
	if err != nil {
		return false
	}

	return getPublic(gc)

}
