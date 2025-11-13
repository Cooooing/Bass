package util

import (
	"common/pkg/constant"
	"common/pkg/model"
	"context"
)

func GetUserInfo(ctx context.Context) (*model.User, bool) {
	user, ok := ctx.Value(constant.UserInfo).(*model.User)
	return user, ok
}

func MustGetUserInfo(ctx context.Context) *model.User {
	if user, ok := ctx.Value(constant.UserInfo).(*model.User); ok {
		return user
	}
	panic("user info not found in context or invalid type")
}
