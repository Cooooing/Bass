package usecase

import (
	"bff_bbs/internal/biz/repo"
	"common/pkg/apperror"
	bbsuserv1 "common/proto/gen/bff_bbs/v1/user"
	bbsuserv1enum "common/proto/gen/bff_bbs/v1/user/enum"
	commonv1 "common/proto/gen/common"
	cerrors "common/proto/gen/common/errors"
	economyv1enum "common/proto/gen/economy/v1/enum"
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"regexp"
	"strings"
)

var assetHashPattern = regexp.MustCompile(`^[a-f0-9]{64}$`)

type AccountUsecase struct {
	accountClient  repo.AccountClient
	assetClient    repo.AssetClient
	economyClient  repo.EconomyClient
	privacyClient  repo.PrivacySettingClient
	locationClient repo.LocationClient
	relationClient repo.RelationClient
}

func NewAccountUsecase(
	accountClient repo.AccountClient,
	assetClient repo.AssetClient,
	economyClient repo.EconomyClient,
	privacyClient repo.PrivacySettingClient,
	locationClient repo.LocationClient,
	relationClient repo.RelationClient,
) *AccountUsecase {
	return &AccountUsecase{
		accountClient:  accountClient,
		assetClient:    assetClient,
		economyClient:  economyClient,
		privacyClient:  privacyClient,
		locationClient: locationClient,
		relationClient: relationClient,
	}
}
func (u *AccountUsecase) GetCurrentAccount(ctx context.Context, userID int64) (*bbsuserv1.GetCurrentAccount_Resp_CurrentAccount, error) {
	reply, err := u.accountClient.GetCurrentAccount(ctx, userID)
	if err != nil {
		return nil, err
	}
	var account *bbsuserv1.GetCurrentAccount_Resp_CurrentAccount
	if reply != nil {
		if reply.Profile != nil {
			avatarURL := "/v1/user/account/avatar?name=" + reply.Profile.Name
			if reply.Profile.AvatarAssetID != nil && *reply.Profile.AvatarAssetID > 0 && u.assetClient != nil {
				asset, err := u.assetClient.Get(ctx, *reply.Profile.AvatarAssetID)
				if err != nil {
					return nil, err
				}
				if asset != nil && asset.URL != "" {
					avatarURL = asset.URL
				}
			}
			reply.Profile.AvatarURL = &avatarURL
		}
		account = &bbsuserv1.GetCurrentAccount_Resp_CurrentAccount{}
		if profile := reply.Profile; profile != nil {
			account.Profile = &bbsuserv1.AccountProfile{
				Id:                profile.ID,
				Name:              profile.Name,
				Nickname:          profile.Nickname,
				Url:               profile.URL,
				AvatarUrl:         profile.AvatarURL,
				BackgroundAssetId: profile.BackgroundAssetID,
				Introduction:      profile.Introduction,
				Status:            bbsuserv1enum.AccountStatus(profile.Status),
				Mbti:              profile.MBTI,
				FollowCount:       profile.FollowCount,
				FollowerCount:     profile.FollowerCount,
			}
			if profile.CreatedAt != nil {
				account.Profile.CreatedAt = timestamppb.New(*profile.CreatedAt)
			}
			if profile.UpdatedAt != nil {
				account.Profile.UpdatedAt = timestamppb.New(*profile.UpdatedAt)
			}
		}
		if contact := reply.Contact; contact != nil {
			account.Contact = &bbsuserv1.AccountContact{
				UserId: contact.UserID,
				Email:  contact.Email,
				Phone:  contact.Phone,
			}
		}
	}
	return account, nil
}

