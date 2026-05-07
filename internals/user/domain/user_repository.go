package domain

import (
	"apart_community/internals/common/repository"
	"context"

	"gorm.io/gorm"
)

type GormUserRepository interface {
	FindAll(ctx context.Context) ([]*User, error)
	FindById(ctx context.Context, id uint) (*User, error)
	Create(ctx context.Context, entity *User) (*User, error)
	Update(ctx context.Context, entity *User) (*User, error)
	Delete(ctx context.Context, id uint) error

	FindByEmail(ctx context.Context, email string) (*User, error)
	WithTrx(tx *gorm.DB) GormUserRepository
}

type gormUserRepository struct {
	*repository.GormBaseRepository[User]
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) GormUserRepository {
	return &gormUserRepository{
		GormBaseRepository: repository.NewGormBaseRepository[User](db),
		db:                 db,
	}
}

func (r *gormUserRepository) WithTrx(tx *gorm.DB) GormUserRepository {
	return &gormUserRepository{
		GormBaseRepository: repository.NewGormBaseRepository[User](tx),
		db:                 tx,
	}
}

func (r *gormUserRepository) FindByEmail(ctx context.Context, email string) (*User, error) {
	var user User

	err := r.db.WithContext(ctx).
		Where("email = ?", email).
		First(&user).
		Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
