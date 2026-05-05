package model

import "gorm.io/gorm"

type UserbelongApartment struct {
	gorm.Model
	UserID      uint
	ApartmentID uint
	RoleID      uint
	Unit        uint      `gorm:"not null"`
	No          uint      `gorm:"not null"`
	IsVerified  bool      `gorm:"default:false"`
	User        User      `gorm:"foreignKey:UserID"`
	Apartment   Apartment `gorm:"foreignKey:ApartmentID"`
	Role        Role      `gorm:"foreignKey:RoleID"`
}
