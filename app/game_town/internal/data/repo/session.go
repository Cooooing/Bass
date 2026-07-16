package repo

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	"context"
	bizrepo "game_town/internal/biz/repo"
	"game_town/internal/data/gen"
	"game_town/internal/data/gen/session"
	"time"
)

type SessionRepo struct{ *baseRepo }

func NewSessionRepo(db *gen.Client) bizrepo.SessionRepo {
	return &SessionRepo{baseRepo: &baseRepo{db: db}}
}

func (r *SessionRepo) StartSession(ctx context.Context, req *bizrepo.StartSessionReq) (*bizrepo.StartSessionResponse, error) {
	now := time.Now()
	row, err := r.db.Session.Create().SetNillablePlayerID(req.PlayerID).SetClientType(req.ClientType).SetStartedAt(now).SetLastSeenAt(now).Save(ctx)
	if err != nil {
		return nil, err
	}
	return &bizrepo.StartSessionResponse{Row: r.session(row)}, nil
}

func (r *SessionRepo) EndSession(ctx context.Context, req *bizrepo.EndSessionReq) (*bizrepo.EndSessionResponse, error) {
	now := time.Now()
	row, err := r.db.Session.UpdateOneID(req.ID).SetEndedAt(now).SetLastSeenAt(now).Save(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_TOWN_SESSION_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return &bizrepo.EndSessionResponse{Row: r.session(row)}, nil
}

func (r *SessionRepo) GetSession(ctx context.Context, req *bizrepo.GetSessionReq) (*bizrepo.GetSessionResponse, error) {
	row, err := r.db.Session.Query().Where(session.ID(req.ID)).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_TOWN_SESSION_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return &bizrepo.GetSessionResponse{Row: r.session(row)}, nil
}

func (r *SessionRepo) UpdateSessionWorld(ctx context.Context, req *bizrepo.UpdateSessionWorldReq) (*bizrepo.UpdateSessionWorldResponse, error) {
	now := time.Now()
	row, err := r.db.Session.UpdateOneID(req.ID).SetPlayerID(req.PlayerID).SetCurrentWorldID(req.WorldID).SetLastSeenAt(now).Save(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_TOWN_SESSION_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return &bizrepo.UpdateSessionWorldResponse{Row: r.session(row)}, nil
}
