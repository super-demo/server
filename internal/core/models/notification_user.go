package models

type NotificationUser struct {
	NotificationUserId int  `json:"notification_user_id" gorm:"column:notification_user_id; primaryKey; autoIncrement"`
	NotificationId     int  `json:"notification_id" gorm:"column:notification_id"`
	UserId             int  `json:"user_id" gorm:"column:user_id"`
	IsRead             bool `json:"is_read" gorm:"column:is_read"`
}
