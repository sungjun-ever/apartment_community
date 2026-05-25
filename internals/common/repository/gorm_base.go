package repository

import (
	"context"

	"gorm.io/gorm"
)

type BaseRepository[T any] interface {
	WithTx(tx *gorm.DB) BaseRepository[T]
	FindByID(ctx context.Context, id uint) (*T, error)
	Create(ctx context.Context, entity *T) (*T, error)
	Update(ctx context.Context, entity *T) (*T, error)
	Delete(ctx context.Context, id uint) error
}

type GormBaseRepository[T any] struct {
	db *gorm.DB
}

func NewGormBaseRepository[T any](db *gorm.DB) *GormBaseRepository[T] {
	return &GormBaseRepository[T]{db: db}
}

func (r *GormBaseRepository[T]) Conn(ctx context.Context) *gorm.DB {
	return r.db.WithContext(ctx)
}

func (r *GormBaseRepository[T]) WithTx(tx *gorm.DB) *GormBaseRepository[T] {
	return &GormBaseRepository[T]{db: tx}
}

func (r *GormBaseRepository[T]) FindAll(ctx context.Context, offset, limit int) ([]*T, int64, error) {
	var entities []*T
	var total int64

	query := r.Conn(ctx).Model(new(T))

	if err := query.Session(&gorm.Session{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Session(&gorm.Session{}).Limit(limit).Offset(offset).Find(&entities).Error; err != nil {
		return nil, 0, err
	}

	return entities, total, nil
}

func (r *GormBaseRepository[T]) FindByID(ctx context.Context, id uint) (*T, error) {
	var entity T
	err := r.Conn(ctx).First(&entity, id).Error
	return &entity, err
}

func (r *GormBaseRepository[T]) Create(ctx context.Context, entity *T) (*T, error) {
	err := r.Conn(ctx).Create(entity).Error
	return entity, err
}

func (r *GormBaseRepository[T]) Update(ctx context.Context, entity *T) (*T, error) {
	//TODO implement me
	panic("implement me")
}

func (r *GormBaseRepository[T]) Delete(ctx context.Context, id uint) error {
	//TODO implement me
	panic("implement me")
}
