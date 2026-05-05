package model

import "gorm.io/gorm"

type Apartment struct {
	gorm.Model
	PublicID string `gorm:"uniqueIndex;type:char(26);not null"`
	KaptCode string `gorm:"not null; unique"`
	KaptName string `gorm:"not null"`
	As1      *string
	As2      *string
	As3      *string
	As4      *string
	BjdCode  *string
}
