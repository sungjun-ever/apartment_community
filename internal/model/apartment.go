package model

import "gorm.io/gorm"

type Apartment struct {
	gorm.Model
	UUID     string `gorm:"unique, not null" json:"uuid"`
	KaptCode string `json:"kapt_code"`
	KaptName string `json:"kapt_name"`
	As1      string `json:"as1"`
	As2      string `json:"as2"`
	As3      string `json:"as3"`
	As4      string `json:"as4"`
	BjdCode  string `json:"bjd_code"`
}
