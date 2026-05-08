package domain

import "gorm.io/gorm"

type Unit struct {
	gorm.Model
	BuildingID uint   `gorm:"not null" json:"buildingId"`
	No         string `gorm:"not null" json:"no"`
	SizeType   string `gorm:"not null" json:"sizeType"`
}
