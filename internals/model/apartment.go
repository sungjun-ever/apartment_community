package model

import "gorm.io/gorm"

type Apartment struct {
	gorm.Model
	PublicID string  `gorm:"uniqueIndex;type:char(26);not null"`
	KaptCode string  `gorm:"not null; unique" json:"kaptCode"`
	KaptName string  `gorm:"not null" json:"kaptName"`
	As1      *string `json:"as1"`
	As2      *string `json:"as2"`
	As3      *string `json:"as3"`
	As4      *string `json:"as4"`
	BjdCode  *string `json:"bjdCode"`
}
