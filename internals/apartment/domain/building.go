package domain

import "gorm.io/gorm"

type Building struct {
	gorm.Model
	ApartmentID uint    `gorm:"uniqueIndex;not null" json:"apartmentId"`
	Name        *string `json:"name"`
	Units       []Unit
}
