package handlers

import (
	"server/internal/core/models"
	"server/internal/core/repositories"
	"server/internal/core/usecases"
	"server/internal/middlewares"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NotificationHandler interface{}

type notificationHandler struct {
	notificationUsecase usecases.NotificationUsecase
}

func NewNotificationHandler(r *gin.Engine, notificationUsecase usecases.NotificationUsecase, globalMiddlewares ...gin.HandlerFunc) NotificationHandler {
	handler := &notificationHandler{notificationUsecase}

	v1 := r.Group("/v1/notifications", globalMiddlewares...)

	createNotification := []gin.HandlerFunc{
		middlewares.ValidateRequestBody(&models.Notification{}),
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{
				repositories.RootUserLevel.UserLevelId,
				repositories.DeveloperUserLevel.UserLevelId,
				repositories.AdminUserLevel.UserLevelId,
				repositories.ViewerUserLevel.UserLevelId,
				repositories.PeopleUserLevel.UserLevelId,
			},
		}),
		handler.CreateNotification,
	}

	getNotificationById := []gin.HandlerFunc{
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{
				repositories.RootUserLevel.UserLevelId,
				repositories.DeveloperUserLevel.UserLevelId,
				repositories.AdminUserLevel.UserLevelId,
				repositories.ViewerUserLevel.UserLevelId,
				repositories.PeopleUserLevel.UserLevelId,
			},
		}),
		handler.GetNotificationById,
	}

	getListNotificationByUserId := []gin.HandlerFunc{
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{
				repositories.RootUserLevel.UserLevelId,
				repositories.DeveloperUserLevel.UserLevelId,
				repositories.AdminUserLevel.UserLevelId,
				repositories.ViewerUserLevel.UserLevelId,
				repositories.PeopleUserLevel.UserLevelId,
			},
		}),
		handler.GetListNotificationByUserId,
	}

	v1.POST("/create", createNotification...)
	v1.GET("/:id", getNotificationById...)
	v1.GET("/list", getListNotificationByUserId...)

	return handler
}

func (h *notificationHandler) CreateNotification(c *gin.Context) {
	notification := &models.Notification{}
	if err := c.ShouldBindJSON(&notification); err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	notification, err := h.notificationUsecase.CreateNotification(notification)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, notification, "Notification created successfully")
}

func (h *notificationHandler) GetNotificationById(c *gin.Context) {
	notificationIdStr := c.Param("id")
	notificationId, err := strconv.Atoi(notificationIdStr)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	notification, err := h.notificationUsecase.GetNotificationById(notificationId)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, notification, "Notification retrieved successfully")
}

func (h *notificationHandler) GetListNotificationByUserId(c *gin.Context) {
	requesterUserId := c.MustGet("user_id").(int)

	notifications, err := h.notificationUsecase.GetListNotificationByUserId(requesterUserId)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, notifications, "List of notifications retrieved successfully")
}
