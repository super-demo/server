package handlers

import (
	"server/internal/core/models"
	"server/internal/core/usecases"
	"server/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type OrganizationCategoryUserHandler interface{}

type organizationCategoryUserHandler struct {
	organizationCategoryUserUsecase usecases.OrganizationCategoryUserUsecase
}

func NewOrganizationCategoryUserHandler(r *gin.Engine, organizationCategoryUserUsecase usecases.OrganizationCategoryUserUsecase, globalMiddlewares ...gin.HandlerFunc) *organizationCategoryUserHandler {
	handler := &organizationCategoryUserHandler{organizationCategoryUserUsecase}

	v1 := r.Group("/v1/organization-category-users", globalMiddlewares...)

	createOrganizationCategoryUser := []gin.HandlerFunc{
		middlewares.ValidateRequestBody(&models.OrganizationCategoryUser{}),
		handler.CreateOrganizationCategoryUser,
	}

	deleteOrganizationCategoryUser := []gin.HandlerFunc{
		middlewares.ValidateRequestBody(&models.OrganizationCategoryUser{}),
		handler.DeleteOrganizationCategoryUser,
	}

	v1.POST("/create", createOrganizationCategoryUser...)
	v1.DELETE("/delete", deleteOrganizationCategoryUser...)

	return handler

}

func (h *organizationCategoryUserHandler) CreateOrganizationCategoryUser(c *gin.Context) {
	requesterUserId := c.MustGet("user_id").(int)

	organizationCategoryUser := &models.OrganizationCategoryUser{}
	if err := c.ShouldBindJSON(organizationCategoryUser); err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	organizationCategoryUser, err := h.organizationCategoryUserUsecase.CreateOrganizationCategoryUser(organizationCategoryUser, requesterUserId)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, organizationCategoryUser, "Organization category user created successfully")
}

func (h *organizationCategoryUserHandler) DeleteOrganizationCategoryUser(c *gin.Context) {
	requesterUserId := c.MustGet("user_id").(int)

	organizationCategoryUser := &models.OrganizationCategoryUser{}
	if err := c.ShouldBindJSON(organizationCategoryUser); err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	err := h.organizationCategoryUserUsecase.DeleteOrganizationCategoryUser(organizationCategoryUser, requesterUserId)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, nil, "Organization category user deleted successfully")
}