func (u *AccountUsecase) GetProfile(ctx context.Context, name string, viewerID int64) (*bbsuserv1.Profile, error) {
	account, lastLogin, err := u.accountClient.GetProfile(ctx, name)
	if err != nil {
		return nil, err
	}
	privacy, err := u.privacyClient.GetCurrentPrivacySetting(ctx, account.ID)
	if err != nil {
		return nil, err
	}
	profile := &bbsuserv1.Profile{
		Account:    u.accountProfile(account),
		Visibility: u.profileVisibility(privacy),
	}
	if lastLogin != nil {
		profile.LastSuccessLoginAt = timestamppb.New(*lastLogin)
	}
	if account.BackgroundAssetID != nil && *account.BackgroundAssetID > 0 {
		asset, err := u.assetClient.Get(ctx, *account.BackgroundAssetID)
		if err != nil {
			return nil, err
		}
		if asset != nil && asset.URL != "" {
			profile.BackgroundUrl = &asset.URL
		}
	}
	if viewerID == account.ID || privacy.PublicLocation == nil || *privacy.PublicLocation {
		location, err := u.locationClient.GetCurrentLocation(ctx, account.ID)
		if err != nil {
			return nil, err
		}
		if location != nil {
			profile.Location = &bbsuserv1.ProfileLocation{Country: location.Country, Province: location.Province, City: location.City}
		}
	}
	if viewerID != 0 {
		status, err := u.relationClient.GetStatus(ctx, &repo.GetStatusRelationReq{ActorID: viewerID, TargetID: account.ID})
		if err != nil {
			return nil, err
		}
		profile.ViewerRelation = u.profileRelation(status)
	}
	return profile, nil
}

func (u *AccountUsecase) ListFollowing(ctx context.Context, name string, viewerID int64, page *repo.PageReq) (*bbsuserv1.ListFollowing_Resp, error) {
	account, _, err := u.accountClient.GetProfile(ctx, name)
	if err != nil {
		return nil, err
	}
	privacy, err := u.privacyClient.GetCurrentPrivacySetting(ctx, account.ID)
	if err != nil {
		return nil, err
	}
	if viewerID != account.ID && privacy.PublicFollowingList != nil && !*privacy.PublicFollowingList {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_PROFILE_FOLLOWING_LIST_PRIVATE)
	}
	result, err := u.relationClient.ListFollowing(ctx, &repo.ListFollowingRelationsReq{ActorID: account.ID, Page: page})
	if err != nil {
		return nil, err
	}
	return u.listFollowingProfileAccounts(ctx, viewerID, result)
}

func (u *AccountUsecase) ListFollowers(ctx context.Context, name string, viewerID int64, page *repo.PageReq) (*bbsuserv1.ListFollowers_Resp, error) {
	account, _, err := u.accountClient.GetProfile(ctx, name)
	if err != nil {
		return nil, err
	}
	privacy, err := u.privacyClient.GetCurrentPrivacySetting(ctx, account.ID)
	if err != nil {
		return nil, err
	}
	if viewerID != account.ID && privacy.PublicFollowerList != nil && !*privacy.PublicFollowerList {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_PROFILE_FOLLOWER_LIST_PRIVATE)
	}
	result, err := u.relationClient.ListFollowers(ctx, &repo.ListFollowersRelationsReq{ActorID: account.ID, Page: page})
	if err != nil {
		return nil, err
	}
	return u.listFollowerProfileAccounts(ctx, viewerID, result)
}

func (u *AccountUsecase) listFollowingProfileAccounts(ctx context.Context, viewerID int64, result *repo.ListFollowingRelationsResp) (*bbsuserv1.ListFollowing_Resp, error) {
	ids := make([]int64, 0, len(result.Rows))
	for _, relation := range result.Rows {
		ids = append(ids, relation.TargetID)
	}
	rows, err := u.profileListItems(ctx, viewerID, ids)
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.ListFollowing_Resp{
		Page: u.profilePage(result.Page),
		Rows: rows,
	}, nil
}

func (u *AccountUsecase) listFollowerProfileAccounts(ctx context.Context, viewerID int64, result *repo.ListFollowersRelationsResp) (*bbsuserv1.ListFollowers_Resp, error) {
	ids := make([]int64, 0, len(result.Rows))
	for _, relation := range result.Rows {
		ids = append(ids, relation.ActorID)
	}
	rows, err := u.profileListItems(ctx, viewerID, ids)
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.ListFollowers_Resp{
		Page: u.profilePage(result.Page),
		Rows: rows,
	}, nil
}

