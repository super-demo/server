package models

import (
	"time"

	"gorm.io/gorm"
)

type SiteUser struct {
	SiteUserId  int            `json:"site_user_id" gorm:"column:site_user_id; primaryKey"`
	SiteId      int            `json:"site_id" gorm:"column:site_id"`
	UserId      int            `json:"user_id" gorm:"column:user_id"`
	UserLevelId int            `json:"user_level_id" gorm:"column:user_level_id"`
	IsActive    bool           `json:"is_active" gorm:"column:is_active"`
	CreatedAt   time.Time      `json:"created_at" gorm:"column:created_at"`
	CreteaBy    int            `json:"created_by" gorm:"column:created_by"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"column:updated_at"`
	UpdateBy    int            `json:"updated_by" gorm:"column:updated_by"`
	DeleteAt    gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}
