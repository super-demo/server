package repositories

import (
	"server/internal/core/models"

	"gorm.io/gorm"
)

type PeopleRoleRepository interface {
	BeginLog() (PeopleRoleRepository, error)
	Commit() error
	Rollback() error
	CreatePeopleRole(role *models.PeopleRole) (*models.PeopleRole, error)
	CheckRoleExistsByName(name string) (bool, error)
	GetRoleListBySiteId(id int) ([]models.PeopleRole, error)
	GetUserByEmail(email string) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
}

type peopleRoleRepository struct {
	db *gorm.DB
	tx *gorm.DB
}

func NewPeopleRoleRepository(db *gorm.DB) PeopleRoleRepository {
	return &peopleRoleRepository{db: db}
}

func (r *peopleRoleRepository) BeginLog() (PeopleRoleRepository, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &peopleRoleRepository{db: r.db, tx: tx}, nil
}

func (r *peopleRoleRepository) Commit() error {
	return r.tx.Commit().Error
}

func (r *peopleRoleRepository) Rollback() error {
	return r.tx.Rollback().Error
}

func (r *peopleRoleRepository) CreatePeopleRole(role *models.PeopleRole) (*models.PeopleRole, error) {
	if err := r.db.Create(role).Error; err != nil {
		return role, err
	}

	return role, nil
}

func (r *peopleRoleRepository) CheckRoleExistsByName(name string) (bool, error) {
	var role models.PeopleRole

	err := r.db.Where("slug = ?", name).First(&role).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// mock

func (r *peopleRoleRepository) GetRoleListBySiteId(id int) ([]models.PeopleRole, error) {
	var roles []models.PeopleRole

	err := r.db.Where("site_id = ?", id).Find(&roles).Error
	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (r *peopleRoleRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *peopleRoleRepository) UpdateUser(user *models.User) (*models.User, error) {
	if err := r.db.Where("user_id = ?", user.UserId).Updates(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