func (u *AccountUsecase) profileListItems(ctx context.Context, viewerID int64, userIDs []int64) ([]*bbsuserv1.AccountProfileListItem, error) {
	accounts, err := u.accountClient.MapProfileAccounts(ctx, userIDs)
	if err != nil {
		return nil, err
	}
	statuses := map[int64]*repo.RelationStatus{}
	if viewerID != 0 {
		statuses, err = u.relationClient.MapStatus(ctx, viewerID, userIDs)
		if err != nil {
			return nil, err
		}
	}
	rows := make([]*bbsuserv1.AccountProfileListItem, 0, len(userIDs))
	for _, userID := range userIDs {
		if account := accounts[userID]; account != nil {
			rows = append(rows, &bbsuserv1.AccountProfileListItem{
				Account:        u.accountProfile(account),
				ViewerRelation: u.profileRelation(statuses[userID]),
			})
		}
	}
	return rows, nil
}

func (*AccountUsecase) accountProfile(account *repo.AccountProfile) *bbsuserv1.AccountProfile {
	if account == nil {
		return nil
	}
	profile := &bbsuserv1.AccountProfile{
		Id:                account.ID,
		Name:              account.Name,
		Nickname:          account.Nickname,
		Url:               account.URL,
		AvatarUrl:         account.AvatarURL,
		BackgroundAssetId: account.BackgroundAssetID,
		Introduction:      account.Introduction,
		Mbti:              account.MBTI,
		Status:            bbsuserv1enum.AccountStatus(account.Status),
		FollowCount:       account.FollowCount,
		FollowerCount:     account.FollowerCount,
	}
	if account.CreatedAt != nil {
		profile.CreatedAt = timestamppb.New(*account.CreatedAt)
	}
	if account.UpdatedAt != nil {
		profile.UpdatedAt = timestamppb.New(*account.UpdatedAt)
	}
	return profile
}

func (*AccountUsecase) profileVisibility(setting *repo.PrivacySetting) *bbsuserv1.ProfileVisibility {
	return &bbsuserv1.ProfileVisibility{
		ArticleList:    setting.PublicArticleList == nil || *setting.PublicArticleList,
		CommentList:    setting.PublicCommentList == nil || *setting.PublicCommentList,
		FollowerList:   setting.PublicFollowerList == nil || *setting.PublicFollowerList,
		FollowingList:  setting.PublicFollowingList == nil || *setting.PublicFollowingList,
		MoonbreezeList: setting.PublicMoonbreezeList == nil || *setting.PublicMoonbreezeList,
	}
}

func (*AccountUsecase) profileRelation(status *repo.RelationStatus) *bbsuserv1.ProfileRelation {
	if status == nil {
		return &bbsuserv1.ProfileRelation{}
	}
	return &bbsuserv1.ProfileRelation{
		Following:  status.Following,
		FollowedBy: status.FollowedBy,
		Blocking:   status.Blocking,
		BlockedBy:  status.BlockedBy,
	}
}

func (*AccountUsecase) profilePage(page *repo.PageResp) *commonv1.PageResp {
	if page == nil {
		return nil
	}
	return &commonv1.PageResp{Page: page.Page, Size: page.Size, Total: page.Total}
}

type UpdateProfileAccountReq struct {
	UserID       int64
	Nickname     *string
	URL          *string
	Introduction *string
	Mbti         *string
}

func (u *AccountUsecase) UpdateProfileAccount(ctx context.Context, req *UpdateProfileAccountReq) (*bbsuserv1.AccountProfile, error) {
	reply, err := u.accountClient.UpdateProfileAccount(ctx, &repo.UpdateProfileAccountReq{
		UserID:       req.UserID,
		Nickname:     req.Nickname,
		URL:          req.URL,
		Introduction: req.Introduction,
		MBTI:         req.Mbti,
	})
	if err != nil {
		return nil, err
	}
	var profile *bbsuserv1.AccountProfile
	if row := reply; row != nil {
		avatarURL := "/v1/user/account/avatar?name=" + row.Name
		if row.AvatarAssetID != nil && *row.AvatarAssetID > 0 && u.assetClient != nil {
			asset, err := u.assetClient.Get(ctx, *row.AvatarAssetID)
			if err != nil {
				return nil, err
			}
			if asset != nil && asset.URL != "" {
				avatarURL = asset.URL
			}
		}
		row.AvatarURL = &avatarURL
		profile = &bbsuserv1.AccountProfile{
			Id:                row.ID,
			Name:              row.Name,
			Nickname:          row.Nickname,
			Url:               row.URL,
			AvatarUrl:         row.AvatarURL,
			BackgroundAssetId: row.BackgroundAssetID,
			Introduction:      row.Introduction,
			Status:            bbsuserv1enum.AccountStatus(row.Status),
			Mbti:              row.MBTI,
			FollowCount:       row.FollowCount,
			FollowerCount:     row.FollowerCount,
		}
		if row.CreatedAt != nil {
			profile.CreatedAt = timestamppb.New(*row.CreatedAt)
		}
		if row.UpdatedAt != nil {
			profile.UpdatedAt = timestamppb.New(*row.UpdatedAt)
		}
	}
	return profile, nil
}

