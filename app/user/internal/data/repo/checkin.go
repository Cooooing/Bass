package repo

import (
	"common/proto/gen/common"
	"context"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/data/gen"
	"user/internal/data/gen/checkinrecord"
	"user/internal/data/gen/checkinstat"

	"common/pkg/server"
	utilent "common/pkg/util/ent"
)

var _ repo.CheckinRecordRepo = (*CheckinRecordRepo)(nil)
var _ repo.CheckinStatRepo = (*CheckinStatRepo)(nil)

type CheckinRecordRepo struct {
	db *gen.Client
}

func NewCheckinRecordRepo(
	db *gen.Client,
) repo.CheckinRecordRepo {
	return &CheckinRecordRepo{
		db: db,
	}
}

func (r *CheckinRecordRepo) getClient(
	ctx context.Context,
) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func (r *CheckinRecordRepo) Get(
	ctx context.Context,
	req *repo.CheckinRecordGetReq,
) (*model.CheckinRecord, error) {
	record, err := r.get(ctx, req)
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (r *CheckinRecordRepo) List(
	ctx context.Context,
	req *repo.CheckinRecordGetReq,
) ([]*model.CheckinRecord, error) {
	rows, err := r.list(ctx, req)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *CheckinRecordRepo) Map(
	ctx context.Context,
	req *repo.CheckinRecordGetReq,
) (map[int64]*model.CheckinRecord, error) {
	rows, err := r.mapRows(ctx, req)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *CheckinRecordRepo) Count(
	ctx context.Context,
	req *repo.CheckinRecordGetReq,
) (int, error) {
	count, err := r.count(ctx, req)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *CheckinRecordRepo) Page(
	ctx context.Context,
	req *repo.CheckinRecordPageReq,
) (*repo.CheckinRecordPageResp, error) {
	rows, page, err := r.page(ctx, &common.PageReq{
		Page: req.Page.Page,
		Size: req.Page.Size,
	}, &req.Query)
	if err != nil {
		return nil, err
	}
	resp := repo.PageResp{}
	if page != nil {
		resp = repo.PageResp{
			Total: page.GetTotal(),
			Page:  page.GetPage(),
			Size:  page.GetSize(),
		}
	}
	return &repo.CheckinRecordPageResp{
		Rows: rows,
		Page: resp,
	}, nil
}

func (r *CheckinRecordRepo) UpsertRecord(
	ctx context.Context,
	record *model.CheckinRecord,
) (*model.CheckinRecord, error) {
	record, err := r.upsertRecord(ctx, record)
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (r *CheckinRecordRepo) get(
	ctx context.Context,
	req *repo.CheckinRecordGetReq,
) (*model.CheckinRecord, error) {
	tx := r.getClient(ctx)
	query := tx.CheckinRecord.Query()
	query = r.getQuery(query, req)
	row, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &model.CheckinRecord{
		ID:            row.ID,
		UserID:        row.UserID,
		Date:          new(row.Date),
		OnlineMinutes: new(row.OnlineMinutes),
		Activity:      new(row.Activity),
		Checked:       row.Checked,
	}, nil
}

func (r *CheckinRecordRepo) list(
	ctx context.Context,
	req *repo.CheckinRecordGetReq,
) ([]*model.CheckinRecord, error) {
	tx := r.getClient(ctx)
	query := tx.CheckinRecord.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.CheckinRecord, 0, len(list))
	for _, row := range list {
		result = append(result, &model.CheckinRecord{
			ID:            row.ID,
			UserID:        row.UserID,
			Date:          new(row.Date),
			OnlineMinutes: new(row.OnlineMinutes),
			Activity:      new(row.Activity),
			Checked:       row.Checked,
		})
	}
	return result, nil
}

func (r *CheckinRecordRepo) mapRows(
	ctx context.Context,
	req *repo.CheckinRecordGetReq,
) (map[int64]*model.CheckinRecord, error) {
	list, err := r.list(ctx, req)
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*model.CheckinRecord, len(list))
	for _, item := range list {
		result[item.ID] = item
	}
	return result, nil
}

