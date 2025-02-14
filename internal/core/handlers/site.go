package handlers

import (
	"server/internal/core/models"
	"server/internal/core/repositories"
	"server/internal/core/usecases"
	"server/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type SiteHandler interface{}

type siteHandler struct {
	siteUsecase usecases.SiteUsecase
}

func NewSiteHandler(r *gin.Engine, siteUsecase usecases.SiteUsecase, globalMiddlewares ...gin.HandlerFunc) SiteHandler {
	handler := &siteHandler{siteUsecase}

	v1 := r.Group("/v1/sites", globalMiddlewares...)

	createSite := []gin.HandlerFunc{
		middlewares.ValidateRequestBody(&models.Site{}),
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{repositories.RootUserLevel.UserLevelId, repositories.StaffUserLevel.UserLevelId},
		}),
		handler.CreateSite,
	}

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
