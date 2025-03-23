package models

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Site struct {
	SiteId           int            `json:"site_id" gorm:"column:site_id; primaryKey"`
	SiteTypeId       int            `json:"site_type_id" gorm:"column:site_type_id"`
	Name             string         `json:"name" gorm:"column:name"`
	Description      string         `json:"description" gorm:"column:description"`
	ShortDescription string         `json:"short_description" gorm:"column:short_description"`
	Url              string         `json:"url" gorm:"column:url"`
	ImageUrl         string         `json:"image_url" gorm:"column:image_url"`
	CreatedAt        time.Time      `json:"created_at" gorm:"column:created_at"`
	CreatedBy        int            `json:"created_by" gorm:"column:created_by"`
	UpdatedAt        time.Time      `json:"updated_at" gorm:"column:updated_at"`
	UpdatedBy        int            `json:"updated_by" gorm:"column:updated_by"`
	DeletedAt        gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}

type Workspace struct {
	SiteId           int            `gorm:"primaryKey;autoIncrement" json:"site_id"`
	SiteParentID     int            `gorm:"index" json:"site_parent_id"`
	SiteTypeId       int            `gorm:"not null;index" json:"site_type_id"`
	Name             string         `gorm:"type:varchar(255);not null" json:"name"`
	Description      string         `gorm:"type:text" json:"description,omitempty"`
	ShortDescription string         `gorm:"type:text" json:"short_description,omitempty"`
	Url              string         `gorm:"type:text" json:"url,omitempty"`
	ImageUrl         string         `gorm:"type:text" json:"image_url,omitempty"`
	CreatedAt        time.Time      `gorm:"autoCreateTime" json:"created_at"`
	CreatedBy        int            `gorm:"not null" json:"created_by"`
	UpdatedAt        time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	UpdatedBy        int            `gorm:"not null" json:"updated_by"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type PeopleRole struct {
	PeopleRoleId int            `gorm:"primaryKey;autoIncrement" json:"people_role_id"`
	Slug         string         `json:"slug" gorm:"column:slug"`
	Description  string         `json:"description" gorm:"column:description"`
	SiteId       int            `json:"site_id" gorm:"column:site_id"`
	CreatedAt    time.Time      `gorm:"autoCreateTime" json:"created_at"`
	CreatedBy    int            `gorm:"not null" json:"created_by"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	UpdatedBy    int            `gorm:"not null" json:"updated_by"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type GetWorkspaceList struct {
	SiteID           int            `gorm:"primaryKey;autoIncrement" json:"site_id"`
	SiteTypeID       int            `gorm:"not null;index" json:"site_type_id"`
	Name             string         `gorm:"type:varchar(255);not null" json:"name"`
	Description      string         `gorm:"type:text" json:"description,omitempty"`
	ShortDescription string         `gorm:"type:text" json:"short_description,omitempty"`
	Url              string         `gorm:"type:text" json:"url,omitempty"`
	ImageUrl         string         `gorm:"type:text" json:"image_url,omitempty"`
	CreatedAt        time.Time      `gorm:"autoCreateTime" json:"created_at"`
	CreatedBy        int            `gorm:"not null" json:"created_by"`
	UpdatedAt        time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	UpdatedBy        int            `gorm:"not null" json:"updated_by"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	SiteParentID     int            `gorm:"index" json:"site_parent_id"`
	SiteParentName   string         `gorm:"type:varchar(255)" json:"site_parent_name"`
	Depth            uint           `gorm:"not null;default:1" json:"depth"`
	Path             pq.Int64Array  `gorm:"type:integer[];not null" json:"path"`
}

type CreateSiteWorkspaceRequest struct {
	Site         Site `json:"site"`
	SiteParentId int  `json:"site_parent_id"`
}

type CreatePeopleRoleRequest struct {
	Slug        string `json:"slug" gorm:"column:slug"`
	Description string `json:"description" gorm:"column:description"`
	SiteId      int    `json:"site_id" gorm:"column:site_id"`
}

type UpdatePeopleRoleRequest struct {
	PeopleRoleId int    `json:"people_role_id"`
	Slug         string `json:"slug" gorm:"column:slug"`
	Description  string `json:"description" gorm:"column:description"`
	SiteId       int    `json:"site_id" gorm:"column:site_id"`
}