func (r *CheckinRecordRepo) count(
	ctx context.Context,
	req *repo.CheckinRecordGetReq,
) (int, error) {
	tx := r.getClient(ctx)
	query := tx.CheckinRecord.Query()
	query = r.getQuery(query, req)
	return query.Count(ctx)
}

func (r *CheckinRecordRepo) page(
	ctx context.Context,
	page *common.PageReq,
	req *repo.CheckinRecordGetReq,
) ([]*model.CheckinRecord, *common.PageResp, error) {
	tx := r.getClient(ctx)
	page = server.PageValid(page)
	query := tx.CheckinRecord.Query()
	query = r.getQuery(query, req)
	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, nil, err
	}
	list, err := query.
		Limit(int(page.Size)).
		Offset(int((page.Page - 1) * page.Size)).
		All(ctx)
	if err != nil {
		return nil, nil, err
	}
	result := make([]*model.CheckinRecord, 0, len(list))
	for _, row := range list {
		result = append(result, &model.CheckinRecord{
			ID:            row.ID,
			UserID:        row.UserID,
			Date:          new(row.Date),
			OnlineMinutes: new(row.OnlineMinutes),
			Activity:      new(row.Activity),
			Checked:       row.Checked,
		})
	}
	return result, &common.PageResp{
		Total: uint32(total),
		Page:  page.Page,
		Size:  page.Size,
	}, nil
}

func (r *CheckinRecordRepo) upsertRecord(
	ctx context.Context,
	record *model.CheckinRecord,
) (*model.CheckinRecord, error) {
	tx := r.getClient(ctx)
	existing, err := tx.CheckinRecord.Query().
		Where(checkinrecord.UserID(record.UserID)).
		Where(checkinrecord.DateEQ(*record.Date)).
		Only(ctx)
	if err != nil && !gen.IsNotFound(err) {
		return nil, err
	}
	if existing != nil {
		update := tx.CheckinRecord.UpdateOneID(existing.ID).
			SetChecked(record.Checked)
		if record.OnlineMinutes != nil {
			update.SetOnlineMinutes(*record.OnlineMinutes)
		}
		if record.Activity != nil {
			update.SetActivity(*record.Activity)
		}
		saved, err := update.Save(ctx)
		if err != nil {
			return nil, err
		}
		return &model.CheckinRecord{
			ID:            saved.ID,
			UserID:        saved.UserID,
			Date:          new(saved.Date),
			OnlineMinutes: new(saved.OnlineMinutes),
			Activity:      new(saved.Activity),
			Checked:       saved.Checked,
		}, nil
	}
	create := tx.CheckinRecord.Create().
		SetUserID(record.UserID).
		SetDate(*record.Date).
		SetChecked(record.Checked)
	if record.OnlineMinutes != nil {
		create.SetOnlineMinutes(*record.OnlineMinutes)
	}
	if record.Activity != nil {
		create.SetActivity(*record.Activity)
	}
	saved, err := create.Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.CheckinRecord{
		ID:            saved.ID,
		UserID:        saved.UserID,
		Date:          new(saved.Date),
		OnlineMinutes: new(saved.OnlineMinutes),
		Activity:      new(saved.Activity),
		Checked:       saved.Checked,
	}, nil
}

func (r *CheckinRecordRepo) getQuery(
	query *gen.CheckinRecordQuery,
	req *repo.CheckinRecordGetReq,
) *gen.CheckinRecordQuery {
	if req == nil {
		return query
	}
	if req.ID != nil {
		query = query.Where(checkinrecord.ID(*req.ID))
	}
	if len(req.IDs) > 0 {
		query = query.Where(checkinrecord.IDIn(req.IDs...))
	}
	if req.UserID != nil {
		query = query.Where(checkinrecord.UserID(*req.UserID))
	}
	if len(req.UserIDs) > 0 {
		query = query.Where(checkinrecord.UserIDIn(req.UserIDs...))
	}
	if req.Date != nil {
		query = query.Where(checkinrecord.DateEQ(*req.Date))
	}
	return query
}

