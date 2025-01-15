package handlers

import (
	"server/internal/core/models"
	"server/internal/core/repositories"
	"server/internal/core/usecases"
	"server/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type OrganizationUserHandler struct{}

type organizationUserHandler struct {
	organizationUserUsecase usecases.OrganizationUserUsecase
}

func NewOrganizationUserHandler(r *gin.Engine, organizationUserUsecase usecases.OrganizationUserUsecase, globalMiddlewares ...gin.HandlerFunc) *organizationUserHandler {
	handler := &organizationUserHandler{organizationUserUsecase}

	v1 := r.Group("/v1/organization-users", globalMiddlewares...)

	createOrganizationUser := []gin.HandlerFunc{
		middlewares.ValidateRequestBody(&models.OrganizationUser{}),
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{
				repositories.OwnerUserLevel.UserLevelId,
				repositories.SuperAdminUserLevel.UserLevelId,
				repositories.AdminUserLevel.UserLevelId,
			},
		}),
		handler.CreateOrganizationUser,
	}

	v1.POST("/create", createOrganizationUser...)

	return handler
}

func (h *organizationUserHandler) CreateOrganizationUser(c *gin.Context) {
	requesterUserId := c.MustGet("user_id").(int)

	organizationUser := &models.OrganizationUser{}
	if err := c.ShouldBindJSON(organizationUser); err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	organizationUser, err := h.organizationUserUsecase.CreateOrganizationUser(organizationUser, requesterUserId)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, organizationUser, "Organization user created successfully")
}
