package domain

import (
	"apart_community/internals/apartment/domain"

	"gorm.io/gorm"
)

type UserBelongApartment struct {
	gorm.Model
	UserID      uint
	ApartmentID uint
	RoleID      uint
	Unit        uint             `gorm:"not null"`
	No          uint             `gorm:"not null"`
	IsVerified  bool             `gorm:"default:false"`
	User        User             `gorm:"foreignKey:UserID"`
	Apartment   domain.Apartment `gorm:"foreignKey:ApartmentID"`
	Role        Role             `gorm:"foreignKey:RoleID"`
}
