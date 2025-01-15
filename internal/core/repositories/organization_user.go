package repositories

import (
	"server/internal/core/models"

	"gorm.io/gorm"
)

type OrganizationUserRepository interface {
	BeginLog() (OrganizationUserRepository, error)
	Commit() error
	Rollback() error
	CreateOrganizationUser(organizationUser *models.OrganizationUser) (*models.OrganizationUser, error)
	CheckOrganizationUserExists(organizationUser *models.OrganizationUser, userId int) (bool, error)
	GetOrganizationUserById(id int) (*models.OrganizationUser, error)
	GetOrganizationUserByEmail(email string) (*models.OrganizationUser, error)
	GetOrganizationUserListByOrganizationId(organizationId int) (*[]models.OrganizationUser, error)
	UpdateOrganizationUser(organizationUser *models.OrganizationUser) (*models.OrganizationUser, error)
	DeleteOrganizationUser(organizationUser *models.OrganizationUser) error
	DeleteOrganizationUserByOrganizationId(organizationUser *models.OrganizationUser) error
}

type organizationUserRepository struct {
	db *gorm.DB
	tx *gorm.DB
}

func NewOrganizationUserRepository(db *gorm.DB) OrganizationUserRepository {
	return &organizationUserRepository{db: db}
}

func (r *organizationUserRepository) BeginLog() (OrganizationUserRepository, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &organizationUserRepository{db: tx, tx: tx}, nil
}

func (r *organizationUserRepository) Commit() error {
	return r.tx.Commit().Error
}

func (r *organizationUserRepository) Rollback() error {
	return r.tx.Rollback().Error
}

func (r *organizationUserRepository) CreateOrganizationUser(organizationUser *models.OrganizationUser) (*models.OrganizationUser, error) {
	if err := r.db.Create(organizationUser).Error; err != nil {
		return organizationUser, err
	}

	return organizationUser, nil
}

func (r *organizationUserRepository) CheckOrganizationUserExists(organizationUser *models.OrganizationUser, userId int) (bool, error) {
	err := r.db.Where("organization_id = ? AND user_id = ?", organizationUser.OrganizationId, userId).First(&organizationUser).Error
	if err != nil {
		return false, err
	}

	return true, nil
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

func (r *organizationUserRepository) UpdateOrganizationUser(organizationUser *models.OrganizationUser) (*models.OrganizationUser, error) {
	if err := r.db.Where("organization_user_id = ?", organizationUser.OrganizationUserId).Updates(organizationUser).Error; err != nil {
		return organizationUser, err
	}

	return organizationUser, nil
}

func (r *organizationUserRepository) DeleteOrganizationUser(organizationUser *models.OrganizationUser) error {
	if err := r.db.Where("organization_user_id = ?", organizationUser.OrganizationUserId).Delete(organizationUser).Error; err != nil {
		return err
	}

	return nil
}

func (r *organizationUserRepository) DeleteOrganizationUserByOrganizationId(organizationUser *models.OrganizationUser) error {
	if err := r.db.Where("organization_id = ?", organizationUser.OrganizationId).Delete(organizationUser).Error; err != nil {
		return err
	}

	return nil
}
