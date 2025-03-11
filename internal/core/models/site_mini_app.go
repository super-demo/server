package models

import "time"

type SiteMiniApp struct {
	SiteMiniAppId int       `json:"site_mini_app_id" gorm:"column:site_mini_app_id; primaryKey"`
	SiteId        int       `json:"site_id" gorm:"column:site_id"`
	Slug          string    `json:"slug" gorm:"column:slug"`
	Description   string    `json:"description" gorm:"column:description"`
	LinkUrl       string    `json:"link_url" gorm:"column:link_url"`
	IsActive      bool      `json:"is_active" gorm:"column:is_active"`
	CreatedAt     time.Time `json:"created_at" gorm:"column:created_at"`
	CreatedBy     int       `json:"created_by" gorm:"column:created_by"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"column:updated_at"`
	UpdatedBy     int       `json:"updated_by" gorm:"column:updated_by"`
	DeletedAt     time.Time `json:"deleted_at" gorm:"column:deleted_at"`
}
