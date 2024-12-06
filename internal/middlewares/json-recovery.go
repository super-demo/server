package middlewares

import (
	"server/infrastructure/app"

	"github.com/gin-gonic/gin"
)

func JSONRecovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if _, ok := recovered.(string); ok {
			c.JSON(app.ErrInternalServerError.ErrHTTPCode, app.ErrInternalServerError.ToJSONResponse())
			c.Abort()
			return
		}
	})
}
