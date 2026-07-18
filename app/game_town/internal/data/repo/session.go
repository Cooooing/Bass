package repo

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	"context"
	"game_town/internal/biz/model"
	bizrepo "game_town/internal/biz/repo"
	"game_town/internal/data/gen"
	"game_town/internal/data/gen/session"
	"time"
)

type SessionRepo struct{ *baseRepo }

func NewSessionRepo(db *gen.Client) bizrepo.SessionRepo {
	return &SessionRepo{baseRepo: &baseRepo{db: db}}
}

func (r *SessionRepo) StartSession(ctx context.Context, req *bizrepo.StartSessionReq) (*model.Session, error) {
	now := time.Now()
	row, err := r.db.Session.Create().SetNillablePlayerID(req.PlayerID).SetClientType(req.ClientType).SetStartedAt(now).SetLastSeenAt(now).Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.session(row), nil
}

func (r *SessionRepo) EndSession(ctx context.Context, id int64) (*model.Session, error) {
	now := time.Now()
	row, err := r.db.Session.UpdateOneID(id).SetEndedAt(now).SetLastSeenAt(now).Save(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_TOWN_SESSION_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return r.session(row), nil
}

func (r *SessionRepo) GetSession(ctx context.Context, id int64) (*model.Session, error) {
	row, err := r.db.Session.Query().Where(session.ID(id)).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_TOWN_SESSION_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return r.session(row), nil
}

func (r *SessionRepo) UpdateSessionWorld(ctx context.Context, req *bizrepo.UpdateSessionWorldReq) (*model.Session, error) {
	now := time.Now()
	row, err := r.db.Session.UpdateOneID(req.ID).SetPlayerID(req.PlayerID).SetCurrentWorldID(req.WorldID).SetLastSeenAt(now).Save(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_TOWN_SESSION_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return r.session(row), nil
}
