package usecases

import (
	"server/infrastructure/app"
	"server/internal/core/models"
	"server/internal/core/repositories"
	"strconv"
)

type OrganizationCategoryUserUsecase interface {
	CreateOrganizationCategoryUser(organizationCategoryUser *models.OrganizationCategoryUser, requesterUserId int) (*models.OrganizationCategoryUser, error)
}

type organizationCategoryUserUsecase struct {
	organizationCategoryRepo     repositories.OrganizationCategoryRepository
	organizationCategoryUserRepo repositories.OrganizationCategoryUserRepository
	organizationLogRepo          repositories.OrganizationLogRepository
}

func NewOrganizationCategoryUserUsecase(organizationCategoryRepo repositories.OrganizationCategoryRepository, organizationCategoryUserRepo repositories.OrganizationCategoryUserRepository, organizationLogRepo repositories.OrganizationLogRepository) OrganizationCategoryUserUsecase {
	return &organizationCategoryUserUsecase{organizationCategoryRepo, organizationCategoryUserRepo, organizationLogRepo}
}

func (u *organizationCategoryUserUsecase) CreateOrganizationCategoryUser(organizationCategoryUser *models.OrganizationCategoryUser, requesterUserId int) (*models.OrganizationCategoryUser, error) {
	txOrganizationCategoryUserRepo, err := u.organizationCategoryUserRepo.BeginLog()
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			txOrganizationCategoryUserRepo.Rollback()
		}
	}()

	exists, err := txOrganizationCategoryUserRepo.CheckOrganizationCategoryUserExistsById(organizationCategoryUser.UserId)
	if err != nil {
		txOrganizationCategoryUserRepo.Rollback()
		return nil, err
	}

	if exists {
		txOrganizationCategoryUserRepo.Rollback()
		return nil, app.ErrOrganizationCategoryUserExists
	}

	organizationCategoryUser.CreatedBy = requesterUserId
	organizationCategoryUser.UpdatedBy = requesterUserId
	newOrganizationCategoryUser, err := txOrganizationCategoryUserRepo.CreateOrganizationCategoryUser(organizationCategoryUser)
	if err != nil {
		txOrganizationCategoryUserRepo.Rollback()
		return nil, err
	}

	if err := txOrganizationCategoryUserRepo.Commit(); err != nil {
		return nil, err
	}

	organizationLog := &models.OrganizationLog{
		OrganizationId: newOrganizationCategoryUser.OrganizationId,
		Action:         "Invited",
		Description:    "Invited user " + strconv.Itoa(newOrganizationCategoryUser.UserId) + " to organization category " + strconv.Itoa(newOrganizationCategoryUser.UserId),
		CreatedBy:      requesterUserId,
	}

	if _, err := u.organizationLogRepo.CreateOrganizationLog(organizationLog); err != nil {
		return nil, err
	}

	return newOrganizationCategoryUser, nil
}
