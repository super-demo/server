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

type SitePeopleHandler interface{}

type sitePeopleHandler struct {
	siteUserUsecase   usecases.SiteUserUsecase
	sitePeopleUsecase usecases.SitePeopleUsecase
}

func NewSitePeopleHandler(r *gin.Engine, siteUserUsecase usecases.SiteUserUsecase, sitePeopleUsecase usecases.SitePeopleUsecase, globalMiddlewares ...gin.HandlerFunc) SitePeopleHandler {
	handler := &sitePeopleHandler{siteUserUsecase, sitePeopleUsecase}

	v1 := r.Group("/v1/site-people", globalMiddlewares...)

	createSitePeople := []gin.HandlerFunc{
		// middlewares.ValidateRequestBody([]models.CreateSiteUserWithoutSignRequest{}),
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{repositories.RootUserLevel.UserLevelId, repositories.DeveloperUserLevel.UserLevelId},
		}),
		handler.CreateSitePeople,
	}

	bulkImportUserWithoutSign := []gin.HandlerFunc{
		middlewares.ValidateCSVOrXLSXFile(),
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{repositories.RootUserLevel.UserLevelId, repositories.DeveloperUserLevel.UserLevelId, repositories.AdminUserLevel.UserLevelId},
		}),
		handler.BulkImportUserWithoutSign,
	}

	getListSitePeopleBySiteId := []gin.HandlerFunc{
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{repositories.RootUserLevel.UserLevelId, repositories.SuperAdminUserLevel.UserLevelId, repositories.AdminUserLevel.UserLevelId},
		}),
		handler.GetListSitePeopleBySiteId,
	}

	deleteSiteUserBySiteIdAndUserId := []gin.HandlerFunc{
		middlewares.ValidateRequestBody(&models.SiteUser{}),
		middlewares.Permission(middlewares.AllowedPermissionConfig{
			AllowedUserLevelIDs: []int{repositories.RootUserLevel.UserLevelId, repositories.SuperAdminUserLevel.UserLevelId, repositories.AdminUserLevel.UserLevelId},
		}),
		handler.DeleteSiteUserBySiteIdAndUserId,
	}

	v1.GET("/list/:siteId", getListSitePeopleBySiteId...)
	v1.POST("/create", createSitePeople...)
	v1.POST("/bulk-import/without/sign/:siteId", bulkImportUserWithoutSign...)
	v1.DELETE("/delete", deleteSiteUserBySiteIdAndUserId...)

	return handler
}

func (h *sitePeopleHandler) CreateSitePeople(c *gin.Context) {
	request := []models.CreateSitePeopleRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		middlewares.ResponseError(c, err)
		return
	}
	requesterUserId := c.MustGet("user_id").(int)

	siteUser, err := h.sitePeopleUsecase.CreateSitePeople(request, requesterUserId)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, siteUser, "Site people created successfully")
}

func (h *sitePeopleHandler) BulkImportUserWithoutSign(c *gin.Context) {
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

func (h *sitePeopleHandler) GetListSitePeopleBySiteId(c *gin.Context) {
	siteId, err := strconv.Atoi(c.Param("siteId"))
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	siteUsers, err := h.sitePeopleUsecase.GetListSitePeopleBySiteId(siteId)
	if err != nil {
		middlewares.ResponseError(c, err)
		return
	}

	middlewares.ResponseSuccess(c, siteUsers, "Site people retrieved successfully")
}

func (h *sitePeopleHandler) DeleteSiteUserBySiteIdAndUserId(c *gin.Context) {
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
