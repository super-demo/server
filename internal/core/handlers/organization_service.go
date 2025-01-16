package handlers

import (
	"server/internal/core/models"
	"server/internal/core/usecases"
	"server/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type OrganizationServiceHandler struct{}

type organizationServiceHandler struct {
	organizationServiceUsecase usecases.OrganizationServiceUsecase
}

func NewOrganizationServiceHandler(r *gin.Engine, organizationServiceUsecase usecases.OrganizationServiceUsecase, globalMiddlewares ...gin.HandlerFunc) *organizationServiceHandler {
	handler := &organizationServiceHandler{organizationServiceUsecase}

	v1 := r.Group("/v1/organization-services", globalMiddlewares...)

	createOrganizationService := []gin.HandlerFunc{
		middlewares.ValidateRequestBody(&models.OrganizationService{}),
		handler.CreateOrganizationService,
	}

	v1.POST("/create", createOrganizationService...)

	return handler
}

func (h *organizationServiceHandler) CreateOrganizationService(c *gin.Context) {
	requesterUserId := c.MustGet("user_id").(int)

	organizationService := &models.OrganizationService{}
	if err := c.ShouldBindJSON(organizationService); err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	organizationService, err := h.organizationServiceUsecase.CreateOrganizationService(organizationService, requesterUserId)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, organizationService, "Organization service created successfully")
}
