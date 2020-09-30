package graph

import (
	"context"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/wheatandcat/PeperomiaBackend/domain"
)

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

// IsPublic Publicか判定する
func isPublic(ctx context.Context) bool {
	gc, err := ginContextFromContext(ctx)
	if err != nil {
		return false
	}

	ok := getRole(gc) == "graphql"

	return ok

}
