package models

import (
	"time"

	"gorm.io/gorm"
)

type SiteTree struct {
	SiteTreeId   int            `json:"site_tree_id" gorm:"column:site_tree_id; primaryKey"`
	SiteParentId int            `json:"site_parent_id" gorm:"column:site_parent_id"`
	SiteChildId  int            `json:"site_child_id" gorm:"column:site_child_id"`
	CreatedAt    time.Time      `json:"created_at" gorm:"column:created_at"`
	CreteaBy     int            `json:"created_by" gorm:"column:created_by"`
	UpdatedAt    time.Time      `json:"updated_at" gorm:"column:updated_at"`
	UpdateBy     int            `json:"updated_by" gorm:"column:updated_by"`
	DeleteAt     gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}
