package data

import (
	"common/pkg/client/rpc"
	badgev1 "common/proto/gen/badge/v1"
	"context"

	"bff_bbs/internal/biz/model"
	bizrepo "bff_bbs/internal/biz/repo"
)

var _ bizrepo.BadgeClient = (*BadgeClient)(nil)

// BadgeClient 适配 Badge 内部 gRPC 契约为 BFF 领域读取端口。
type BadgeClient struct {
	badgeClient *rpc.BadgeClient
}

func NewBadgeClient(badgeClient *rpc.BadgeClient) bizrepo.BadgeClient {
	return &BadgeClient{badgeClient: badgeClient}
}

func (r *BadgeClient) RenderMaker(ctx context.Context, designerUserID int64, spec *badgev1.RenderSpec) (*model.BadgePreview, error) {
	resp, err := r.badgeClient.Badge.RenderMaker(ctx, &badgev1.RenderMakerBadge_Req{DesignerUserId: designerUserID, RenderSpec: spec})
	if err != nil {
		return nil, err
	}
	return &model.BadgePreview{
		Data:        []byte(resp.GetSvg()),
		ContentType: "image/svg+xml; charset=utf-8",
	}, nil
}

func (r *BadgeClient) RenderPublished(ctx context.Context, badgeCode string) (*model.BadgePreview, error) {
	resp, err := r.badgeClient.Badge.RenderPublished(ctx, &badgev1.RenderPublishedBadge_Req{BadgeCode: badgeCode})
	if err != nil {
		return nil, err
	}
	return &model.BadgePreview{Data: []byte(resp.GetSvg()), ContentType: "image/svg+xml; charset=utf-8"}, nil
}

func (r *BadgeClient) UpsertDraft(ctx context.Context, req *badgev1.UpsertBadgeDraft_Req) (*badgev1.Badge, error) {
	resp, err := r.badgeClient.Badge.UpsertDraft(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.GetBadge(), nil
}

func (r *BadgeClient) ListDrafts(ctx context.Context, designerUserID int64) ([]*badgev1.Badge, error) {
	resp, err := r.badgeClient.Badge.ListDesignerDrafts(ctx, &badgev1.ListDesignerBadgeDrafts_Req{DesignerUserId: designerUserID})
	if err != nil {
		return nil, err
	}
	return resp.GetRows(), nil
}

func (r *BadgeClient) DeleteDraft(ctx context.Context, badgeID, designerUserID int64) error {
	_, err := r.badgeClient.Badge.DeleteDraft(ctx, &badgev1.DeleteBadgeDraft_Req{BadgeId: badgeID, DesignerUserId: designerUserID})
	return err
}
