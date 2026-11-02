package service

import (
	"badge/internal/biz/model"
	"badge/internal/biz/usecase"
	"common/pkg/apperror"
	v1 "common/proto/gen/badge/v1"
	badgeenum "common/proto/gen/badge/v1/enum"
	cerrors "common/proto/gen/common/errors"
	"context"
	"time"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	transporthttp "github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// BadgeService exposes internal gRPC only; public HTTP belongs to BFF.
type BadgeService struct {
	v1.UnimplementedBadgeServiceServer
	badgeUsecase  *usecase.BadgeUsecase
	renderUsecase *usecase.BadgeRenderUsecase
}

func NewBadgeService(badgeUsecase *usecase.BadgeUsecase, renderUsecase *usecase.BadgeRenderUsecase) *BadgeService {
	return &BadgeService{badgeUsecase: badgeUsecase, renderUsecase: renderUsecase}
}
func (s *BadgeService) RegisterGrpc(server *grpc.Server) { v1.RegisterBadgeServiceServer(server, s) }
func (*BadgeService) RegisterHTTP(*transporthttp.Server) {}

// RenderMaker renders a typed Maker request after BFF has authenticated its designer.
func (s *BadgeService) RenderMaker(ctx context.Context, req *v1.RenderMakerBadge_Req) (*v1.RenderMakerBadge_Resp, error) {
	svg, err := s.renderUsecase.RenderSpec(ctx, req.GetRenderSpec())
	if err != nil {
		return nil, err
	}
	return &v1.RenderMakerBadge_Resp{Svg: svg}, nil
}

// RenderPublished loads a published badge and renders its stored parameters.
func (s *BadgeService) RenderPublished(ctx context.Context, req *v1.RenderPublishedBadge_Req) (*v1.RenderPublishedBadge_Resp, error) {
	item, err := s.badgeUsecase.PublishedByCode(ctx, req.GetBadgeCode())
	if err != nil {
		return nil, err
	}
	if item == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_NOT_FOUND)
	}
	spec := &v1.RenderSpec{}
	if err = protojson.Unmarshal(item.RenderParams, spec); err != nil {
		return nil, err
	}
	svg, err := s.renderUsecase.RenderSpec(ctx, spec)
	if err != nil {
		return nil, err
	}
	return &v1.RenderPublishedBadge_Resp{Svg: svg}, nil
}

func (s *BadgeService) UpsertDraft(ctx context.Context, req *v1.UpsertBadgeDraft_Req) (*v1.UpsertBadgeDraft_Resp, error) {
	if req == nil || req.GetRenderSpec() == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	params, err := protojson.Marshal(req.GetRenderSpec())
	if err != nil {
		return nil, err
	}
	item, err := s.badgeUsecase.UpsertDraft(
		ctx,
		req.GetDesignerUserId(),
		req.BadgeId,
		req.GetName(),
		req.GetDescription(),
		params,
	)
	if err != nil {
		return nil, err
	}
	return &v1.UpsertBadgeDraft_Resp{Badge: s.badge(item)}, nil
}

func (s *BadgeService) DeleteDraft(ctx context.Context, req *v1.DeleteBadgeDraft_Req) (*v1.DeleteBadgeDraft_Resp, error) {
	if req == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if err := s.badgeUsecase.DeleteDraft(ctx, req.GetBadgeId(), req.GetDesignerUserId()); err != nil {
		return nil, err
	}
	return &v1.DeleteBadgeDraft_Resp{}, nil
}

// Publish promotes a validated draft after the calling BFF has checked administrator permission.
func (s *BadgeService) Publish(ctx context.Context, req *v1.PublishBadgeDraft_Req) (*v1.PublishBadgeDraft_Resp, error) {
	if req == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	item, err := s.badgeUsecase.Publish(ctx, req.GetBadgeId(), req.GetCode())
	if err != nil {
		return nil, err
	}
	return &v1.PublishBadgeDraft_Resp{Badge: s.badge(item)}, nil
}

// ListDesignerDrafts exposes only the requested designer's remaining editable drafts.
func (s *BadgeService) ListDesignerDrafts(ctx context.Context, req *v1.ListDesignerBadgeDrafts_Req) (*v1.ListDesignerBadgeDrafts_Resp, error) {
	items, err := s.badgeUsecase.ListDrafts(ctx, req.GetDesignerUserId())
	if err != nil {
		return nil, err
	}
	rows := make([]*v1.Badge, 0, len(items))
	for _, item := range items {
		rows = append(rows, s.badge(item))
	}
	return &v1.ListDesignerBadgeDrafts_Resp{Rows: rows}, nil
}

