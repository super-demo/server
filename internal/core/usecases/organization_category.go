package usecases

import (
	"server/infrastructure/app"
	"server/internal/core/models"
	"server/internal/core/repositories"
	"time"
)

type OrganizationCategoryUsecase interface {
	CreateOrganizationCategory(organizationCategory *models.OrganizationCategory, requesterUserId int) (*models.OrganizationCategory, error)
	UpdateOrganizationCategory(organizationCategory *models.OrganizationCategory, requesterUserId int) (*models.OrganizationCategory, error)
	DeleteOrganizationCategory(organizationCategory *models.OrganizationCategory, requesterUserId int) error
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

func (u *organizationCategoryUsecase) UpdateOrganizationCategory(organizationCategory *models.OrganizationCategory, requesterUserId int) (*models.OrganizationCategory, error) {
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

	oldOrganizationCategory, err := txOrganizationCategoryRepo.GetOrganizationCategoryById(organizationCategory.OrganizationCategoryId)
	if err != nil {
		txOrganizationCategoryRepo.Rollback()
		return nil, err
	}

	newOrganizationCategory := &models.OrganizationCategory{
		OrganizationCategoryId: oldOrganizationCategory.OrganizationCategoryId,
		OrganizationId:         oldOrganizationCategory.OrganizationId,
		Name:                   organizationCategory.Name,
		Description:            organizationCategory.Description,
		UpdatedBy:              requesterUserId,
		UpdatedAt:              time.Now(),
	}

	newOrganizationCategory, err = txOrganizationCategoryRepo.UpdateOrganizationCategory(newOrganizationCategory)
	if err != nil {
		txOrganizationCategoryRepo.Rollback()
		return nil, err
	}

	if err := txOrganizationCategoryRepo.Commit(); err != nil {
		return nil, err
	}

	organizationLog := &models.OrganizationLog{
		OrganizationId: newOrganizationCategory.OrganizationId,
		Action:         "Updated Category",
		Description:    "Category" + newOrganizationCategory.Name + " Updated in Organization",
		CreatedBy:      requesterUserId,
	}

	if _, err := u.organizationLogRepo.CreateOrganizationLog(organizationLog); err != nil {
		return nil, err
	}

	return newOrganizationCategory, nil
}

func (u *organizationCategoryUsecase) DeleteOrganizationCategory(organizationCategory *models.OrganizationCategory, requesterUserId int) error {
	txOrganizationCategoryRepo, err := u.organizationCategoryRepo.BeginLog()
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			txOrganizationCategoryRepo.Rollback()
		}
	}()

	oldOrganizationCategory, err := txOrganizationCategoryRepo.GetOrganizationCategoryById(organizationCategory.OrganizationCategoryId)
	if err != nil {
		txOrganizationCategoryRepo.Rollback()
		return err
	}

	if err := txOrganizationCategoryRepo.DeleteOrganizationCategory(oldOrganizationCategory); err != nil {
		txOrganizationCategoryRepo.Rollback()
		return err
	}

	if err := txOrganizationCategoryRepo.Commit(); err != nil {
		return err
	}

	organizationLog := &models.OrganizationLog{
		OrganizationId: oldOrganizationCategory.OrganizationId,
		Action:         "Deleted Category",
		Description:    "Category" + oldOrganizationCategory.Name + " Deleted in Organization",
		CreatedBy:      requesterUserId,
	}

	if _, err := u.organizationLogRepo.CreateOrganizationLog(organizationLog); err != nil {
		return err
	}

	return nil
}
