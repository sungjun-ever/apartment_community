package repository

import (
	"apart_community/internals/common/utils"
	"apart_community/internals/user/domain"
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisAuthRepository interface {
	GetSession(ctx context.Context, publicID string) (string, error)
	SaveSession(ctx context.Context, publicID string, session *domain.UserSession, duration time.Duration) error
	DeleteSession(ctx context.Context, publicID string) error
	SaveBlacklist(ctx context.Context, token string, duration time.Duration) error
	IsBlacklisted(ctx context.Context, token string) (int64, error)
}

type redisAuthRepository struct {
	rds *redis.Client
}

func NewRedisAuthRepository(redis *redis.Client) RedisAuthRepository {
	return &redisAuthRepository{
		rds: redis,
	}
}

func (r redisAuthRepository) GetSession(ctx context.Context, publicID string) (string, error) {
	key := utils.SessionKey(publicID)
	result, err := r.rds.Get(ctx, key).Result()

	if err != nil {
		return "", err
	}

	return result, nil
}

func (r redisAuthRepository) SaveSession(ctx context.Context, publicID string, session *domain.UserSession,
	duration time.Duration) error {
	key := utils.SessionKey(publicID)
	data, _ := json.Marshal(session)

	err := r.rds.Set(ctx, key, data, duration).Err()

	if err != nil {
		return err
	}

	return nil
}

func (r redisAuthRepository) DeleteSession(ctx context.Context, publicID string) error {
	key := utils.SessionKey(publicID)

	err := r.rds.Del(ctx, key).Err()

	if err != nil {
		return err
	}

	return nil
}

func (r redisAuthRepository) IsBlacklisted(ctx context.Context, token string) (int64, error) {
	key := utils.BlacklistAccessTokenKey(token)
	result, err := r.rds.Exists(ctx, key).Result()

	if err != nil {
		return 0, err
	}

	if result == 1 {
		return result, nil
	}

	return 0, nil
}

func (r redisAuthRepository) SaveBlacklist(ctx context.Context, token string, duration time.Duration) error {
	key := utils.BlacklistAccessTokenKey(token)

	err := r.rds.Set(ctx, key, token, duration).Err()

	if err != nil {
		return err
	}

	return nil
}
