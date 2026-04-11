package repo

import (
	cv1 "common/api/gen/common/v1"
	"common/pkg/constant"
	"context"
	"infra/internal/biz/model"
	"infra/internal/biz/repo"
	"infra/internal/data/base"
	"infra/internal/data/ent/gen"
	"infra/internal/data/ent/gen/objectstorage"
)

type ObjectStorageRepo struct {
	*base.BaseData
}

func NewObjectStorageRepo(base *base.BaseData) repo.ObjectStorageRepo {
	return &ObjectStorageRepo{
		BaseData: base,
	}
}

func (r *ObjectStorageRepo) Save(ctx context.Context, tx *gen.Client, u *model.ObjectStorage) (*model.ObjectStorage, error) {
	save, err := tx.ObjectStorage.Create().
		SetProvider(u.Provider).
		SetBucket(u.Bucket).
		SetKey(u.Key).
		SetMimeType(u.MimeType).
		SetSize(u.Size).
		SetHash(u.Hash).
		SetUploadBy(u.UploadBy).
		SetUploadByName(u.UploadByName).
		Save(ctx)
	return &model.ObjectStorage{ObjectStorage: save}, err
}

func (r *ObjectStorageRepo) UpdateAudit(ctx context.Context, tx *gen.Client, u *model.ObjectStorage) error {
	_, err := tx.ObjectStorage.Update().
		Where(objectstorage.KeyEQ(u.Key)).
		SetNillableAuditCallbackReply(u.AuditCallbackReply).
		SetBlocked(u.Blocked).
		SetNillableBlockedReason(u.BlockedReason).
		SetNillableBlockedAt(u.BlockedAt).
		SetNillableBlockedBy(u.BlockedBy).
		SetNillableBlockedByName(u.BlockedByName).
		Save(ctx)
	return err
}

func (r *ObjectStorageRepo) Delete(ctx context.Context, tx *gen.Client, u *model.ObjectStorage) (int, error) {
	if u == nil {
		return 0, nil
	}
	if u.ID != 0 {
		return 1, tx.ObjectStorage.DeleteOneID(u.ID).Exec(ctx)
	}
	return tx.ObjectStorage.Delete().
		Where(objectstorage.KeyEQ(u.Key)).
		Exec(ctx)
}

func (r *ObjectStorageRepo) Exist(ctx context.Context, tx *gen.Client, req *repo.ObjectStorageGetReq) (bool, error) {
	query := tx.ObjectStorage.Query()
	query = r.getQuery(query, req)
	return query.Exist(ctx)
}

func (r *ObjectStorageRepo) GetList(ctx context.Context, tx *gen.Client, req *repo.ObjectStorageGetReq) ([]*model.ObjectStorage, error) {
	var (
		records []*model.ObjectStorage
		err     error
	)
	query := tx.ObjectStorage.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	for i := range list {
		records = append(records, &model.ObjectStorage{ObjectStorage: list[i]})
	}
	return records, nil
}

func (r *ObjectStorageRepo) GetPage(ctx context.Context, tx *gen.Client, page *cv1.PageRequest, req *repo.ObjectStorageGetReq) ([]*model.ObjectStorage, *cv1.PageReply, error) {
	var (
		notificationRecords []*model.ObjectStorage
		err                 error
	)
	page = constant.PageValid(page)
	query := tx.ObjectStorage.Query()
	query = r.getQuery(query, req)
	countQuery := query.Clone()
	count, err := countQuery.Count(ctx)
	if err != nil {
		return nil, nil, err
	}
	list, err := query.Limit(int(page.Size)).Offset(int((page.Page - 1) * page.Size)).All(ctx)
	if err != nil {
		return nil, nil, err
	}

	for _, item := range list {
		notificationRecords = append(notificationRecords, &model.ObjectStorage{ObjectStorage: item})
	}
	return notificationRecords, &cv1.PageReply{
		Total: uint32(count),
		Size:  page.Size,
		Page:  page.Page,
	}, nil
}

func (r *ObjectStorageRepo) getQuery(query *gen.ObjectStorageQuery, req *repo.ObjectStorageGetReq) *gen.ObjectStorageQuery {
	if req.Provider != nil {
		query = query.Where(objectstorage.ProviderEQ(*req.Provider))
	}
	if req.Bucket != nil {
		query = query.Where(objectstorage.BucketEQ(*req.Bucket))
	}
	if req.Key != nil {
		query = query.Where(objectstorage.KeyEQ(*req.Key))
	}
	if req.MimeType != nil {
		query = query.Where(objectstorage.MimeTypeEQ(*req.MimeType))
	}
	if req.Size != nil {
		if req.Size.Start != nil {
			query = query.Where(objectstorage.SizeGTE(*req.Size.Start))
		}
		if req.Size.End != nil {
			query = query.Where(objectstorage.SizeLTE(*req.Size.End))
		}
	}
	if req.Blocked != nil {
		query = query.Where(objectstorage.BlockedEQ(*req.Blocked))
	}
	if req.BlockedByName != nil {
		query = query.Where(objectstorage.BlockedByNameEQ(*req.BlockedByName))
	}
	return query
}
