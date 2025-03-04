package repositories

import (
	"server/internal/core/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	BeginLog() (UserRepository, error)
	Commit() error
	Rollback() error
	CreateUser(user *models.User) (*models.User, error)
	CheckUserExistsByEmail(email string) (bool, error)
	GetUserById(id int) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
	tx *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) BeginLog() (UserRepository, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &userRepository{db: r.db, tx: tx}, nil
}

func (r *userRepository) Commit() error {
	return r.tx.Commit().Error
}

func (r *userRepository) Rollback() error {
	return r.tx.Rollback().Error
}

func (r *userRepository) CreateUser(user *models.User) (*models.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) CheckUserExistsByEmail(email string) (bool, error) {
	var user models.User

	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (r *userRepository) GetUserById(id int) (*models.User, error) {
	user := new(models.User)

	if err := r.db.Where("user_id = ?", id).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) UpdateUser(user *models.User) (*models.User, error) {
	if err := r.db.Where("user_id = ?", user.UserId).Updates(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
