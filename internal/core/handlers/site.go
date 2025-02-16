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

	getSiteByName := []gin.HandlerFunc{
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{repositories.RootUserLevel.UserLevelId, repositories.DeveloperUserLevel.UserLevelId},
		}),
		handler.GetSiteByName,
	}

	getListSite := []gin.HandlerFunc{
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{repositories.RootUserLevel.UserLevelId, repositories.DeveloperUserLevel.UserLevelId},
		}),
		handler.GetListSite,
	}

	getListSiteBySiteTypeId := []gin.HandlerFunc{
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{repositories.RootUserLevel.UserLevelId, repositories.DeveloperUserLevel.UserLevelId},
		}),
		handler.GetListSiteBySiteTypeId,
	}

	getListSiteWithoutBySiteTypeId := []gin.HandlerFunc{
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{repositories.RootUserLevel.UserLevelId, repositories.DeveloperUserLevel.UserLevelId},
		}),
		handler.GetListSiteWithoutBySiteTypeId,
	}

	createSite := []gin.HandlerFunc{
		middlewares.ValidateRequestBody(&models.Site{}),
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{repositories.RootUserLevel.UserLevelId, repositories.DeveloperUserLevel.UserLevelId},
		}),
		handler.CreateSite,
	}

	v1.GET("/:name", getSiteByName...)
	v1.GET("/list", getListSite...)
	v1.GET("/list/:site_type_id", getListSiteBySiteTypeId...)
	v1.GET("/list/without/:site_type_id", getListSiteWithoutBySiteTypeId...)
	v1.POST("/create", createSite...)

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

func (h *siteHandler) GetSiteByName(c *gin.Context) {
	name := c.Param("name")
	site, err := h.siteUsecase.GetSiteByName(name)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, site, "Site retrieved successfully")
}
