package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	UUID        string `gorm:"unique, not null" json:"uuid"`
	BoardID     uint   `gorm:"not null" json:"board_id"`
	ApartmentID uint   `gorm:"not null" json:"apartment_id"`
	UserID      uint   `gorm:"not null" json:"user_id"`
	Title       string `gorm:"not null" json:"title"`
	Content     string `gorm:"not null" json:"content"`
	ViewCount   uint
}
