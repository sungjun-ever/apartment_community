package repository

import (
	"apart_community/internal/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	BaseRepository[model.User]
	WithTrx(tx *gorm.DB) UserRepository
	FindByEmail(email string) (model.User, error)
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

func (u userRepository) WithTrx(tx *gorm.DB) UserRepository {
	return &userRepository{
		BaseRepository: NewBaseRepository[model.User](tx),
		db:             tx,
	}
}

func (u userRepository) FindByID(id uint) (model.User, error) {
	var user model.User
	err := u.db.Preload("Profile").First(&user, id).Error
	return user, err
}

func (u userRepository) FindByEmail(email string) (model.User, error) {
	var user model.User
	err := u.db.Preload("Profiles").Where("email = ?", email).First(&user).Error
	return user, err
}
