package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type BoardSetting struct {
	AllowComment    bool     `json:"allow_comment"`
	AllowAttachment bool     `json:"allow_attachment"`
	MaxFileSizeMB   uint     `json:"max_file_size_mb"`
	UserRoles       []string `json:"user_roles"`
}

type Board struct {
	gorm.Model
	ApartmentID uint           `gorm:"not null" json:"apartment_id"`
	Name        string         `gorm:"not null" json:"name"`
	Slug        string         `gorm:"unique,not null" json:"slug"`
	Description string         `json:"description"`
	Setting     datatypes.JSON `gorm:"type:json" json:"setting"`
}
