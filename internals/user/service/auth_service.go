package service

import (
	"apart_community/internals/common/errUtils"
	"apart_community/internals/common/utils"
	"apart_community/internals/common/utils/token"
	"apart_community/internals/user/domain"
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type AuthService struct {
	userRepo domain.GormUserRepository
	db       *gorm.DB
}

func NewAuthService(userRepo domain.GormUserRepository, db *gorm.DB) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		db:       db,
	}
}

func (s *AuthService) Auth(c context.Context, usb domain.UserAuthBase) (*domain.User, error) {
	user, err := s.userRepo.FindByEmail(c, usb.Email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errUtils.NewAppError(errors.New("존재하지 않는 사용자"), 404, "U001")
		}

		return nil, errUtils.NewAppError(err, 500, "S001")
	}

	isVerified := utils.VerifyPassword(user.Password, usb.Password)

	if !isVerified {
		return nil, errUtils.NewAppError(errors.New("로그인 정보가 일치하지 않음"), 400, "U004")
	}

	return user, nil
}

func (s *AuthService) IssueToken(user *domain.User) (*string, *string, error) {
	accessClaims := &domain.AccessClaims{
		PublicID: user.PublicID,
		Email:    user.Email,
		Nickname: user.Profile.Nickname,
		//Roles:    roles,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.PublicID,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	authToken, err := token.CreateAccessToken(accessClaims)

	if err != nil {
		return nil, nil, errUtils.NewAppError(err, 500, "S001")
	}

	refreshClaims := &domain.RefreshClaims{
		PublicID: user.PublicID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.PublicID,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	refreshToken, err := token.CreateRefreshToken(refreshClaims)

	if err != nil {
		return nil, nil, errUtils.NewAppError(err, 500, "S001")
	}

	return &authToken, &refreshToken, nil
}
