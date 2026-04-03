package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	RoleName string `gorm:"unique,not null" json:"role_name" binding:"required,min=2,max=20"`
	RoleCode string `gorm:"unique,not null" json:"role_code" binding:"required"`
}
