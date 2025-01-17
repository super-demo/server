package handlers

import (
	"server/internal/core/models"
	"server/internal/core/usecases"
	"server/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type OrganizationCategoryServiceHandler struct{}

type organizationCategoryServiceHandler struct {
	organizationCategoryServiceUsecase usecases.OrganizationCategoryServiceUsecase
}

func NewOrganizationCategoryServiceHandler(r *gin.Engine, organizationCategoryService usecases.OrganizationCategoryServiceUsecase, globalMiddlewares ...gin.HandlerFunc) *organizationCategoryServiceHandler {
	handler := &organizationCategoryServiceHandler{organizationCategoryService}

	v1 := r.Group("/v1/organization-category-services", globalMiddlewares...)

	createOrganizationCategoryService := []gin.HandlerFunc{
		handler.CreateOrganizationCategoryService,
	}

	v1.POST("/create", createOrganizationCategoryService...)

	return handler

}

func (h *organizationCategoryServiceHandler) CreateOrganizationCategoryService(c *gin.Context) {
	requesterUserId := c.MustGet("user_id").(int)

	organizationCategoryService := &models.OrganizationCategoryService{}
	if err := c.ShouldBindJSON(organizationCategoryService); err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	organizationCategoryService, err := h.organizationCategoryServiceUsecase.CreateOrganizationCategoryService(organizationCategoryService, requesterUserId)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, organizationCategoryService, "Organization category service created successfully")
}
