package repo

import (
	"common/proto/gen/common"
	"context"
	"platform/internal/biz/model"
	"platform/internal/biz/repo"
	"platform/internal/data/gen"
	"platform/internal/data/gen/objectstorage"
	"platform/internal/enum"

	"common/pkg/server"
	utilent "common/pkg/util/ent"
)

var _ repo.ObjectStorageRepo = (*ObjectStorageRepo)(nil)

type ObjectStorageRepo struct {
	db *gen.Client
}

func NewObjectStorageRepo(db *gen.Client) repo.ObjectStorageRepo {
	return &ObjectStorageRepo{db: db}
}

func (r *ObjectStorageRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func (r *ObjectStorageRepo) Save(ctx context.Context, row *model.ObjectStorage) (*model.ObjectStorage, error) {
	saved, err := r.getClient(ctx).ObjectStorage.Create().
		SetProvider(objectstorage.Provider(row.Provider)).
		SetBucket(row.Bucket).
		SetKey(row.Key).
		SetMimeType(row.MimeType).
		SetSize(row.Size).
		SetHash(row.Hash).
		SetUploadBy(row.UploadBy).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.ObjectStorage{
		ID:                 saved.ID,
		Provider:           enum.ObjectStorageProvider(saved.Provider),
		Bucket:             saved.Bucket,
		Key:                saved.Key,
		MimeType:           saved.MimeType,
		Size:               saved.Size,
		Hash:               saved.Hash,
		UploadBy:           saved.UploadBy,
		AuditCallbackReply: saved.AuditCallbackReply,
		Blocked:            saved.Blocked,
		BlockedReason:      saved.BlockedReason,
		BlockedAt:          saved.BlockedAt,
		BlockedBy:          saved.BlockedBy,
		CreatedAt:          saved.CreatedAt,
		UpdatedAt:          saved.UpdatedAt,
	}, nil
}

func (r *ObjectStorageRepo) UpdateAudit(ctx context.Context, row *model.ObjectStorage) error {
	_, err := r.getClient(ctx).ObjectStorage.Update().
		Where(objectstorage.KeyEQ(row.Key)).
		SetNillableAuditCallbackReply(row.AuditCallbackReply).
		SetBlocked(row.Blocked).
		SetNillableBlockedReason(row.BlockedReason).
		SetNillableBlockedAt(row.BlockedAt).
		SetNillableBlockedBy(row.BlockedBy).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *ObjectStorageRepo) Delete(ctx context.Context, row *model.ObjectStorage) (int, error) {
	if row == nil {
		return 0, nil
	}
	if row.ID != 0 {
		err := r.getClient(ctx).ObjectStorage.DeleteOneID(row.ID).Exec(ctx)
		if err != nil {
			return 0, err
		}
		return 1, nil
	}
	count, err := r.getClient(ctx).ObjectStorage.Delete().
		Where(objectstorage.KeyEQ(row.Key)).
		Exec(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *ObjectStorageRepo) Exist(ctx context.Context, req *repo.ObjectStorageGetReq) (bool, error) {
	query := r.getClient(ctx).ObjectStorage.Query()
	query = r.getQuery(query, req)
	exists, err := query.Exist(ctx)
	if err != nil {
		return false, err
	}
	return exists, nil
}
func (r *ObjectStorageRepo) Get(ctx context.Context, req *repo.ObjectStorageGetReq) (*model.ObjectStorage, error) {
	query := r.getClient(ctx).ObjectStorage.Query()
	query = r.getQuery(query, req)
	row, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &model.ObjectStorage{
		ID:                 row.ID,
		Provider:           enum.ObjectStorageProvider(row.Provider),
		Bucket:             row.Bucket,
		Key:                row.Key,
		MimeType:           row.MimeType,
		Size:               row.Size,
		Hash:               row.Hash,
		UploadBy:           row.UploadBy,
		AuditCallbackReply: row.AuditCallbackReply,
		Blocked:            row.Blocked,
		BlockedReason:      row.BlockedReason,
		BlockedAt:          row.BlockedAt,
		BlockedBy:          row.BlockedBy,
		CreatedAt:          row.CreatedAt,
		UpdatedAt:          row.UpdatedAt,
	}, nil
}

func (r *ObjectStorageRepo) List(ctx context.Context, req *repo.ObjectStorageGetReq) ([]*model.ObjectStorage, error) {
	query := r.getClient(ctx).ObjectStorage.Query()
	query = r.getQuery(query, req)
	rows, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.ObjectStorage, 0, len(rows))
	for _, row := range rows {
		result = append(result, &model.ObjectStorage{
			ID:                 row.ID,
			Provider:           enum.ObjectStorageProvider(row.Provider),
			Bucket:             row.Bucket,
			Key:                row.Key,
			MimeType:           row.MimeType,
			Size:               row.Size,
			Hash:               row.Hash,
			UploadBy:           row.UploadBy,
			AuditCallbackReply: row.AuditCallbackReply,
			Blocked:            row.Blocked,
			BlockedReason:      row.BlockedReason,
			BlockedAt:          row.BlockedAt,
			BlockedBy:          row.BlockedBy,
			CreatedAt:          row.CreatedAt,
			UpdatedAt:          row.UpdatedAt,
		})
	}
	return result, nil
}

func (r *ObjectStorageRepo) Map(ctx context.Context, req *repo.ObjectStorageGetReq) (map[int64]*model.ObjectStorage, error) {
	rows, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*model.ObjectStorage, len(rows))
	for _, row := range rows {
		result[row.ID] = row
	}
	return result, nil
}

