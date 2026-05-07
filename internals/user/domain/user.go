package domain

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	PublicID string  `gorm:"uniqueIndex;type:char(26);not null"`
	Email    string  `gorm:"uniqueIndex; not null" json:"email"`
	Password string  `josn:"password; not null" json:"password"`
	Profile  Profile `gorm:"foreignKey:UserID"`
	Roles    []Role  `gorm:"many2many:user_belong_apartments"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	u.PublicID = ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()
	return
}

func (u *User) AfterDelete(tx *gorm.DB) error {
	if err := tx.Model(&User{}).Where("id = ?", u.ID).Delete(&User{}).Error; err != nil {
		return err
	}

	if err := tx.Model(&UserBelongApartment{}).Where("user_id = ?", u.ID).Delete(&UserBelongApartment{}).Error; err != nil {
		return err
	}

	return nil
}
