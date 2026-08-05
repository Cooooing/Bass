package repo

import (
	cerrors "common/proto/gen/common/errors"
	"context"
	"economy/internal/biz/base"
	"economy/internal/biz/model"
	bizrepo "economy/internal/biz/repo"
	"economy/internal/data/gen"
	accountent "economy/internal/data/gen/account"

	"common/pkg/apperror"
	utilent "common/pkg/util/ent"
	"github.com/samber/lo"
)

var _ bizrepo.AccountRepo = (*AccountRepo)(nil)

type AccountRepo struct {
	pageNormalizer
	db *gen.Client
}

func NewAccountRepo(db *gen.Client) bizrepo.AccountRepo {
	return &AccountRepo{db: db}
}

func (r *AccountRepo) getClient(ctx context.Context) *gen.Client {
	if tx, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return tx
	}
	return r.db
}

func (r *AccountRepo) Save(ctx context.Context, account *model.Account) (*model.Account, error) {
	save, err := r.getClient(ctx).Account.Create().
		SetUserID(account.UserID).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.model(save), nil
}

func (r *AccountRepo) UpdateBalance(ctx context.Context, req *bizrepo.AccountUpdateBalanceReq) (*model.Account, error) {
	update := r.getClient(ctx).Account.Update().Where(accountent.UserIDEQ(req.UserID), accountent.DeletedAtIsNil())
	if req.BalanceDelta != 0 {
		update = update.AddBalance(req.BalanceDelta)
	}
	if req.IncomeDelta != 0 {
		update = update.AddTotalIncome(req.IncomeDelta)
	}
	if req.ExpenseDelta != 0 {
		update = update.AddTotalExpense(req.ExpenseDelta)
	}
	if _, err := update.Save(ctx); err != nil {
		return nil, err
	}
	return r.Get(ctx, &bizrepo.AccountGetReq{UserID: new(req.UserID)})
}

func (r *AccountRepo) Get(ctx context.Context, req *bizrepo.AccountGetReq) (*model.Account, error) {
	query := r.getClient(ctx).Account.Query()
	query = r.getQuery(query, req)
	row, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return r.model(row), nil
}

func (r *AccountRepo) List(ctx context.Context, req *bizrepo.AccountGetReq) ([]*model.Account, error) {
	query := r.getClient(ctx).Account.Query()
	query = r.getQuery(query, req)
	rows, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	return lo.Map(rows, func(row *gen.Account, _ int) *model.Account { return r.model(row) }), nil
}

func (r *AccountRepo) Map(ctx context.Context, req *bizrepo.AccountGetReq) (map[int64]*model.Account, error) {
	rows, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	return lo.SliceToMap(rows, func(row *model.Account) (int64, *model.Account) { return row.UserID, row }), nil
}

func (r *AccountRepo) Count(ctx context.Context, req *bizrepo.AccountGetReq) (int, error) {
	query := r.getClient(ctx).Account.Query()
	query = r.getQuery(query, req)
	return query.Count(ctx)
}

func (r *AccountRepo) Page(ctx context.Context, req *bizrepo.AccountGetReq) (*bizrepo.AccountPageResp, error) {
	page := r.normalizePage(req.Page)
	query := r.getClient(ctx).Account.Query()
	query = r.getQuery(query, req)
	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := query.Limit(int(page.Size)).Offset(int((page.Page - 1) * page.Size)).All(ctx)
	if err != nil {
		return nil, err
	}
	return &bizrepo.AccountPageResp{
		Rows: lo.Map(rows, func(row *gen.Account, _ int) *model.Account { return r.model(row) }),
		Page: &base.PageResp{Total: int64(total), Page: page.Page, Size: page.Size},
	}, nil
}

func (r *AccountRepo) getQuery(query *gen.AccountQuery, req *bizrepo.AccountGetReq) *gen.AccountQuery {
	query = query.Where(accountent.DeletedAtIsNil())
	if req == nil {
		return query
	}
	if req.ID != nil {
		query = query.Where(accountent.IDEQ(*req.ID))
	}
	if len(req.IDs) > 0 {
		query = query.Where(accountent.IDIn(req.IDs...))
	}
	if req.UserID != nil {
		query = query.Where(accountent.UserIDEQ(*req.UserID))
	}
	if len(req.UserIDs) > 0 {
		query = query.Where(accountent.UserIDIn(req.UserIDs...))
	}
	return query
}

func (r *AccountRepo) model(row *gen.Account) *model.Account {
	return &model.Account{ID: row.ID, UserID: row.UserID, Balance: row.Balance, TotalIncome: row.TotalIncome, TotalExpense: row.TotalExpense, CreatedAt: row.CreatedAt, UpdatedAt: row.UpdatedAt, DeletedAt: row.DeletedAt}
}
