package models

import (
	"time"

	"gorm.io/gorm"
)

type Site struct {
	SiteId      int            `json:"site_id" gorm:"column:site_id; primaryKey"`
	SiteTypeId  int            `json:"site_type_id" gorm:"column:site_type_id"`
	Name        string         `json:"name" gorm:"column:name"`
	Description string         `json:"description" gorm:"column:description"`
	CreatedAt   time.Time      `json:"created_at" gorm:"column:created_at"`
	CreteaBy    int            `json:"created_by" gorm:"column:created_by"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"column:updated_at"`
	UpdateBy    int            `json:"updated_by" gorm:"column:updated_by"`
	DeleteAt    gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}
