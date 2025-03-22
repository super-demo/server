package handlers

import (
	"server/internal/core/models"
	"server/internal/core/repositories"
	"server/internal/core/usecases"
	"server/internal/middlewares"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AnnouncementHandler interface{}

type announcementHandler struct {
	announcementUsecase usecases.AnnouncementUsecase
}

func NewAnnouncementHandler(r *gin.Engine, announcementUsecase usecases.AnnouncementUsecase, globalMiddlewares ...gin.HandlerFunc) AnnouncementHandler {
	handler := &announcementHandler{announcementUsecase}

	v1 := r.Group("/v1/announcements", globalMiddlewares...)

	createAnnouncement := []gin.HandlerFunc{
		middlewares.ValidateRequestBody(&models.Announcement{}),
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{
				repositories.RootUserLevel.UserLevelId,
				repositories.SuperAdminUserLevel.UserLevelId,
				repositories.AdminUserLevel.UserLevelId,
			},
		}),
		handler.CreateAnnouncement,
	}

	getListAnnouncementBySiteId := []gin.HandlerFunc{
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
		handler.GetListAnnouncementBySiteId,
	}

	getAnnouncementById := []gin.HandlerFunc{
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
		handler.GetAnnouncementById,
	}

	updateAnnouncement := []gin.HandlerFunc{
		middlewares.ValidateRequestBody(&models.Announcement{}),
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{
				repositories.RootUserLevel.UserLevelId,
				repositories.SuperAdminUserLevel.UserLevelId,
				repositories.AdminUserLevel.UserLevelId,
			},
		}),
		handler.UpdateAnnouncement,
	}

	deleteAnnouncement := []gin.HandlerFunc{
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{
				repositories.RootUserLevel.UserLevelId,
				repositories.SuperAdminUserLevel.UserLevelId,
				repositories.AdminUserLevel.UserLevelId,
			},
		}),
		handler.DeleteAnnouncement,
	}

	v1.POST("/create", createAnnouncement...)
	v1.GET("/:id", getAnnouncementById...)
	v1.GET("/list/:id", getListAnnouncementBySiteId...)
	v1.PUT("/update", updateAnnouncement...)
	v1.DELETE("/delete/:id", deleteAnnouncement...)

	return handler
}

func (h *announcementHandler) CreateAnnouncement(c *gin.Context) {
	announcement := &models.Announcement{}
	if err := c.ShouldBindJSON(announcement); err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	requesterUserId := c.MustGet("user_id").(int)

	announcement, err := h.announcementUsecase.CreateAnnouncement(announcement, requesterUserId)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, announcement, "Create announcement successfully")
}

func (h *announcementHandler) GetAnnouncementById(c *gin.Context) {
	announcementIdStr := c.Param("id")
	announcementId, err := strconv.Atoi(announcementIdStr)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	announcement, err := h.announcementUsecase.GetAnnouncementById(announcementId)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, announcement, "Get announcement successfully")
}

func (h *announcementHandler) GetListAnnouncementBySiteId(c *gin.Context) {
	siteIdStr := c.Param("id")
	siteId, err := strconv.Atoi(siteIdStr)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	announcements, err := h.announcementUsecase.GetListAnnouncementBySiteId(siteId)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, announcements, "Get list announcement successfully")
}

func (h *announcementHandler) UpdateAnnouncement(c *gin.Context) {
	announcement := &models.Announcement{}
	if err := c.ShouldBindJSON(announcement); err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	announcement, err := h.announcementUsecase.UpdateAnnouncement(announcement)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, announcement, "Update announcement successfully")
}

func (h *announcementHandler) DeleteAnnouncement(c *gin.Context) {
	announcementIdStr := c.Param("id")
	announcementId, err := strconv.Atoi(announcementIdStr)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	err = h.announcementUsecase.DeleteAnnouncement(announcementId)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, nil, "Delete announcement successfully")
}
