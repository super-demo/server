package handlers

import (
	"server/internal/core/models"
	"server/internal/core/repositories"
	"server/internal/core/usecases"
	"server/internal/middlewares"
	"server/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SiteMiniAppHandler interface{}

type siteMiniAppHandler struct {
	siteMiniAppUsecase usecases.SiteMiniAppUsecase
}

func NewSiteMiniAppHandler(r *gin.Engine, siteMiniAppUsecase usecases.SiteMiniAppUsecase, globalMiddlewares ...gin.HandlerFunc) SiteMiniAppHandler {
	handler := &siteMiniAppHandler{siteMiniAppUsecase}

	v1 := r.Group("/v1/site-mini-apps", globalMiddlewares...)

	createSiteMiniApp := []gin.HandlerFunc{
		middlewares.ValidateRequestBody(&models.SiteMiniApp{}),
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{
				repositories.RootUserLevel.UserLevelId,
				repositories.SuperAdminUserLevel.UserLevelId,
				repositories.AdminUserLevel.UserLevelId,
			},
		}),
		handler.CreateSiteMiniApp,
	}

	getListSiteMiniAppBySiteId := []gin.HandlerFunc{
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{
				repositories.RootUserLevel.UserLevelId,
				repositories.SuperAdminUserLevel.UserLevelId,
				repositories.AdminUserLevel.UserLevelId,
				repositories.MemberUserLevel.UserLevelId,
			},
		}),
		handler.GetListSiteMiniAppBySiteId,
	}

	getSiteMiniAppById := []gin.HandlerFunc{
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{
				repositories.RootUserLevel.UserLevelId,
				repositories.SuperAdminUserLevel.UserLevelId,
				repositories.AdminUserLevel.UserLevelId,
				repositories.MemberUserLevel.UserLevelId,
			},
		}),
		handler.GetSiteMiniAppById,
	}

	updateSiteMiniApp := []gin.HandlerFunc{
		middlewares.ValidateRequestBody(&models.SiteMiniApp{}),
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{
				repositories.RootUserLevel.UserLevelId,
				repositories.SuperAdminUserLevel.UserLevelId,
				repositories.AdminUserLevel.UserLevelId,
			},
		}),
		handler.UpdateSiteMiniApp,
	}

	deleteSiteMiniApp := []gin.HandlerFunc{
		middlewares.ValidateRequestBody(&models.SiteMiniApp{}),
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{
				repositories.RootUserLevel.UserLevelId,
				repositories.SuperAdminUserLevel.UserLevelId,
				repositories.AdminUserLevel.UserLevelId,
				repositories.MemberUserLevel.UserLevelId,
			},
		}),
		handler.DeleteSiteMiniApp,
	}

	v1.GET("/:id", getSiteMiniAppById...)
	v1.GET("/list/:site_id", getListSiteMiniAppBySiteId...)
	v1.POST("/create", createSiteMiniApp...)
	v1.PUT("/update", updateSiteMiniApp...)
	v1.DELETE("/delete", deleteSiteMiniApp...)

	return handler
}

func (h *siteMiniAppHandler) CreateSiteMiniApp(c *gin.Context) {
	siteMiniApp := &models.SiteMiniApp{}
	if err := c.ShouldBindJSON(siteMiniApp); err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	requesterUserId := c.MustGet("user_id").(int)

	data, err := h.siteMiniAppUsecase.CreateSiteMiniApp(siteMiniApp, requesterUserId)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, data, "Site mini app created successfully")
}

func (h *siteMiniAppHandler) GetListSiteMiniAppBySiteId(c *gin.Context) {
	siteIdStr, _ := c.Params.Get("site_id")
	siteId, _ := strconv.Atoi(siteIdStr)

	siteMiniApps, err := h.siteMiniAppUsecase.GetListSiteMiniAppBySiteId(siteId)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, siteMiniApps, "List of site mini apps")
}

func (h *siteMiniAppHandler) GetSiteMiniAppById(c *gin.Context) {
	siteMiniAppId := utils.GetIdFromParams(c)

	siteMiniApp, err := h.siteMiniAppUsecase.GetSiteMiniAppById(siteMiniAppId)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, siteMiniApp, "Site mini app retrieved successfully")
}

func (h *siteMiniAppHandler) UpdateSiteMiniApp(c *gin.Context) {
	siteMiniApp := &models.SiteMiniApp{}
	if err := c.ShouldBindJSON(siteMiniApp); err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	requesterUserId := c.MustGet("user_id").(int)

	data, err := h.siteMiniAppUsecase.UpdateSiteMiniApp(siteMiniApp, requesterUserId)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, data, "Site mini app updated successfully")
}

func (h *siteMiniAppHandler) DeleteSiteMiniApp(c *gin.Context) {
	siteMiniApp := &models.SiteMiniApp{}
	if err := c.ShouldBindJSON(siteMiniApp); err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	requesterUserId := c.MustGet("user_id").(int)

	err := h.siteMiniAppUsecase.DeleteSiteMiniApp(siteMiniApp, requesterUserId)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, nil, "Site mini app deleted successfully")
}
