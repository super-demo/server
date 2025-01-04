package models

import "time"

type User struct {
	UserId      int       `json:"user_id" gorm:"primaryKey;column:user_id;not null"`
	GoogleToken string    `json:"google_token" gorm:"column:google_token"`
	AvatarUrl   string    `json:"avatar_url" gorm:"column:avatar_url"`
	Name        string    `json:"name" gorm:"column:name;not null"`
	Nickname    string    `json:"nickname" gorm:"column:nickname"`
	Role        string    `json:"role" gorm:"column:role"`
	Email       string    `json:"email" gorm:"column:email;uniqueIndex;not null"`
	PhoneNumber string    `json:"phone_number" gorm:"column:phone_number"`
	Birthday    time.Time `json:"birthday" gorm:"column:birthday"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at;not null"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at;not null"`
}

type UserInfoResponse struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture"`
	HD            string `json:"hd"`
}
