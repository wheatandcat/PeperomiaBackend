package middleware

import (
	"github.com/gin-gonic/gin"
)

// AdminMiddleWare Admin Middleware
func (m *Middleware) AdminMiddleWare(gc *gin.Context) {
	// FIXME: 後でミドルウェアで制御を追加

	gc.Next()
}
