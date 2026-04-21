package repository

import (
	"apart_community/internal/model"
	"context"

	"gorm.io/gorm"
)

type ProfileRepository interface {
	BaseRepository[model.Profile]
	WithTrx(tx *gorm.DB) ProfileRepository
	UpdateByUserID(ctx context.Context, entity model.Profile, userId uint) error
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

func (p profileRepository) conn(ctx context.Context) *gorm.DB {
	return p.db.WithContext(ctx)
}

func (p profileRepository) WithTrx(tx *gorm.DB) ProfileRepository {
	return &profileRepository{
		BaseRepository: NewBaseRepository[model.Profile](tx),
		db:             tx,
	}
}

func (p profileRepository) UpdateByUserID(ctx context.Context, entity model.Profile, userId uint) error {
	err := p.conn(ctx).Model(&model.Profile{}).Where("user_id = ?", userId).Updates(&entity).Error

	return err
}
