package handlers

import (
	"server/internal/core/models"
	"server/internal/core/repositories"
	"server/internal/core/usecases"
	"server/internal/middlewares"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SiteTreeHandler interface{}

type siteTreeHandler struct {
	siteTreeUsecase usecases.SiteTreeUsecase
}

func NewSiteTreeHandler(r *gin.Engine, siteTreeUsecase usecases.SiteTreeUsecase, globalMiddlewares ...gin.HandlerFunc) SiteTreeHandler {
	handler := &siteTreeHandler{siteTreeUsecase}

	v1 := r.Group("/v1/site-trees", globalMiddlewares...)

	createSiteTree := []gin.HandlerFunc{
		middlewares.ValidateRequestBody(&models.SiteTree{}),
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{
				repositories.RootUserLevel.UserLevelId,
				repositories.SuperAdminUserLevel.UserLevelId,
				repositories.AdminUserLevel.UserLevelId,
			},
		}),
		handler.CreateSiteTree,
	}

	getListSiteTreeBySiteId := []gin.HandlerFunc{
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{
				repositories.RootUserLevel.UserLevelId,
				repositories.DeveloperUserLevel.UserLevelId,
				repositories.SuperAdminUserLevel.UserLevelId,
				repositories.AdminUserLevel.UserLevelId,
				repositories.MemberUserLevel.UserLevelId,
			},
		}),
		handler.GetListSiteTreeBySiteId,
	}

	updateSiteTree := []gin.HandlerFunc{
		middlewares.ValidateRequestBody(&models.SiteTree{}),
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{
				repositories.RootUserLevel.UserLevelId,
				repositories.SuperAdminUserLevel.UserLevelId,
				repositories.AdminUserLevel.UserLevelId,
			},
		}),
		handler.UpdateSiteTree,
	}

	deleteSiteTree := []gin.HandlerFunc{
		middlewares.ValidateRequestBody(&models.SiteTree{}),
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{
				repositories.RootUserLevel.UserLevelId,
				repositories.SuperAdminUserLevel.UserLevelId,
				repositories.AdminUserLevel.UserLevelId,
			},
		}),
		handler.DeleteSiteTree,
	}

	v1.GET("/list/:site_id", getListSiteTreeBySiteId...)
	v1.POST("/create", createSiteTree...)
	v1.PUT("/update", updateSiteTree...)
	v1.DELETE("/delete", deleteSiteTree...)

	return handler
}

func (h *siteTreeHandler) CreateSiteTree(c *gin.Context) {
	siteTree := &models.SiteTree{}
	if err := c.ShouldBindJSON(siteTree); err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	requesterUserId := c.MustGet("user_id").(int)

	siteTree, err := h.siteTreeUsecase.CreateSiteTree(siteTree, requesterUserId)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, siteTree, "Site tree created successfully")
}

func (h *siteTreeHandler) GetListSiteTreeBySiteId(c *gin.Context) {
	siteIdStr, _ := c.Params.Get("site_id")
	siteId, _ := strconv.Atoi(siteIdStr)

	requesterUserId := c.MustGet("user_id").(int)

	siteList, err := h.siteTreeUsecase.GetListSiteTreeBySiteId(siteId, requesterUserId)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, siteList, "Site tree retrieved successfully")
}

func (h *siteTreeHandler) UpdateSiteTree(c *gin.Context) {
	siteTree := &models.SiteTree{}
	if err := c.ShouldBindJSON(siteTree); err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	requesterUserId := c.MustGet("user_id").(int)

	siteTree, err := h.siteTreeUsecase.UpdateSiteTree(siteTree, requesterUserId)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, siteTree, "Site tree updated successfully")
}

func (h *siteTreeHandler) DeleteSiteTree(c *gin.Context) {
	siteTree := &models.SiteTree{}
	if err := c.ShouldBindJSON(siteTree); err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	requesterUserId := c.MustGet("user_id").(int)

	err := h.siteTreeUsecase.DeleteSiteTree(siteTree, requesterUserId)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, nil, "Site tree deleted successfully")
}
