package models

import "time"

type User struct {
	UserId      int       `json:"user_id" gorm:"primaryKey;column:user_id;not null"`
	UserLevelId int       `json:"user_level_id" gorm:"column:user_level_id;not null"`
	SubRoleId   int       `json:"sub_role_id" gorm:"column:sub_role_id;not null"`
	SiteId      int       `json:"site_id" gorm:"column:site_id"`
	GoogleToken string    `json:"google_token" gorm:"column:google_token"`
	AvatarUrl   string    `json:"avatar_url" gorm:"column:avatar_url"`
	Name        string    `json:"name" gorm:"column:name;not null"`
	Nickname    string    `json:"nickname" gorm:"column:nickname"`
	Email       string    `json:"email" gorm:"column:email;uniqueIndex;not null"`
	PhoneNumber string    `json:"phone_number" gorm:"column:phone_number"`
	BirthDate   time.Time `json:"birth_date" gorm:"column:birth_date"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at;not null"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at;not null"`
}

type BulkImportUser struct {
	Name            string `json:"name"`
	Nickname        string `json:"nickname"`
	Email           string `json:"email"`
	SiteUserLevelId string `json:"site_user_level_id"`
}

type BulkImportFailure struct {
	Name            string `json:"name"`
	Nickname        string `json:"nickname"`
	Email           string `json:"email"`
	SiteUserLevelId string `json:"site_user_level_id"`
	Message         string `json:"message"`
}

type UserInfoResponse struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture"`
	HD            string `json:"hd"`
}

type BulkImportResponse struct {
	SuccessCount int                 `json:"success_count"`
	FailedCount  int                 `json:"failed_count"`
	Failures     []BulkImportFailure `json:"failures"`
}
