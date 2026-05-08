package domain

import "gorm.io/gorm"

type Apartment struct {
	gorm.Model
	PublicID  string  `gorm:"uniqueIndex;type:char(26);not null"`
	Code      string  `gorm:"not null; unique" json:"Code"`
	Name      string  `gorm:"not null" json:"Name"`
	Address   *string `json:"address"`
	BjdCode   *string `json:"bjdCode"`
	Buildings []Building
}
