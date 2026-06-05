package usecase

import (
	"context"

	"common/api/gen/common"
	commonenums "common/api/gen/common/enums"
	cerrors "common/api/gen/common/errors"
	v1 "common/api/gen/user/v1"
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

// UpdateRelation 创建或移除一条有向关系，并为关注关系维护两侧冗余计数。
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

		if relationType == v1.RelationType_RELATION_TYPE_FOLLOW || relationType == v1.RelationType_RELATION_TYPE_BLOCK {
			if _, err = d.accountRepo.Get(ctx, &repo.AccountGetReq{UserID: &actorID}); err != nil {
				return err
			}
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
			if _, err = d.accountRepo.AddStat(ctx, targetID, enum.AccountStatTypeFollower, delta); err != nil {
				return err
			}
			if isAdd {
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
		case v1.RelationType_RELATION_TYPE_BLOCK:
			if isAdd {
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
		default:
			return cerrors.ErrorBadRequest("unknown relation type")
		}
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
