package models

import "time"

type Organization struct {
	OrganizationId int       `json:"organization_id" gorm:"primaryKey;column:organization_id;not null"`
	Name           string    `json:"name" gorm:"column:name;not null"`
	Description    string    `json:"description" gorm:"column:description"`
	Url            string    `json:"url" gorm:"column:url"`
	ImageUrl       string    `json:"image_url" gorm:"column:image_url"`
	CreatedAt      time.Time `json:"created_at" gorm:"column:created_at;not null"`
	CreatedBy      int       `json:"created_by" gorm:"column:created_by;not null"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"column:updated_at;not null"`
	UpdatedBy      int       `json:"updated_by" gorm:"column:updated_by;not null"`
}
