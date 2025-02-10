package models

type UserLevel struct {
	UserLevelId int    `json:"user_level_id" gorm:"column:user_level_id; primaryKey"`
	Slug        string `json:"slug" gorm:"column:slug"`
	Description string `json:"description" gorm:"column:description"`
}
