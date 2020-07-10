package middleware

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wheatandcat/PeperomiaBackend/domain"
	repository "github.com/wheatandcat/PeperomiaBackend/repository"
)

// AdminMiddleWare Admin Middleware
func (m *Middleware) AdminMiddleWare(gc *gin.Context) {

	ctx := context.Background()
	fuid, _ := gc.Get("firebaseUID")
	uid := fuid.(string)
	ur := repository.NewUserRepository()
	fc, err := m.FirebaseApp.Firestore(ctx)
	if err != nil {
		gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	u, _ := ur.FindByUID(ctx, fc, uid)
	if u.Role != domain.UserRoleAdmin {
		gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "role is not admin"})
		return
	}

	gc.Next()
}
