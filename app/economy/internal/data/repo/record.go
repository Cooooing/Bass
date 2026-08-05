package repo

import (
	cerrors "common/proto/gen/common/errors"
	"context"
	"economy/internal/biz/base"
	"economy/internal/biz/model"
	bizrepo "economy/internal/biz/repo"
	"economy/internal/data/gen"
	recordent "economy/internal/data/gen/record"
	"economy/internal/enum"

	"common/pkg/apperror"
	utilent "common/pkg/util/ent"
	"github.com/samber/lo"
)

var _ bizrepo.RecordRepo = (*RecordRepo)(nil)

type RecordRepo struct {
	pageNormalizer
	db *gen.Client
}

func NewRecordRepo(db *gen.Client) bizrepo.RecordRepo {
	return &RecordRepo{db: db}
}

func (r *RecordRepo) getClient(ctx context.Context) *gen.Client {
	if tx, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return tx
	}
	return r.db
}

func (r *RecordRepo) Save(ctx context.Context, record *model.Record) (*model.Record, error) {
	save, err := r.getClient(ctx).Record.Create().
		SetTransactionNo(record.TransactionNo).
		SetUserID(record.UserID).
		SetRecordType(recordent.RecordType(record.RecordType)).
		SetDirection(recordent.Direction(record.Direction)).
		SetAmount(record.Amount).
		SetBalanceBefore(record.BalanceBefore).
		SetBalanceAfter(record.BalanceAfter).
		SetNillableSourceID(record.SourceID).
		SetIdempotencyKey(record.IdempotencyKey).
		SetNillableRemark(record.Remark).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.model(save), nil
}

func (r *RecordRepo) Get(ctx context.Context, req *bizrepo.RecordGetReq) (*model.Record, error) {
	query := r.getClient(ctx).Record.Query()
	query = r.getQuery(query, req)
	row, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_RECORD_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return r.model(row), nil
}

func (r *RecordRepo) List(ctx context.Context, req *bizrepo.RecordGetReq) ([]*model.Record, error) {
	query := r.getClient(ctx).Record.Query()
	query = r.getQuery(query, req)
	rows, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	return lo.Map(rows, func(row *gen.Record, _ int) *model.Record { return r.model(row) }), nil
}

func (r *RecordRepo) Map(ctx context.Context, req *bizrepo.RecordGetReq) (map[int64]*model.Record, error) {
	rows, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	return lo.SliceToMap(rows, func(row *model.Record) (int64, *model.Record) { return row.ID, row }), nil
}

func (r *RecordRepo) Count(ctx context.Context, req *bizrepo.RecordGetReq) (int, error) {
	query := r.getClient(ctx).Record.Query()
	query = r.getQuery(query, req)
	return query.Count(ctx)
}

func (r *RecordRepo) Page(ctx context.Context, req *bizrepo.RecordGetReq) (*bizrepo.RecordPageResp, error) {
	page := r.normalizePage(req.Page)
	query := r.getClient(ctx).Record.Query()
	query = r.getQuery(query, req)
	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := query.Limit(int(page.Size)).Offset(int((page.Page - 1) * page.Size)).All(ctx)
	if err != nil {
		return nil, err
	}
	return &bizrepo.RecordPageResp{
		Rows: lo.Map(rows, func(row *gen.Record, _ int) *model.Record { return r.model(row) }),
		Page: &base.PageResp{Total: int64(total), Page: page.Page, Size: page.Size},
	}, nil
}

func (r *RecordRepo) getQuery(query *gen.RecordQuery, req *bizrepo.RecordGetReq) *gen.RecordQuery {
	if req == nil {
		return query.Order(gen.Desc(recordent.FieldCreatedAt, recordent.FieldID))
	}
	if req.ID != nil {
		query = query.Where(recordent.IDEQ(*req.ID))
	}
	if len(req.IDs) > 0 {
		query = query.Where(recordent.IDIn(req.IDs...))
	}
	if req.TransactionNo != nil {
		query = query.Where(recordent.TransactionNoEQ(*req.TransactionNo))
	}
	if req.UserID != nil {
		query = query.Where(recordent.UserIDEQ(*req.UserID))
	}
	if len(req.UserIDs) > 0 {
		query = query.Where(recordent.UserIDIn(req.UserIDs...))
	}
	if req.RecordType != nil {
		query = query.Where(recordent.RecordTypeEQ(recordent.RecordType(*req.RecordType)))
	}
	if req.Direction != nil {
		query = query.Where(recordent.DirectionEQ(recordent.Direction(*req.Direction)))
	}
	if req.SourceID != nil {
		query = query.Where(recordent.SourceIDEQ(*req.SourceID))
	}
	if req.IdempotencyKey != nil {
		query = query.Where(recordent.IdempotencyKeyEQ(*req.IdempotencyKey))
	}
	return query.Order(gen.Desc(recordent.FieldCreatedAt, recordent.FieldID))
}

func (r *RecordRepo) model(row *gen.Record) *model.Record {
	return &model.Record{
		ID: row.ID, TransactionNo: row.TransactionNo, UserID: row.UserID, RecordType: enum.EconomyRecordType(row.RecordType), Direction: enum.EconomyRecordDirection(row.Direction),
		Amount: row.Amount, BalanceBefore: row.BalanceBefore, BalanceAfter: row.BalanceAfter, SourceID: row.SourceID, IdempotencyKey: row.IdempotencyKey, Remark: row.Remark, CreatedAt: row.CreatedAt,
	}
}
