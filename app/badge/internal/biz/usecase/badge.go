package usecase

import (
	"badge/internal/biz/model"
	"badge/internal/biz/repo"
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	"context"
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

type BadgeUsecase struct {
	badges          repo.BadgeRepo
	userBadges      repo.UserBadgeRepo
	maxDrafts       int
	draftStatus     string
	publishedStatus string
}

func NewBadgeUsecase(badges repo.BadgeRepo, userBadges repo.UserBadgeRepo) *BadgeUsecase {
	return &BadgeUsecase{
		badges:          badges,
		userBadges:      userBadges,
		maxDrafts:       5,
		draftStatus:     "draft",
		publishedStatus: "published",
	}
}

// UpsertDraft creates a draft or updates its owner's existing draft.
func (u *BadgeUsecase) UpsertDraft(ctx context.Context, designerUserID int64, badgeID *int64, name, description string, renderParams []byte) (*model.Badge, error) {
	if designerUserID <= 0 || name == "" || !json.Valid(renderParams) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if badgeID != nil {
		if *badgeID <= 0 {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		current, err := u.badges.Get(ctx, &repo.BadgeGetReq{BadgeID: badgeID})
		if err != nil {
			return nil, err
		}
		if current == nil || current.Status != u.draftStatus || current.DesignerUserID != designerUserID {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
		}
		current.Name, current.Description, current.RenderParams = name, description, renderParams
		return u.badges.Update(ctx, current)
	}
	count, err := u.badges.Page(ctx, &repo.BadgeGetReq{
		DesignerUserID: &designerUserID,
		Status:         &u.draftStatus,
	})
	if err != nil {
		return nil, err
	}
	if count.Page.Total >= int64(u.maxDrafts) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	return u.badges.Create(ctx, &model.Badge{
		Name:           name,
		Description:    description,
		RenderParams:   renderParams,
		DesignerUserID: designerUserID,
		Status:         u.draftStatus,
	})
}

// DeleteDraft removes an unpublished definition. Caller authorization belongs
// to the BFF boundary; this domain operation only protects published records.
func (u *BadgeUsecase) DeleteDraft(ctx context.Context, badgeID, designerUserID int64) error {
	// ID zero was written by an earlier create-path bug. It is not a valid
	// newly generated ID, but accepting it here lets the owner remove that
	// historical corrupt draft. Negative IDs remain invalid.
	if badgeID < 0 || designerUserID <= 0 {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	item, err := u.badges.Get(ctx, &repo.BadgeGetReq{BadgeID: &badgeID})
	if err != nil {
		return err
	}
	if item == nil {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_NOT_FOUND)
	}
	if item.Status != u.draftStatus || item.DesignerUserID != designerUserID {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
	}
	return u.badges.Delete(ctx, item.ID)
}

func (u *BadgeUsecase) PublishedByCode(ctx context.Context, code string) (*model.Badge, error) {
	if strings.TrimSpace(code) == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}

	publishedStatus := u.publishedStatus
	normalizedCode := strings.TrimSpace(code)
	return u.badges.Get(ctx, &repo.BadgeGetReq{
		Code:   &normalizedCode,
		Status: &publishedStatus,
	})
}

func (u *BadgeUsecase) Published(ctx context.Context, badgeID int64) (*model.Badge, error) {
	if badgeID <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}

	publishedStatus := u.publishedStatus
	return u.badges.Get(ctx, &repo.BadgeGetReq{
		BadgeID: &badgeID,
		Status:  &publishedStatus,
	})
}

// Publish promotes a draft to a code-addressable badge. Authorization belongs
// to the BFF administrator boundary; this domain command enforces state only.
func (u *BadgeUsecase) Publish(ctx context.Context, badgeID int64, code string) (*model.Badge, error) {
	if badgeID <= 0 || strings.TrimSpace(code) == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}

	item, err := u.badges.Get(ctx, &repo.BadgeGetReq{BadgeID: &badgeID})
	if err != nil {
		return nil, err
	}
	if item == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_NOT_FOUND)
	}
	if item.Status != u.draftStatus {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
	}

	normalizedCode := strings.TrimSpace(code)
	existing, err := u.badges.Get(ctx, &repo.BadgeGetReq{Code: &normalizedCode})
	if err != nil {
		return nil, err
	}
	if existing != nil && existing.ID != item.ID {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}

	item.Code = &normalizedCode
	item.Status = u.publishedStatus
	now := time.Now()
	item.PublishedAt = &now
	return u.badges.Update(ctx, item)
}

