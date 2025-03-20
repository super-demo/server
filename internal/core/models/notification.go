package models

import "time"

type Notification struct {
	NotificationId int       `json:"notification_id" gorm:"column:notification_id; primaryKey"`
	SiteId         int       `json:"site_id" gorm:"column:site_id"`
	Action         string    `json:"action" gorm:"column:action"`
	Detail         string    `json:"detail" gorm:"column:detail"`
	IamgeUrl       string    `json:"image_url" gorm:"column:image_url"`
	CreatedAt      time.Time `json:"created_at" gorm:"column:created_at"`
	CreatedBy      int       `json:"created_by" gorm:"column:created_by"`
}
