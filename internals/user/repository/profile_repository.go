package repository

import (
	"apart_community/internals/common/repository"
	"apart_community/internals/user/domain"
	"context"

	"gorm.io/gorm"
)

type ProfileRepository interface {
	Create(ctx context.Context, entity *domain.Profile) (*domain.Profile, error)

	WithTx(tx *gorm.DB) ProfileRepository
}

type profileRepository struct {
	*repository.GormBaseRepository[domain.Profile]
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) ProfileRepository {
	return &profileRepository{
		GormBaseRepository: repository.NewGormBaseRepository[domain.Profile](db),
		db:                 db,
	}
}

func (r *profileRepository) WithTx(tx *gorm.DB) ProfileRepository {
	return &profileRepository{
		GormBaseRepository: repository.NewGormBaseRepository[domain.Profile](tx),
		db:                 tx,
	}
}
