package service

import (
	"context"

	"common/pkg/apperror"
	"common/pkg/constant"
	commonmodel "common/pkg/model"
	"common/pkg/util"
	bffbbsv1 "common/proto/gen/bff_bbs/v1"
	cerrors "common/proto/gen/common/errors"
	imv1 "common/proto/gen/im/v1"

	"bff_bbs/internal/biz/model"
	"bff_bbs/internal/biz/usecase"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

// EmojiService 是浏览器表情分组请求的 BFF 协议适配层。
type EmojiService struct {
	bffbbsv1.UnimplementedEmojiServiceServer
	emojis *usecase.EmojiUsecase
}

func NewEmojiService(emojis *usecase.EmojiUsecase) *EmojiService {
	return &EmojiService{emojis: emojis}
}

func (*EmojiService) RegisterGrpc(*grpc.Server) {}

func (s *EmojiService) RegisterHTTP(server *http.Server) {
	bffbbsv1.RegisterEmojiServiceHTTPServer(server, s)
}

func (s *EmojiService) List(ctx context.Context, req *bffbbsv1.ListEmojiGroups_Req) (*bffbbsv1.ListEmojiGroups_Resp, error) {
	user, err := s.currentUser(ctx)
	if err != nil {
		return nil, err
	}
	groups, err := s.emojis.List(ctx, user.ID, req.GetQuery())
	if err != nil {
		return nil, err
	}
	rows, err := s.emojis.Groups(ctx, groups.GetRows())
	if err != nil {
		return nil, err
	}
	return &bffbbsv1.ListEmojiGroups_Resp{Rows: s.groups(rows)}, nil
}

func (s *EmojiService) CreateGroup(ctx context.Context, req *bffbbsv1.CreateEmojiGroup_Req) (*bffbbsv1.CreateEmojiGroup_Resp, error) {
	user, err := s.currentUser(ctx)
	if err != nil {
		return nil, err
	}
	group, err := s.emojis.CreateGroup(ctx, user.ID, req.GetName())
	if err != nil {
		return nil, err
	}
	rows, err := s.emojis.Groups(ctx, []*imv1.EmojiGroup{group})
	if err != nil {
		return nil, err
	}
	return &bffbbsv1.CreateEmojiGroup_Resp{Group: s.groups(rows)[0]}, nil
}

func (s *EmojiService) RenameGroup(ctx context.Context, req *bffbbsv1.RenameEmojiGroup_Req) (*bffbbsv1.RenameEmojiGroup_Resp, error) {
	user, err := s.currentUser(ctx)
	if err != nil {
		return nil, err
	}
	group, err := s.emojis.RenameGroup(ctx, user.ID, req.GetGroupId(), req.GetName())
	if err != nil {
		return nil, err
	}
	rows, err := s.emojis.Groups(ctx, []*imv1.EmojiGroup{group})
	if err != nil {
		return nil, err
	}
	return &bffbbsv1.RenameEmojiGroup_Resp{Group: s.groups(rows)[0]}, nil
}

func (s *EmojiService) DeleteGroup(ctx context.Context, req *bffbbsv1.DeleteEmojiGroup_Req) (*bffbbsv1.DeleteEmojiGroup_Resp, error) {
	user, err := s.currentUser(ctx)
	if err != nil {
		return nil, err
	}
	if err = s.emojis.DeleteGroup(ctx, user.ID, req.GetGroupId()); err != nil {
		return nil, err
	}
	return &bffbbsv1.DeleteEmojiGroup_Resp{}, nil
}

func (s *EmojiService) AddItem(ctx context.Context, req *bffbbsv1.AddEmojiItem_Req) (*bffbbsv1.AddEmojiItem_Resp, error) {
	user, err := s.currentUser(ctx)
	if err != nil {
		return nil, err
	}
	item, err := s.emojis.AddItem(ctx, user.ID, req.GroupId, req.GetAssetId(), req.GetName())
	if err != nil {
		return nil, err
	}
	rows, err := s.emojis.Groups(ctx, []*imv1.EmojiGroup{{Items: []*imv1.EmojiItem{item}}})
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 || len(rows[0].Items) == 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_NOT_FOUND)
	}
	return &bffbbsv1.AddEmojiItem_Resp{Item: s.item(rows[0].Items[0])}, nil
}

func (s *EmojiService) DeleteItem(ctx context.Context, req *bffbbsv1.DeleteEmojiItem_Req) (*bffbbsv1.DeleteEmojiItem_Resp, error) {
	user, err := s.currentUser(ctx)
	if err != nil {
		return nil, err
	}
	if err = s.emojis.DeleteItem(ctx, user.ID, req.GetItemId()); err != nil {
		return nil, err
	}
	return &bffbbsv1.DeleteEmojiItem_Resp{}, nil
}

