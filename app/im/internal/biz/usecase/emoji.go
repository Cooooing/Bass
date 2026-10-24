package usecase

import (
	"context"
	cryptorand "crypto/rand"
	"encoding/base64"
	"strconv"
	"strings"
	"time"

	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"

	"im/internal/biz/base"
	"im/internal/biz/model"
	"im/internal/biz/repo"
	"im/internal/config"
)

// EmojiUsecase 管理账号表情分组及临时分享快照。
type EmojiUsecase struct {
	tx       base.Tx
	groups   repo.EmojiGroupRepo
	assets   repo.EmojiAssetClient
	shares   repo.EmojiShareCache
	shareTTL time.Duration
}

func NewEmojiUsecase(
	tx base.Tx,
	groups repo.EmojiGroupRepo,
	assets repo.EmojiAssetClient,
	shares repo.EmojiShareCache,
	conf *config.Bootstrap,
) *EmojiUsecase {
	shareTTL := 7 * 24 * time.Hour
	if conf != nil && conf.GetBusiness() != nil && conf.GetBusiness().GetEmojiShareTtl() != nil && conf.GetBusiness().GetEmojiShareTtl().AsDuration() > 0 {
		shareTTL = conf.GetBusiness().GetEmojiShareTtl().AsDuration()
	}
	return &EmojiUsecase{tx: tx, groups: groups, assets: assets, shares: shares, shareTTL: shareTTL}
}

func (u *EmojiUsecase) List(ctx context.Context, userID int64, query string) ([]*model.EmojiGroup, error) {
	if userID <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	return u.groups.List(ctx, userID, strings.TrimSpace(query))
}

func (u *EmojiUsecase) CreateGroup(ctx context.Context, userID int64, name string) (*model.EmojiGroup, error) {
	name = strings.TrimSpace(name)
	if userID <= 0 || name == "" || len([]rune(name)) > 32 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	return u.groups.Create(ctx, &model.EmojiGroup{UserID: userID, Name: name})
}

func (u *EmojiUsecase) RenameGroup(ctx context.Context, userID, groupID int64, name string) (*model.EmojiGroup, error) {
	name = strings.TrimSpace(name)
	if userID <= 0 || groupID <= 0 || name == "" || len([]rune(name)) > 32 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	group, err := u.groups.Rename(ctx, userID, groupID, name)
	if err != nil {
		return nil, err
	}
	if group == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_NOT_FOUND)
	}
	return group, nil
}

func (u *EmojiUsecase) DeleteGroup(ctx context.Context, userID, groupID int64) error {
	if userID <= 0 || groupID <= 0 {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	group, err := u.groups.Get(ctx, userID, groupID)
	if err != nil {
		return err
	}
	if group == nil {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_NOT_FOUND)
	}
	return u.tx(ctx, func(txCtx context.Context) error { return u.groups.Delete(txCtx, userID, groupID) })
}

func (u *EmojiUsecase) AddItem(ctx context.Context, userID int64, groupID *int64, assetID int64, name string) (*model.EmojiItem, error) {
	name = strings.TrimSpace(name)
	if userID <= 0 || assetID <= 0 || (groupID != nil && *groupID <= 0) || len([]rune(name)) > 64 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if err := u.assets.ValidateAvailable(ctx, assetID); err != nil {
		return nil, err
	}
	var item *model.EmojiItem
	err := u.tx(ctx, func(txCtx context.Context) error {
		resolvedGroupID := groupID
		if resolvedGroupID == nil {
			group, groupErr := u.groups.GetDefault(txCtx, userID)
			if groupErr != nil {
				return groupErr
			}
			if group == nil {
				group, groupErr = u.groups.Create(txCtx, &model.EmojiGroup{UserID: userID, Name: "默认分组"})
				if groupErr != nil {
					return groupErr
				}
			}
			resolvedGroupID = &group.ID
		}
		var createErr error
		item, createErr = u.groups.CreateItem(txCtx, &model.EmojiItem{UserID: userID, GroupID: resolvedGroupID, AssetID: assetID, Name: name})
		return createErr
	})
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (u *EmojiUsecase) RenameItem(ctx context.Context, userID, itemID int64, name string) (*model.EmojiItem, error) {
	name = strings.TrimSpace(name)
	if userID <= 0 || itemID <= 0 || name == "" || len([]rune(name)) > 64 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	item, err := u.groups.RenameItem(ctx, userID, itemID, name)
	if err != nil {
		return nil, err
	}
	if item == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_NOT_FOUND)
	}
	return item, nil
}

