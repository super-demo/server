package handlers

import (
	"server/internal/core/models"
	"server/internal/core/repositories"
	"server/internal/core/usecases"
	"server/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type SiteTypeHandler interface{}

type siteTypeHandler struct {
	siteTypeUsecase usecases.SiteTypeUsecase
}

func NewSiteTypeHandler(r *gin.Engine, siteTypeUsecase usecases.SiteTypeUsecase, globalMiddlewares ...gin.HandlerFunc) SiteTypeHandler {
	handler := &siteTypeHandler{siteTypeUsecase}

	v1 := r.Group("/v1/site-types", globalMiddlewares...)

	createSiteType := []gin.HandlerFunc{
		middlewares.ValidateRequestBody(&models.SiteType{}),
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{repositories.RootUserLevel.UserLevelId, repositories.StaffUserLevel.UserLevelId},
		}),
		handler.CreateSiteType,
	}

	getListSiteType := []gin.HandlerFunc{
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{repositories.RootUserLevel.UserLevelId, repositories.StaffUserLevel.UserLevelId},
		}),
		handler.GetListSiteType,
	}

	deleteSiteType := []gin.HandlerFunc{
		middlewares.ValidateRequestBody(&models.SiteType{}),
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{repositories.RootUserLevel.UserLevelId, repositories.StaffUserLevel.UserLevelId},
		}),
		handler.DeleteSiteType,
	}

	v1.GET("/list", getListSiteType...)
	v1.POST("/create", createSiteType...)
	v1.DELETE("/delete", deleteSiteType...)

	return handler
}

func (h *siteTypeHandler) CreateSiteType(c *gin.Context) {
	siteType := &models.SiteType{}
	if err := c.ShouldBindJSON(siteType); err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	requesterUserId := c.MustGet("user_id").(int)

	siteType, err := h.siteTypeUsecase.CreateSiteType(siteType, requesterUserId)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, siteType, "Site type created successfully")
}

func (h *siteTypeHandler) GetListSiteType(c *gin.Context) {
	siteTypes, err := h.siteTypeUsecase.GetListSiteType()
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, siteTypes, "Site types retrieved successfully")
}

func (h *siteTypeHandler) DeleteSiteType(c *gin.Context) {
	siteType := &models.SiteType{}
	if err := c.ShouldBindJSON(siteType); err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	requesterUserId := c.MustGet("user_id").(int)

	err := h.siteTypeUsecase.DeleteSiteType(siteType, requesterUserId)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, nil, "Site type deleted successfully")
}
