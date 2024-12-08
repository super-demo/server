package models

import "time"

type User struct {
	UserId      int       `json:"user_id" gorm:"primaryKey;column:user_id;not null"`
	GoogleToken string    `json:"google_token" gorm:"column:google_token"`
	AvatarUrl   string    `json:"avatar_url" gorm:"column:avatar_url"`
	Name        string    `json:"name" gorm:"column:name;not null"`
	Nickname    string    `json:"nickname" gorm:"column:nickname"`
	Email       string    `json:"email" gorm:"column:email;uniqueIndex;not null"`
	PhoneNumber string    `json:"phone_number" gorm:"column:phone_number"`
	Role        string    `json:"role" gorm:"column:role"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at;not null"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at;not null"`
}
