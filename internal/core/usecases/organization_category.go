package usecases

import (
	"server/infrastructure/app"
	"server/internal/core/models"
	"server/internal/core/repositories"
)

type OrganizationCategoryUsecase interface {
	CreateOrganizationCategory(organizationCategory *models.OrganizationCategory, requesterUserId int) (*models.OrganizationCategory, error)
}

type organizationCategoryUsecase struct {
	organizationRepo         repositories.OrganizationRepository
	organizationCategoryRepo repositories.OrganizationCategoryRepository
	organizationLogRepo      repositories.OrganizationLogRepository
}

func NewOrganizationCategoryUsecase(organizationRepo repositories.OrganizationRepository, organizationCategoryRepo repositories.OrganizationCategoryRepository, organizationLogRepo repositories.OrganizationLogRepository) OrganizationCategoryUsecase {
	return &organizationCategoryUsecase{organizationRepo, organizationCategoryRepo, organizationLogRepo}
}

func (u *organizationCategoryUsecase) CreateOrganizationCategory(organizationCategory *models.OrganizationCategory, requesterUserId int) (*models.OrganizationCategory, error) {
	txOrganizationCategoryRepo, err := u.organizationCategoryRepo.BeginLog()
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			txOrganizationCategoryRepo.Rollback()
		}
	}()

	exists, err := txOrganizationCategoryRepo.CheckOrganizationCategoryExistsByName(organizationCategory.Name)
	if err != nil {
		txOrganizationCategoryRepo.Rollback()
		return nil, err
	}

	if exists {
		txOrganizationCategoryRepo.Rollback()
		return nil, app.ErrOrganizationCategoryNameExists
	}

	organizationCategory.CreatedBy = requesterUserId
	organizationCategory.UpdatedBy = requesterUserId
	newOrganizationCategory, err := txOrganizationCategoryRepo.CreateOrganizationCategory(organizationCategory)
	if err != nil {
		txOrganizationCategoryRepo.Rollback()
		return nil, err
	}

	if err := txOrganizationCategoryRepo.Commit(); err != nil {
		return nil, err
	}

	organizationLog := &models.OrganizationLog{
		OrganizationId: newOrganizationCategory.OrganizationId,
		Action:         "Created Organization Category",
		Description:    "Category Created in Organization",
		CreatedBy:      requesterUserId,
	}

	if _, err := u.organizationLogRepo.CreateOrganizationLog(organizationLog); err != nil {
		return nil, err
	}

	return newOrganizationCategory, nil
}
