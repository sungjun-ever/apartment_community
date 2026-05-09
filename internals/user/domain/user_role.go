package domain

import "time"

// UserRole
/**
System Admin - 플랫폼 전체 관리자
Apartment Admin - 아파트 관리자
Building Admin - 단지 관리자
GUEST - 비회원(미인증 사용자)
*/
type UserRole struct {
	UserID     uint       `gorm:"primaryKey;index:idx_user_role"`
	TargetType string     `gorm:"primaryKey;index:idx_target" json:"targetType"`
	TargetID   uint       `gorm:"primaryKey;index:idx_target" json:"targetID"`
	RoleName   string     `gorm:"not null;size:255" json:"roleName"`
	IsVerified bool       `gorm:"default:false"`
	CreatedAt  time.Time  `gorm:"autoCreateTime"`
	UpdatedAt  time.Time  `gorm:"autoUpdateTime"`
	DeletedAt  *time.Time `gorm:"index"`
}
