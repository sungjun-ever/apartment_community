package service

import (
	"apart_community/internal/dto/user"
	"apart_community/internal/model"
	"apart_community/internal/repository"
	"apart_community/internal/utils"
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type UserService struct {
	userRepo        repository.UserRepository
	profileRepo     repository.ProfileRepository
	belongApartRepo repository.BelongApartRepository
	db              *gorm.DB
}

func NewUserService(
	userRepo repository.UserRepository,
	profileRepo repository.ProfileRepository,
	belongApartRepo repository.BelongApartRepository,
	db *gorm.DB,
) *UserService {
	return &UserService{
		userRepo:    userRepo,
		profileRepo: profileRepo,
		db:          db,
	}
}

func (u *UserService) FindAllUsers(ctx context.Context) ([]model.User, error) {
	traceID := utils.GetTraceID(ctx)

	users, err := u.userRepo.FindAll(ctx)

	if err != nil {
		utils.ErrorLogWithContext(ctx, "FindAllUsers", err.Error(), traceID)
		return []model.User{}, err
	}

	return users, nil
}

func (u *UserService) FindUser(ctx context.Context, id uint) (*model.User, error) {
	traceID := utils.GetTraceID(ctx)

	getUser, err := u.userRepo.FindByID(ctx, id)

	if err != nil {
		utils.ErrorLogWithContext(ctx, err.Error(), "FindUser", traceID)
		return nil, err
	}

	return &getUser, nil
}

func (u *UserService) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	traceID := utils.GetTraceID(ctx)

	user, err := u.userRepo.FindByEmail(ctx, email)

	if err != nil {
		utils.ErrorLogWithContext(ctx, err.Error(), "GetUserByEmail", traceID)
		return nil, err
	}

	return &user, nil
}

func (u *UserService) CreateUser(ctx context.Context, req user.RegisterRequest) (createdUser *model.User, err error) {
	traceID := utils.GetTraceID(ctx)

	defer func() {
		if err != nil {
			utils.ErrorLogWithContext(ctx, err.Error(), "CreateUser", traceID)
		}
	}()

	userEntity := req.ToUserEntity()
	profileEntity := req.ToProfileEntity()

	if userEntity == nil || profileEntity == nil {
		return nil, errors.New("엔티티 생성 데이터가 없습니다")
	}

	err = u.db.Transaction(func(tx *gorm.DB) error {
		txUserRepo := u.userRepo.WithTrx(ctx, tx)
		txProfileRepo := u.profileRepo.WithTrx(tx)

		exist, err := txUserRepo.FindByEmail(ctx, userEntity.Email)

		if exist.ID != 0 {
			return fmt.Errorf("이미 존재하는 사용자: %w", err)
		}

		hashedPassword, err := utils.HashPassword(userEntity.Password)
		if err != nil {
			return fmt.Errorf("비밀번호 암호화 실패: %w", err)
		}

		userEntity.Password = hashedPassword

		if _, err = txUserRepo.Create(ctx, userEntity); err != nil {
			return fmt.Errorf("create user transaction 실패: %w", err)
		}

		profileEntity.UserID = userEntity.ID

		if _, err = txProfileRepo.Create(ctx, profileEntity); err != nil {
			return fmt.Errorf("create user profile transaction 실패: %w", err)
		}

		userEntity.Profile = profileEntity
		createdUser = userEntity

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("UserService.CreateUser: %w", err)
	}

	return createdUser, nil

}

func (u *UserService) UpdateUser(ctx context.Context, user model.User) (*model.User, error) {
	traceID := utils.GetTraceID(ctx)

	updatedUser, err := u.userRepo.Update(ctx, &user)

	if err != nil {
		utils.ErrorLogWithContext(ctx, err.Error(), "UpdateUser", traceID)
		return nil, err
	}

	return &updatedUser, nil
}

func (u *UserService) UpdateBelongApart(ctx context.Context, data model.UserBelongApartment) error {
	traceID := utils.GetTraceID(ctx)

	_, err := u.belongApartRepo.Create(ctx, &data)

	if err != nil {
		utils.ErrorLogWithContext(ctx, err.Error(), "UpdateBelongApart", traceID)
		return err
	}

	return nil
}

func (u *UserService) DeleteUser(ctx context.Context, id uint) error {
	traceID := utils.GetTraceID(ctx)

	_, err := u.userRepo.FindByID(ctx, id)

	if err != nil {
		utils.ErrorLogWithContext(ctx, err.Error(), "DeleteUser", traceID)
		return err
	}

	return u.userRepo.Delete(ctx, id)
}
