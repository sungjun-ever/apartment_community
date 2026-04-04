package repository

import "gorm.io/gorm"

type BaseRepository[T any] interface {
	FindAll() ([]T, error)
	FindByID(id uint) (T, error)
	Create(entity *T) (T, error)
	Update(entity *T) (T, error)
	Delete(id uint) error
}

type baseRepository[T any] struct {
	db *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) BaseRepository[T] {
	return &baseRepository[T]{db: db}
}

func (b baseRepository[T]) FindAll() ([]T, error) {
	var entities []T
	err := b.db.Find(&entities).Error
	return entities, err
}

func (b baseRepository[T]) FindByID(id uint) (T, error) {
	var entity T
	err := b.db.First(&entity, id).Error
	return entity, err
}

func (b baseRepository[T]) Create(entity *T) (T, error) {
	err := b.db.Create(entity).Error
	return *entity, err
}

func (b baseRepository[T]) Update(entity *T) (T, error) {
	err := b.db.Save(entity).Error
	return *entity, err
}

func (b baseRepository[T]) Delete(id uint) error {
	var entity T
	return b.db.Delete(&entity, id).Error
}
