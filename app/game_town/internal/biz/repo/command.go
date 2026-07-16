package repo

import (
	"common/proto/gen/common"
	"context"
	"game_town/internal/biz/model"
)

type CommandRepo interface {
	CreateCommand(ctx context.Context, req *CreateCommandReq) (*CreateCommandResponse, error)
	FinishCommand(ctx context.Context, req *FinishCommandReq) (*FinishCommandResponse, error)
	Page(ctx context.Context, req *CommandPageReq) (*CommandPageResponse, error)
	List(ctx context.Context, req *CommandListReq) (*CommandListResponse, error)
}

type CreateCommandReq struct {
	Row *model.Command
}

type CreateCommandResponse struct {
	Row *model.Command
}

type FinishCommandReq struct {
	ID        int64
	Status    string
	Summary   string
	ErrorCode *int32
	WorldID   *int64
}

type FinishCommandResponse struct {
	Row *model.Command
}

type CommandPageReq struct {
	Page  *common.PageRequest
	Query CommandQuery
}

type CommandPageResponse struct {
	Rows []*model.Command
	Page *common.PageResponse
}

type CommandListReq struct {
	SessionID int64
	PlayerID  int64
}

type CommandListResponse struct {
	Rows []*model.Command
}

type CommandQuery struct {
	WorldID   *int64
	SessionID *int64
	PlayerID  *int64
}
