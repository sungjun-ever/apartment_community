package repository

import (
	"apart_community/internal/model"

	"gorm.io/gorm"
)

type ProfileRepository interface {
	BaseRepository[model.Profile]
	WithTrx(tx *gorm.DB) ProfileRepository
}

type profileRepository struct {
	BaseRepository[model.Profile]
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) ProfileRepository {
	return &profileRepository{
		BaseRepository: NewBaseRepository[model.Profile](db),
		db:             db,
	}
}

func (p profileRepository) WithTrx(tx *gorm.DB) ProfileRepository {
	return &profileRepository{
		BaseRepository: NewBaseRepository[model.Profile](tx),
		db:             tx,
	}
}
