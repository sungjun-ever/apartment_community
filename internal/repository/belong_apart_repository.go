package repository

import (
	"apart_community/internal/model"

	"gorm.io/gorm"
)

type BelongApartRepository interface {
	BaseRepository[model.UserBelongApartment]
}

type belongApartRepository struct {
	BaseRepository[model.UserBelongApartment]
	db *gorm.DB
}

func NewBelongApartRepository(db *gorm.DB) BelongApartRepository {
	return &belongApartRepository{
		BaseRepository: NewBaseRepository[model.UserBelongApartment](db),
		db:             db,
	}
}
