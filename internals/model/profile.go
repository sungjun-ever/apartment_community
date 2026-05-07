package model

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	UserID   uint
	Nickname string      `gorm:"default:anonymous; unique; not null" json:"nickname"`
	ImageID  *Attachment `gorm:"polymorphic:Target" json:"imageId"`
}
