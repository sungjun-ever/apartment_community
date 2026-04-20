package model

import "gorm.io/gorm"

type UserBelongApartment struct {
	gorm.Model
	UserID      uint  `gorm:"unique,not null" json:"user_id" binding:"required"`
	ApartmentID uint  `gorm:"unique,not null" json:"apartment_id" binding:"required"`
	RoleID      uint  `gorm:"unique,not null" json:"role_id" binding:"required"`
	Unit        *uint `json:"unit"`
	No          *uint `json:"no"`
	IsVerified  bool  `gorm:"default:false" json:"is_verified"`
}
