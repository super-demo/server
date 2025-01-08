package models

import "time"

type OrganizationLog struct {
	OrganizationLogId int       `json:"organization_log_id" gorm:"primaryKey;column:organization_log_id;not null"`
	OrganizationId    int       `json:"organization_id" gorm:"column:organization_id;not null"`
	Action            string    `json:"action" gorm:"column:action;not null"`
	Description       string    `json:"description" gorm:"column:description"`
	CreatedAt         time.Time `json:"created_at" gorm:"column:created_at;not null"`
	CreatedBy         int       `json:"created_by" gorm:"column:created_by;not null"`
}
