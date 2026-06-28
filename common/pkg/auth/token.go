package auth

import (
	"common/pkg/apperror"
	"common/pkg/client"
	"common/pkg/constant"
	"common/pkg/model"
	"common/pkg/util"
	cerrors "common/proto/gen/common/errors"
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
	"log/slog"
)

type TokenCache struct {
	log         *util.LogHelper
	redisClient *client.RedisClient
}

func NewTokenCache(logger *slog.Logger, redisClient *client.RedisClient) *TokenCache {
	return &TokenCache{
		log:         util.NewLogHelper(logger),
		redisClient: redisClient,
	}
}

type verityCodeTokenData struct {
	Code    string          `json:"code"`
	Payload json.RawMessage `json:"payload"`
}

func (r *TokenCache) SaveVerityCode(ctx context.Context, verityCodeType constant.VerifyCodeType, account string, code string, payload any, expires time.Duration) error {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	value, err := json.Marshal(&verityCodeTokenData{
		Code:    code,
		Payload: payloadBytes,
	})
	if err != nil {
		return err
	}
	return r.redisClient.Client.Set(ctx, constant.GetKeyTokenVerityCode(verityCodeType, account), value, expires).Err()
}

func (r *TokenCache) GetVerityCode(ctx context.Context, verityCodeType constant.VerifyCodeType, account string, payload any) (string, error) {
	value, err := r.redisClient.Client.Get(ctx, constant.GetKeyTokenVerityCode(verityCodeType, account)).Result()
	if errors.Is(err, redis.Nil) {
		return "", nil
	}
	if err != nil {
		return "", err
	}
	var data verityCodeTokenData
	if err := json.Unmarshal([]byte(value), &data); err != nil {
		return "", err
	}
	if err := json.Unmarshal(data.Payload, payload); err != nil {
		return "", err
	}
	return data.Code, nil
}

func (r *TokenCache) ExistVerityCode(ctx context.Context, verityCodeType constant.VerifyCodeType, account string) (bool, error) {
	result, err := r.redisClient.Client.Exists(ctx, constant.GetKeyTokenVerityCode(verityCodeType, account)).Result()
	if err != nil {
		return false, err
	}
	return result == 1, nil
}

func (r *TokenCache) DelVerityCode(ctx context.Context, verityCodeType constant.VerifyCodeType, account string) error {
	return r.redisClient.Client.Del(ctx, constant.GetKeyTokenVerityCode(verityCodeType, account)).Err()
}

func (r *TokenCache) SaveToken(ctx context.Context, token string, user *model.User, expires time.Duration) error {
	value, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return r.redisClient.Client.Set(ctx, constant.GetKeyToken(token), value, expires).Err()
}

func (r *TokenCache) GetToken(ctx context.Context, token string) (*model.User, error) {
	value, err := r.redisClient.Client.Get(ctx, constant.GetKeyToken(token)).Result()
	if errors.Is(err, redis.Nil) {
		return nil, nil
	}
	if err != nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_INVALID)
	}
	var user model.User
	return &user, json.Unmarshal([]byte(value), &user)
}

func (r *TokenCache) DelToken(ctx context.Context, token string) error {
	return r.redisClient.Client.Del(ctx, constant.GetKeyToken(token)).Err()
}
