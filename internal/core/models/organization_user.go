package models

import (
	"time"

	"gorm.io/gorm"
)

type OrganizationUser struct {
	OrganizationUserId int            `json:"organization_user_id" gorm:"primaryKey;column:organization_user_id;not null"`
	OrganizationId     int            `json:"organization_id" gorm:"column:organization_id;not null"`
	UserId             int            `json:"user_id" gorm:"column:user_id;not null"`
	UserLevelId        int            `json:"user_level_id" gorm:"column:user_level_id;not null"`
	IsActive           bool           `json:"is_active" gorm:"column:is_active;default:true;not null"`
	CreatedAt          time.Time      `json:"created_at" gorm:"column:created_at;not null"`
	CreatedBy          int            `json:"created_by" gorm:"column:created_by;not null"`
	UpdatedAt          time.Time      `json:"updated_at" gorm:"column:updated_at;not null"`
	UpdatedBy          int            `json:"updated_by" gorm:"column:updated_by;not null"`
	DeletedAt          gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at;type:timestamp;not null"`
}
