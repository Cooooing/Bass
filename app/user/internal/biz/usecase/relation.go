package usecase

import (
	cerrors "common/proto/gen/common/errors"
	"context"

	"common/pkg/apperror"
	"common/proto/gen/common"
	commonenums "common/proto/gen/common/enums"
	v1 "common/proto/gen/user/v1"
	base "user/internal/biz/base"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/enum"
)

type RelationUsecase struct {
	tx           base.Tx
	relationRepo repo.RelationRepo
	accountRepo  repo.AccountRepo
	outboxRepo   repo.OutboxEventRepo
}

func NewRelationUsecase(
	tx base.Tx,
	relationRepo repo.RelationRepo,
	accountRepo repo.AccountRepo,
	outboxRepo repo.OutboxEventRepo,
) (*RelationUsecase, error) {
	return &RelationUsecase{
		tx:           tx,
		relationRepo: relationRepo,
		accountRepo:  accountRepo,
		outboxRepo:   outboxRepo,
	}, nil
}

func (d *RelationUsecase) Follow(ctx context.Context, actorID int64, targetID int64) error {
	err := d.tx(ctx, func(ctx context.Context) error {
		exist, err := d.relationRepo.Exists(ctx, &repo.RelationGetReq{
			ActorId:  &actorID,
			TargetId: &targetID,
			Type:     new(v1.RelationType_RELATION_TYPE_FOLLOW),
		})
		if err != nil {
			return err
		}
		if _, err = d.accountRepo.Get(ctx, &repo.AccountGetReq{UserID: &actorID}); err != nil {
			return err
		}
		if exist {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_RELATION_ALREADY_EXISTS)
		}
		if _, err = d.relationRepo.Create(ctx, &model.Relation{
			ActorID:  actorID,
			TargetID: targetID,
			Type:     enum.RelationTypeFollow,
		}); err != nil {
			return err
		}
		if _, err = d.accountRepo.AddStat(ctx, actorID, enum.AccountStatTypeFollow, 1); err != nil {
			return err
		}
		if _, err = d.accountRepo.AddStat(ctx, targetID, enum.AccountStatTypeFollower, 1); err != nil {
			return err
		}
		return d.outboxRepo.Save(ctx, &repo.OutboxEventSave{
			Event: &commonenums.Event{
				Type:    commonenums.EventType_EVENT_TYPE_USER_FOLLOW,
				Subject: commonenums.EventSubject_EVENT_SUBJECT_USER_FOLLOW,
				Payload: &commonenums.Event_UserFollow{
					UserFollow: &commonenums.UserFollowPayload{
						SenderId:   actorID,
						FollowedId: targetID,
					},
				},
			},
		})
	})
	if err != nil {
		return err
	}
	return nil
}

func (d *RelationUsecase) Unfollow(ctx context.Context, actorID int64, targetID int64) error {
	err := d.tx(ctx, func(ctx context.Context) error {
		exist, err := d.relationRepo.Exists(ctx, &repo.RelationGetReq{
			ActorId:  &actorID,
			TargetId: &targetID,
			Type:     new(v1.RelationType_RELATION_TYPE_FOLLOW),
		})
		if err != nil {
			return err
		}
		if !exist {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_RELATION_NOT_FOUND)
		}
		deleted, err := d.relationRepo.Delete(ctx, &repo.RelationDeleteReq{
			ActorID:  actorID,
			TargetID: targetID,
			Type:     enum.RelationTypeFollow,
		})
		if err != nil {
			return err
		}
		if deleted == 0 {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_RELATION_NOT_FOUND)
		}
		if _, err = d.accountRepo.AddStat(ctx, actorID, enum.AccountStatTypeFollow, -1); err != nil {
			return err
		}
		if _, err = d.accountRepo.AddStat(ctx, targetID, enum.AccountStatTypeFollower, -1); err != nil {
			return err
		}
		return d.outboxRepo.Save(ctx, &repo.OutboxEventSave{
			Event: &commonenums.Event{
				Type:    commonenums.EventType_EVENT_TYPE_USER_UNFOLLOW,
				Subject: commonenums.EventSubject_EVENT_SUBJECT_USER_UNFOLLOW,
				Payload: &commonenums.Event_UserUnfollow{
					UserUnfollow: &commonenums.UserUnfollowPayload{
						SenderId:   actorID,
						FollowedId: targetID,
					},
				},
			},
		})
	})
	if err != nil {
		return err
	}
	return nil
}

