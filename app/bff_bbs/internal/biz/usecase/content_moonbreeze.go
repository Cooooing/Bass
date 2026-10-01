package usecase

import (
	"bff_bbs/internal/biz/repo"
	"common/pkg/apperror"
	commonmodel "common/pkg/model"
	cerrors "common/proto/gen/common/errors"
	"context"
	"strings"
)

const moonbreezeMaxPageSize uint32 = 50

type ContentMoonbreezeUsecase struct {
	moonbreezeClient   repo.ContentMoonbreezeClient
	accountClient      repo.AccountClient
	assetClient        repo.AssetClient
	privacyClient      repo.PrivacySettingClient
	relationClient     repo.RelationClient
	ipResolutionClient repo.IPResolutionClient
}

func NewContentMoonbreezeUsecase(
	moonbreezeClient repo.ContentMoonbreezeClient,
	accountClient repo.AccountClient,
	assetClient repo.AssetClient,
	privacyClient repo.PrivacySettingClient,
	relationClient repo.RelationClient,
	ipResolutionClient repo.IPResolutionClient,
) *ContentMoonbreezeUsecase {
	return &ContentMoonbreezeUsecase{
		moonbreezeClient:   moonbreezeClient,
		accountClient:      accountClient,
		assetClient:        assetClient,
		privacyClient:      privacyClient,
		relationClient:     relationClient,
		ipResolutionClient: ipResolutionClient,
	}
}

type CreateMoonbreezeReq struct {
	UserID  int64
	Content string
	IP      string
}

func (u *ContentMoonbreezeUsecase) Create(ctx context.Context, req *CreateMoonbreezeReq) (*repo.Moonbreeze, error) {
	if req == nil || req.UserID <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_MOONBREEZE_INVALID)
	}
	var city *string
	if ip := strings.TrimSpace(req.IP); ip != "" && u.ipResolutionClient != nil {
		resolved, err := u.ipResolutionClient.Resolve(ctx, ip)
		if err == nil && resolved != nil {
			if value := strings.TrimSpace(resolved.City); value != "" {
				city = &value
			}
		}
	}
	row, err := u.moonbreezeClient.CreateMoonbreeze(ctx, &repo.CreateMoonbreezeReq{
		Content:  req.Content,
		AuthorID: req.UserID,
		City:     city,
	})
	if err != nil {
		return nil, err
	}
	if err = u.hydrateRows(ctx, []*repo.Moonbreeze{row}, req.UserID); err != nil {
		return nil, err
	}
	return row, nil
}

type PageMoonbreezesReq struct {
	ViewerID  int64
	AuthorIDs []int64
	Cursor    *string
	Size      uint32
}

func (u *ContentMoonbreezeUsecase) Page(
	ctx context.Context,
	req *PageMoonbreezesReq,
) (*repo.PageMoonbreezesResp, error) {
	if req == nil {
		req = &PageMoonbreezesReq{}
	}
	pageSize := req.Size
	if pageSize == 0 {
		pageSize = 20
	}
	if pageSize > moonbreezeMaxPageSize {
		pageSize = moonbreezeMaxPageSize
	}
	// Content owns the stream's ordering and filtering inputs. Personal-list
	// visibility is enforced only by PageMember, so global and following feeds
	// remain independent from a member's profile-list preference.
	cursor := req.Cursor
	visible := make([]*repo.Moonbreeze, 0, pageSize)
	var next *string
	for len(visible) < int(pageSize) {
		remaining := pageSize - uint32(len(visible))
		batch, err := u.moonbreezeClient.PageMoonbreezes(ctx, &repo.PageMoonbreezesReq{
			AuthorIDs: req.AuthorIDs,
			Cursor:    cursor,
			Size:      remaining,
		})
		if err != nil {
			return nil, err
		}
		if len(batch.Rows) == 0 {
			next = nil
			break
		}
		for _, row := range batch.Rows {
			if row == nil {
				continue
			}
			visible = append(visible, row)
		}
		if len(visible) == int(pageSize) {
			next = batch.NextCursor
			break
		}
		if batch.NextCursor == nil {
			break
		}
		cursor = batch.NextCursor
	}
	if err := u.hydrateRows(ctx, visible, req.ViewerID); err != nil {
		return nil, err
	}
	return &repo.PageMoonbreezesResp{Rows: visible, NextCursor: next}, nil
}

