package middlewares

import (
	"server/infrastructure/app"

	"github.com/gin-gonic/gin"
)

func ValidateAppSecret() gin.HandlerFunc {
	return func(c *gin.Context) {
		var appSecret string = c.GetHeader("App-Secret")
		if appSecret != app.Config.AppSecret {
			ResponseError(c, app.ErrInvalidAppSecret)
			c.Abort()
			return
		}
		c.Next()
	}
}