type PrepareProfileImageUploadAccountReq struct {
	UserID   int64
	Purpose  bbsuserv1.ProfileImagePurpose
	Hash     string
	MimeType string
	Size     int64
}

type PrepareProfileImageUploadAccountResp struct {
	AssetID    *int64
	UploadURL  string
	FormFields map[string]string
}

// PrepareProfileImageUploadAccount applies profile-specific file limits before
// asking platform for a content-addressed MinIO upload policy.
func (u *AccountUsecase) PrepareProfileImageUploadAccount(ctx context.Context, req *PrepareProfileImageUploadAccountReq) (*PrepareProfileImageUploadAccountResp, error) {
	if req == nil || req.UserID <= 0 || !u.profileImagePurposeValid(req.Purpose) || !u.validProfileImageHashAndMime(req.Hash, req.MimeType) || req.Size <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_PROFILE_IMAGE_INVALID)
	}
	if req.Size > maxProfileImageSize {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_PROFILE_IMAGE_TOO_LARGE)
	}
	hash := u.normalizeProfileImageHash(req.Hash)
	result, err := u.assetClient.PrepareDirectUpload(ctx, &repo.PrepareDirectAssetUploadReq{
		Hash:       hash,
		MimeType:   strings.ToLower(strings.TrimSpace(req.MimeType)),
		Size:       req.Size,
		UploadByID: req.UserID,
	})
	if err != nil {
		return nil, err
	}
	return &PrepareProfileImageUploadAccountResp{AssetID: result.AssetID, UploadURL: result.UploadURL, FormFields: result.FormFields}, nil
}

type CompleteProfileImageUploadAccountReq struct {
	UserID  int64
	Purpose bbsuserv1.ProfileImagePurpose
	Hash    string
}
type CompleteProfileImageUploadAccountResp struct {
	Profile  *bbsuserv1.AccountProfile
	ImageURL string
}

// CompleteProfileImageUploadAccount is safely repeatable while MinIO's event
// callback races with the browser. It updates only the user-owned reference.
func (u *AccountUsecase) CompleteProfileImageUploadAccount(ctx context.Context, req *CompleteProfileImageUploadAccountReq) (*CompleteProfileImageUploadAccountResp, error) {
	if req == nil || !u.profileImagePurposeValid(req.Purpose) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_PROFILE_IMAGE_INVALID)
	}
	hash := u.normalizeProfileImageHash(req.Hash)
	if !assetHashPattern.MatchString(hash) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_PROFILE_IMAGE_INVALID)
	}
	asset, err := u.assetClient.GetAvailableByHash(ctx, hash)
	if err != nil {
		return nil, err
	}
	if asset == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	// Complete does not trust the original browser declaration. The completed
	// Asset metadata is the provider-confirmed fact used for the final bind.
	if !u.validProfileImageHashAndMime(hash, asset.MimeType) || asset.Size <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_PROFILE_IMAGE_INVALID)
	}
	if asset.Size > maxProfileImageSize {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_PROFILE_IMAGE_TOO_LARGE)
	}
	update := &repo.UpdateProfileAccountReq{UserID: req.UserID}
	if req.Purpose == bbsuserv1.ProfileImagePurpose_PROFILE_IMAGE_PURPOSE_AVATAR {
		update.AvatarAssetID = &asset.ID
	} else {
		update.BackgroundAssetID = &asset.ID
	}
	row, err := u.accountClient.UpdateProfileAccount(ctx, update)
	if err != nil {
		return nil, err
	}
	profile := u.accountProfile(row)
	if profile != nil {
		if req.Purpose == bbsuserv1.ProfileImagePurpose_PROFILE_IMAGE_PURPOSE_AVATAR {
			profile.AvatarUrl = &asset.URL
		} else {
			profile.AvatarUrl = row.AvatarURL
		}
	}
	return &CompleteProfileImageUploadAccountResp{Profile: profile, ImageURL: asset.URL}, nil
}

