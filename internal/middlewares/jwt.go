package middlewares

import (
	"server/infrastructure/app"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	UserId      int    `json:"user_id"`
	UserLevelId int    `json:"user_level_id"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	jwt.StandardClaims
}

func JWT() gin.HandlerFunc {
	jwtKey := []byte(app.Config.Jwt.JwtSecretKey)

	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			ResponseError(c, app.ErrAccessTokenEmpty)
			c.Abort()
			return
		}

		tokenString = ExtractTokenFromHeader(tokenString)

		if tokenString == "" {
			ResponseError(c, app.ErrInvalidToken)
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			ResponseError(c, app.ErrInvalidToken)
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			c.Set("user_id", claims.UserId)
			c.Set("user_level_id", claims.UserLevelId)
			c.Set("email", claims.Email)
			c.Set("name", claims.Name)
			c.Next()
			return
		} else {
			ResponseError(c, app.ErrInvalidToken)
			c.Abort()
			return
		}
	}
}

// extractTokenFromHeader extracts the JWT token from an "Authorization" header.
func ExtractTokenFromHeader(headerValue string) string {
	parts := strings.Split(headerValue, " ")
	if len(parts) == 2 && parts[0] == "Bearer" {
		return parts[1]
	}
	return ""
}
