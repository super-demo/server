package models

import (
	"time"

	"gorm.io/gorm"
)

type SitePeople struct {
	SitePeopleId int            `json:"site_user_id" gorm:"column:site_people_id; primaryKey"`
	SiteId       int            `json:"site_id" gorm:"column:site_id"`
	UserId       int            `json:"user_id" gorm:"column:user_id"`
	IsActive     bool           `json:"is_active" gorm:"column:is_active"`
	CreatedAt    time.Time      `json:"created_at" gorm:"column:created_at"`
	CreatedBy    int            `json:"created_by" gorm:"column:created_by"`
	UpdatedAt    time.Time      `json:"updated_at" gorm:"column:updated_at"`
	UpdatedBy    int            `json:"updated_by" gorm:"column:updated_by"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}

type SitePeopleJoinTable struct {
	SitePeopleId int            `json:"site_people_id"`
	SiteId       int            `json:"site_id"`
	UserId       int            `json:"user_id"`
	User         User           `json:"user" gorm:"foreignKey:UserId;references:UserId"`
	IsActive     bool           `json:"is_active"`
	CreatedAt    time.Time      `json:"created_at"`
	CreatedBy    int            `json:"created_by"`
	UpdatedAt    time.Time      `json:"updated_at"`
	UpdatedBy    int            `json:"updated_by"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}

type CreateSitePeopleRequest struct {
	SiteId    int    `json:"site_id" binding:"required"`
	Email     string `json:"email" binding:"required"`
	SubRoleId int    `json:"sub_role_id" binding:"required"`
}
