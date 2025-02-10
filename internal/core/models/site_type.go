package models

type SiteType struct {
	SiteTypeId  int    `json:"site_type_id" gorm:"column:site_type_id; primaryKey"`
	Slug        string `json:"slug" gorm:"column:slug"`
	Description string `json:"description" gorm:"column:description"`
}
