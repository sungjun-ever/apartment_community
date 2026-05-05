package model

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	UserID   uint
	nickname string      `gorm:"default:anonymous; unique; not null"`
	ImageID  *Attachment `gorm:"polymorphic:Target"`
}
