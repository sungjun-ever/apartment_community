package repository

import (
	"context"

	"gorm.io/gorm"
)

type BaseRepository[T any] interface {
	FindAll(ctx context.Context) ([]T, error)
	FindByID(ctx context.Context, id uint) (T, error)
	Create(ctx context.Context, entity *T) (T, error)
	Update(ctx context.Context, entity *T) (T, error)
	Delete(ctx context.Context, id uint) error
}

type baseRepository[T any] struct {
	db *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) BaseRepository[T] {
	return &baseRepository[T]{db: db}
}

func (b baseRepository[T]) conn(ctx context.Context) *gorm.DB {
	return b.db.WithContext(ctx)
}

func (b baseRepository[T]) FindAll(ctx context.Context) ([]T, error) {
	var entities []T
	err := b.conn(ctx).Find(&entities).Error
	return entities, err
}

func (b baseRepository[T]) FindByID(ctx context.Context, id uint) (T, error) {
	var entity T
	err := b.conn(ctx).First(&entity, id).Error
	return entity, err
}

func (b baseRepository[T]) Create(ctx context.Context, entity *T) (T, error) {
	err := b.conn(ctx).Create(entity).Error
	return *entity, err
}

func (b baseRepository[T]) Update(ctx context.Context, entity *T) (T, error) {
	err := b.conn(ctx).Save(entity).Error
	return *entity, err
}

func (b baseRepository[T]) Delete(ctx context.Context, id uint) error {
	var entity T
	return b.conn(ctx).Delete(&entity, id).Error
}
