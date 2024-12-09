package middlewares

import (
	"server/infrastructure/app"

	"github.com/gin-gonic/gin"
)

func BearerAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")
		if authorizationHeader == "" {
			ResponseError(c, app.ErrAccessTokenEmpty)
			c.Abort()
			return
		}

		token := ExtractTokenFromHeader(authorizationHeader)
		if token == "" {
			ResponseError(c, app.ErrAccessTokenEmpty)
		}
		c.Set("token", token)

		c.Next()
	}
}