func (*AccountUsecase) profileImagePurposeValid(p bbsuserv1.ProfileImagePurpose) bool {
	switch p {
	case bbsuserv1.ProfileImagePurpose_PROFILE_IMAGE_PURPOSE_AVATAR,
		bbsuserv1.ProfileImagePurpose_PROFILE_IMAGE_PURPOSE_BACKGROUND:
		return true
	}
	return false
}

const maxProfileImageSize = 2 * 1024 * 1024

func (u *AccountUsecase) validProfileImageHashAndMime(hash, mimeType string) bool {
	if !assetHashPattern.MatchString(u.normalizeProfileImageHash(hash)) {
		return false
	}
	switch strings.ToLower(strings.TrimSpace(mimeType)) {
	case "image/gif", "image/jpeg", "image/png", "image/webp":
		return true
	}
	return false
}

func (*AccountUsecase) normalizeProfileImageHash(value string) string {
	return strings.ToLower(strings.TrimSpace(value))
}

type UpdatePasswordAccountReq struct {
	UserID      int64
	OldPassword string
	NewPassword string
}

func (u *AccountUsecase) UpdatePasswordAccount(ctx context.Context, req *UpdatePasswordAccountReq) error {
	return u.accountClient.UpdatePasswordAccount(ctx, &repo.UpdatePasswordAccountReq{
		UserID:      req.UserID,
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	})
}

type UpdateEmailAccountReq struct {
	UserID int64
	Email  string
	Code   string
}

func (u *AccountUsecase) UpdateEmailAccount(ctx context.Context, req *UpdateEmailAccountReq) error {
	return u.accountClient.UpdateEmailAccount(ctx, &repo.UpdateEmailAccountReq{
		UserID: req.UserID,
		Email:  req.Email,
		Code:   req.Code,
	})
}

type UpdatePhoneAccountReq struct {
	UserID int64
	Phone  string
	Code   string
}

func (u *AccountUsecase) UpdatePhoneAccount(ctx context.Context, req *UpdatePhoneAccountReq) error {
	return u.accountClient.UpdatePhoneAccount(ctx, &repo.UpdatePhoneAccountReq{
		UserID: req.UserID,
		Phone:  req.Phone,
		Code:   req.Code,
	})
}

type AvatarAccountResp struct {
	Data        []byte
	ContentType string
}

func (u *AccountUsecase) AvatarAccount(ctx context.Context, name string) (*AvatarAccountResp, error) {
	reply, err := u.accountClient.AvatarAccount(ctx, name)
	if err != nil {
		return nil, err
	}
	return &AvatarAccountResp{
		Data:        reply.Data,
		ContentType: reply.ContentType,
	}, nil
}

type AccountEconomyResp struct {
	Balance      int64
	TotalIncome  int64
	TotalExpense int64
}

func (u *AccountUsecase) GetEconomyAccount(ctx context.Context, userID int64) (*AccountEconomyResp, error) {
	account, err := u.economyClient.GetAccount(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &AccountEconomyResp{Balance: account.Balance, TotalIncome: account.TotalIncome, TotalExpense: account.TotalExpense}, nil
}

type ListAccountEconomyRecordsReq struct {
	UserID     int64
	Page       *repo.PageReq
	Direction  *economyv1enum.EconomyRecordDirection
	RecordType *economyv1enum.EconomyRecordType
}

type ListAccountEconomyRecordsResp struct {
	Rows []*repo.EconomyRecord
	Page *repo.PageResp
}

func (u *AccountUsecase) ListEconomyRecords(ctx context.Context, req *ListAccountEconomyRecordsReq) (*ListAccountEconomyRecordsResp, error) {
	resp, err := u.economyClient.ListRecords(ctx, &repo.ListEconomyRecordsReq{UserID: req.UserID, Page: req.Page, Direction: req.Direction, RecordType: req.RecordType})
	if err != nil {
		return nil, err
	}
	return &ListAccountEconomyRecordsResp{Rows: resp.Rows, Page: resp.Page}, nil
}
