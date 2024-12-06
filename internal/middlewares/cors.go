package middlewares

import (
	"server/infrastructure/app"
	"strings"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	allowedOrigins := strings.Split(app.Config.ALLOW_ORIGINS, ",")
	return func(ctx *gin.Context) {
		origin := ctx.Request.Header.Get("Origin")

		if Contains(allowedOrigins, origin) {
			ctx.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		} else if Contains(allowedOrigins, "*") {
			ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		}

		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, App-Secret, Content-Disposition")
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
		ctx.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Type, Authorization, Content-Disposition")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}

		ctx.Next()
	}
}

func Contains(list []string, str string) bool {
	for _, v := range list {
		if strings.TrimSpace(v) == str {
			return true
		}
	}
	return false
}
