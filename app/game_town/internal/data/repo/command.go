package repo

import (
	"common/pkg/server"
	"common/proto/gen/common"
	"context"
	"game_town/internal/biz/model"
	bizrepo "game_town/internal/biz/repo"
	"game_town/internal/data/gen"
	"game_town/internal/data/gen/command"
	"time"
)

type CommandRepo struct{ *baseRepo }

func NewCommandRepo(db *gen.Client) bizrepo.CommandRepo {
	return &CommandRepo{baseRepo: &baseRepo{db: db}}
}

func (r *CommandRepo) CreateCommand(ctx context.Context, row *model.Command) (*model.Command, error) {
	now := time.Now()
	created, err := r.db.Command.Create().SetNillableWorldID(row.WorldID).SetSessionID(row.SessionID).SetNillablePlayerID(row.PlayerID).SetRawText(row.RawText).SetType(row.Type).SetParsedPayload(row.ParsedPayload).SetStatus(row.Status).SetNillableErrorCode(row.ErrorCode).SetResultSummary(row.ResultSummary).SetCreatedAt(now).SetNillableHandledAt(row.HandledAt).Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.command(created), nil
}

func (r *CommandRepo) FinishCommand(ctx context.Context, req *bizrepo.FinishCommandReq) (*model.Command, error) {
	now := time.Now()
	row, err := r.db.Command.UpdateOneID(req.ID).SetStatus(req.Status).SetResultSummary(req.Summary).SetNillableErrorCode(req.ErrorCode).SetNillableWorldID(req.WorldID).SetHandledAt(now).Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.command(row), nil
}

func (r *CommandRepo) Page(ctx context.Context, req *bizrepo.CommandPageReq) (*bizrepo.CommandPageResp, error) {
	pageReq := server.PageValid(req.Page)
	queryReq := req.Query
	query := r.db.Command.Query()
	if queryReq.WorldID != nil {
		query = query.Where(command.WorldID(*queryReq.WorldID))
	}
	if queryReq.SessionID != nil {
		query = query.Where(command.SessionID(*queryReq.SessionID))
	}
	if queryReq.PlayerID != nil {
		query = query.Where(command.PlayerID(*queryReq.PlayerID))
	}
	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := query.Order(gen.Desc(command.FieldCreatedAt)).Limit(int(pageReq.Size)).Offset(int((pageReq.Page - 1) * pageReq.Size)).All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.Command, 0, len(rows))
	for _, row := range rows {
		result = append(result, r.command(row))
	}
	return &bizrepo.CommandPageResp{Rows: result, Page: &common.PageResp{Total: uint32(total), Page: pageReq.Page, Size: pageReq.Size}}, nil
}

func (r *CommandRepo) List(ctx context.Context, req *bizrepo.CommandListReq) ([]*model.Command, error) {
	rows, err := r.db.Command.Query().Where(command.SessionID(req.SessionID), command.PlayerID(req.PlayerID)).Order(command.ByCreatedAt()).All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.Command, 0, len(rows))
	for _, row := range rows {
		result = append(result, r.command(row))
	}
	return result, nil
}
