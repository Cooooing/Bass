package repo

import (
	"context"
	"fmt"
	"platform/internal/biz/model"
	"platform/internal/biz/repo"
	"platform/internal/data/gen"
	"platform/internal/data/gen/asset"
	"platform/internal/enum"
	"time"

	utilent "common/pkg/util/ent"
	entsql "entgo.io/ent/dialect/sql"
)

var _ repo.AssetRepo = (*AssetRepo)(nil)

type AssetRepo struct {
	db *gen.Client
}

func NewAssetRepo(db *gen.Client) repo.AssetRepo {
	return &AssetRepo{db: db}
}

func (r *AssetRepo) client(ctx context.Context) *gen.Client {
	if client, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return client
	}
	return r.db
}

func (r *AssetRepo) CreateOrGet(ctx context.Context, row *model.Asset) (*model.Asset, error) {
	if row == nil || row.Hash == nil || *row.Hash == "" {
		return nil, fmt.Errorf("asset hash is required")
	}
	err := r.client(ctx).Asset.Create().
		SetProvider(asset.Provider(row.Provider)).
		SetBucket(row.Bucket).
		SetObjectKey(row.ObjectKey).
		SetHash(*row.Hash).
		SetNillableUploadBy(row.UploadByID).
		SetNillableProviderEtag(row.ProviderETag).
		SetMimeType(row.MimeType).
		SetSize(row.Size).
		SetStatus(asset.Status(row.Status)).
		OnConflict(entsql.ConflictColumns(asset.FieldHash)).
		Ignore().
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return r.GetByHash(ctx, *row.Hash)
}

func (r *AssetRepo) Get(ctx context.Context, id int64) (*model.Asset, error) {
	if id <= 0 {
		return nil, nil
	}
	entity, err := r.client(ctx).Asset.Query().Where(asset.IDEQ(id)).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return assetModel(entity), nil
}

func (r *AssetRepo) GetByHash(ctx context.Context, hash string) (*model.Asset, error) {
	entity, err := r.client(ctx).Asset.Query().Where(asset.HashEQ(hash)).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return assetModel(entity), nil
}

func (r *AssetRepo) Map(ctx context.Context, ids []int64) (map[int64]*model.Asset, error) {
	result := make(map[int64]*model.Asset)
	if len(ids) == 0 {
		return result, nil
	}
	entities, err := r.client(ctx).Asset.Query().Where(asset.IDIn(ids...)).All(ctx)
	if err != nil {
		return nil, err
	}
	for _, entity := range entities {
		result[entity.ID] = assetModel(entity)
	}
	return result, nil
}

func (r *AssetRepo) Block(ctx context.Context, id int64, reason string, operatorID int64) error {
	now := time.Now().UTC()
	_, err := r.client(ctx).Asset.UpdateOneID(id).
		SetStatus(asset.Status(enum.AssetStatusBlocked)).
		SetBlockedReason(reason).
		SetBlockedAt(now).
		SetBlockedBy(operatorID).
		Save(ctx)
	return err
}

func assetModel(entity *gen.Asset) *model.Asset {
	if entity == nil {
		return nil
	}
	return &model.Asset{
		ID:            entity.ID,
		Provider:      enum.AssetProvider(entity.Provider),
		Bucket:        entity.Bucket,
		ObjectKey:     entity.ObjectKey,
		Hash:          entity.Hash,
		UploadByID:    entity.UploadBy,
		ProviderETag:  entity.ProviderEtag,
		MimeType:      entity.MimeType,
		Size:          entity.Size,
		Status:        enum.AssetStatus(entity.Status),
		BlockedReason: entity.BlockedReason,
		BlockedAt:     entity.BlockedAt,
		BlockedBy:     entity.BlockedBy,
		CreatedAt:     entity.CreatedAt,
		UpdatedAt:     entity.UpdatedAt,
	}
}
