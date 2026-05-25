package repository

import (
	"apart_community/internals/common/repository"
	"apart_community/internals/user/domain"
	"context"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll(ctx context.Context, offset, limit int) ([]*domain.User, int64, error)
	FindByID(ctx context.Context, id uint) (*domain.User, error)
	Create(ctx context.Context, entity *domain.User) (*domain.User, error)
	Update(ctx context.Context, entity *domain.User) (*domain.User, error)
	Delete(ctx context.Context, id uint) error

	FindByPublicID(ctx context.Context, publicId string) (*domain.User, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	WithTx(tx *gorm.DB) UserRepository
}

type userRepository struct {
	*repository.GormBaseRepository[domain.User]
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		GormBaseRepository: repository.NewGormBaseRepository[domain.User](db),
		db:                 db,
	}
}

func (r *userRepository) WithTx(tx *gorm.DB) UserRepository {
	return &userRepository{
		GormBaseRepository: repository.NewGormBaseRepository[domain.User](tx),
		db:                 tx,
	}
}

func (r *userRepository) FindAll(ctx context.Context, offset, limit int) ([]*domain.User, int64, error) {
	var users []*domain.User
	var total int64

	query := r.Conn(ctx).Model(&domain.User{})

	if err := query.Session(&gorm.Session{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Session(&gorm.Session{}).Preload("Profile").
		Limit(limit).Offset(offset).Order("created_at desc").Find(&users).
		Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (r *userRepository) FindByPublicID(ctx context.Context, publicId string) (*domain.User, error) {
	var user domain.User

	err := r.Conn(ctx).Preload("Profile").
		Where("public_id = ?", publicId).
		First(&user).
		Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User

	err := r.Conn(ctx).
		Where("email = ?", email).
		First(&user).
		Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
