package domain

import (
	userDomain "apart_community/internals/user/domain"

	"gorm.io/gorm"
)

type Unit struct {
	gorm.Model
	BuildingID uint                      `gorm:"not null;index" json:"buildingId"`
	No         string                    `gorm:"not null" json:"no"`
	UnitsRoles []userDomain.UserUnitRole `gorm:"foreignKey:UnitID"`
}
