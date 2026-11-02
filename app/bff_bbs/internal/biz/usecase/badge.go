package usecase

import (
	"context"

	badgev1 "common/proto/gen/badge/v1"

	"bff_bbs/internal/biz/model"
	"bff_bbs/internal/biz/repo"
)

// BadgePreviewUsecase 编排公开徽章预览，不涉及账号或持有关系控制。
type BadgePreviewUsecase struct {
	badgeRepo repo.BadgeClient
}

func NewBadgePreviewUsecase(badgeRepo repo.BadgeClient) *BadgePreviewUsecase {
	return &BadgePreviewUsecase{badgeRepo: badgeRepo}
}

func (u *BadgePreviewUsecase) RenderMaker(ctx context.Context, designerUserID int64, spec *badgev1.RenderSpec) (*model.BadgePreview, error) {
	return u.badgeRepo.RenderMaker(ctx, designerUserID, spec)
}
func (u *BadgePreviewUsecase) RenderPublished(ctx context.Context, badgeCode string) (*model.BadgePreview, error) {
	return u.badgeRepo.RenderPublished(ctx, badgeCode)
}

func (u *BadgePreviewUsecase) UpsertDraft(ctx context.Context, req *badgev1.UpsertBadgeDraft_Req) (*badgev1.Badge, error) {
	return u.badgeRepo.UpsertDraft(ctx, req)
}

func (u *BadgePreviewUsecase) ListDrafts(ctx context.Context, designerUserID int64) ([]*badgev1.Badge, error) {
	return u.badgeRepo.ListDrafts(ctx, designerUserID)
}

func (u *BadgePreviewUsecase) DeleteDraft(ctx context.Context, badgeID, designerUserID int64) error {
	return u.badgeRepo.DeleteDraft(ctx, badgeID, designerUserID)
}
