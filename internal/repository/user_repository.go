package repository

import (
	"apart_community/internal/model"
	"context"

	"gorm.io/gorm"
)

type UserRepository interface {
	BaseRepository[model.User]
	WithTrx(ctx context.Context, tx *gorm.DB) UserRepository
	FindByEmail(ctx context.Context, email string) (model.User, error)
}

type userRepository struct {
	BaseRepository[model.User]
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		BaseRepository: NewBaseRepository[model.User](db),
		db:             db,
	}
}

func (u userRepository) conn(ctx context.Context) *gorm.DB {
	return u.db.WithContext(ctx)
}

func (u userRepository) WithTrx(ctx context.Context, tx *gorm.DB) UserRepository {
	return &userRepository{
		BaseRepository: NewBaseRepository[model.User](tx),
		db:             tx,
	}
}

func (u userRepository) FindByID(ctx context.Context, id uint) (model.User, error) {
	var user model.User
	err := u.conn(ctx).Preload("Profile").First(&user, id).Error
	return user, err
}

func (u userRepository) FindByEmail(ctx context.Context, email string) (model.User, error) {
	var user model.User
	err := u.conn(ctx).Preload("Profiles").Where("email = ?", email).First(&user).Error
	return user, err
}
