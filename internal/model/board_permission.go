package model

import "gorm.io/gorm"

type BoardPermission struct {
	gorm.Model
	BoardID      uint `gorm:"not null" json:"board_id"`
	RoleID       uint `gorm:"not null" json:"role_id"`
	CanRead      bool `gorm:"default:false" json:"can_read"`
	CanWrite     bool `gorm:"default:false" json:"can_write"`
	CanComment   bool `gorm:"default:false" json:"can_comment"`
	CanDeleteOwn bool `gorm:"default:true" json:"can_delete_own"`
	CanDeleteAll bool `gorm:"default:false" json:"can_delete_all"`
}
