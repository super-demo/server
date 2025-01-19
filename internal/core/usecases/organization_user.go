package usecases

import (
	"server/infrastructure/app"
	"server/internal/core/models"
	"server/internal/core/repositories"
	"strconv"
)

type OrganizationUserUsecase interface {
	CreateOrganizationUser(organizationUser *models.OrganizationUser, requesterUserId int) (*models.OrganizationUser, error)
	DeleteOrganizationUser(organizationUser *models.OrganizationUser, requesterUserId, requesterUserLevelId int) error
}

type organizationUserUsecase struct {
	organizationRepo             repositories.OrganizationRepository
	organizationUserRepo         repositories.OrganizationUserRepository
	organizationCategoryUserRepo repositories.OrganizationCategoryUserRepository
	organizationLogRepo          repositories.OrganizationLogRepository
}

func NewOrganizationUserUsecase(organizationRepo repositories.OrganizationRepository, organizationUserRepo repositories.OrganizationUserRepository, organizationCategoryUserRepo repositories.OrganizationCategoryUserRepository, organizationLogRepo repositories.OrganizationLogRepository) OrganizationUserUsecase {
	return &organizationUserUsecase{organizationRepo, organizationUserRepo, organizationCategoryUserRepo, organizationLogRepo}
}

func (u *organizationUserUsecase) CreateOrganizationUser(organizationUser *models.OrganizationUser, requesterUserId int) (*models.OrganizationUser, error) {
	txOrganizationUserRepo, err := u.organizationUserRepo.BeginLog()
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			txOrganizationUserRepo.Rollback()
		}
	}()

	exists, err := txOrganizationUserRepo.CheckOrganizationUserExists(organizationUser.OrganizationId, organizationUser.UserId)
	if err != nil {
		txOrganizationUserRepo.Rollback()
		return nil, err
	}

	if exists {
		txOrganizationUserRepo.Rollback()
		return nil, app.ErrOrganizationUserExists
	}

	organizationUser.CreatedBy = requesterUserId
	organizationUser.UpdatedBy = requesterUserId
	newOrganizationUser, err := txOrganizationUserRepo.CreateOrganizationUser(organizationUser)
	if err != nil {
		txOrganizationUserRepo.Rollback()
		return nil, err
	}

	if err := txOrganizationUserRepo.Commit(); err != nil {
		return nil, err
	}

	organizationLog := &models.OrganizationLog{
		OrganizationId: newOrganizationUser.OrganizationId,
		Action:         "Invited",
		Description:    "Invited user " + strconv.Itoa(newOrganizationUser.UserId) + " to the organization",
		CreatedBy:      requesterUserId,
	}

	if _, err := u.organizationLogRepo.CreateOrganizationLog(organizationLog); err != nil {
		return nil, err
	}

	return newOrganizationUser, nil
}

func (u *organizationUserUsecase) DeleteOrganizationUser(organizationUser *models.OrganizationUser, requesterUserId, requesterUserLevelId int) error {
	if organizationUser.UserId == requesterUserId {
		return app.ErrOrganizationUserDeleteMyself
	}

	if repositories.OwnerUserLevel.UserLevelId == requesterUserLevelId {
		return app.ErrOrganizationDeleteOwner
	}

	txOrganizationUserRepo, err := u.organizationUserRepo.BeginLog()
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			txOrganizationUserRepo.Rollback()
		}
	}()

	exists, err := txOrganizationUserRepo.CheckOrganizationUserExists(organizationUser.OrganizationId, organizationUser.UserId)
	if err != nil {
		txOrganizationUserRepo.Rollback()
		return err
	}

	if !exists {
		txOrganizationUserRepo.Rollback()
		return app.ErrUserNotFound
	}

	organizationCategoryUser, err := u.organizationCategoryUserRepo.GetOrganizationCategoryUserByUserId(organizationUser.UserId)
	if err != nil && organizationCategoryUser != nil {
		txOrganizationUserRepo.Rollback()
		return err
	}

	if organizationCategoryUser != nil {
		if err := u.organizationCategoryUserRepo.DeleteOrganizationCategoryUser(organizationCategoryUser); err != nil {
			txOrganizationUserRepo.Rollback()
			return err
		}
	}

	if err := txOrganizationUserRepo.DeleteOrganizationUser(organizationUser); err != nil {
		txOrganizationUserRepo.Rollback()
		return err
	}

	if err := txOrganizationUserRepo.Commit(); err != nil {
		return err
	}

	organizationLog := &models.OrganizationLog{
		OrganizationId: organizationUser.OrganizationId,
		Action:         "Removed",
		Description:    "Removed user " + strconv.Itoa(organizationUser.UserId) + " from the organization",
		CreatedBy:      requesterUserId,
	}

	if _, err := u.organizationLogRepo.CreateOrganizationLog(organizationLog); err != nil {
		return err
	}

	return nil
}
