package domain

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	RoleName string `gorm:"not null" json:"role_name"`
	RoleCode string `gorm:"not null; unique" json:"role_code"`
}
