package jwt

import (
	cerrors "common/api/gen/common/errors"
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

type TokenCache struct {
	log   *log.Helper
	redis *client.RedisClient
}

func NewTokenCache(log *log.Helper, redis *client.RedisClient) *TokenCache {
	return &TokenCache{
		log:   log,
		redis: redis,
	}
}

type verityCodeTokenData struct {
	Code string      `json:"code"`
	User *model.User `json:"user"`
}

func (r *TokenCache) SaveVerityCode(ctx context.Context, verityCodeType constant.VerifyCodeType, account string, code string, user *model.User, expires time.Duration) error {
	value, err := json.Marshal(&verityCodeTokenData{
		Code: code,
		User: user,
	})
	if err != nil {
		return err
	}
	return r.redis.Client.Set(ctx, constant.GetKeyTokenVerityCode(verityCodeType, account), value, expires).Err()
}

func (r *TokenCache) GetVerityCode(ctx context.Context, verityCodeType constant.VerifyCodeType, account string) (string, *model.User, error) {
	value, err := r.redis.Client.Get(ctx, constant.GetKeyTokenVerityCode(verityCodeType, account)).Result()
	if errors.Is(err, redis.Nil) {
		return "", nil, errors.New("email code invalid")
	}
	if err != nil {
		return "", nil, err
	}
	var data verityCodeTokenData
	if err := json.Unmarshal([]byte(value), &data); err != nil {
		return "", nil, err
	}
	return data.Code, data.User, nil
}

func (r *TokenCache) ExistVerityCode(ctx context.Context, verityCodeType constant.VerifyCodeType, account string) (bool, error) {
	result, err := r.redis.Client.Exists(ctx, constant.GetKeyTokenVerityCode(verityCodeType, account)).Result()
	if err != nil {
		return false, err
	}
	return result == 1, nil
}

func (r *TokenCache) DelVerityCode(ctx context.Context, verityCodeType constant.VerifyCodeType, account string) error {
	return r.redis.Client.Del(ctx, constant.GetKeyTokenVerityCode(verityCodeType, account)).Err()
}

func (r *TokenCache) SaveToken(ctx context.Context, token string, user *model.User, expires time.Duration) error {
	value, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return r.redis.Client.Set(ctx, constant.GetKeyToken(token), value, expires).Err()
}

func (r *TokenCache) GetToken(ctx context.Context, token string) (*model.User, error) {
	value, err := r.redis.Client.Get(ctx, constant.GetKeyToken(token)).Result()
	if errors.Is(err, redis.Nil) {
		return nil, nil
	}
	if err != nil {
		return nil, cerrors.ErrorUnauthorized("token is invalid")
	}
	var user model.User
	return &user, json.Unmarshal([]byte(value), &user)
}

func (r *TokenCache) DelToken(ctx context.Context, token string) error {
	return r.redis.Client.Del(ctx, constant.GetKeyToken(token)).Err()
}
