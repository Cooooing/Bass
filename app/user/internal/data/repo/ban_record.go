package repo

import (
	"context"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/data/gen"
	"user/internal/data/gen/banrecord"
	"user/internal/enum"

	utilent "common/pkg/util/ent"
)

var _ repo.BanRecordRepo = (*BanRecordRepo)(nil)

type BanRecordRepo struct {
	db *gen.Client
}

func NewBanRecordRepo(db *gen.Client) repo.BanRecordRepo {
	return &BanRecordRepo{db: db}
}

func (r *BanRecordRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func (r *BanRecordRepo) Create(ctx context.Context, record *model.BanRecord) (*model.BanRecord, error) {
	create := r.getClient(ctx).BanRecord.Create().
		SetUserID(record.UserID).
		SetOperatorID(record.OperatorID).
		SetOperatorRealm(banrecord.OperatorRealm(record.OperatorRealm)).
		SetReason(record.Reason).
		SetRemark(record.Remark).
		SetStartedAt(record.StartedAt).
		SetNillableBannedUntil(record.BannedUntil)
	row, err := create.Save(ctx)
	if err != nil {
		return nil, err
	}
	return banRecordToModel(row), nil
}

func (r *BanRecordRepo) Get(ctx context.Context, id int64) (*model.BanRecord, error) {
	row, err := r.getClient(ctx).BanRecord.Query().Where(banrecord.ID(id)).First(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return banRecordToModel(row), nil
}

func (r *BanRecordRepo) LatestByUserID(ctx context.Context, userID int64) (*model.BanRecord, error) {
	row, err := r.getClient(ctx).BanRecord.Query().Where(banrecord.UserID(userID)).Order(gen.Desc(banrecord.FieldCreatedAt)).First(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return banRecordToModel(row), nil
}

func banRecordToModel(row *gen.BanRecord) *model.BanRecord {
	if row == nil {
		return nil
	}
	return &model.BanRecord{
		ID:            row.ID,
		UserID:        row.UserID,
		OperatorID:    row.OperatorID,
		OperatorRealm: enum.LoginRealm(row.OperatorRealm),
		Reason:        row.Reason,
		Remark:        row.Remark,
		StartedAt:     row.StartedAt,
		BannedUntil:   row.BannedUntil,
		CreatedAt:     row.CreatedAt,
		UpdatedAt:     row.UpdatedAt,
	}
}
