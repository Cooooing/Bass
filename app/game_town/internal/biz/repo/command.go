package repo

import (
	"common/proto/gen/common"
	"context"
	"game_town/internal/biz/model"
)

type CommandRepo interface {
	CreateCommand(ctx context.Context, row *model.Command) (*model.Command, error)
	FinishCommand(ctx context.Context, req *FinishCommandReq) (*model.Command, error)
	Page(ctx context.Context, req *CommandPageReq) (*CommandPageResp, error)
	List(ctx context.Context, req *CommandListReq) ([]*model.Command, error)
}

type FinishCommandReq struct {
	ID        int64
	Status    string
	Summary   string
	ErrorCode *int32
	WorldID   *int64
}

type CommandPageReq struct {
	Page  *common.PageReq
	Query CommandQuery
}

type CommandPageResp struct {
	Rows []*model.Command
	Page *common.PageResp
}

type CommandListReq struct {
	SessionID int64
	PlayerID  int64
}

type CommandQuery struct {
	WorldID   *int64
	SessionID *int64
	PlayerID  *int64
}
