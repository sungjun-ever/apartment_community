package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     string  `gorm:"unique, not null" json:"uuid"`
	Email    string  `gorm:"unique, not null" json:"email" binding:"required,email"`
	Password string  `gorm:"not null" json:"password" binding:"required,min=6,max=20"`
	Status   int     `gorm:"default:0" json:"status"`
	Profile  Profile `gorm:"foreignKey:UserID" json:"profile"`
}
