package repo

import (
	commonClient "common/pkg/client"
	"common/proto/gen/common"
	"context"
	"platform/internal/biz/model"
	"platform/internal/biz/repo"
	"platform/internal/conf"
	"platform/internal/data/gen"
	"platform/internal/data/gen/objectstorage"

	"common/pkg/server"
	utilent "common/pkg/util/ent"
	"github.com/go-kratos/kratos/v2/log"
)

var _ repo.ObjectStorageRepo = (*ObjectStorageRepo)(nil)

type ObjectStorageRepo struct {
	conf         *conf.Bootstrap
	log          *log.Helper
	db           *gen.Client
	consulClient *commonClient.ConsulClient
	redisClient  *commonClient.RedisClient
}

func NewObjectStorageRepo(
	conf *conf.Bootstrap,
	logger log.Logger,
	db *gen.Client,
	consulClient *commonClient.ConsulClient,
	redisClient *commonClient.RedisClient,
) repo.ObjectStorageRepo {
	return &ObjectStorageRepo{
		conf:         conf,
		log:          log.NewHelper(logger),
		db:           db,
		consulClient: consulClient,
		redisClient:  redisClient,
	}
}

func (r *ObjectStorageRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func (r *ObjectStorageRepo) Save(ctx context.Context, u *model.ObjectStorage) (*model.ObjectStorage, error) {
	save, err := r.getClient(ctx).ObjectStorage.Create().
		SetProvider(u.Provider).
		SetBucket(u.Bucket).
		SetKey(u.Key).
		SetMimeType(u.MimeType).
		SetSize(u.Size).
		SetHash(u.Hash).
		SetUploadBy(u.UploadBy).
		SetUploadByName(u.UploadByName).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.ObjectStorage{
		ID:                 save.ID,
		Provider:           save.Provider,
		Bucket:             save.Bucket,
		Key:                save.Key,
		MimeType:           save.MimeType,
		Size:               save.Size,
		Hash:               save.Hash,
		UploadBy:           save.UploadBy,
		UploadByName:       save.UploadByName,
		AuditCallbackReply: save.AuditCallbackReply,
		Blocked:            save.Blocked,
		BlockedReason:      save.BlockedReason,
		BlockedAt:          save.BlockedAt,
		BlockedBy:          save.BlockedBy,
		BlockedByName:      save.BlockedByName,
		CreatedAt:          save.CreatedAt,
		UpdatedAt:          save.UpdatedAt,
	}, nil
}

func (r *ObjectStorageRepo) UpdateAudit(ctx context.Context, u *model.ObjectStorage) error {
	_, err := r.getClient(ctx).ObjectStorage.Update().
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

func (r *ObjectStorageRepo) Delete(ctx context.Context, u *model.ObjectStorage) (int, error) {
	if u == nil {
		return 0, nil
	}
	if u.ID != 0 {
		return 1, r.getClient(ctx).ObjectStorage.DeleteOneID(u.ID).Exec(ctx)
	}
	return r.getClient(ctx).ObjectStorage.Delete().
		Where(objectstorage.KeyEQ(u.Key)).
		Exec(ctx)
}

func (r *ObjectStorageRepo) Exist(ctx context.Context, req *repo.ObjectStorageGetReq) (bool, error) {
	query := r.getClient(ctx).ObjectStorage.Query()
	query = r.getQuery(query, req)
	return query.Exist(ctx)
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
		Provider:           row.Provider,
		Bucket:             row.Bucket,
		Key:                row.Key,
		MimeType:           row.MimeType,
		Size:               row.Size,
		Hash:               row.Hash,
		UploadBy:           row.UploadBy,
		UploadByName:       row.UploadByName,
		AuditCallbackReply: row.AuditCallbackReply,
		Blocked:            row.Blocked,
		BlockedReason:      row.BlockedReason,
		BlockedAt:          row.BlockedAt,
		BlockedBy:          row.BlockedBy,
		BlockedByName:      row.BlockedByName,
		CreatedAt:          row.CreatedAt,
		UpdatedAt:          row.UpdatedAt,
	}, nil
}

func (r *ObjectStorageRepo) List(ctx context.Context, req *repo.ObjectStorageGetReq) ([]*model.ObjectStorage, error) {
	var (
		records []*model.ObjectStorage
		err     error
	)
	query := r.getClient(ctx).ObjectStorage.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	for i := range list {
		records = append(records, &model.ObjectStorage{
			ID:                 list[i].ID,
			Provider:           list[i].Provider,
			Bucket:             list[i].Bucket,
			Key:                list[i].Key,
			MimeType:           list[i].MimeType,
			Size:               list[i].Size,
			Hash:               list[i].Hash,
			UploadBy:           list[i].UploadBy,
			UploadByName:       list[i].UploadByName,
			AuditCallbackReply: list[i].AuditCallbackReply,
			Blocked:            list[i].Blocked,
			BlockedReason:      list[i].BlockedReason,
			BlockedAt:          list[i].BlockedAt,
			BlockedBy:          list[i].BlockedBy,
			BlockedByName:      list[i].BlockedByName,
			CreatedAt:          list[i].CreatedAt,
			UpdatedAt:          list[i].UpdatedAt,
		})
	}
	return records, nil
}

func (r *ObjectStorageRepo) Map(ctx context.Context, req *repo.ObjectStorageGetReq) (map[int64]*model.ObjectStorage, error) {
	list, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*model.ObjectStorage, len(list))
	for _, item := range list {
		result[item.ID] = item
	}
	return result, nil
}

func (r *ObjectStorageRepo) Count(ctx context.Context, req *repo.ObjectStorageGetReq) (int, error) {
	query := r.getClient(ctx).ObjectStorage.Query()
	query = r.getQuery(query, req)
	return query.Count(ctx)
}

func (r *ObjectStorageRepo) Page(ctx context.Context, page *common.PageRequest, req *repo.ObjectStorageGetReq) ([]*model.ObjectStorage, *common.PageReply, error) {
	var (
		items []*model.ObjectStorage
		err   error
	)
	page = server.PageValid(page)
	query := r.getClient(ctx).ObjectStorage.Query()
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
		items = append(items, &model.ObjectStorage{
			ID:                 item.ID,
			Provider:           item.Provider,
			Bucket:             item.Bucket,
			Key:                item.Key,
			MimeType:           item.MimeType,
			Size:               item.Size,
			Hash:               item.Hash,
			UploadBy:           item.UploadBy,
			UploadByName:       item.UploadByName,
			AuditCallbackReply: item.AuditCallbackReply,
			Blocked:            item.Blocked,
			BlockedReason:      item.BlockedReason,
			BlockedAt:          item.BlockedAt,
			BlockedBy:          item.BlockedBy,
			BlockedByName:      item.BlockedByName,
			CreatedAt:          item.CreatedAt,
			UpdatedAt:          item.UpdatedAt,
		})
	}
	return items, &common.PageReply{
		Total: uint32(count),
		Size:  page.Size,
		Page:  page.Page,
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
