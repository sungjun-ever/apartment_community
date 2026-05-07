package service

import (
	"apart_community/internals/common/errUtils"
	"apart_community/internals/user/domain"
	"errors"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	userRepo    domain.GormUserRepository
	profileRepo domain.GormProfileRepository
	db          *gorm.DB
}

func NewService(
	userRepo domain.GormUserRepository,
	profileRepo domain.GormProfileRepository,
	db *gorm.DB,
) *UserService {
	return &UserService{
		userRepo:    userRepo,
		profileRepo: profileRepo,
		db:          db,
	}
}

func (us *UserService) CreateUser(gc *gin.Context, rq domain.RegisterRequest) (*domain.User, error) {
	userEntity := rq.ToUserEntity()
	profileEntity := rq.ToProfileEntity()
	createdUser := &domain.User{}

	if userEntity == nil || profileEntity == nil {
		_ = gc.Error(errUtils.NewAppError(errors.New("엔티티 생성 데이터가 없음"), 500, "S001"))
		return nil, nil
	}

	err := us.db.Transaction(func(tx *gorm.DB) error {
		txUserRepo := us.userRepo.WithTrx(tx)
		txProfileRepo := us.profileRepo.WithTrx(tx)

		existUser, err := txUserRepo.FindByEmail(gc, userEntity.Email)

		if err != nil {
			return errUtils.NewAppError(err, 500, "S001")
		}

		if existUser != nil {
			return errUtils.NewAppError(errors.New("이미 존재하는 이메일"), 400, "U002")
		}

		hashed, err := bcrypt.GenerateFromPassword([]byte(userEntity.Password), bcrypt.DefaultCost)

		if err != nil {
			return errUtils.NewAppError(err, 500, "S001")
		}

		userEntity.Password = string(hashed)

		user, err := txUserRepo.Create(gc, userEntity)

		if err != nil {
			return errUtils.NewAppError(err, 500, "S001")
		}

		profileEntity.UserID = user.ID

		profile, err := txProfileRepo.Create(gc, profileEntity)

		if err != nil {
			return errUtils.NewAppError(err, 500, "S001")
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
