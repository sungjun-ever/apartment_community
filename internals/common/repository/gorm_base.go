package repository

import (
	"context"

	"gorm.io/gorm"
)

type GormBaseRepository[T any] struct {
	db *gorm.DB
}

func NewGormBaseRepository[T any](db *gorm.DB) *GormBaseRepository[T] {
	return &GormBaseRepository[T]{db: db}
}

func (r *GormBaseRepository[T]) conn(ctx context.Context) *gorm.DB {
	return r.db.WithContext(ctx)
}

func (r *GormBaseRepository[T]) WithTrx(ctx context.Context, tx *gorm.DB) *GormBaseRepository[T] {
	return &GormBaseRepository[T]{db: tx}
}

func (r *GormBaseRepository[T]) FindAll(ctx context.Context) ([]*T, error) {
	//TODO implement me
	panic("implement me")
}

func (r *GormBaseRepository[T]) FindById(ctx context.Context, id uint) (*T, error) {
	var entity T
	err := r.conn(ctx).First(&entity, id).Error
	return &entity, err
}

func (r *GormBaseRepository[T]) Create(ctx context.Context, entity *T) (*T, error) {
	err := r.conn(ctx).Create(entity).Error
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
