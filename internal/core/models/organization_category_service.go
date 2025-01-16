package models

import (
	"time"

	"gorm.io/gorm"
)

type OrganizationCategoryService struct {
	OrganizationCategoryServiceId int            `json:"organization_category_service_id" gorm:"primaryKey;column:organization_category_service_id;not null"`
	OrganizationCategoryId        int            `json:"organization_category_id" gorm:"column:organization_category_id;not null"`
	OrganizationServiceId         int            `json:"organization_service_id" gorm:"column:organization_service_id;not null"`
	OrganizationId                int            `json:"organization_id" gorm:"column:organization_id;not null"`
	CreatedAt                     time.Time      `json:"created_at" gorm:"column:created_at;not null"`
	CreatedBy                     int            `json:"created_by" gorm:"column:created_by;not null"`
	UpdatedAt                     time.Time      `json:"updated_at" gorm:"column:updated_at"`
	UpdatedBy                     int            `json:"updated_by" gorm:"column:updated_by"`
	DeletedAt                     gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}
