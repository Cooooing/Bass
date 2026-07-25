package service

import (
	"common/pkg/apperror"
	"common/pkg/constant"
	commonmodel "common/pkg/model"
	"common/pkg/util"
	cerrors "common/proto/gen/common/errors"
	"context"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type contextReader struct{}

func (contextReader) currentUser(ctx context.Context) (*commonmodel.User, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	return user, nil
}

func (r contextReader) currentUserID(ctx context.Context) (int64, error) {
	user, err := r.currentUser(ctx)
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (contextReader) currentToken(ctx context.Context) (string, error) {
	token, ok := util.GetContextValue[string](ctx, constant.CtxToken)
	if !ok || token == "" {
		return "", apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	return token, nil
}

func (contextReader) protoTime(value string) *timestamppb.Timestamp {
	if value == "" {
		return nil
	}
	parsed, err := time.Parse(time.RFC3339Nano, value)
	if err != nil {
		return nil
	}
	return timestamppb.New(parsed)
}
