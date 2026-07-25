package service

import (
	"common/pkg/apperror"
	commonenum "common/pkg/enum"
	cerrors "common/proto/gen/common/errors"
	v1 "common/proto/gen/user/v1"
	"context"
	"time"
	"user/internal/biz/model"
	"user/internal/biz/usecase"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type RbacService struct {
	v1.UnimplementedRbacServiceServer
	rbacUsecase *usecase.RbacUsecase
}

func NewRbacService(rbacUsecase *usecase.RbacUsecase) *RbacService {
	return &RbacService{rbacUsecase: rbacUsecase}
}
func (s *RbacService) RegisterGrpc(gs *grpc.Server) { v1.RegisterRbacServiceServer(gs, s) }
func (s *RbacService) RegisterHttp(hs *http.Server) {}

func (s *RbacService) UpsertRole(ctx context.Context, req *v1.UpsertRbacRole_Req) (*v1.UpsertRbacRole_Resp, error) {
	realm, ok := commonenum.LoginRealmMap.ToEnum(req.GetRealm())
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	enabled := true
	if req.Enabled != nil {
		enabled = req.GetEnabled()
	}
	row, err := s.rbacUsecase.UpsertRole(ctx, &model.RbacRole{ID: req.GetId(), Realm: realm, Code: req.GetCode(), Name: req.GetName(), Description: req.GetDescription(), BuiltIn: req.GetBuiltIn(), Enabled: enabled})
	if err != nil {
		return nil, err
	}
	var role *v1.UpsertRbacRole_Resp_Role
	if row != nil {
		role = &v1.UpsertRbacRole_Resp_Role{Id: row.ID, Realm: commonenum.LoginRealmMap.MustToProto(row.Realm), Code: row.Code, Name: row.Name, Description: row.Description, BuiltIn: row.BuiltIn, Enabled: row.Enabled}
		if row.CreatedAt != nil {
			role.CreatedAt = timestamppb.New(*row.CreatedAt)
		}
		if row.UpdatedAt != nil {
			role.UpdatedAt = timestamppb.New(*row.UpdatedAt)
		}
	}
	return &v1.UpsertRbacRole_Resp{Row: role}, nil
}

func (s *RbacService) UpsertPermission(ctx context.Context, req *v1.UpsertRbacPermission_Req) (*v1.UpsertRbacPermission_Resp, error) {
	realm, ok := commonenum.LoginRealmMap.ToEnum(req.GetRealm())
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	enabled := true
	if req.Enabled != nil {
		enabled = req.GetEnabled()
	}
	row, err := s.rbacUsecase.UpsertPermission(ctx, &model.RbacPermission{ID: req.GetId(), Realm: realm, Code: req.GetCode(), Name: req.GetName(), Description: req.GetDescription(), Enabled: enabled})
	if err != nil {
		return nil, err
	}
	var permission *v1.UpsertRbacPermission_Resp_Permission
	if row != nil {
		permission = &v1.UpsertRbacPermission_Resp_Permission{Id: row.ID, Realm: commonenum.LoginRealmMap.MustToProto(row.Realm), Code: row.Code, Name: row.Name, Description: row.Description, Enabled: row.Enabled}
		if row.CreatedAt != nil {
			permission.CreatedAt = timestamppb.New(*row.CreatedAt)
		}
		if row.UpdatedAt != nil {
			permission.UpdatedAt = timestamppb.New(*row.UpdatedAt)
		}
	}
	return &v1.UpsertRbacPermission_Resp{Row: permission}, nil
}

func (s *RbacService) BindRolePermission(ctx context.Context, req *v1.BindRbacRolePermission_Req) (*v1.BindRbacRolePermission_Resp, error) {
	return &v1.BindRbacRolePermission_Resp{}, s.rbacUsecase.BindRolePermission(ctx, req.GetRoleId(), req.GetPermissionId())
}
func (s *RbacService) UnbindRolePermission(ctx context.Context, req *v1.UnbindRbacRolePermission_Req) (*v1.UnbindRbacRolePermission_Resp, error) {
	return &v1.UnbindRbacRolePermission_Resp{}, s.rbacUsecase.UnbindRolePermission(ctx, req.GetRoleId(), req.GetPermissionId())
}
func (s *RbacService) GrantRole(ctx context.Context, req *v1.GrantRbacRole_Req) (*v1.GrantRbacRole_Resp, error) {
	var expiresAt *time.Time
	if req.ExpiresAt != nil {
		value := req.ExpiresAt.AsTime()
		expiresAt = &value
	}
	return &v1.GrantRbacRole_Resp{}, s.rbacUsecase.GrantRole(ctx, req.GetUserId(), req.GetRoleId(), req.GetGrantedBy(), expiresAt)
}
func (s *RbacService) RevokeRole(ctx context.Context, req *v1.RevokeRbacRole_Req) (*v1.RevokeRbacRole_Resp, error) {
	return &v1.RevokeRbacRole_Resp{}, s.rbacUsecase.RevokeRole(ctx, req.GetUserId(), req.GetRoleId())
}
func (s *RbacService) CheckPermission(ctx context.Context, req *v1.CheckRbacPermission_Req) (*v1.CheckRbacPermission_Resp, error) {
	realm, ok := commonenum.LoginRealmMap.ToEnum(req.GetRealm())
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	allowed, err := s.rbacUsecase.CheckPermission(ctx, req.GetUserId(), realm, req.GetPermissionCode())
	return &v1.CheckRbacPermission_Resp{Allowed: allowed}, err
}