func (s *EmojiService) MoveItem(ctx context.Context, req *bffbbsv1.MoveEmojiItem_Req) (*bffbbsv1.MoveEmojiItem_Resp, error) {
	user, err := s.currentUser(ctx)
	if err != nil {
		return nil, err
	}
	item, err := s.emojis.MoveItem(ctx, user.ID, req.GetItemId(), req.GetGroupId())
	if err != nil {
		return nil, err
	}
	rows, err := s.emojis.Groups(ctx, []*imv1.EmojiGroup{{Items: []*imv1.EmojiItem{item}}})
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 || len(rows[0].Items) == 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_NOT_FOUND)
	}
	return &bffbbsv1.MoveEmojiItem_Resp{Item: s.item(rows[0].Items[0])}, nil
}

func (s *EmojiService) RenameItem(ctx context.Context, req *bffbbsv1.RenameEmojiItem_Req) (*bffbbsv1.RenameEmojiItem_Resp, error) {
	user, err := s.currentUser(ctx)
	if err != nil {
		return nil, err
	}
	item, err := s.emojis.RenameItem(ctx, user.ID, req.GetItemId(), req.GetName())
	if err != nil {
		return nil, err
	}
	rows, err := s.emojis.Groups(ctx, []*imv1.EmojiGroup{{Items: []*imv1.EmojiItem{item}}})
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 || len(rows[0].Items) == 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_NOT_FOUND)
	}
	return &bffbbsv1.RenameEmojiItem_Resp{Item: s.item(rows[0].Items[0])}, nil
}

func (s *EmojiService) ReorderGroups(ctx context.Context, req *bffbbsv1.ReorderEmojiGroups_Req) (*bffbbsv1.ReorderEmojiGroups_Resp, error) {
	user, err := s.currentUser(ctx)
	if err != nil {
		return nil, err
	}
	if err = s.emojis.ReorderGroups(ctx, user.ID, req.GetGroupIds()); err != nil {
		return nil, err
	}
	return &bffbbsv1.ReorderEmojiGroups_Resp{}, nil
}

func (s *EmojiService) ReorderItems(ctx context.Context, req *bffbbsv1.ReorderEmojiItems_Req) (*bffbbsv1.ReorderEmojiItems_Resp, error) {
	user, err := s.currentUser(ctx)
	if err != nil {
		return nil, err
	}
	if err = s.emojis.ReorderItems(ctx, user.ID, req.GroupId, req.GetItemIds()); err != nil {
		return nil, err
	}
	return &bffbbsv1.ReorderEmojiItems_Resp{}, nil
}

func (s *EmojiService) CreateShare(ctx context.Context, req *bffbbsv1.CreateEmojiShare_Req) (*bffbbsv1.CreateEmojiShare_Resp, error) {
	user, err := s.currentUser(ctx)
	if err != nil {
		return nil, err
	}
	code, expiresInSeconds, err := s.emojis.CreateShare(ctx, user.ID, req.GetGroupId())
	if err != nil {
		return nil, err
	}
	return &bffbbsv1.CreateEmojiShare_Resp{Code: code, ExpiresInSeconds: expiresInSeconds}, nil
}

func (s *EmojiService) ImportShare(ctx context.Context, req *bffbbsv1.ImportEmojiShare_Req) (*bffbbsv1.ImportEmojiShare_Resp, error) {
	user, err := s.currentUser(ctx)
	if err != nil {
		return nil, err
	}
	group, skipped, err := s.emojis.ImportShare(ctx, user.ID, req.GetCode())
	if err != nil {
		return nil, err
	}
	rows, err := s.emojis.Groups(ctx, []*imv1.EmojiGroup{group})
	if err != nil {
		return nil, err
	}
	return &bffbbsv1.ImportEmojiShare_Resp{Group: s.groups(rows)[0], SkippedCount: skipped}, nil
}

func (s *EmojiService) groups(groups []*model.EmojiGroup) []*bffbbsv1.EmojiGroup {
	rows := make([]*bffbbsv1.EmojiGroup, 0, len(groups))
	for _, group := range groups {
		row := &bffbbsv1.EmojiGroup{Id: group.ID, Name: group.Name, Items: make([]*bffbbsv1.EmojiItem, 0, len(group.Items))}
		for _, item := range group.Items {
			row.Items = append(row.Items, s.item(item))
		}
		rows = append(rows, row)
	}
	return rows
}

func (*EmojiService) item(item *model.EmojiItem) *bffbbsv1.EmojiItem {
	return &bffbbsv1.EmojiItem{Id: item.ID, AssetId: item.AssetID, Name: item.Name, ImageUrl: item.ImageURL}
}

func (*EmojiService) currentUser(ctx context.Context) (*commonmodel.User, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil || user.ID <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	return user, nil
}
