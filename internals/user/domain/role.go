package domain

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	RoleName string `gorm:"not null" json:"roleName"`
	RoleCode string `gorm:"not null; unique" json:"roleCode"`
}