func (u *EmojiUsecase) DeleteItem(ctx context.Context, userID, itemID int64) error {
	if userID <= 0 || itemID <= 0 {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	return u.groups.DeleteItem(ctx, userID, itemID)
}

func (u *EmojiUsecase) MoveItem(ctx context.Context, userID, itemID, groupID int64) (*model.EmojiItem, error) {
	if userID <= 0 || itemID <= 0 || groupID <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	var item *model.EmojiItem
	err := u.tx(ctx, func(txCtx context.Context) error {
		var moveErr error
		item, moveErr = u.groups.MoveItem(txCtx, userID, itemID, groupID)
		return moveErr
	})
	if err != nil {
		return nil, err
	}
	if item == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_NOT_FOUND)
	}
	return item, nil
}

func (u *EmojiUsecase) ReorderGroups(ctx context.Context, userID int64, ids []int64) error {
	if userID <= 0 {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	return u.tx(ctx, func(txCtx context.Context) error { return u.groups.ReorderGroups(txCtx, userID, ids) })
}

func (u *EmojiUsecase) ReorderItems(ctx context.Context, userID int64, groupID *int64, ids []int64) error {
	if userID <= 0 || (groupID != nil && *groupID <= 0) {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	return u.tx(ctx, func(txCtx context.Context) error { return u.groups.ReorderItems(txCtx, userID, groupID, ids) })
}

func (u *EmojiUsecase) CreateShare(ctx context.Context, userID, groupID int64) (string, time.Duration, error) {
	if userID <= 0 || groupID <= 0 {
		return "", 0, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	groups, err := u.groups.List(ctx, userID, "")
	if err != nil {
		return "", 0, err
	}
	var source *model.EmojiGroup
	for _, group := range groups {
		if group.ID == groupID {
			source = group
			break
		}
	}
	if source == nil {
		return "", 0, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_NOT_FOUND)
	}
	if len(source.Items) == 0 {
		return "", 0, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_IM_EMOJI_GROUP_EMPTY)
	}
	bytes := make([]byte, 24)
	if _, err = cryptorand.Read(bytes); err != nil {
		return "", 0, err
	}
	code := base64.RawURLEncoding.EncodeToString(bytes)
	if err = u.shares.Save(ctx, code, &model.EmojiShare{OwnerUserID: userID, GroupName: source.Name, Items: source.Items}, u.shareTTL); err != nil {
		return "", 0, err
	}
	return code, u.shareTTL, nil
}

func (u *EmojiUsecase) ImportShare(ctx context.Context, userID int64, code string) (*model.EmojiGroup, uint32, error) {
	if userID <= 0 || strings.TrimSpace(code) == "" {
		return nil, 0, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	snapshot, err := u.shares.Get(ctx, strings.TrimSpace(code))
	if err != nil {
		return nil, 0, err
	}
	if snapshot == nil || snapshot.GroupName == "" {
		return nil, 0, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_IM_EMOJI_SHARE_INVALID_OR_EXPIRED)
	}
	validItems := make([]*model.EmojiItem, 0, len(snapshot.Items))
	skipped := uint32(0)
	for _, item := range snapshot.Items {
		if item == nil || item.AssetID <= 0 || u.assets.ValidateAvailable(ctx, item.AssetID) != nil {
			skipped++
			continue
		}
		validItems = append(validItems, item)
	}
	if len(validItems) == 0 {
		return nil, skipped, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_IM_EMOJI_SHARE_INVALID_OR_EXPIRED)
	}
	var group *model.EmojiGroup
	err = u.tx(ctx, func(txCtx context.Context) error {
		groups, listErr := u.groups.List(txCtx, userID, "")
		if listErr != nil {
			return listErr
		}
		baseName := strings.TrimSpace(snapshot.GroupName) + "（导入）"
		name := baseName
		for suffix := 2; ; suffix++ {
			used := false
			for _, current := range groups {
				if current.Name == name {
					used = true
					break
				}
			}
			if !used {
				break
			}
			name = baseName + " " + strconv.Itoa(suffix)
		}
		var createErr error
		group, createErr = u.groups.Create(txCtx, &model.EmojiGroup{UserID: userID, Name: name})
		if createErr != nil {
			return createErr
		}
		groupID := group.ID
		for index, item := range validItems {
			created, itemErr := u.groups.CreateItem(txCtx, &model.EmojiItem{UserID: userID, GroupID: &groupID, AssetID: item.AssetID, Name: item.Name, Sort: int64(len(validItems) - index)})
			if itemErr != nil {
				createErr = itemErr
				return createErr
			}
			group.Items = append(group.Items, created)
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}
	return group, skipped, nil
}
