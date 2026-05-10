package domain

import (
	"apart_community/internals/common/utils/idgen"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	PublicID   string         `gorm:"uniqueIndex;type:char(26);not null"`
	Email      string         `gorm:"uniqueIndex; not null" json:"email"`
	Password   string         `gorm:"not null" json:"password"`
	Profile    Profile        `gorm:"foreignKey:UserID"`
	Roles      []UserRole     `gorm:"foreignKey:UserID"`
	UnitsRoles []UserUnitRole `gorm:"foreignKey:UserID"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.PublicID = idgen.GeneratePublicId()
	return
}

func (u *User) AfterDelete(tx *gorm.DB) error {
	if err := tx.Model(&User{}).Where("id = ?", u.ID).Delete(&User{}).Error; err != nil {
		return err
	}

	return nil
}
