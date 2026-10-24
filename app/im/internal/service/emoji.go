package service

import (
	"context"

	imv1 "common/proto/gen/im/v1"

	"im/internal/biz/model"
	"im/internal/biz/usecase"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	transporthttp "github.com/go-kratos/kratos/v3/transport/http"
)

// EmojiService 是 IM 表情分组的 gRPC 协议适配层。
type EmojiService struct {
	imv1.UnimplementedIMEmojiServiceServer
	emojis *usecase.EmojiUsecase
}

func NewEmojiService(emojis *usecase.EmojiUsecase) *EmojiService {
	return &EmojiService{emojis: emojis}
}

func (s *EmojiService) RegisterGrpc(server *grpc.Server) {
	imv1.RegisterIMEmojiServiceServer(server, s)
}

func (*EmojiService) RegisterHTTP(*transporthttp.Server) {}

func (s *EmojiService) List(ctx context.Context, req *imv1.ListEmojiGroups_Req) (*imv1.ListEmojiGroups_Resp, error) {
	groups, err := s.emojis.List(ctx, req.GetUserId(), req.GetQuery())
	if err != nil {
		return nil, err
	}
	rows := make([]*imv1.EmojiGroup, 0, len(groups))
	for _, group := range groups {
		rows = append(rows, s.group(group))
	}
	return &imv1.ListEmojiGroups_Resp{Rows: rows}, nil
}

func (s *EmojiService) CreateGroup(ctx context.Context, req *imv1.CreateEmojiGroup_Req) (*imv1.CreateEmojiGroup_Resp, error) {
	group, err := s.emojis.CreateGroup(ctx, req.GetUserId(), req.GetName())
	if err != nil {
		return nil, err
	}
	return &imv1.CreateEmojiGroup_Resp{Group: s.group(group)}, nil
}

func (s *EmojiService) RenameGroup(ctx context.Context, req *imv1.RenameEmojiGroup_Req) (*imv1.RenameEmojiGroup_Resp, error) {
	group, err := s.emojis.RenameGroup(ctx, req.GetUserId(), req.GetGroupId(), req.GetName())
	if err != nil {
		return nil, err
	}
	return &imv1.RenameEmojiGroup_Resp{Group: s.group(group)}, nil
}

func (s *EmojiService) DeleteGroup(ctx context.Context, req *imv1.DeleteEmojiGroup_Req) (*imv1.DeleteEmojiGroup_Resp, error) {
	if err := s.emojis.DeleteGroup(ctx, req.GetUserId(), req.GetGroupId()); err != nil {
		return nil, err
	}
	return &imv1.DeleteEmojiGroup_Resp{}, nil
}

func (s *EmojiService) AddItem(ctx context.Context, req *imv1.AddEmojiItem_Req) (*imv1.AddEmojiItem_Resp, error) {
	item, err := s.emojis.AddItem(ctx, req.GetUserId(), req.GroupId, req.GetAssetId(), req.GetName())
	if err != nil {
		return nil, err
	}
	return &imv1.AddEmojiItem_Resp{Item: s.item(item)}, nil
}

func (s *EmojiService) DeleteItem(ctx context.Context, req *imv1.DeleteEmojiItem_Req) (*imv1.DeleteEmojiItem_Resp, error) {
	if err := s.emojis.DeleteItem(ctx, req.GetUserId(), req.GetItemId()); err != nil {
		return nil, err
	}
	return &imv1.DeleteEmojiItem_Resp{}, nil
}

func (s *EmojiService) MoveItem(ctx context.Context, req *imv1.MoveEmojiItem_Req) (*imv1.MoveEmojiItem_Resp, error) {
	item, err := s.emojis.MoveItem(ctx, req.GetUserId(), req.GetItemId(), req.GetGroupId())
	if err != nil {
		return nil, err
	}
	return &imv1.MoveEmojiItem_Resp{Item: s.item(item)}, nil
}

func (s *EmojiService) RenameItem(ctx context.Context, req *imv1.RenameEmojiItem_Req) (*imv1.RenameEmojiItem_Resp, error) {
	item, err := s.emojis.RenameItem(ctx, req.GetUserId(), req.GetItemId(), req.GetName())
	if err != nil {
		return nil, err
	}
	return &imv1.RenameEmojiItem_Resp{Item: s.item(item)}, nil
}

func (s *EmojiService) ReorderGroups(ctx context.Context, req *imv1.ReorderEmojiGroups_Req) (*imv1.ReorderEmojiGroups_Resp, error) {
	if err := s.emojis.ReorderGroups(ctx, req.GetUserId(), req.GetGroupIds()); err != nil {
		return nil, err
	}
	return &imv1.ReorderEmojiGroups_Resp{}, nil
}

func (s *EmojiService) ReorderItems(ctx context.Context, req *imv1.ReorderEmojiItems_Req) (*imv1.ReorderEmojiItems_Resp, error) {
	if err := s.emojis.ReorderItems(ctx, req.GetUserId(), req.GroupId, req.GetItemIds()); err != nil {
		return nil, err
	}
	return &imv1.ReorderEmojiItems_Resp{}, nil
}

func (s *EmojiService) CreateShare(ctx context.Context, req *imv1.CreateEmojiShare_Req) (*imv1.CreateEmojiShare_Resp, error) {
	code, ttl, err := s.emojis.CreateShare(ctx, req.GetUserId(), req.GetGroupId())
	if err != nil {
		return nil, err
	}
	return &imv1.CreateEmojiShare_Resp{Code: code, ExpiresInSeconds: int64(ttl.Seconds())}, nil
}

func (s *EmojiService) ImportShare(ctx context.Context, req *imv1.ImportEmojiShare_Req) (*imv1.ImportEmojiShare_Resp, error) {
	group, skipped, err := s.emojis.ImportShare(ctx, req.GetUserId(), req.GetCode())
	if err != nil {
		return nil, err
	}
	return &imv1.ImportEmojiShare_Resp{Group: s.group(group), SkippedCount: skipped}, nil
}

func (s *EmojiService) group(group *model.EmojiGroup) *imv1.EmojiGroup {
	if group == nil {
		return nil
	}
	items := make([]*imv1.EmojiItem, 0, len(group.Items))
	for _, item := range group.Items {
		items = append(items, s.item(item))
	}
	return &imv1.EmojiGroup{Id: group.ID, Name: group.Name, Sort: group.Sort, Items: items}
}

func (*EmojiService) item(item *model.EmojiItem) *imv1.EmojiItem {
	if item == nil {
		return nil
	}
	return &imv1.EmojiItem{Id: item.ID, AssetId: item.AssetID, Name: item.Name, Sort: item.Sort}
}
