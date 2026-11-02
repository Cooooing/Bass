package repo

import (
	"context"

	badgev1 "common/proto/gen/badge/v1"

	"bff_bbs/internal/biz/model"
)

// BadgeClient 从内部 Badge 服务获取 SVG 预览结果。
type BadgeClient interface {
	RenderMaker(ctx context.Context, designerUserID int64, spec *badgev1.RenderSpec) (*model.BadgePreview, error)
	RenderPublished(ctx context.Context, badgeCode string) (*model.BadgePreview, error)
	UpsertDraft(ctx context.Context, req *badgev1.UpsertBadgeDraft_Req) (*badgev1.Badge, error)
	ListDrafts(ctx context.Context, designerUserID int64) ([]*badgev1.Badge, error)
	DeleteDraft(ctx context.Context, badgeID, designerUserID int64) error
}
