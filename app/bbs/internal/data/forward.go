package data

import (
	"common/pkg/constant"
	commonModel "common/pkg/model"
	"common/pkg/util"
	"context"
	"strings"
	"time"

	cerrors "common/api/gen/common/errors"

	"github.com/go-kratos/kratos/v2/transport"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func forwardAuth(ctx context.Context) context.Context {
	tr, ok := transport.FromServerContext(ctx)
	if !ok || tr.RequestHeader() == nil {
		return ctx
	}
	auth := tr.RequestHeader().Get(constant.HeaderAuthentication)
	if auth == "" {
		return ctx
	}
	return metadata.AppendToOutgoingContext(ctx, strings.ToLower(constant.HeaderAuthentication), auth)
}

func currentUserID(ctx context.Context) (int64, error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return 0, cerrors.ErrorUnauthorized("user not login")
	}
	return user.ID, nil
}

func formatProtoTime(ts *timestamppb.Timestamp) string {
	if ts == nil {
		return ""
	}
	return ts.AsTime().Format(time.RFC3339Nano)
}
