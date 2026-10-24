package usecase

import (
	"context"

	"common/pkg/client/rpc"
	imv1 "common/proto/gen/im/v1"

	"bff_bbs/internal/biz/model"
	"bff_bbs/internal/biz/repo"
)

// EmojiUsecase 聚合 IM 表情分组事实与 Platform 的公开图片地址。
type EmojiUsecase struct {
	emojis imv1.IMEmojiServiceClient
	assets repo.AssetClient
}

func NewEmojiUsecase(client *rpc.IMClient, assets repo.AssetClient) *EmojiUsecase {
	return &EmojiUsecase{emojis: client.Emoji, assets: assets}
}

func (u *EmojiUsecase) List(ctx context.Context, userID int64, query string) (*imv1.ListEmojiGroups_Resp, error) {
	return u.emojis.List(ctx, &imv1.ListEmojiGroups_Req{UserId: userID, Query: query})
}

func (u *EmojiUsecase) RenameItem(ctx context.Context, userID, itemID int64, name string) (*imv1.EmojiItem, error) {
	response, err := u.emojis.RenameItem(ctx, &imv1.RenameEmojiItem_Req{UserId: userID, ItemId: itemID, Name: name})
	if err != nil {
		return nil, err
	}
	return response.GetItem(), nil
}

func (u *EmojiUsecase) Groups(ctx context.Context, groups []*imv1.EmojiGroup) ([]*model.EmojiGroup, error) {
	ids := make([]int64, 0)
	for _, group := range groups {
		for _, item := range group.GetItems() {
			if item.GetAssetId() > 0 {
				ids = append(ids, item.GetAssetId())
			}
		}
	}
	assets, err := u.assets.Map(ctx, &repo.AssetGetReq{IDs: ids})
	if err != nil {
		return nil, err
	}
	rows := make([]*model.EmojiGroup, 0, len(groups))
	for _, group := range groups {
		row := &model.EmojiGroup{ID: group.GetId(), Name: group.GetName(), Items: make([]*model.EmojiItem, 0, len(group.GetItems()))}
		for _, item := range group.GetItems() {
			asset := assets[item.GetAssetId()]
			if asset == nil || asset.URL == "" {
				continue
			}
			row.Items = append(row.Items, &model.EmojiItem{ID: item.GetId(), AssetID: item.GetAssetId(), Name: item.GetName(), ImageURL: asset.URL})
		}
		rows = append(rows, row)
	}
	return rows, nil
}

func (u *EmojiUsecase) CreateGroup(ctx context.Context, userID int64, name string) (*imv1.EmojiGroup, error) {
	response, err := u.emojis.CreateGroup(ctx, &imv1.CreateEmojiGroup_Req{UserId: userID, Name: name})
	if err != nil {
		return nil, err
	}
	return response.GetGroup(), nil
}

func (u *EmojiUsecase) RenameGroup(ctx context.Context, userID, groupID int64, name string) (*imv1.EmojiGroup, error) {
	response, err := u.emojis.RenameGroup(ctx, &imv1.RenameEmojiGroup_Req{UserId: userID, GroupId: groupID, Name: name})
	if err != nil {
		return nil, err
	}
	return response.GetGroup(), nil
}

func (u *EmojiUsecase) DeleteGroup(ctx context.Context, userID, groupID int64) error {
	_, err := u.emojis.DeleteGroup(ctx, &imv1.DeleteEmojiGroup_Req{UserId: userID, GroupId: groupID})
	return err
}

func (u *EmojiUsecase) AddItem(ctx context.Context, userID int64, groupID *int64, assetID int64, name string) (*imv1.EmojiItem, error) {
	response, err := u.emojis.AddItem(ctx, &imv1.AddEmojiItem_Req{UserId: userID, GroupId: groupID, AssetId: assetID, Name: name})
	if err != nil {
		return nil, err
	}
	return response.GetItem(), nil
}

func (u *EmojiUsecase) ReorderGroups(ctx context.Context, userID int64, ids []int64) error {
	_, err := u.emojis.ReorderGroups(ctx, &imv1.ReorderEmojiGroups_Req{UserId: userID, GroupIds: ids})
	return err
}

func (u *EmojiUsecase) ReorderItems(ctx context.Context, userID int64, groupID *int64, ids []int64) error {
	_, err := u.emojis.ReorderItems(ctx, &imv1.ReorderEmojiItems_Req{UserId: userID, GroupId: groupID, ItemIds: ids})
	return err
}

func (u *EmojiUsecase) DeleteItem(ctx context.Context, userID, itemID int64) error {
	_, err := u.emojis.DeleteItem(ctx, &imv1.DeleteEmojiItem_Req{UserId: userID, ItemId: itemID})
	return err
}

func (u *EmojiUsecase) MoveItem(ctx context.Context, userID, itemID, groupID int64) (*imv1.EmojiItem, error) {
	response, err := u.emojis.MoveItem(ctx, &imv1.MoveEmojiItem_Req{UserId: userID, ItemId: itemID, GroupId: groupID})
	if err != nil {
		return nil, err
	}
	return response.GetItem(), nil
}

func (u *EmojiUsecase) CreateShare(ctx context.Context, userID, groupID int64) (string, int64, error) {
	response, err := u.emojis.CreateShare(ctx, &imv1.CreateEmojiShare_Req{UserId: userID, GroupId: groupID})
	if err != nil {
		return "", 0, err
	}
	return response.GetCode(), response.GetExpiresInSeconds(), nil
}

func (u *EmojiUsecase) ImportShare(ctx context.Context, userID int64, code string) (*imv1.EmojiGroup, uint32, error) {
	response, err := u.emojis.ImportShare(ctx, &imv1.ImportEmojiShare_Req{UserId: userID, Code: code})
	if err != nil {
		return nil, 0, err
	}
	return response.GetGroup(), response.GetSkippedCount(), nil
}
