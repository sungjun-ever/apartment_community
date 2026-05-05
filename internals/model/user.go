package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	PublicID string `gorm:"uniqueIndex;type:char(26);not null"`
	Email    string `gorm:"unqiueIndex; not null" json:"email"`
	Password string `josn:"password; not null"`
	Profile  Profile
	Roles    []Role `gorm:"many2many:user_belong_apartments"`
}
