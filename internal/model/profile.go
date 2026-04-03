package model

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	UserID         uint   `gorm:"unique,not null" json:"user_id" binding:"required"`
	Nickname       string `gorm:"unique" json:"nickname" binding:"required"`
	ProfileImageId uint   `gorm:"unique" json:"profile_image_id"`
}
