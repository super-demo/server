package handlers

import (
	"server/internal/core/models"
	"server/internal/core/repositories"
	"server/internal/core/usecases"
	"server/internal/middlewares"
	"server/pkg/utils"
	"strconv"
	"strings"

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
		// middlewares.ValidateRequestBody([]models.CreateSiteUserWithoutSignRequest{}),
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{repositories.RootUserLevel.UserLevelId, repositories.SuperAdminUserLevel.UserLevelId, repositories.AdminUserLevel.UserLevelId},
		}),
		handler.CreateSiteUserWithoutSign,
	}

	bulkImportUserWithoutSign := []gin.HandlerFunc{
		middlewares.ValidateCSVOrXLSXFile(),
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{repositories.RootUserLevel.UserLevelId, repositories.SuperAdminUserLevel.UserLevelId, repositories.AdminUserLevel.UserLevelId},
		}),
		handler.BulkImportUserWithoutSign,
	}

	getListSiteUserBySiteId := []gin.HandlerFunc{
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{repositories.RootUserLevel.UserLevelId, repositories.SuperAdminUserLevel.UserLevelId, repositories.AdminUserLevel.UserLevelId},
		}),
		handler.GetListSiteUserBySiteId,
	}

	updateSiteUser := []gin.HandlerFunc{
		middlewares.ValidateRequestBody(&models.SiteUser{}),
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{repositories.RootUserLevel.UserLevelId, repositories.SuperAdminUserLevel.UserLevelId, repositories.AdminUserLevel.UserLevelId},
		}),
		handler.UpdateSiteUser,
	}

	deleteSiteUserBySiteIdAndUserId := []gin.HandlerFunc{
		middlewares.ValidateRequestBody(&models.SiteUser{}),
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{repositories.RootUserLevel.UserLevelId, repositories.SuperAdminUserLevel.UserLevelId, repositories.AdminUserLevel.UserLevelId},
		}),
		handler.DeleteSiteUserBySiteIdAndUserId,
	}

	v1.GET("/list/:siteId", getListSiteUserBySiteId...)
	v1.POST("/create/without/sign", createSiteUserWithoutSign...)
	v1.POST("/bulk-import/without/sign/:siteId", bulkImportUserWithoutSign...)
	v1.PUT("/update", updateSiteUser...)
	v1.DELETE("/delete", deleteSiteUserBySiteIdAndUserId...)

	return handler
}

func (h *siteUserHandler) CreateSiteUserWithoutSign(c *gin.Context) {
	requests := []models.CreateSiteUserWithoutSignRequest{}
	if err := c.ShouldBindJSON(&requests); err != nil {
		middlewares.ResponseError(c, err)
		return
	}
	requesterUserId := c.MustGet("user_id").(int)

	siteUser, err := h.siteUserUsecase.CreateSiteUserWithoutSign(requests, requesterUserId)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, siteUser, "Site user created successfully")
}

func (h *siteUserHandler) BulkImportUserWithoutSign(c *gin.Context) {
	siteId, err := strconv.Atoi(c.Param("siteId"))
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	requesterUserId := c.MustGet("user_id").(int)

	file, err := c.FormFile("file")
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	openedFile, err := file.Open()
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}
	defer openedFile.Close()

	var parseResult *utils.ParseCsvOrXlsxResult

	if strings.HasSuffix(file.Filename, ".csv") {
		parseResult, err = utils.ParseCsv(openedFile)
		if err != nil {
			middlewares.ResponseError(c, err)
			return
		}
	} else if strings.HasSuffix(file.Filename, ".xlsx") {
		parseResult, err = utils.ParseXlsx(openedFile)
		if err != nil {
			middlewares.ResponseError(c, err)
			return
		}
	}

	bulkImportResult, err := h.siteUserUsecase.BulkImportUserWithoutSign(siteId, parseResult.User, requesterUserId)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	for _, recordError := range parseResult.Failures {
		bulkImportResult.FailedCount++
		bulkImportResult.Failures = append(bulkImportResult.Failures, recordError)
	}

	if bulkImportResult.SuccessCount == 0 {
		middlewares.ResponseSuccess(c, bulkImportResult, "Bulk import user created with all failures")
		return
	}

	if bulkImportResult.FailedCount > 0 {
		middlewares.ResponseSuccess(c, bulkImportResult, "Bulk import user created with some failures")
		return
	}

	middlewares.ResponseSuccess(c, bulkImportResult, "Bulk import user created successfully")
}

func (h *siteUserHandler) GetListSiteUserBySiteId(c *gin.Context) {
	siteId, err := strconv.Atoi(c.Param("siteId"))
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	siteUsers, err := h.siteUserUsecase.GetListSiteUserBySiteId(siteId)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, siteUsers, "Site users retrieved successfully")
}

func (h *siteUserHandler) UpdateSiteUser(c *gin.Context) {
	siteUser := &models.SiteUser{}
	if err := c.ShouldBindJSON(siteUser); err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	siteUser, err := h.siteUserUsecase.UpdateSiteUser(siteUser)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, siteUser, "Site user updated successfully")
}

func (h *siteUserHandler) DeleteSiteUserBySiteIdAndUserId(c *gin.Context) {
	siteUser := &models.SiteUser{}
	if err := c.ShouldBindJSON(siteUser); err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	requesterUserId := c.MustGet("user_id").(int)

	err := h.siteUserUsecase.DeleteSiteUserBySiteIdAndUserId(siteUser, requesterUserId)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, nil, "Site user deleted successfully")
}
