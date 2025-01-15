package handlers

import (
	"server/internal/core/models"
	"server/internal/core/repositories"
	"server/internal/core/usecases"
	"server/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type OrganizationCategoryHandler interface{}

type organizationCategoryHandler struct {
	organizationCategoryUsecase usecases.OrganizationCategoryUsecase
}

func NewOrganizationCategoryHandler(r *gin.Engine, organizationCategoryUsecase usecases.OrganizationCategoryUsecase, globalMiddlewares ...gin.HandlerFunc) *organizationCategoryHandler {
	handler := &organizationCategoryHandler{organizationCategoryUsecase}

	v1 := r.Group("/v1/organization-categories", globalMiddlewares...)

	createOrganizationCategory := []gin.HandlerFunc{
		middlewares.ValidateRequestBody(&models.OrganizationCategory{}),
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{
				repositories.OwnerUserLevel.UserLevelId,
				repositories.SuperAdminUserLevel.UserLevelId,
				repositories.AdminUserLevel.UserLevelId,
			},
		}),
		handler.CreateOrganizationCategory,
	}

	updateOrganizationCategory := []gin.HandlerFunc{
		middlewares.ValidateRequestBody(&models.OrganizationCategory{}),
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{
				repositories.OwnerUserLevel.UserLevelId,
				repositories.SuperAdminUserLevel.UserLevelId,
				repositories.AdminUserLevel.UserLevelId,
			},
		}),
		handler.UpdateOrganizationCategory,
	}

	deleteOrganizationCategory := []gin.HandlerFunc{
		middlewares.ValidateRequestBody(&models.OrganizationCategory{}),
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{
				repositories.OwnerUserLevel.UserLevelId,
				repositories.SuperAdminUserLevel.UserLevelId,
				repositories.AdminUserLevel.UserLevelId,
			},
		}),
		handler.DeleteOrganizationCategory,
	}

	v1.POST("/create", createOrganizationCategory...)
	v1.PUT("/update", updateOrganizationCategory...)
	v1.DELETE("/delete", deleteOrganizationCategory...)

	return handler
}

func (h *organizationCategoryHandler) CreateOrganizationCategory(c *gin.Context) {
	organizationCategory := &models.OrganizationCategory{}
	if err := c.ShouldBindJSON(organizationCategory); err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	requesterUserId := c.MustGet("user_id").(int)
	organizationCategory, err := h.organizationCategoryUsecase.CreateOrganizationCategory(organizationCategory, requesterUserId)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, organizationCategory, "Organization Category Created Successfully")
}

func (h *organizationCategoryHandler) UpdateOrganizationCategory(c *gin.Context) {
	organizationCategory := &models.OrganizationCategory{}
	if err := c.ShouldBindJSON(organizationCategory); err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	requesterUserId := c.MustGet("user_id").(int)
	organizationCategory, err := h.organizationCategoryUsecase.UpdateOrganizationCategory(organizationCategory, requesterUserId)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, organizationCategory, "Organization Category Updated Successfully")
}

func (h *organizationCategoryHandler) DeleteOrganizationCategory(c *gin.Context) {
	organizationCategory := &models.OrganizationCategory{}
	if err := c.ShouldBindJSON(organizationCategory); err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	requesterUserId := c.MustGet("user_id").(int)
	err := h.organizationCategoryUsecase.DeleteOrganizationCategory(organizationCategory, requesterUserId)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, "Organization Category Deleted Successfully")
}
