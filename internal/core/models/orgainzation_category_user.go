package models

import (
	"time"

	"gorm.io/gorm"
)

type OrganizationCategoryUser struct {
	OrganizationCategoryUserId int            `json:"organization_category_user_id" gorm:"primaryKey;column:organization_category_user_id;not null"`
	OrganizationCategoryId     int            `json:"organization_category_id" gorm:"column:organization_category_id;not null"`
	UserId                     int            `json:"user_id" gorm:"column:user_id;not null"`
	UserLevelId                int            `json:"user_level_id" gorm:"column:user_level_id;not null"`
	IsActive                   bool           `json:"is_active" gorm:"column:is_active;not null"`
	CreatedAt                  time.Time      `json:"created_at" gorm:"column:created_at;not null"`
	CreatedBy                  int            `json:"created_by" gorm:"column:created_by;not null"`
	UpdatedAt                  time.Time      `json:"updated_at" gorm:"column:updated_at"`
	UpdatedBy                  int            `json:"updated_by" gorm:"column:updated_by"`
	DeletedAt                  gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}
