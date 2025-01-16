package usecases

import (
	"server/infrastructure/app"
	"server/internal/core/models"
	"server/internal/core/repositories"
	"strconv"
)

type OrganizationCategoryUserUsecase interface {
	CreateOrganizationCategoryUser(organizationCategoryUser *models.OrganizationCategoryUser, requesterUserId int) (*models.OrganizationCategoryUser, error)
	DeleteOrganizationCategoryUser(organizationCategoryUser *models.OrganizationCategoryUser, requesterUserId int) error
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

func (u *organizationCategoryUserUsecase) DeleteOrganizationCategoryUser(organizationCategoryUser *models.OrganizationCategoryUser, requesterUserId int) error {
	txOrganizationCategoryUserRepo, err := u.organizationCategoryUserRepo.BeginLog()
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			txOrganizationCategoryUserRepo.Rollback()
		}
	}()

	if err := txOrganizationCategoryUserRepo.DeleteOrganizationCategoryUser(organizationCategoryUser); err != nil {
		txOrganizationCategoryUserRepo.Rollback()
		return err
	}

	if err := txOrganizationCategoryUserRepo.Commit(); err != nil {
		return err
	}

	organizationLog := &models.OrganizationLog{
		OrganizationId: organizationCategoryUser.OrganizationId,
		Action:         "Removed",
		Description:    "Removed user " + strconv.Itoa(organizationCategoryUser.UserId) + " from organization category " + strconv.Itoa(organizationCategoryUser.UserId),
		CreatedBy:      requesterUserId,
	}

	if _, err := u.organizationLogRepo.CreateOrganizationLog(organizationLog); err != nil {
		return err
	}

	return nil
}
