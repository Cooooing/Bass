package repo

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	"context"
	bizrepo "game_town/internal/biz/repo"
	"game_town/internal/data/gen"
	"game_town/internal/data/gen/location"
	"game_town/internal/data/gen/world"
	"game_town/internal/data/gen/worldmember"
	"time"
)

type WorldMemberRepo struct{ *baseRepo }

func NewWorldMemberRepo(db *gen.Client) bizrepo.WorldMemberRepo {
	return &WorldMemberRepo{baseRepo: &baseRepo{db: db}}
}

func (r *WorldMemberRepo) JoinWorld(ctx context.Context, req *bizrepo.JoinWorldReq) (*bizrepo.JoinWorldResponse, error) {
	worldRow, err := r.db.World.Query().Where(world.Code(req.WorldCode), world.DeletedAtIsNil()).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_TOWN_WORLD_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	if worldRow.DefaultLocationID == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_TOWN_WORLD_INVALID)
	}
	now := time.Now()
	memberRow, err := r.db.WorldMember.Query().Where(worldmember.WorldID(worldRow.ID), worldmember.PlayerID(req.PlayerID)).Only(ctx)
	if gen.IsNotFound(err) {
		memberRow, err = r.db.WorldMember.Create().SetWorldID(worldRow.ID).SetPlayerID(req.PlayerID).SetCurrentLocationID(*worldRow.DefaultLocationID).SetRole("member").SetJoinedAt(now).SetLastSeenAt(now).Save(ctx)
	} else if err == nil {
		memberRow, err = r.db.WorldMember.UpdateOneID(memberRow.ID).SetLastSeenAt(now).Save(ctx)
	}
	if err != nil {
		return nil, err
	}
	locationRow, err := r.db.Location.Query().Where(location.ID(memberRow.CurrentLocationID), location.DeletedAtIsNil()).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_TOWN_LOCATION_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return &bizrepo.JoinWorldResponse{World: r.world(worldRow), Member: r.member(memberRow), Location: r.location(locationRow)}, nil
}

func (r *WorldMemberRepo) GetMember(ctx context.Context, req *bizrepo.GetMemberReq) (*bizrepo.GetMemberResponse, error) {
	row, err := r.db.WorldMember.Query().Where(worldmember.WorldID(req.WorldID), worldmember.PlayerID(req.PlayerID)).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
	}
	if err != nil {
		return nil, err
	}
	return &bizrepo.GetMemberResponse{Row: r.member(row)}, nil
}

func (r *WorldMemberRepo) MoveMember(ctx context.Context, req *bizrepo.MoveMemberReq) (*bizrepo.MoveMemberResponse, error) {
	current, err := r.db.WorldMember.Query().Where(worldmember.WorldID(req.WorldID), worldmember.PlayerID(req.PlayerID)).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
	}
	if err != nil {
		return nil, err
	}
	row, err := r.db.WorldMember.UpdateOneID(current.ID).SetCurrentLocationID(req.LocationID).SetLastSeenAt(time.Now()).Save(ctx)
	if err != nil {
		return nil, err
	}
	return &bizrepo.MoveMemberResponse{Row: r.member(row)}, nil
}
