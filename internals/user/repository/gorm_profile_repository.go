package repository

import (
	"apart_community/internals/common/repository"
	"apart_community/internals/user/domain"
	"context"

	"gorm.io/gorm"
)

type GormProfileRepository interface {
	Create(ctx context.Context, entity *domain.Profile) (*domain.Profile, error)

	WithTrx(tx *gorm.DB) GormProfileRepository
}

type gormProfileRepository struct {
	*repository.GormBaseRepository[domain.Profile]
	db *gorm.DB
}

func NewGormProfileRepository(db *gorm.DB) GormProfileRepository {
	return &gormProfileRepository{
		GormBaseRepository: repository.NewGormBaseRepository[domain.Profile](db),
		db:                 db,
	}
}

func (r *gormProfileRepository) WithTrx(tx *gorm.DB) GormProfileRepository {
	return &gormProfileRepository{
		GormBaseRepository: repository.NewGormBaseRepository[domain.Profile](tx),
		db:                 tx,
	}
}