func (d *RelationUsecase) Block(ctx context.Context, actorID int64, targetID int64) error {
	err := d.tx(ctx, func(ctx context.Context) error {
		exist, err := d.relationRepo.Exists(ctx, &repo.RelationGetReq{
			ActorId:  &actorID,
			TargetId: &targetID,
			Type:     new(v1.RelationType_RELATION_TYPE_BLOCK),
		})
		if err != nil {
			return err
		}
		if _, err = d.accountRepo.Get(ctx, &repo.AccountGetReq{UserID: &actorID}); err != nil {
			return err
		}
		if exist {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_RELATION_ALREADY_EXISTS)
		}
		if _, err = d.relationRepo.Create(ctx, &model.Relation{
			ActorID:  actorID,
			TargetID: targetID,
			Type:     enum.RelationTypeBlock,
		}); err != nil {
			return err
		}
		return d.outboxRepo.Save(ctx, &repo.OutboxEventSave{
			Event: &commonenums.Event{
				Type:    commonenums.EventType_EVENT_TYPE_USER_BLOCK,
				Subject: commonenums.EventSubject_EVENT_SUBJECT_USER_BLOCK,
				Payload: &commonenums.Event_UserBlock{
					UserBlock: &commonenums.UserBlockPayload{
						SenderId:  actorID,
						BlockedId: targetID,
					},
				},
			},
		})
	})
	if err != nil {
		return err
	}
	return nil
}

func (d *RelationUsecase) Unblock(ctx context.Context, actorID int64, targetID int64) error {
	err := d.tx(ctx, func(ctx context.Context) error {
		exist, err := d.relationRepo.Exists(ctx, &repo.RelationGetReq{
			ActorId:  &actorID,
			TargetId: &targetID,
			Type:     new(v1.RelationType_RELATION_TYPE_BLOCK),
		})
		if err != nil {
			return err
		}
		if !exist {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_RELATION_NOT_FOUND)
		}
		deleted, err := d.relationRepo.Delete(ctx, &repo.RelationDeleteReq{
			ActorID:  actorID,
			TargetID: targetID,
			Type:     enum.RelationTypeBlock,
		})
		if err != nil {
			return err
		}
		if deleted == 0 {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_RELATION_NOT_FOUND)
		}
		return d.outboxRepo.Save(ctx, &repo.OutboxEventSave{
			Event: &commonenums.Event{
				Type:    commonenums.EventType_EVENT_TYPE_USER_UNBLOCK,
				Subject: commonenums.EventSubject_EVENT_SUBJECT_USER_UNBLOCK,
				Payload: &commonenums.Event_UserUnblock{
					UserUnblock: &commonenums.UserUnblockPayload{
						SenderId:    actorID,
						UnblockedId: targetID,
					},
				},
			},
		})
	})
	if err != nil {
		return err
	}
	return nil
}

func (d *RelationUsecase) ListFollowing(ctx context.Context, page *common.PageRequest, actorID int64) ([]*model.Relation, *common.PageReply, error) {
	return d.relationRepo.Page(ctx, page, &repo.RelationGetReq{ActorId: &actorID, Type: new(v1.RelationType_RELATION_TYPE_FOLLOW)})
}

func (d *RelationUsecase) ListFollowers(ctx context.Context, page *common.PageRequest, targetID int64) ([]*model.Relation, *common.PageReply, error) {
	return d.relationRepo.Page(ctx, page, &repo.RelationGetReq{TargetId: &targetID, Type: new(v1.RelationType_RELATION_TYPE_FOLLOW)})
}

func (d *RelationUsecase) ListBlocked(ctx context.Context, page *common.PageRequest, actorID int64) ([]*model.Relation, *common.PageReply, error) {
	return d.relationRepo.Page(ctx, page, &repo.RelationGetReq{ActorId: &actorID, Type: new(v1.RelationType_RELATION_TYPE_BLOCK)})
}

func (d *RelationUsecase) MapStatus(ctx context.Context, actorID int64, targetIDs []int64) (map[int64]*model.RelationStatus, error) {
	statuses := make(map[int64]*model.RelationStatus, len(targetIDs))
	for _, targetID := range targetIDs {
		statuses[targetID] = &model.RelationStatus{TargetID: targetID}
	}
	rows, err := d.relationRepo.List(ctx, &repo.RelationGetReq{ActorId: &actorID})
	if err != nil {
		return nil, err
	}
	for _, row := range rows {
		if status, ok := statuses[row.TargetID]; ok {
			switch row.Type {
			case enum.RelationTypeFollow:
				status.Following = true
			case enum.RelationTypeBlock:
				status.Blocking = true
			}
		}
	}
	rows, err = d.relationRepo.List(ctx, &repo.RelationGetReq{TargetId: &actorID})
	if err != nil {
		return nil, err
	}
	for _, row := range rows {
		if status, ok := statuses[row.ActorID]; ok {
			switch row.Type {
			case enum.RelationTypeFollow:
				status.FollowedBy = true
			case enum.RelationTypeBlock:
				status.BlockedBy = true
			}
		}
	}
	return statuses, nil
}