// Grant records every requested award as one atomic batch.
func (s *BadgeService) Grant(ctx context.Context, req *v1.GrantBadge_Req) (*v1.GrantBadge_Resp, error) {
	if req == nil || len(req.GetItems()) == 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}

	requests := make([]*usecase.BadgeGrantReq, 0, len(req.GetItems()))
	for _, item := range req.GetItems() {
		var expiresAt *time.Time
		if timestamp := item.GetExpiresAt(); timestamp != nil {
			value := timestamp.AsTime()
			expiresAt = &value
		}
		requests = append(requests, &usecase.BadgeGrantReq{
			UserID:    item.GetUserId(),
			BadgeCode: item.GetBadgeCode(),
			Variables: item.GetVariablesJson(),
			ExpiresAt: expiresAt,
		})
	}

	items, err := s.badgeUsecase.Grant(ctx, requests)
	if err != nil {
		return nil, err
	}
	rows := make([]*v1.UserBadge, 0, len(items))
	for _, item := range items {
		userBadge, mapErr := s.userBadge(ctx, item)
		if mapErr != nil {
			return nil, mapErr
		}
		rows = append(rows, userBadge)
	}
	return &v1.GrantBadge_Resp{Rows: rows}, nil
}

// Revoke removes every requested user badge relation as one atomic batch.
func (s *BadgeService) Revoke(ctx context.Context, req *v1.RevokeBadge_Req) (*v1.RevokeBadge_Resp, error) {
	if req == nil || len(req.GetItems()) == 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}

	requests := make([]*usecase.BadgeRevokeReq, 0, len(req.GetItems()))
	for _, item := range req.GetItems() {
		requests = append(requests, &usecase.BadgeRevokeReq{
			UserID:    item.GetUserId(),
			BadgeCode: item.GetBadgeCode(),
		})
	}

	count, err := s.badgeUsecase.Revoke(ctx, requests)
	if err != nil {
		return nil, err
	}
	return &v1.RevokeBadge_Resp{Count: int64(count)}, nil
}

// UpdatePresentation persists a user's public visibility and ordering choice.
func (s *BadgeService) UpdatePresentation(ctx context.Context, req *v1.UpdateUserBadgePresentation_Req) (*v1.UpdateUserBadgePresentation_Resp, error) {
	item, err := s.badgeUsecase.UpdatePresentation(
		ctx,
		req.GetUserId(),
		req.GetBadgeCode(),
		req.GetDisplayed(),
		req.GetDisplayOrder(),
	)
	if err != nil {
		return nil, err
	}
	userBadge, err := s.userBadge(ctx, item)
	if err != nil {
		return nil, err
	}
	return &v1.UpdateUserBadgePresentation_Resp{UserBadge: userBadge}, nil
}

// ListUserBadges returns only currently valid awards and their published badge codes.
func (s *BadgeService) ListUserBadges(ctx context.Context, req *v1.ListUserBadges_Req) (*v1.ListUserBadges_Resp, error) {
	items, err := s.badgeUsecase.ListUserBadges(ctx, req.GetUserId(), req.GetDisplayedOnly())
	if err != nil {
		return nil, err
	}
	rows := make([]*v1.UserBadge, 0, len(items))
	for _, item := range items {
		userBadge, mapErr := s.userBadge(ctx, item)
		if mapErr != nil {
			return nil, mapErr
		}
		rows = append(rows, userBadge)
	}
	return &v1.ListUserBadges_Resp{Rows: rows}, nil
}

func (*BadgeService) badge(item *model.Badge) *v1.Badge {
	if item == nil {
		return nil
	}
	spec := &v1.RenderSpec{}
	_ = protojson.Unmarshal(item.RenderParams, spec)
	response := &v1.Badge{
		Id:             item.ID,
		Name:           item.Name,
		Description:    item.Description,
		RenderSpec:     spec,
		DesignerUserId: item.DesignerUserID,
	}
	if item.CreatedAt != nil {
		response.CreatedAt = timestamppb.New(*item.CreatedAt)
	}
	if item.UpdatedAt != nil {
		response.UpdatedAt = timestamppb.New(*item.UpdatedAt)
	}
	if item.Code != nil {
		response.Code = *item.Code
	}
	if item.Status == "published" {
		response.Status = badgeenum.BadgeStatus_BADGE_STATUS_PUBLISHED
	} else {
		response.Status = badgeenum.BadgeStatus_BADGE_STATUS_DRAFT
	}
	if item.PublishedAt != nil {
		response.PublishedAt = timestamppb.New(*item.PublishedAt)
	}
	return response
}

func (s *BadgeService) userBadge(ctx context.Context, item *model.UserBadge) (*v1.UserBadge, error) {
	if item == nil {
		return nil, nil
	}

	response := &v1.UserBadge{
		Id:            item.ID,
		UserId:        item.UserID,
		VariablesJson: string(item.Variables),
		ObtainedAt:    timestamppb.New(item.ObtainedAt),
		Displayed:     item.Displayed,
		DisplayOrder:  item.DisplayOrder,
	}
	if item.CreatedAt != nil {
		response.CreatedAt = timestamppb.New(*item.CreatedAt)
	}
	if item.UpdatedAt != nil {
		response.UpdatedAt = timestamppb.New(*item.UpdatedAt)
	}
	if item.ExpiresAt != nil {
		response.ExpiresAt = timestamppb.New(*item.ExpiresAt)
	}

	badge, err := s.badgeUsecase.Published(ctx, item.BadgeID)
	if err != nil {
		return nil, err
	}
	if badge == nil || badge.Code == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_NOT_FOUND)
	}
	response.BadgeCode = *badge.Code
	return response, nil
}
