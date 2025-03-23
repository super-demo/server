package handlers

import (
	"server/internal/core/models"
	"server/internal/core/repositories"
	"server/internal/core/usecases"
	"server/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type NotificationUserHandler interface{}

type notificationUserHandler struct {
	notificationUserUsecase usecases.NotificationUserUsecase
}

func NewNotificationUserHandler(r *gin.Engine, notificationUserUsecase usecases.NotificationUserUsecase, globalMiddlewares ...gin.HandlerFunc) NotificationUserHandler {
	handler := &notificationUserHandler{notificationUserUsecase}

	v1 := r.Group("/v1/notification-users", globalMiddlewares...)

	updateNotificationUser := []gin.HandlerFunc{
		middlewares.ValidateRequestBody(&models.NotificationUser{}),
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{
				repositories.RootUserLevel.UserLevelId,
				repositories.DeveloperUserLevel.UserLevelId,
				repositories.SuperAdminUserLevel.UserLevelId,
				repositories.AdminUserLevel.UserLevelId,
				repositories.ViewerUserLevel.UserLevelId,
				repositories.PeopleUserLevel.UserLevelId,
			},
		}),
		handler.UpdateNotificationUser,
	}

	v1.PUT("update", updateNotificationUser...)

	return handler
}

func (h *notificationUserHandler) UpdateNotificationUser(c *gin.Context) {
	notificationUser := &models.NotificationUser{}
	if err := c.ShouldBindJSON(notificationUser); err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	notificationUser, err := h.notificationUserUsecase.UpdateNotificationUser(notificationUser)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, notificationUser, "Notification user updated successfully")
}
