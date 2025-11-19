package util

import (
	v1 "common/api/common/v1"
	"common/pkg/client"
	"common/pkg/constant"
	"common/pkg/model"
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
)

type TokenRepo struct {
	log   *log.Helper
	redis *client.RedisClient
}

func NewTokenRepo(log *log.Helper, redis *client.RedisClient) *TokenRepo {
	return &TokenRepo{
		log:   log,
		redis: redis,
	}
}

type emailTokenData struct {
	Code string      `json:"code"`
	User *model.User `json:"user"`
}

func (r *TokenRepo) SaveEmailVerificationCode(ctx context.Context, email string, code string, user *model.User, expires time.Duration) error {
	value, err := json.Marshal(&emailTokenData{
		Code: code,
		User: user,
	})
	if err != nil {
		return err
	}
	return r.redis.Client.Set(ctx, constant.GetKeyTokenEmailCode(email), value, expires).Err()
}

func (r *TokenRepo) GetEmailVerificationCode(ctx context.Context, email string) (string, *model.User, error) {
	value, err := r.redis.Client.Get(ctx, constant.GetKeyTokenEmailCode(email)).Result()
	if errors.Is(err, redis.Nil) {
		return "", nil, errors.New("email code invalid")
	}
	if err != nil {
		return "", nil, err
	}
	var data emailTokenData
	if err := json.Unmarshal([]byte(value), &data); err != nil {
		return "", nil, err
	}
	return data.Code, data.User, nil
}

func (r *TokenRepo) ExistEmailVerificationCode(ctx context.Context, email string) (bool, error) {
	result, err := r.redis.Client.Exists(ctx, constant.GetKeyTokenEmailCode(email)).Result()
	if err != nil {
		return false, err
	}
	return result == 1, nil
}

func (r *TokenRepo) DelEmailVerificationCode(ctx context.Context, email string) error {
	return r.redis.Client.Del(ctx, constant.GetKeyTokenEmailCode(email)).Err()
}

func (r *TokenRepo) SaveToken(ctx context.Context, token string, user *model.User, expires time.Duration) error {
	value, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return r.redis.Client.Set(ctx, constant.GetKeyToken(token), value, expires).Err()
}

func (r *TokenRepo) GetToken(ctx context.Context, token string) (*model.User, error) {
	value, err := r.redis.Client.Get(ctx, constant.GetKeyToken(token)).Result()
	if errors.Is(err, redis.Nil) {
		return nil, nil
	}
	if err != nil {
		return nil, v1.ErrorUnauthorized("token is invalid")
	}
	var user model.User
	return &user, json.Unmarshal([]byte(value), &user)
}

func (r *TokenRepo) DelToken(ctx context.Context, token string) error {
	return r.redis.Client.Del(ctx, constant.GetKeyToken(token)).Err()
}
