package domain

import (
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	UserID   uint
	Nickname string `gorm:"default:anonymous; unique; not null" json:"nickname"`
	ImageID  *uint  `gorm:"unique" json:"imageId"`
}
