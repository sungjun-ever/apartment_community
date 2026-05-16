package service

import (
	"apart_community/internals/common/errUtils"
	"apart_community/internals/common/utils"
	"apart_community/internals/user/domain"
	"apart_community/internals/user/repository"
	"context"
	"errors"

	"gorm.io/gorm"
)

type UserService struct {
	userRepo    repository.GormUserRepository
	profileRepo repository.GormProfileRepository
	db          *gorm.DB
}

func NewService(
	userRepo repository.GormUserRepository,
	profileRepo repository.GormProfileRepository,
	db *gorm.DB,
) *UserService {
	return &UserService{
		userRepo:    userRepo,
		profileRepo: profileRepo,
		db:          db,
	}
}

func (us *UserService) FindUsers(c context.Context, rq domain.PaginationRequest) ([]*domain.User, int64, error) {
	offset := (rq.Page - 1) * rq.Size
	users, total, err := us.userRepo.FindAll(c, offset, rq.Size)

	if err != nil {
		return nil, 0, errUtils.NewAppError(err, 500, errUtils.S001, errUtils.LevelInfo)
	}

	return users, total, nil
}

func (us *UserService) FindUser(c context.Context, rq domain.PublicIdUriRequest) (*domain.User, error) {
	user, err := us.userRepo.FindByPublicId(c, rq.PublicID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errUtils.NewAppError(errors.New("사용자가 존재하지 않음"), 404, errUtils.U001, errUtils.LevelInfo)
		}

		return nil, errUtils.NewAppError(err, 500, errUtils.S001, errUtils.LevelWarn)
	}

	return user, nil
}

func (us *UserService) CreateUser(c context.Context, rq domain.RegisterRequest) (*domain.User, error) {
	userEntity := rq.ToUserEntity()
	profileEntity := rq.ToProfileEntity()
	createdUser := &domain.User{}

	if userEntity == nil || profileEntity == nil {
		return nil, errUtils.NewAppError(errors.New("엔티티 생성 데이터가 없음"), 500, errUtils.S001, errUtils.LevelWarn)
	}

	err := us.db.Transaction(func(tx *gorm.DB) error {
		txUserRepo := us.userRepo.WithTrx(tx)
		txProfileRepo := us.profileRepo.WithTrx(tx)

		existUser, err := txUserRepo.FindByEmail(c, userEntity.Email)

		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return errUtils.NewAppError(err, 500, errUtils.S001, errUtils.LevelWarn)
			}
		}

		if existUser != nil {
			return errUtils.NewAppError(errors.New("이미 존재하는 이메일"), 400, errUtils.U002, errUtils.LevelInfo)
		}

		hashed, err := utils.HashPassword(userEntity.Password)

		if err != nil {
			return errUtils.NewAppError(err, 500, errUtils.S001, errUtils.LevelWarn)
		}

		userEntity.Password = string(hashed)

		user, err := txUserRepo.Create(c, userEntity)

		if err != nil {
			return errUtils.NewAppError(err, 500, errUtils.S001, errUtils.LevelWarn)
		}

		profileEntity.UserID = user.ID

		profile, err := txProfileRepo.Create(c, profileEntity)

		if err != nil {
			return errUtils.NewAppError(err, 500, errUtils.S001, errUtils.LevelWarn)
		}

		userEntity.Profile = *profile
		createdUser = userEntity

		return nil
	})

	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
