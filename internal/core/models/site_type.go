package models

import (
	"time"

	"gorm.io/gorm"
)

type SiteType struct {
	SiteTypeId  int            `json:"site_type_id" gorm:"column:site_type_id; primaryKey"`
	Slug        string         `json:"slug" gorm:"column:slug"`
	Description string         `json:"description" gorm:"column:description"`
	CreatedAt   time.Time      `json:"created_at" gorm:"column:created_at"`
	CreatedBy   int            `json:"created_by" gorm:"column:created_by"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"column:updated_at"`
	UpdatedBy   int            `json:"updated_by" gorm:"column:updated_by"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}
