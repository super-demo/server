package middlewares

import (
	"server/infrastructure/app"
	"server/internal/core/repositories"

	"slices"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AllowedPermissionConfig struct {
	AllowedUserLevelIDs []int
}

func Permission(config AllowedPermissionConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		userLevelIDStr, _ := c.Get("user_level_id")
		userIDStr, _ := c.Get("user_id")
		userLevelID, _ := userLevelIDStr.(int)
		jwtUserID, _ := userIDStr.(int)
		hasUserLevel := slices.Contains(config.AllowedUserLevelIDs, userLevelID)

		if paramUserIDStr, exists := c.Params.Get("user_id"); exists && userLevelID == int(repositories.PeopleUserLevel.UserLevelId) {
			paramUserID, err := strconv.Atoi(paramUserIDStr)
			if err != nil {
				c.AbortWithStatusJSON(app.ErrInvalidUserId.ErrHTTPCode, app.ErrInvalidUserId.ToJSONResponse())
				return
			}

			if jwtUserID != paramUserID {
				c.AbortWithStatusJSON(app.ErrAccessAnotherUser.ErrHTTPCode, app.ErrAccessAnotherUser.ToJSONResponse())
				return
			}
		}

		if hasUserLevel {
			c.Next()
			return
		}

		c.AbortWithStatusJSON(app.ErrInsufficientPermissions.ErrHTTPCode, app.ErrInsufficientPermissions.ToJSONResponse())
	}
}
