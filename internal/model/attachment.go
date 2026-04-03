package model

import "gorm.io/gorm"

type Attachment struct {
	gorm.Model
	UUID          string `gorm:"unique, not null" json:"uuid"`
	TargetType    string `gorm:"not null" json:"target_type"`
	TargetID      uint   `gorm:"not null" json:"target_id"`
	UploaderID    uint   `gorm:"not null" json:"uploader_id"`
	SavedPath     string `gorm:"not null" json:"saved_path"`
	FileSize      uint64 `gorm:"not null" json:"file_size"`
	MimeType      string `gorm:"not null" json:"mime_type"`
	OriginalName  string `gorm:"not null" json:"original_name"`
	Extension     string `gorm:"not null" json:"extension"`
	DownloadCount uint   `gorm:"default:0" json:"download_count"`
	IsPublic      bool   `gorm:"default:false" json:"is_public"`
}
