package models

import (
	"time"

	"gorm.io/gorm"
)

type OrganizationService struct {
	OrganizationServiceId int            `json:"organization_service_id" gorm:"primaryKey;column:organization_service_id;not null"`
	OrganizationId        int            `json:"organization_id" gorm:"column:organization_id;not null"`
	ServiceId             int            `json:"service_id" gorm:"column:service_id;not null"`
	Slug                  string         `json:"slug" gorm:"column:slug;not null"`
	Description           string         `json:"description" gorm:"column:description"`
	WebHookUrl            string         `json:"web_hook_url" gorm:"column:web_hook_url"`
	IsActive              bool           `json:"is_active" gorm:"column:is_active;default:true;not null"`
	CreatedAt             time.Time      `json:"created_at" gorm:"column:created_at;not null"`
	CreatedBy             int            `json:"created_by" gorm:"column:created_by;not null"`
	UpdatedAt             time.Time      `json:"updated_at" gorm:"column:updated_at;not null"`
	UpdatedBy             int            `json:"updated_by" gorm:"column:updated_by;not null"`
	DeletedAt             gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at;type:timestamp;not null"`
}
