package handlers

import (
	"server/internal/core/models"
	"server/internal/core/repositories"
	"server/internal/core/usecases"
	"server/internal/middlewares"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SiteHandler interface{}

type siteHandler struct {
	siteUsecase usecases.SiteUsecase
}

func NewSiteHandler(r *gin.Engine, siteUsecase usecases.SiteUsecase, globalMiddlewares ...gin.HandlerFunc) SiteHandler {
	handler := &siteHandler{siteUsecase}

	v1 := r.Group("/v1/sites", globalMiddlewares...)

	getSiteById := []gin.HandlerFunc{
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{
				repositories.RootUserLevel.UserLevelId,
				repositories.DeveloperUserLevel.UserLevelId,
				repositories.SuperAdminUserLevel.UserLevelId,
				repositories.AdminUserLevel.UserLevelId,
				repositories.MemberUserLevel.UserLevelId,
			},
		}),
		handler.GetSiteById,
	}

	getSiteByName := []gin.HandlerFunc{
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{
				repositories.RootUserLevel.UserLevelId,
				repositories.DeveloperUserLevel.UserLevelId,
			},
		}),
		handler.GetSiteByName,
	}

	getListSite := []gin.HandlerFunc{
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{
				repositories.RootUserLevel.UserLevelId,
				repositories.DeveloperUserLevel.UserLevelId,
			},
		}),
		handler.GetListSite,
	}

	getListSiteBySiteTypeId := []gin.HandlerFunc{
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{
				repositories.RootUserLevel.UserLevelId,
				repositories.DeveloperUserLevel.UserLevelId,
			},
		}),
		handler.GetListSiteBySiteTypeId,
	}

	getListSiteWithoutBySiteTypeId := []gin.HandlerFunc{
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{
				repositories.RootUserLevel.UserLevelId,
				repositories.DeveloperUserLevel.UserLevelId,
			},
		}),
		handler.GetListSiteWithoutBySiteTypeId,
	}

	createSite := []gin.HandlerFunc{
		middlewares.ValidateRequestBody(&models.Site{}),
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{
				repositories.RootUserLevel.UserLevelId,
				repositories.DeveloperUserLevel.UserLevelId,
			},
		}),
		handler.CreateSite,
	}

	createSiteWorkspace := []gin.HandlerFunc{
		middlewares.ValidateRequestBody(&models.CreateSiteWorkspaceRequest{}),
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{
				repositories.RootUserLevel.UserLevelId,
				repositories.SuperAdminUserLevel.UserLevelId,
				repositories.AdminUserLevel.UserLevelId,
			},
		}),
		handler.CreateSiteWorkspace,
	}

	updateSiteWorkspace := []gin.HandlerFunc{
		middlewares.ValidateRequestBody(&models.Site{}),
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{
				repositories.RootUserLevel.UserLevelId,
				repositories.SuperAdminUserLevel.UserLevelId,
				repositories.AdminUserLevel.UserLevelId,
			},
		}),
		handler.UpdatedSiteWorkspace,
	}

	deleteSiteWorkspace := []gin.HandlerFunc{
		middlewares.ValidateRequestBody(&models.Site{}),
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{
				repositories.RootUserLevel.UserLevelId,
				repositories.SuperAdminUserLevel.UserLevelId,
				repositories.AdminUserLevel.UserLevelId,
			},
		}),
		handler.DeleteSiteWorkspace,
	}

	v1.GET("/:id", getSiteById...)

	v1.GET("/:name", getSiteByName...)
	v1.GET("/list", getListSite...)
	v1.GET("/list/:site_type_id", getListSiteBySiteTypeId...)
	v1.GET("/list/without/:site_type_id", getListSiteWithoutBySiteTypeId...)
	v1.POST("/create", createSite...)

	v1.POST("/create/workspace", createSiteWorkspace...)
	v1.PUT("/update/workspace", updateSiteWorkspace...)
	v1.DELETE("/delete/workspace", deleteSiteWorkspace...)

	return handler
}

func (h *siteHandler) CreateSite(c *gin.Context) {
	site := &models.Site{}
	if err := c.ShouldBindJSON(site); err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	requesterUserId := c.MustGet("user_id").(int)

	site, err := h.siteUsecase.CreateSite(site, requesterUserId)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, site, "Site created successfully")
}

func (h *siteHandler) GetListSite(c *gin.Context) {
	sites, err := h.siteUsecase.GetListSite()
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, sites, "List of sites")
}

func (h *siteHandler) GetListSiteBySiteTypeId(c *gin.Context) {
	siteTypeIdStr := c.Param("site_type_id")
	siteTypeId, err := strconv.Atoi(siteTypeIdStr)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}
	sites, err := h.siteUsecase.GetListSiteBySiteTypeId(siteTypeId)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, sites, "List of sites by site type id")
}

func (h *siteHandler) GetListSiteWithoutBySiteTypeId(c *gin.Context) {
	siteTypeIdStr := c.Param("site_type_id")
	siteTypeId, err := strconv.Atoi(siteTypeIdStr)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}
	sites, err := h.siteUsecase.GetListSiteWithoutBySiteTypeId(siteTypeId)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, sites, "List of sites without site type id")
}

func (h *siteHandler) GetSiteById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}
	site, err := h.siteUsecase.GetSiteById(id)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, site, "Site retrieved successfully")
}

func (h *siteHandler) GetSiteByName(c *gin.Context) {
	name := c.Param("name")
	site, err := h.siteUsecase.GetSiteByName(name)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, site, "Site retrieved successfully")
}

func (h *siteHandler) CreateSiteWorkspace(c *gin.Context) {
	siteWorkspace := &models.CreateSiteWorkspaceRequest{}
	if err := c.ShouldBindJSON(siteWorkspace); err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	requesterUserId := c.MustGet("user_id").(int)

	site, err := h.siteUsecase.CreateSiteWorkspace(siteWorkspace, requesterUserId)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, site, "Site workspace created successfully")
}

func (h *siteHandler) UpdatedSiteWorkspace(c *gin.Context) {
	siteWorkspace := &models.Site{}
	if err := c.ShouldBindJSON(siteWorkspace); err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	requesterUserId := c.MustGet("user_id").(int)

	site, err := h.siteUsecase.UpdateSiteWorkspace(siteWorkspace, requesterUserId)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, site, "Site workspace updated successfully")
}

func (h *siteHandler) DeleteSiteWorkspace(c *gin.Context) {
	siteWorkspace := &models.Site{}
	if err := c.ShouldBindJSON(siteWorkspace); err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	requesterUserId := c.MustGet("user_id").(int)

	err := h.siteUsecase.DeleteSiteWorkspace(siteWorkspace, requesterUserId)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, nil, "Site workspace deleted successfully")
}
