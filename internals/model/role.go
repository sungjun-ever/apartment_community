package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	RoleName string `gorm:"not null"`
	RoleCode string `gorm:"not null; unique"`
}