type CheckinStatRepo struct {
	db *gen.Client
}

func NewCheckinStatRepo(
	db *gen.Client,
) repo.CheckinStatRepo {
	return &CheckinStatRepo{
		db: db,
	}
}

func (r *CheckinStatRepo) getClient(
	ctx context.Context,
) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func (r *CheckinStatRepo) Get(
	ctx context.Context,
	req *repo.CheckinStatGetReq,
) (*model.CheckinStat, error) {
	stat, err := r.get(ctx, req)
	if err != nil {
		return nil, err
	}
	return stat, nil
}

func (r *CheckinStatRepo) List(
	ctx context.Context,
	req *repo.CheckinStatGetReq,
) ([]*model.CheckinStat, error) {
	rows, err := r.list(ctx, req)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *CheckinStatRepo) Map(
	ctx context.Context,
	req *repo.CheckinStatGetReq,
) (map[int64]*model.CheckinStat, error) {
	rows, err := r.mapRows(ctx, req)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *CheckinStatRepo) Count(
	ctx context.Context,
	req *repo.CheckinStatGetReq,
) (int, error) {
	count, err := r.count(ctx, req)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *CheckinStatRepo) Page(
	ctx context.Context,
	req *repo.CheckinStatPageReq,
) (*repo.CheckinStatPageResp, error) {
	rows, page, err := r.page(ctx, &common.PageReq{
		Page: req.Page.Page,
		Size: req.Page.Size,
	}, &req.Query)
	if err != nil {
		return nil, err
	}
	resp := repo.PageResp{}
	if page != nil {
		resp = repo.PageResp{
			Total: page.GetTotal(),
			Page:  page.GetPage(),
			Size:  page.GetSize(),
		}
	}
	return &repo.CheckinStatPageResp{
		Rows: rows,
		Page: resp,
	}, nil
}

func (r *CheckinStatRepo) UpsertStat(
	ctx context.Context,
	stat *model.CheckinStat,
) (*model.CheckinStat, error) {
	stat, err := r.upsertStat(ctx, stat)
	if err != nil {
		return nil, err
	}
	return stat, nil
}

func (r *CheckinStatRepo) get(
	ctx context.Context,
	req *repo.CheckinStatGetReq,
) (*model.CheckinStat, error) {
	tx := r.getClient(ctx)
	query := tx.CheckinStat.Query()
	query = r.getQuery(query, req)
	s, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &model.CheckinStat{
		ID:                 s.ID,
		UserID:             s.UserID,
		TotalOnlineMinutes: new(s.TotalOnlineMinutes),
		CurrentStreak:      new(s.CurrentStreak),
		LongestStreak:      new(s.LongestStreak),
	}, nil
}

func (r *CheckinStatRepo) list(
	ctx context.Context,
	req *repo.CheckinStatGetReq,
) ([]*model.CheckinStat, error) {
	tx := r.getClient(ctx)
	query := tx.CheckinStat.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.CheckinStat, 0, len(list))
	for _, s := range list {
		result = append(result, &model.CheckinStat{
			ID:                 s.ID,
			UserID:             s.UserID,
			TotalOnlineMinutes: new(s.TotalOnlineMinutes),
			CurrentStreak:      new(s.CurrentStreak),
			LongestStreak:      new(s.LongestStreak),
		})
	}
	return result, nil
}

func (r *CheckinStatRepo) mapRows(
	ctx context.Context,
	req *repo.CheckinStatGetReq,
) (map[int64]*model.CheckinStat, error) {
	list, err := r.list(ctx, req)
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*model.CheckinStat, len(list))
	for _, item := range list {
		result[item.ID] = item
	}
	return result, nil
}

func (r *CheckinStatRepo) count(
	ctx context.Context,
	req *repo.CheckinStatGetReq,
) (int, error) {
	tx := r.getClient(ctx)
	query := tx.CheckinStat.Query()
	query = r.getQuery(query, req)
	return query.Count(ctx)
}

