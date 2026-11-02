package service

import (
	"common/pkg/apperror"
	"common/pkg/constant"
	commonmodel "common/pkg/model"
	"common/pkg/util"
	badgev1 "common/proto/gen/badge/v1"
	bffbadgev1 "common/proto/gen/bff_bbs/v1/badge"
	"common/proto/gen/common"
	cerrors "common/proto/gen/common/errors"
	"context"

	"bff_bbs/internal/biz/usecase"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

// BadgePreviewService adapts authenticated Maker and public published rendering.
type BadgePreviewService struct {
	bffbadgev1.UnimplementedBadgeServiceServer
	badgeUsecase *usecase.BadgePreviewUsecase
}

func NewBadgePreviewService(badgeUsecase *usecase.BadgePreviewUsecase) *BadgePreviewService {
	return &BadgePreviewService{badgeUsecase: badgeUsecase}
}
func (*BadgePreviewService) RegisterGrpc(*grpc.Server) {}
func (s *BadgePreviewService) RegisterHTTP(server *http.Server) {
	bffbadgev1.RegisterBadgeServiceHTTPServer(server, s)
}

func (s *BadgePreviewService) RenderMaker(ctx context.Context, req *bffbadgev1.RenderMakerBadge_Req) (*common.ImageResp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil || user.ID <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	if req == nil || req.GetRenderSpec() == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if req.GetRenderSpec().GetImageAssetId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	preview, err := s.badgeUsecase.RenderMaker(ctx, user.ID, s.renderSpec(req.GetRenderSpec()))
	if err != nil {
		return nil, err
	}
	return &common.ImageResp{Data: preview.Data, ContentType: preview.ContentType}, nil
}

func (s *BadgePreviewService) RenderPublished(ctx context.Context, req *bffbadgev1.RenderPublishedBadge_Req) (*common.ImageResp, error) {
	if req == nil || req.GetCode() == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	preview, err := s.badgeUsecase.RenderPublished(ctx, req.GetCode())
	if err != nil {
		return nil, err
	}
	return &common.ImageResp{Data: preview.Data, ContentType: preview.ContentType}, nil
}

func (s *BadgePreviewService) UpsertDraft(ctx context.Context, req *bffbadgev1.UpsertBadgeDraft_Req) (*bffbadgev1.UpsertBadgeDraft_Resp, error) {
	user, err := s.currentUser(ctx)
	if err != nil {
		return nil, err
	}
	if req == nil || req.GetRenderSpec() == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if req.GetRenderSpec().GetImageAssetId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}

	// An omitted optional draft ID creates a new draft. Some generated clients
	// serialize a freshly initialized optional integer as zero, so normalize it
	// back to absence instead of forwarding an invalid update ID to Badge.
	var badgeID *int64
	if req.DraftId != nil && req.GetDraftId() > 0 {
		badgeID = req.DraftId
	}
	draft, err := s.badgeUsecase.UpsertDraft(ctx, &badgev1.UpsertBadgeDraft_Req{
		DesignerUserId: user.ID,
		BadgeId:        badgeID,
		Name:           req.GetName(),
		Description:    req.GetDescription(),
		RenderSpec:     s.renderSpec(req.GetRenderSpec()),
	})
	if err != nil {
		return nil, err
	}
	return &bffbadgev1.UpsertBadgeDraft_Resp{Draft: s.draft(draft)}, nil
}

func (s *BadgePreviewService) ListDrafts(ctx context.Context, req *bffbadgev1.ListBadgeDrafts_Req) (*bffbadgev1.ListBadgeDrafts_Resp, error) {
	user, err := s.currentUser(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := s.badgeUsecase.ListDrafts(ctx, user.ID)
	if err != nil {
		return nil, err
	}
	result := make([]*bffbadgev1.BadgeDraft, 0, len(rows))
	for _, row := range rows {
		result = append(result, s.draft(row))
	}
	return &bffbadgev1.ListBadgeDrafts_Resp{Rows: result}, nil
}

func (s *BadgePreviewService) DeleteDraft(ctx context.Context, req *bffbadgev1.DeleteBadgeDraft_Req) (*bffbadgev1.DeleteBadgeDraft_Resp, error) {
	user, err := s.currentUser(ctx)
	if err != nil {
		return nil, err
	}
	// Keep zero deletable only to clean drafts written by the previous client
	// and repository zero-ID bugs. Upsert still treats zero as a new draft.
	if req == nil || req.GetDraftId() < 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if err := s.badgeUsecase.DeleteDraft(ctx, req.GetDraftId(), user.ID); err != nil {
		return nil, err
	}
	return &bffbadgev1.DeleteBadgeDraft_Resp{}, nil
}

func (*BadgePreviewService) renderSpec(req *bffbadgev1.RenderSpec) *badgev1.RenderSpec {
	return &badgev1.RenderSpec{
		ImageAssetId:        req.GetImageAssetId(),
		Text:                req.GetText(),
		BackgroundColors:    req.GetBackgroundColors(),
		TextColors:          req.GetTextColors(),
		BackgroundDirection: req.GetBackgroundDirection(),
		TextDirection:       req.GetTextDirection(),
		Size:                req.GetSize(),
		Border:              req.GetBorder(),
		TextBarLength:       req.GetTextBarLength(),
		TextSize:            req.GetTextSize(),
		TextBarRadius:       req.GetTextBarRadius(),
		Scale:               req.GetScale(),
		Shadow:              req.GetShadow(),
		AnimationSeconds:    req.GetAnimationSeconds(),
		FontAssetId:         req.GetFontAssetId(),
	}
}

func (s *BadgePreviewService) draft(item *badgev1.Badge) *bffbadgev1.BadgeDraft {
	if item == nil {
		return nil
	}

	return &bffbadgev1.BadgeDraft{
		Id:             item.GetId(),
		Name:           item.GetName(),
		Description:    item.GetDescription(),
		RenderSpec:     s.renderSpecForBFF(item.GetRenderSpec()),
		DesignerUserId: item.GetDesignerUserId(),
	}
}

func (*BadgePreviewService) renderSpecForBFF(spec *badgev1.RenderSpec) *bffbadgev1.RenderSpec {
	if spec == nil {
		return nil
	}
	return &bffbadgev1.RenderSpec{
		ImageAssetId:        spec.GetImageAssetId(),
		Text:                spec.GetText(),
		BackgroundColors:    spec.GetBackgroundColors(),
		TextColors:          spec.GetTextColors(),
		BackgroundDirection: spec.GetBackgroundDirection(),
		TextDirection:       spec.GetTextDirection(),
		Size:                spec.GetSize(),
		Border:              spec.GetBorder(),
		TextBarLength:       spec.GetTextBarLength(),
		TextSize:            spec.GetTextSize(),
		TextBarRadius:       spec.GetTextBarRadius(),
		Scale:               spec.GetScale(),
		Shadow:              spec.GetShadow(),
		AnimationSeconds:    spec.GetAnimationSeconds(),
		FontAssetId:         spec.GetFontAssetId(),
	}
}

func (*BadgePreviewService) currentUser(ctx context.Context) (*commonmodel.User, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil || user.ID <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	return user, nil
}
