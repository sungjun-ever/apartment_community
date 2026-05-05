package model

import "gorm.io/gorm"

type Attachment struct {
	gorm.Model
	PublicID      string `gorm:"uniqueIndex;type:char(26);not null"`
	TargetID      uint   `gorm:"not null"`
	TargetType    string `gorm:"not null"`
	UploaderID    uint   `gorm:"not null"`
	SavePath      string `gorm:"not null"`
	FileSize      uint   `gorm:"not null"`
	MimeType      string `gorm:"not null"`
	OriginalName  string `gorm:"not null"`
	Extension     string `gorm:"not null"`
	DownloadCount uint   `gorm:"default:0"`
	IsPublic      bool   `gorm:"not null; default:false"`
}
