package handlers

import (
	"server/internal/core/models"
	"server/internal/core/repositories"
	"server/internal/core/usecases"
	"server/internal/middlewares"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SiteMiniAppUserHandler interface{}

type siteMiniAppUserHandler struct {
	siteMiniAppUserHandler usecases.SiteMiniAppUserUsecase
}

func NewSiteMiniAppUserHandler(r *gin.Engine, siteMiniAppUserUsecase usecases.SiteMiniAppUserUsecase, globalMiddlewares ...gin.HandlerFunc) SiteMiniAppUserHandler {
	handler := &siteMiniAppUserHandler{siteMiniAppUserUsecase}

	v1 := r.Group("/v1/site-mini-app-users", globalMiddlewares...)

	createSiteMiniAppUser := []gin.HandlerFunc{
		middlewares.ValidateRequestBody([]models.SiteMiniAppUser{}),
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{repositories.RootUserLevel.UserLevelId, repositories.DeveloperUserLevel.UserLevelId},
		}),
		handler.CreateSiteMiniAppUser,
	}

	getListSiteMiniAppUserBySiteId := []gin.HandlerFunc{
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{repositories.RootUserLevel.UserLevelId, repositories.SuperAdminUserLevel.UserLevelId, repositories.AdminUserLevel.UserLevelId},
		}),
		handler.GetListSiteMiniAppUserBySiteId,
	}

	deleteSiteMiniAppUserBySiteIdAndUserId := []gin.HandlerFunc{
		middlewares.ValidateRequestBody(&models.SiteMiniAppUser{}),
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{repositories.RootUserLevel.UserLevelId, repositories.SuperAdminUserLevel.UserLevelId, repositories.AdminUserLevel.UserLevelId},
		}),
		handler.DeleteSiteMiniAppUserBySiteIdAndUserId,
	}

	v1.POST("/create", createSiteMiniAppUser...)
	v1.GET("/list/:siteId", getListSiteMiniAppUserBySiteId...)
	v1.DELETE("/delete", deleteSiteMiniAppUserBySiteIdAndUserId...)

	return handler
}

func (h *siteMiniAppUserHandler) CreateSiteMiniAppUser(c *gin.Context) {
	siteMiniAppUser := []models.SiteMiniAppUser{}
	if err := c.ShouldBindJSON(siteMiniAppUser); err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	siteMiniAppUser, err := h.siteMiniAppUserHandler.CreateSiteMiniAppUser(siteMiniAppUser)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, siteMiniAppUser, "Site mini app user created successfully")
}

func (h *siteMiniAppUserHandler) GetListSiteMiniAppUserBySiteId(c *gin.Context) {
	siteId, err := strconv.Atoi(c.Param("siteId"))
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	siteMiniAppUsers, err := h.siteMiniAppUserHandler.GetListSiteMiniAppUserBySiteId(siteId)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, siteMiniAppUsers, "Site mini app user retrieved successfully")
}

func (h *siteMiniAppUserHandler) DeleteSiteMiniAppUserBySiteIdAndUserId(c *gin.Context) {
	siteMiniAppUser := []models.SiteMiniAppUser{}
	if err := c.ShouldBindJSON(siteMiniAppUser); err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	err := h.siteMiniAppUserHandler.DeleteSiteMiniAppUserBySiteIdAndUserId(siteMiniAppUser)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, nil, "Site mini app user deleted successfully")
}