func (r *CheckinStatRepo) page(
	ctx context.Context,
	page *common.PageReq,
	req *repo.CheckinStatGetReq,
) ([]*model.CheckinStat, *common.PageResp, error) {
	tx := r.getClient(ctx)
	page = server.PageValid(page)
	query := tx.CheckinStat.Query()
	query = r.getQuery(query, req)
	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, nil, err
	}
	list, err := query.
		Limit(int(page.Size)).
		Offset(int((page.Page - 1) * page.Size)).
		All(ctx)
	if err != nil {
		return nil, nil, err
	}
	result := make([]*model.CheckinStat, 0, len(list))
	for _, s := range list {
		result = append(result, &model.CheckinStat{
			ID:                 s.ID,
			UserID:             s.UserID,
			TotalOnlineMinutes: new(s.TotalOnlineMinutes),
			CurrentStreak:      new(s.CurrentStreak),
			LongestStreak:      new(s.LongestStreak),
		})
	}
	return result, &common.PageResp{
		Total: uint32(total),
		Page:  page.Page,
		Size:  page.Size,
	}, nil
}

func (r *CheckinStatRepo) upsertStat(
	ctx context.Context,
	stat *model.CheckinStat,
) (*model.CheckinStat, error) {
	existing, err := r.get(ctx, &repo.CheckinStatGetReq{
		UserID: &stat.UserID,
	})
	if err != nil {
		return nil, err
	}
	if existing == nil {
		tx := r.getClient(ctx)
		create := tx.CheckinStat.Create().
			SetUserID(stat.UserID)
		if stat.TotalOnlineMinutes != nil {
			create.SetTotalOnlineMinutes(*stat.TotalOnlineMinutes)
		}
		if stat.CurrentStreak != nil {
			create.SetCurrentStreak(*stat.CurrentStreak)
		}
		if stat.LongestStreak != nil {
			create.SetLongestStreak(*stat.LongestStreak)
		}
		saved, err := create.Save(ctx)
		if err != nil {
			return nil, err
		}
		return &model.CheckinStat{
			ID:                 saved.ID,
			UserID:             saved.UserID,
			TotalOnlineMinutes: new(saved.TotalOnlineMinutes),
			CurrentStreak:      new(saved.CurrentStreak),
			LongestStreak:      new(saved.LongestStreak),
		}, nil
	}
	if stat.TotalOnlineMinutes == nil && stat.CurrentStreak == nil && stat.LongestStreak == nil {
		return existing, nil
	}
	stat.ID = existing.ID
	tx := r.getClient(ctx)
	update := tx.CheckinStat.UpdateOneID(stat.ID)
	if stat.TotalOnlineMinutes != nil {
		update.SetTotalOnlineMinutes(*stat.TotalOnlineMinutes)
	}
	if stat.CurrentStreak != nil {
		update.SetCurrentStreak(*stat.CurrentStreak)
	}
	if stat.LongestStreak != nil {
		update.SetLongestStreak(*stat.LongestStreak)
	}
	saved, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.CheckinStat{
		ID:                 saved.ID,
		UserID:             saved.UserID,
		TotalOnlineMinutes: new(saved.TotalOnlineMinutes),
		CurrentStreak:      new(saved.CurrentStreak),
		LongestStreak:      new(saved.LongestStreak),
	}, nil
}

func (r *CheckinStatRepo) getQuery(
	query *gen.CheckinStatQuery,
	req *repo.CheckinStatGetReq,
) *gen.CheckinStatQuery {
	if req == nil {
		return query
	}
	if req.ID != nil {
		query = query.Where(checkinstat.ID(*req.ID))
	}
	if len(req.IDs) > 0 {
		query = query.Where(checkinstat.IDIn(req.IDs...))
	}
	if req.UserID != nil {
		query = query.Where(checkinstat.UserID(*req.UserID))
	}
	if len(req.UserIDs) > 0 {
		query = query.Where(checkinstat.UserIDIn(req.UserIDs...))
	}
	return query
}
