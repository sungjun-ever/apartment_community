package domain

import "gorm.io/gorm"

type Building struct {
	gorm.Model
	ApartmentID uint    `gorm:"index" json:"apartmentId"`
	Name        *string `json:"name"`
	Units       []Unit  `gorm:"foreignKey:BuildingID"`
}
