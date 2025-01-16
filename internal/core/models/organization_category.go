package models

import (
	"time"

	"gorm.io/gorm"
)

type OrganizationCategory struct {
	OrganizationCategoryId int            `json:"organization_category_id" gorm:"primaryKey;column:organization_category_id;not null"`
	OrganizationId         int            `json:"organization_id" gorm:"column:organization_id;not null"`
	Name                   string         `json:"name" gorm:"column:name;not null"`
	Description            string         `json:"description" gorm:"column:description"`
	CreatedAt              time.Time      `json:"created_at" gorm:"column:created_at;not null"`
	CreatedBy              int            `json:"created_by" gorm:"column:created_by;not null"`
	UpdatedAt              time.Time      `json:"updated_at" gorm:"column:updated_at"`
	UpdatedBy              int            `json:"updated_by" gorm:"column:updated_by"`
	DeletedAt              gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}
