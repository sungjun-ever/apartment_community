package domain

import "gorm.io/gorm"

// UserRole
/**
System Admin - 플랫폼 전체 관리자
Apartment Admin - 아파트 관리자
Building Admin - 단지 관리자
*/
type UserRole struct {
	gorm.Model
	RoleName string `gorm:"not null" json:"roleName"`
}
