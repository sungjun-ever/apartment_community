package domain

import "time"

// UserUnitRole
/**
OWNER: 집 주인
RESIDENT: 거주자
*/
type UserUnitRole struct {
	UserID     uint       `gorm:"not null;index:idx_user_unit_role"`
	UnitID     uint       `gorm:"not null;index:idx_user_unit_role"`
	RoleName   string     `gorm:"not null" json:"roleName"`
	IsVerified bool       `gorm:"default:false"`
	CreatedAt  time.Time  `gorm:"autoCreateTime"`
	UpdatedAt  time.Time  `gorm:"autoUpdateTime"`
	DeletedAt  *time.Time `gorm:"index"`
}
