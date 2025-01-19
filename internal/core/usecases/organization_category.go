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
	organizationRepo                repositories.OrganizationRepository
	organizationCategoryRepo        repositories.OrganizationCategoryRepository
	organizationCategoryServiceRepo repositories.OrganizationCategoryServiceRepository
	organizationCategoryUserRepo    repositories.OrganizationCategoryUserRepository
	organizationLogRepo             repositories.OrganizationLogRepository
}

func NewOrganizationCategoryUsecase(organizationRepo repositories.OrganizationRepository, organizationCategoryRepo repositories.OrganizationCategoryRepository, organizationCategoryServiceRepo repositories.OrganizationCategoryServiceRepository, organizationCategoryUserRepo repositories.OrganizationCategoryUserRepository, organizationLogRepo repositories.OrganizationLogRepository) OrganizationCategoryUsecase {
	return &organizationCategoryUsecase{organizationRepo, organizationCategoryRepo, organizationCategoryServiceRepo, organizationCategoryUserRepo, organizationLogRepo}
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
		Action:         "Created",
		Description:    "Created category " + newOrganizationCategory.Name + " in Organization",
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
		Action:         "Updated",
		Description:    "Updated category " + newOrganizationCategory.Name + " in Organization",
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

	exists, err := txOrganizationCategoryRepo.CheckOrganizationCategoryExists(organizationCategory.OrganizationId, organizationCategory.OrganizationCategoryId)
	if err != nil {
		txOrganizationCategoryRepo.Rollback()
		return err
	}

	if !exists {
		txOrganizationCategoryRepo.Rollback()
		return app.ErrCategoryNotFound
	}

	organizationService, err := u.organizationCategoryServiceRepo.GetOrganizationCategoryServiceByOrganizationCategoryId(organizationCategory.OrganizationCategoryId)
	if err != nil {
		txOrganizationCategoryRepo.Rollback()
		return err
	}

	organizationUser, err := u.organizationCategoryUserRepo.GetOrganizationCategoryUserByOrganizationCategoryId(organizationCategory.OrganizationCategoryId)
	if err != nil {
		txOrganizationCategoryRepo.Rollback()
		return err
	}

	if organizationService != nil {
		organizationCategoryService, err := u.organizationCategoryServiceRepo.GetOrganizationCategoryServiceByOrganizationServiceId(organizationService.OrganizationServiceId)
		if err != nil {
			txOrganizationCategoryRepo.Rollback()
			return err
		}

		if organizationCategoryService != nil {
			if err := u.organizationCategoryServiceRepo.DeleteOrganizationCategoryService(organizationCategoryService); err != nil {
				txOrganizationCategoryRepo.Rollback()
				return err
			}
		}
	}

	if organizationUser != nil {
		organizationCategoryUser, err := u.organizationCategoryUserRepo.GetOrganizationCategoryUserByUserId(organizationUser.UserId)
		if err != nil {
			txOrganizationCategoryRepo.Rollback()
			return err
		}

		if organizationCategoryUser != nil {
			if err := u.organizationCategoryUserRepo.DeleteOrganizationCategoryUser(organizationCategoryUser); err != nil {
				txOrganizationCategoryRepo.Rollback()
				return err
			}
		}
	}

	if err := txOrganizationCategoryRepo.DeleteOrganizationCategory(organizationCategory); err != nil {
		txOrganizationCategoryRepo.Rollback()
		return err
	}

	if err := txOrganizationCategoryRepo.Commit(); err != nil {
		return err
	}

	organizationLog := &models.OrganizationLog{
		OrganizationId: organizationCategory.OrganizationId,
		Action:         "Deleted",
		Description:    "Deleted Category " + organizationCategory.Name + " in Organization",
		CreatedBy:      requesterUserId,
	}

	if _, err := u.organizationLogRepo.CreateOrganizationLog(organizationLog); err != nil {
		return err
	}

	return nil
}
