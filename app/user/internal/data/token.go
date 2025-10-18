package data

import (
	"common/pkg/constant"
	"context"
	"encoding/json"
	"user/internal/biz/model"
	"user/internal/biz/repo"
)

type TokenRepo struct {
	*BaseRepo
}

func NewTokenRepo(baseRepo *BaseRepo) repo.TokenRepo {
	return &TokenRepo{BaseRepo: baseRepo}
}

type emailTokenData struct {
	Code string      `json:"code"`
	User *model.User `json:"user"`
}

func (r *TokenRepo) SaveEmailToken(ctx context.Context, token string, code string, user *model.User) error {
	value, err := json.Marshal(&emailTokenData{
		Code: code,
		User: user,
	})
	if err != nil {
		return err
	}
	return r.redis.Client.Set(ctx, constant.GetKeyTokenEmailCode(token), value, r.conf.Jwt.EmailExpire.AsDuration()).Err()
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

func (r *TokenRepo) SaveToken(ctx context.Context, token string, user *model.User) error {
	value, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return r.redis.Client.Set(ctx, constant.GetKeyToken(token), value, r.conf.Jwt.Expires.AsDuration()).Err()
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
