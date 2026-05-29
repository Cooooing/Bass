package data

import (
	"common/pkg/constant"
	commonModel "common/pkg/model"
	"common/pkg/util"
	"context"
	"time"

	cerrors "common/api/gen/common/errors"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func currentUserID(ctx context.Context) (int64, error) {
	user, err := currentUser(ctx)
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

func currentUser(ctx context.Context) (*commonModel.User, error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
	return user, nil
}

func currentToken(ctx context.Context) (string, error) {
	token, ok := util.GetContextValue[string](ctx, constant.CtxToken)
	if !ok || token == "" {
		return "", cerrors.ErrorUnauthorized("user not login")
	}
	return token, nil
}

func formatProtoTime(ts *timestamppb.Timestamp) string {
	if ts == nil {
		return ""
	}
	return ts.AsTime().Format(time.RFC3339Nano)
}