func (r *ObjectStorageRepo) Count(ctx context.Context, req *repo.ObjectStorageGetReq) (int, error) {
	query := r.getClient(ctx).ObjectStorage.Query()
	query = r.getQuery(query, req)
	count, err := query.Count(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *ObjectStorageRepo) Page(ctx context.Context, req *repo.ObjectStoragePageReq) (*repo.ObjectStoragePageResp, error) {
	if req == nil {
		req = &repo.ObjectStoragePageReq{}
	}
	page := server.PageValid(req.Page)
	query := r.getClient(ctx).ObjectStorage.Query()
	query = r.getQuery(query, &req.ObjectStorageGetReq)
	countQuery := query.Clone()
	count, err := countQuery.Count(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := query.Limit(int(page.Size)).Offset(int((page.Page - 1) * page.Size)).All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.ObjectStorage, 0, len(rows))
	for _, row := range rows {
		result = append(result, &model.ObjectStorage{
			ID:                 row.ID,
			Provider:           enum.ObjectStorageProvider(row.Provider),
			Bucket:             row.Bucket,
			Key:                row.Key,
			MimeType:           row.MimeType,
			Size:               row.Size,
			Hash:               row.Hash,
			UploadBy:           row.UploadBy,
			AuditCallbackReply: row.AuditCallbackReply,
			Blocked:            row.Blocked,
			BlockedReason:      row.BlockedReason,
			BlockedAt:          row.BlockedAt,
			BlockedBy:          row.BlockedBy,
			CreatedAt:          row.CreatedAt,
			UpdatedAt:          row.UpdatedAt,
		})
	}
	return &repo.ObjectStoragePageResp{
		Rows: result,
		Page: &common.PageResp{
			Total: uint32(count),
			Size:  page.Size,
			Page:  page.Page,
		},
	}, nil
}

func (r *ObjectStorageRepo) getQuery(query *gen.ObjectStorageQuery, req *repo.ObjectStorageGetReq) *gen.ObjectStorageQuery {
	if req == nil {
		return query
	}
	if req.ID != nil {
		query = query.Where(objectstorage.IDEQ(*req.ID))
	}
	if len(req.IDs) > 0 {
		query = query.Where(objectstorage.IDIn(req.IDs...))
	}
	if req.Provider != nil {
		query = query.Where(objectstorage.ProviderEQ(objectstorage.Provider(*req.Provider)))
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
	return query
}
