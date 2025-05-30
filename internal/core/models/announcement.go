package models

import (
	"time"

	"gorm.io/gorm"
)

type Announcement struct {
	AnnouncementId   int            `json:"announcement_id" gorm:"column:announcement_id; primaryKey;"`
	SiteId           int            `json:"site_id" gorm:"column:site_id"`
	Title            string         `json:"title" gorm:"column:title"`
	ShortDescription string         `json:"short_description" gorm:"column:short_description"`
	ImageUrl         string         `json:"image_url" gorm:"column:image_url"`
	LinkUrl          string         `json:"link_url" gorm:"column:link_url"`
	IsPin            bool           `json:"is_pin" gorm:"column:is_pin"`
	CreatedAt        time.Time      `json:"created_at" gorm:"column:created_at"`
	CreatedBy        int            `json:"created_by" gorm:"column:created_by"`
	DeletedAt        gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}
