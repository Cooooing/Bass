package util

import (
	"common/pkg/client"
	"common/pkg/constant"
	"common/pkg/model"
	"context"
	"encoding/json"
	"time"

	"github.com/go-kratos/kratos/v2/log"
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

func (r *TokenRepo) SaveEmailToken(ctx context.Context, token string, code string, user *model.User, expires time.Duration) error {
	value, err := json.Marshal(&emailTokenData{
		Code: code,
		User: user,
	})
	if err != nil {
		return err
	}
	return r.redis.Client.Set(ctx, constant.GetKeyTokenEmailCode(token), value, expires).Err()
}

func (r *TokenRepo) GetEmailToken(ctx context.Context, token string) (string, *model.User, error) {
	value, err := r.redis.Client.Get(ctx, constant.GetKeyTokenEmailCode(token)).Result()
	if err != nil {
		return "", nil, err
	}
	var data emailTokenData
	if err := json.Unmarshal([]byte(value), &data); err != nil {
		return "", nil, err
	}
	return data.Code, data.User, nil
}

func (r *TokenRepo) DelEmailToken(ctx context.Context, token string) error {
	return r.redis.Client.Del(ctx, constant.GetKeyTokenEmailCode(token)).Err()
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
	if err != nil {
		return nil, err
	}
	var user model.User
	return &user, json.Unmarshal([]byte(value), &user)
}

func (r *TokenRepo) DelToken(ctx context.Context, token string) error {
	return r.redis.Client.Del(ctx, constant.GetKeyToken(token)).Err()
}
