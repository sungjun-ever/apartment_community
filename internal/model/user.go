package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     string               `gorm:"unique, not null" json:"uuid"`
	Email    string               `gorm:"unique, not null" json:"email" binding:"required,email"`
	Password string               `gorm:"not null" json:"-" binding:"required,min=6,max=20"`
	Status   int                  `gorm:"default:0" json:"status"`
	Profile  *Profile             `gorm:"foreignKey:UserID" json:"profile"`
	Uba      *UserBelongApartment `gorm:"foreignKey:UserID" json: uba`
}

func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
	if err := tx.Where("user_id = ?", u.ID).Delete(&Profile{}).Error; err != nil {
		return err
	}

	if err := tx.Where("user_id = ?", u.ID).Delete(&UserBelongApartment{}).Error; err != nil {
		return err
	}

	return nil
}
