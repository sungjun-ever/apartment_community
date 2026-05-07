package domain

import (
	"apart_community/internals/common/repository"
	"context"

	"gorm.io/gorm"
)

type GormProfileRepository interface {
	Create(ctx context.Context, entity *Profile) (*Profile, error)

	WithTrx(tx *gorm.DB) GormProfileRepository
}

type gormProfileRepository struct {
	*repository.GormBaseRepository[Profile]
	db *gorm.DB
}

func NewGormProfileRepository(db *gorm.DB) GormProfileRepository {
	return &gormProfileRepository{
		GormBaseRepository: repository.NewGormBaseRepository[Profile](db),
		db:                 db,
	}
}

func (r *gormProfileRepository) WithTrx(tx *gorm.DB) GormProfileRepository {
	return &gormProfileRepository{
		GormBaseRepository: repository.NewGormBaseRepository[Profile](tx),
		db:                 tx,
	}
}
