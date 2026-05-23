package usecase

import (
	"context"

	"common/api/gen/common"
	cerrors "common/api/gen/common/errors"
	v1 "common/api/gen/user/v1"
	"common/pkg/util"
	base "user/internal/biz/base"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/enum"
)

type RelationUsecase struct {
	tx           base.Tx
	eventPool    *util.EventPool
	relationRepo repo.RelationRepo
	accountRepo  repo.AccountRepo
}

func NewRelationUsecase(
	tx base.Tx,
	eventPool *util.EventPool,
	relationRepo repo.RelationRepo,
	accountRepo repo.AccountRepo,
) (*RelationUsecase, error) {
	return &RelationUsecase{
		tx:           tx,
		eventPool:    eventPool,
		relationRepo: relationRepo,
		accountRepo:  accountRepo,
	}, nil
}

// UpdateRelation 创建或移除一条有向关系，并在同一事务中维护两侧冗余计数。
func (d *RelationUsecase) UpdateRelation(ctx context.Context, relationType v1.RelationType, isAdd bool, actorID int64, targetID int64) error {
	err := d.tx(ctx, func(ctx context.Context) error {
		var delta int32

		dbRelationType, ok := enum.RelationTypeMap.ToEnum(relationType)
		if !ok {
			return cerrors.ErrorBadRequest("unknown relation type")
		}

		exist, err := d.relationRepo.Exists(ctx, &repo.RelationGetReq{
			ActorId:  &actorID,
			TargetId: &targetID,
			Type:     &relationType,
		})
		if err != nil {
			return err
		}

		if isAdd {
			if exist {
				return cerrors.ErrorBadRequest("relation already exists")
			}
			delta = 1
			_, err = d.relationRepo.Create(ctx, &model.Relation{
				ActorID:  actorID,
				TargetID: targetID,
				Type:     dbRelationType,
			})
			if err != nil {
				return err
			}
		} else {
			if !exist {
				return cerrors.ErrorBadRequest("relation not exists")
			}
			delta = -1
			deleted, err := d.relationRepo.Delete(ctx, &repo.RelationDeleteReq{
				ActorID:  actorID,
				TargetID: targetID,
				Type:     dbRelationType,
			})
			if err != nil {
				return err
			}
			if deleted == 0 {
				return cerrors.ErrorBadRequest("relation not exists")
			}
		}

		switch relationType {
		case v1.RelationType_RELATION_TYPE_FOLLOW:
			if _, err = d.accountRepo.AddStat(ctx, actorID, enum.AccountStatTypeFollow, delta); err != nil {
				return err
			}
			_, err = d.accountRepo.AddStat(ctx, targetID, enum.AccountStatTypeFollower, delta)
			return err
		case v1.RelationType_RELATION_TYPE_BLOCK:
			if _, err = d.accountRepo.AddStat(ctx, actorID, enum.AccountStatTypeBlock, delta); err != nil {
				return err
			}
			_, err = d.accountRepo.AddStat(ctx, targetID, enum.AccountStatTypeBlocked, delta)
			return err
		default:
			return cerrors.ErrorBadRequest("unknown relation type")
		}
	})
	if err != nil {
		return err
	}

	// 关注通知属于 outbox 职责；第一版 schema 先保留空 hook，等待分发器和事件构建器接入。
	if relationType == v1.RelationType_RELATION_TYPE_FOLLOW {
		return d.eventPool.Submit(func() {})
	}
	return nil
}

func (d *RelationUsecase) Page(ctx context.Context, page *common.PageRequest, req *repo.RelationGetReq) ([]*model.Relation, *common.PageReply, error) {
	return d.relationRepo.Page(ctx, page, req)
}
