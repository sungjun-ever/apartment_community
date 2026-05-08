package domain

import (
	"apart_community/internals/common/repository"
	"context"

	"gorm.io/gorm"
)

type GormUserRepository interface {
	FindAll(ctx context.Context, offset, limit int) ([]*User, int64, error)
	FindById(ctx context.Context, id uint) (*User, error)
	Create(ctx context.Context, entity *User) (*User, error)
	Update(ctx context.Context, entity *User) (*User, error)
	Delete(ctx context.Context, id uint) error

	FindByPublicId(ctx context.Context, publicId string) (*User, error)
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

func (r *gormUserRepository) FindAll(ctx context.Context, offset, limit int) ([]*User, int64, error) {
	var users []*User
	var total int64

	query := r.Conn(ctx).Preload("Profile").Model(&User{})

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Limit(limit).Offset(offset).Order("created_at desc").Find(&users).Error

	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (r *gormUserRepository) FindByPublicId(ctx context.Context, publicId string) (*User, error) {
	var user User

	err := r.Conn(ctx).Preload("Profile").
		Where("public_id = ?", publicId).
		First(&user).
		Error

	if err != nil {
		return nil, err
	}

	return &user, nil
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
