package domain

import (
	"gorm.io/gorm"
)

type UserApartmentRole struct {
	gorm.Model
	UserID     uint `gorm:"not null"`
	UnitID     uint `gorm:"not null"`
	No         uint `gorm:"not null"`
	IsVerified bool `gorm:"default:false"`
}
