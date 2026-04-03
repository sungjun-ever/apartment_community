package model

import "gorm.io/gorm"

type PostComment struct {
	gorm.Model
	PostID  uint   `gorm:"not null" json:"post_id"`
	UserID  uint   `gorm:"not null" json:"user_id"`
	Content string `gorm:"not null" json:"content"`
}
