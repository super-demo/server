package repositories

import (
	"server/internal/core/models"

	"gorm.io/gorm"
)

type OrganizationUserRepository interface {
	CreateOrganizationUser(organizationUser *models.OrganizationUser) (*models.OrganizationUser, error)
	UpdateOrganizationUser(organizationUser *models.OrganizationUser) (*models.OrganizationUser, error)
	DeleteOrganizationUser(id int) error
	GetOrganizationUserById(id int) (*models.OrganizationUser, error)
	GetOrganizationUserByEmail(email string) (*models.OrganizationUser, error)
	GetOrganizationUserListByOrganizationId(organizationId int) (*[]models.OrganizationUser, error)
	DeleteOrganizationUserByOrganizationId(organizationId int) error
}

type organizationUserRepository struct {
	db *gorm.DB
}

func NewOrganizationUserRepository(db *gorm.DB) OrganizationUserRepository {
	return &organizationUserRepository{db}
}

func (r *organizationUserRepository) CreateOrganizationUser(organizationUser *models.OrganizationUser) (*models.OrganizationUser, error) {
	if err := r.db.Create(organizationUser).Error; err != nil {
		return organizationUser, err
	}

	return organizationUser, nil
}

func (r *organizationUserRepository) UpdateOrganizationUser(organizationUser *models.OrganizationUser) (*models.OrganizationUser, error) {
	if err := r.db.Save(organizationUser).Error; err != nil {
		return organizationUser, err
	}

	return organizationUser, nil
}

func (r *organizationUserRepository) DeleteOrganizationUser(id int) error {
	if err := r.db.Where("organization_user_id = ?", id).Delete(&models.OrganizationUser{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *organizationUserRepository) GetOrganizationUserById(id int) (*models.OrganizationUser, error) {
	organizationUser := new(models.OrganizationUser)

	if err := r.db.Where("organization_user_id = ?", id).First(organizationUser).Error; err != nil {
		return nil, err
	}

	return organizationUser, nil
}

func (r *organizationUserRepository) GetOrganizationUserByEmail(email string) (*models.OrganizationUser, error) {
	var organizationUser models.OrganizationUser

	if err := r.db.Where("email = ?", email).First(&organizationUser).Error; err != nil {
		return nil, err
	}

	return &organizationUser, nil
}

func (r *organizationUserRepository) GetOrganizationUserListByOrganizationId(organizationId int) (*[]models.OrganizationUser, error) {
	organizationUsers := new([]models.OrganizationUser)

	if err := r.db.Where("organization_id = ?", organizationId).Find(organizationUsers).Error; err != nil {
		return nil, err
	}

	return organizationUsers, nil
}

func (r *organizationUserRepository) DeleteOrganizationUserByOrganizationId(organizationId int) error {
	if err := r.db.Where("organization_id = ?", organizationId).Delete(&models.OrganizationUser{}).Error; err != nil {
		return err
	}
	return nil
}
