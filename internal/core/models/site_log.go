package models

import "time"

type SiteLog struct {
	SiteLogId int       `json:"site_log_id" gorm:"column:site_log_id; primaryKey"`
	SiteId    int       `json:"site_id" gorm:"column:site_id"`
	Action    string    `json:"action" gorm:"column:action"`
	Detail    string    `json:"detail" gorm:"column:detail"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	CreatedBy int       `json:"created_by" gorm:"column:created_by"`
}