// ListDrafts returns only the caller-owned drafts, which are never public.
func (u *BadgeUsecase) ListDrafts(ctx context.Context, designerUserID int64) ([]*model.Badge, error) {
	if designerUserID <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	return u.badges.List(ctx, &repo.BadgeGetReq{
		DesignerUserID: &designerUserID,
		Status:         &u.draftStatus,
	})
}

// BadgeGrantReq describes one award in an atomic batch command.
type BadgeGrantReq struct {
	UserID    int64
	BadgeCode string
	Variables string
	ExpiresAt *time.Time
}

// Grant creates or refreshes all requested awards atomically.
func (u *BadgeUsecase) Grant(ctx context.Context, requests []*BadgeGrantReq) ([]*model.UserBadge, error) {
	if len(requests) == 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}

	now := time.Now()
	badgesByCode := make(map[string]*model.Badge)
	items := make([]*model.UserBadge, 0, len(requests))
	seen := make(map[string]struct{}, len(requests))
	for _, request := range requests {
		if request == nil || request.UserID <= 0 || strings.TrimSpace(request.BadgeCode) == "" ||
			!json.Valid([]byte(request.Variables)) || request.ExpiresAt != nil && !request.ExpiresAt.After(now) {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}

		badgeCode := strings.TrimSpace(request.BadgeCode)
		badge := badgesByCode[badgeCode]
		if badge == nil {
			var err error
			badge, err = u.PublishedByCode(ctx, badgeCode)
			if err != nil {
				return nil, err
			}
			if badge == nil {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_NOT_FOUND)
			}
			badgesByCode[badgeCode] = badge
		}

		key := strconv.FormatInt(request.UserID, 10) + ":" + strconv.FormatInt(badge.ID, 10)
		if _, exists := seen[key]; exists {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		seen[key] = struct{}{}
		items = append(items, &model.UserBadge{
			UserID:     request.UserID,
			BadgeID:    badge.ID,
			Variables:  []byte(request.Variables),
			ObtainedAt: now,
			ExpiresAt:  request.ExpiresAt,
			Displayed:  true,
		})
	}
	return u.userBadges.Grant(ctx, items)
}

// ListUserBadges excludes expired awards whether or not the caller filters for display.
func (u *BadgeUsecase) ListUserBadges(ctx context.Context, userID int64, displayedOnly bool) ([]*model.UserBadge, error) {
	if userID <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	request := &repo.UserBadgeGetReq{
		UserID:     &userID,
		ActiveOnly: true,
	}
	if displayedOnly {
		displayed := true
		request.Displayed = &displayed
	}
	return u.userBadges.List(ctx, request)
}

// BadgeRevokeReq identifies one user badge relation in an atomic batch command.
type BadgeRevokeReq struct {
	UserID    int64
	BadgeCode string
}

// Revoke removes all requested user badge instances without deleting definitions.
func (u *BadgeUsecase) Revoke(ctx context.Context, requests []*BadgeRevokeReq) (int, error) {
	if len(requests) == 0 {
		return 0, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}

	badgesByCode := make(map[string]*model.Badge)
	items := make([]*model.UserBadge, 0, len(requests))
	seen := make(map[string]struct{}, len(requests))
	for _, request := range requests {
		if request == nil || request.UserID <= 0 || strings.TrimSpace(request.BadgeCode) == "" {
			return 0, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}

		badgeCode := strings.TrimSpace(request.BadgeCode)
		badge := badgesByCode[badgeCode]
		if badge == nil {
			var err error
			badge, err = u.PublishedByCode(ctx, badgeCode)
			if err != nil {
				return 0, err
			}
			if badge == nil {
				return 0, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_NOT_FOUND)
			}
			badgesByCode[badgeCode] = badge
		}

		key := strconv.FormatInt(request.UserID, 10) + ":" + strconv.FormatInt(badge.ID, 10)
		if _, exists := seen[key]; exists {
			return 0, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		seen[key] = struct{}{}
		items = append(items, &model.UserBadge{UserID: request.UserID, BadgeID: badge.ID})
	}
	return u.userBadges.Revoke(ctx, items)
}

// SetDisplay changes only the owner's public display preference.
func (u *BadgeUsecase) UpdatePresentation(ctx context.Context, userID int64, badgeCode string, displayed bool, displayOrder int32) (*model.UserBadge, error) {
	if userID <= 0 || badgeCode == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	badge, err := u.PublishedByCode(ctx, badgeCode)
	if err != nil {
		return nil, err
	}
	if badge == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_NOT_FOUND)
	}
	award, err := u.userBadges.UpdatePresentation(ctx, userID, badge.ID, displayed, displayOrder)
	if err != nil {
		return nil, err
	}
	if award == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_NOT_FOUND)
	}
	return award, nil
}
