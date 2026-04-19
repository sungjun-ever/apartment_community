package service

import (
	"apart_community/internal/model"
	"apart_community/internal/repository"
	"apart_community/internal/utils"
	"errors"

	"gorm.io/gorm"
)

type UserService struct {
	userRepo    repository.UserRepository
	profileRepo repository.ProfileRepository
	db          *gorm.DB
}

func NewUserService(
	userRepo repository.UserRepository,
	profileRepo repository.ProfileRepository,
	db *gorm.DB,
) *UserService {
	return &UserService{
		userRepo:    userRepo,
		profileRepo: profileRepo,
		db:          db,
	}
}

func (u *UserService) GetAllUsers() ([]model.User, error) {
	return u.userRepo.FindAll()
}

func (u *UserService) GetUser(id uint) (*model.User, error) {
	getUser, err := u.userRepo.FindByID(id)

	if err != nil {
		return nil, err
	}

	return &getUser, err
}

func (u *UserService) GetUserByEmail(email string) (model.User, error) {
	return u.userRepo.FindByEmail(email)
}

func (u *UserService) CreateUser(user model.User, profile model.Profile) (*model.User, error) {
	var createdUser *model.User

	err := u.db.Transaction(func(tx *gorm.DB) error {
		txUserRepo := u.userRepo.WithTrx(tx)
		txProfileRepo := u.profileRepo.WithTrx(tx)

		exist, err := txUserRepo.FindByEmail(user.Email)

		if exist.ID != 0 {
			return errors.New("존재하는 사용자")
		}

		hashedPassword, _ := utils.HashPassword(user.Password)
		user.Password = hashedPassword

		if _, err = txUserRepo.Create(&user); err != nil {
			return err
		}

		profile.UserID = user.ID

		if _, err = txProfileRepo.Create(&profile); err != nil {
			return err
		}

		user.Profile = profile
		createdUser = &user

		return nil
	})

	if err != nil {
		return nil, err
	}

	return createdUser, nil

}

func (u *UserService) UpdateUser(user model.User) (model.User, error) {
	return u.userRepo.Update(&user)
}

func (u *UserService) DeleteUser(id uint) error {
	_, err := u.userRepo.FindByID(id)

	if err != nil {
		return err
	}

	return u.userRepo.Delete(id)
}
