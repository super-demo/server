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
		middlewares.ValidateRequestBody(&models.OrganizationCategoryService{}),
		handler.CreateOrganizationCategoryService,
	}

	deleteOrganizationCategoryService := []gin.HandlerFunc{
		middlewares.ValidateRequestBody(&models.OrganizationCategoryService{}),
		handler.DeleteOrganizationCategoryService,
	}

	v1.POST("/create", createOrganizationCategoryService...)
	v1.DELETE("/delete", deleteOrganizationCategoryService...)

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

func (h *organizationCategoryServiceHandler) DeleteOrganizationCategoryService(c *gin.Context) {
	requesterUserId := c.MustGet("user_id").(int)

	organizationCategoryService := &models.OrganizationCategoryService{}
	if err := c.ShouldBindJSON(organizationCategoryService); err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	err := h.organizationCategoryServiceUsecase.DeleteOrganizationCategoryService(organizationCategoryService, requesterUserId)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, nil, "Organization category service deleted successfully")
}
