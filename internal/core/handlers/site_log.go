package handlers

import (
	"server/internal/core/repositories"
	"server/internal/core/usecases"
	"server/internal/middlewares"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SiteLogHandler interface{}

type siteLogHandler struct {
	siteLogUsecase usecases.SiteLogUsecase
}

func NewSiteLogHandler(r *gin.Engine, siteLogUsecase usecases.SiteLogUsecase, globalMiddlewares ...gin.HandlerFunc) SiteLogHandler {
	handler := &siteLogHandler{siteLogUsecase}

	v1 := r.Group("/v1/site-logs", globalMiddlewares...)

	getListSiteLog := []gin.HandlerFunc{
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{
				repositories.RootUserLevel.UserLevelId,
			},
		}),
		handler.GetListSiteLog,
	}

	getSiteLogById := []gin.HandlerFunc{
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{
				repositories.RootUserLevel.UserLevelId,
				repositories.DeveloperUserLevel.UserLevelId,
				repositories.SuperAdminUserLevel.UserLevelId,
				repositories.AdminUserLevel.UserLevelId,
			},
		}),
		handler.GetSiteLogById,
	}

	v1.GET("/list", getListSiteLog...)
	v1.GET("/:siteLogId", getSiteLogById...)

	return handler
}

func (h *siteLogHandler) GetListSiteLog(c *gin.Context) {
	siteLogs, err := h.siteLogUsecase.GetListSiteLog()
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, siteLogs, "List site logs successfully")
}

func (h *siteLogHandler) GetSiteLogById(c *gin.Context) {
	siteLogIdStr := c.Param("siteLogId")
	siteLogId, err := strconv.Atoi(siteLogIdStr)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	siteLog, err := h.siteLogUsecase.GetSiteLogBySiteId(siteLogId)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, siteLog, "Get site log successfully")
}
