package service

import (
	"apart_community/internals/common/errUtils"
	"apart_community/internals/common/utils"
	"apart_community/internals/common/utils/token"
	"apart_community/internals/user/domain"
	"apart_community/internals/user/repository"
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type AuthService struct {
	sessionRepo repository.RedisAuthRepository
	userRepo    repository.GormUserRepository
	db          *gorm.DB
	redis       *redis.Client
}

func NewAuthService(
	sessionRepo repository.RedisAuthRepository,
	userRepo repository.GormUserRepository,
	db *gorm.DB,
	redis *redis.Client,
) *AuthService {
	return &AuthService{
		sessionRepo: sessionRepo,
		userRepo:    userRepo,
		db:          db,
		redis:       redis,
	}
}

func (s *AuthService) Auth(c context.Context, usb domain.UserAuthBase) (*domain.User, error) {
	user, err := s.userRepo.FindByEmail(c, usb.Email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errUtils.NewAppError(errors.New("존재하지 않는 사용자"), 404, errUtils.U001, errUtils.LevelInfo)
		}

		return nil, errUtils.NewAppError(err, 500, errUtils.S001, errUtils.LevelWarn)
	}

	isVerified := utils.VerifyPassword(user.Password, usb.Password)

	if !isVerified {
		return nil, errUtils.NewAppError(errors.New("로그인 정보가 일치하지 않음"), 400, errUtils.U004, errUtils.LevelInfo)
	}

	return user, nil
}

func (s *AuthService) IssueToken(user *domain.User, duration time.Duration) (*string, *string, error) {
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
		return nil, nil, errUtils.NewAppError(err, 500, errUtils.S001, errUtils.LevelWarn)
	}

	refreshClaims := &domain.RefreshClaims{
		PublicID: user.PublicID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.PublicID,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	refreshToken, err := token.CreateRefreshToken(refreshClaims)

	if err != nil {
		return nil, nil, errUtils.NewAppError(err, 500, errUtils.S001, errUtils.LevelWarn)
	}

	return &authToken, &refreshToken, nil
}

func (s *AuthService) CreateSession(
	ctx context.Context,
	publicID string,
	session *domain.UserSession,
	duration time.Duration,
) error {
	err := s.sessionRepo.SaveSession(ctx, publicID, session, duration)

	if err != nil {
		return errUtils.NewAppError(err, 500, errUtils.S001, errUtils.LevelWarn)
	}

	return nil
}

func (s *AuthService) DestroySession(ctx context.Context, token string, claims *domain.AccessClaims) error {
	err := s.sessionRepo.DeleteSession(ctx, claims.PublicID)

	if err != nil {
		return errUtils.NewAppError(err, 500, errUtils.S001, errUtils.LevelWarn)
	}

	remainingTime := time.Until(claims.ExpiresAt.Time)

	if remainingTime > 0 {
		err = s.sessionRepo.SaveBlacklist(ctx, token, remainingTime)

		if err != nil {
			return errUtils.NewAppError(err, 500, errUtils.S001, errUtils.LevelWarn)
		}
	}

	return nil
}
