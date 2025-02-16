package handlers

import (
	"server/internal/core/models"
	"server/internal/core/repositories"
	"server/internal/core/usecases"
	"server/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type SiteUserHandler interface{}

type siteUserHandler struct {
	siteUserUsecase usecases.SiteUserUsecase
}

func NewSiteUserHandler(r *gin.Engine, siteUserUsecase usecases.SiteUserUsecase, globalMiddlewares ...gin.HandlerFunc) SiteUserHandler {
	handler := &siteUserHandler{siteUserUsecase}

	v1 := r.Group("/v1/site-users", globalMiddlewares...)

	createSiteUserWithoutSign := []gin.HandlerFunc{
		middlewares.ValidateRequestBody(&models.CreateSiteUserWithoutSignRequest{}),
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{repositories.RootUserLevel.UserLevelId, repositories.DeveloperUserLevel.UserLevelId},
		}),
		handler.CreateSiteUserWithoutSign,
	}

	v1.POST("/create/without/sign", createSiteUserWithoutSign...)

	return handler
}

func (h *siteUserHandler) CreateSiteUserWithoutSign(c *gin.Context) {
	request := &models.CreateSiteUserWithoutSignRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		middlewares.ResponseError(c, err)
		return
	}
	requesterUserId := c.MustGet("user_id").(int)

	siteUser, err := h.siteUserUsecase.CreateSiteUserWithoutSign(request, requesterUserId)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, siteUser, "Site user created successfully")
}
