package data

import (
	"common/pkg/apperror"
	"common/pkg/constant"
	commonModel "common/pkg/model"
	"common/pkg/util"
	cerrors "common/proto/gen/common/errors"
	"context"
	"time"

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
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	return user, nil
}

func currentToken(ctx context.Context) (string, error) {
	token, ok := util.GetContextValue[string](ctx, constant.CtxToken)
	if !ok || token == "" {
		return "", apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	return token, nil
}

func formatProtoTime(ts *timestamppb.Timestamp) string {
	if ts == nil {
		return ""
	}
	return ts.AsTime().Format(time.RFC3339Nano)
}
