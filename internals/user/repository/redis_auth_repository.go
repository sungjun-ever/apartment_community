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
	GetSession(ctx context.Context, refreshToken string) (string, error)
	SaveSession(ctx context.Context, publicID string, session *domain.UserSession, duration time.Duration) error
	DeleteSession(ctx context.Context, publicID string) error
}

type redisAuthRepository struct {
	rds *redis.Client
}

func NewRedisAuthRepository(redis *redis.Client) RedisAuthRepository {
	return &redisAuthRepository{
		rds: redis,
	}
}

func (r redisAuthRepository) GetSession(ctx context.Context, refreshToken string) (string, error) {
	key := utils.SessionKey(refreshToken)
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
	panic("implement me")
}
