package handlers

import (
	"server/internal/core/models"
	"server/internal/core/usecases"
	"server/internal/middlewares"
	"server/pkg/utils"

	"github.com/gin-gonic/gin"
)

type OrganizationHandler interface{}

type organizationHandler struct {
	organizationUsecase usecases.OrganizationUsecase
}

func NewOrganizationHandler(r *gin.Engine, organizationUsecase usecases.OrganizationUsecase, globalMiddlewares ...gin.HandlerFunc) OrganizationHandler {
	handler := &organizationHandler{organizationUsecase}

	v1 := r.Group("/v1/organizations", globalMiddlewares...)

	getOrganizationById := []gin.HandlerFunc{
		handler.GetOrganizationById,
	}

	getOrganizationListByUserId := []gin.HandlerFunc{
		handler.GetOrganizationListByUserId,
	}

	createOrganization := []gin.HandlerFunc{
		middlewares.ValidateRequestBody(&models.Organization{}),
		handler.CreateOrganization,
	}

	deleteOrganization := []gin.HandlerFunc{
		middlewares.ValidateRequestBody(&models.Organization{}),
		handler.DeleteOrganization,
	}

	v1.GET("/:id", getOrganizationById...)
	v1.GET("/list", getOrganizationListByUserId...)
	v1.POST("/create", createOrganization...)
	v1.DELETE("/delete", deleteOrganization...)

	return handler
}

func (h *organizationHandler) CreateOrganization(c *gin.Context) {
	organization := &models.Organization{}
	if err := c.ShouldBindJSON(organization); err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	requesterUserId := c.MustGet("user_id").(int)

	organization, err := h.organizationUsecase.CreateOrganization(organization, requesterUserId)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, organization, "Organization created successfully")
}

func (h *organizationHandler) GetOrganizationById(c *gin.Context) {
	organizationId := utils.GetIdFromParams(c)
	requesterUserId := c.MustGet("user_id").(int)

	organization, err := h.organizationUsecase.GetOrganizationById(organizationId, requesterUserId)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, organization, "Organization retrieved successfully")
}

func (h *organizationHandler) GetOrganizationListByUserId(c *gin.Context) {
	requesterUserId := c.MustGet("user_id").(int)

	organization, err := h.organizationUsecase.GetOrganizationListByUserId(requesterUserId)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, organization, "Organization list retrieved successfully")
}

func (h *organizationHandler) DeleteOrganization(c *gin.Context) {
	organization := &models.Organization{}
	requesterUserId := c.MustGet("user_id").(int)

	err := h.organizationUsecase.DeleteOrganization(organization, requesterUserId)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, nil, "Organization deleted successfully")
}