func (u *ContentMoonbreezeUsecase) PageWatching(
	ctx context.Context,
	viewerID int64,
	cursor *string,
	size uint32,
) (*repo.PageMoonbreezesResp, error) {
	if viewerID <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	following, err := u.followingIDs(ctx, viewerID)
	if err != nil {
		return nil, err
	}
	if len(following) == 0 {
		return &repo.PageMoonbreezesResp{}, nil
	}
	return u.Page(ctx, &PageMoonbreezesReq{
		ViewerID:  viewerID,
		AuthorIDs: following,
		Cursor:    cursor,
		Size:      size,
	})
}

func (u *ContentMoonbreezeUsecase) PageMember(
	ctx context.Context,
	viewerID int64,
	name string,
	cursor *string,
	size uint32,
) (*repo.PageMoonbreezesResp, error) {
	account, _, err := u.accountClient.GetProfile(ctx, name)
	if err != nil {
		return nil, err
	}
	if account == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_ACCOUNT_NOT_FOUND)
	}
	if viewerID != account.ID {
		privacy, err := u.privacyClient.GetCurrentPrivacySetting(ctx, account.ID)
		if err != nil {
			return nil, err
		}
		if privacy.PublicMoonbreezeList != nil && !*privacy.PublicMoonbreezeList {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_PROFILE_MOONBREEZE_LIST_PRIVATE)
		}
	}
	return u.Page(ctx, &PageMoonbreezesReq{
		ViewerID:  viewerID,
		AuthorIDs: []int64{account.ID},
		Cursor:    cursor,
		Size:      size,
	})
}

// followingIDs loads the caller's complete following set before querying the
// content service. The relation API is page based, so this remains correct
// until a dedicated ID-only query is needed for significantly larger sets.
func (u *ContentMoonbreezeUsecase) followingIDs(ctx context.Context, userID int64) ([]int64, error) {
	var ids []int64
	for page := commonmodel.DefaultPage; ; page++ {
		result, err := u.relationClient.ListFollowing(ctx, &repo.ListFollowingRelationsReq{
			ActorID: userID,
			Page: &commonmodel.PageReq{
				Page: page,
				Size: 100,
			},
		})
		if err != nil {
			return nil, err
		}
		for _, row := range result.Rows {
			if row != nil && row.TargetID > 0 {
				ids = append(ids, row.TargetID)
			}
		}
		if result.Page == nil || result.Page.Size == 0 || len(result.Rows) == 0 {
			break
		}
		if page >= (result.Page.Total+result.Page.Size-1)/result.Page.Size {
			break
		}
	}
	return ids, nil
}

// hydrateRows loads presentation-only author, avatar and privacy data in
// batches. It never changes the content records returned by the content domain.
func (u *ContentMoonbreezeUsecase) hydrateRows(ctx context.Context, rows []*repo.Moonbreeze, viewerID int64) error {
	ids := make([]int64, 0, len(rows))
	for _, row := range rows {
		if row != nil {
			ids = append(ids, row.AuthorID)
		}
	}
	profiles, err := u.accountClient.MapProfileAccounts(ctx, ids)
	if err != nil {
		return err
	}
	privacy, err := u.privacyClient.MapPrivacySettings(ctx, ids)
	if err != nil {
		return err
	}
	assetIDs := make([]int64, 0, len(profiles))
	for _, profile := range profiles {
		if profile != nil && profile.AvatarAssetID != nil && *profile.AvatarAssetID > 0 {
			assetIDs = append(assetIDs, *profile.AvatarAssetID)
		}
	}
	assets, err := u.assetClient.Map(ctx, &repo.AssetGetReq{IDs: assetIDs})
	if err != nil {
		return err
	}
	for _, row := range rows {
		if row == nil {
			continue
		}
		profile := profiles[row.AuthorID]
		if profile != nil {
			avatarURL := "/v1/user/account/avatar?name=" + profile.Name
			if profile.AvatarAssetID != nil {
				if asset := assets[*profile.AvatarAssetID]; asset != nil && asset.URL != "" {
					avatarURL = asset.URL
				}
			}
			profile.AvatarURL = &avatarURL
		}
		row.Author = profile
		setting := privacy[row.AuthorID]
		row.ShowCity = row.AuthorID == viewerID || setting == nil || setting.PublicLocation == nil || *setting.PublicLocation
	}
	return nil
}
