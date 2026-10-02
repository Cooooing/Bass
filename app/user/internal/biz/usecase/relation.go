package usecase

import (
	"common/pkg/constant"
	commonmodel "common/pkg/model"
	cerrors "common/proto/gen/common/errors"
	"context"
	"log/slog"

	"common/pkg/apperror"
	commonenums "common/proto/gen/common/enums"
	"user/internal/biz/base"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/enum"
)

type RelationUsecase struct {
	logger        *slog.Logger
	tx            base.Tx
	relationRepo  repo.RelationRepo
	accountRepo   repo.AccountRepo
	outboxRepo    repo.OutboxEventRepo
	outboxUsecase *OutboxUsecase
}

func NewRelationUsecase(
	logger *slog.Logger,
	tx base.Tx,
	relationRepo repo.RelationRepo,
	accountRepo repo.AccountRepo,
	outboxRepo repo.OutboxEventRepo,
	outboxUsecase *OutboxUsecase,
) (*RelationUsecase, error) {
	return &RelationUsecase{
		logger:        logger,
		tx:            tx,
		relationRepo:  relationRepo,
		accountRepo:   accountRepo,
		outboxRepo:    outboxRepo,
		outboxUsecase: outboxUsecase,
	}, nil
}

type FollowRelationReq struct {
	ActorID  int64
	TargetID int64
}

func (d *RelationUsecase) Follow(ctx context.Context, req *FollowRelationReq) error {
	if req.ActorID <= 0 || req.TargetID <= 0 {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_RELATION_INVALID)
	}
	if req.ActorID == req.TargetID {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_SELF_OPERATION_NOT_ALLOWED)
	}
	var outboxEvent *model.OutboxEvent
	err := d.tx(ctx, func(ctx context.Context) error {
		if err := d.ensureNormalAccounts(ctx, req.ActorID, req.TargetID); err != nil {
			return err
		}
		statuses, err := d.MapStatus(ctx, &MapRelationStatusReq{
			ActorID:   req.ActorID,
			TargetIDs: []int64{req.TargetID},
		})
		if err != nil {
			return err
		}
		if status := statuses[req.TargetID]; status != nil && (status.Blocking || status.BlockedBy) {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_ACCOUNT_NOT_FOUND)
		}
		exists, err := d.relationRepo.Exists(ctx, &repo.RelationGetReq{
			ActorId:  &req.ActorID,
			TargetId: &req.TargetID,
			Type:     new(enum.RelationTypeFollow),
		})
		if err != nil {
			return err
		}
		if exists {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_RELATION_ALREADY_EXISTS)
		}
		if _, err = d.relationRepo.Create(ctx, &model.Relation{
			ActorID:  req.ActorID,
			TargetID: req.TargetID,
			Type:     enum.RelationTypeFollow,
		}); err != nil {
			return err
		}
		if _, err = d.accountRepo.AddStat(ctx, &repo.AccountAddStatReq{
			UserID:   req.ActorID,
			StatType: enum.AccountStatTypeFollow,
			Num:      1,
		}); err != nil {
			return err
		}
		if _, err = d.accountRepo.AddStat(ctx, &repo.AccountAddStatReq{
			UserID:   req.TargetID,
			StatType: enum.AccountStatTypeFollower,
			Num:      1,
		}); err != nil {
			return err
		}
		outboxEvent, err = d.outboxRepo.Save(ctx, &repo.OutboxEventSave{
			Event: &commonenums.Event{
				Type:    commonenums.EventType_EVENT_TYPE_USER_FOLLOW,
				Subject: commonenums.EventSubject_EVENT_SUBJECT_USER_FOLLOW,
				Payload: &commonenums.Event_UserFollow{
					UserFollow: &commonenums.UserFollowPayload{
						SenderId:   req.ActorID,
						FollowedId: req.TargetID,
					},
				},
			},
		})
		return err
	})
	if err != nil {
		return err
	}
	if outboxEvent != nil {
		if _, err := d.outboxUsecase.Publish(ctx, &PublishOutboxEventReq{
			ID: outboxEvent.ID,
		}); err != nil {
			d.logger.WarnContext(ctx, "publish outbox event best effort failed", constant.LogFieldEventID, outboxEvent.EventID, constant.LogFieldErr, err)
		}
	}
	return nil
}

type UnfollowRelationReq struct {
	ActorID  int64
	TargetID int64
}

func (d *RelationUsecase) Unfollow(ctx context.Context, req *UnfollowRelationReq) error {
	if req.ActorID <= 0 || req.TargetID <= 0 {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_RELATION_INVALID)
	}
	if req.ActorID == req.TargetID {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_SELF_OPERATION_NOT_ALLOWED)
	}
	var outboxEvent *model.OutboxEvent
	err := d.tx(ctx, func(ctx context.Context) error {
		if _, err := d.accountRepo.Get(ctx, &repo.AccountGetReq{UserID: &req.ActorID}); err != nil {
			return err
		}
		if _, err := d.accountRepo.Get(ctx, &repo.AccountGetReq{UserID: &req.TargetID}); err != nil {
			return err
		}
		exists, err := d.relationRepo.Exists(ctx, &repo.RelationGetReq{
			ActorId:  &req.ActorID,
			TargetId: &req.TargetID,
			Type:     new(enum.RelationTypeFollow),
		})
		if err != nil {
			return err
		}
		if !exists {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_RELATION_NOT_FOUND)
		}
		deleted, err := d.relationRepo.Delete(ctx, &repo.RelationDeleteReq{
			ActorID:  req.ActorID,
			TargetID: req.TargetID,
			Type:     enum.RelationTypeFollow,
		})
		if err != nil {
			return err
		}
		if deleted == 0 {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_RELATION_NOT_FOUND)
		}
		if _, err = d.accountRepo.AddStat(ctx, &repo.AccountAddStatReq{
			UserID:   req.ActorID,
			StatType: enum.AccountStatTypeFollow,
			Num:      -1,
		}); err != nil {
			return err
		}
		if _, err = d.accountRepo.AddStat(ctx, &repo.AccountAddStatReq{
			UserID:   req.TargetID,
			StatType: enum.AccountStatTypeFollower,
			Num:      -1,
		}); err != nil {
			return err
		}
		outboxEvent, err = d.outboxRepo.Save(ctx, &repo.OutboxEventSave{
			Event: &commonenums.Event{
				Type:    commonenums.EventType_EVENT_TYPE_USER_UNFOLLOW,
				Subject: commonenums.EventSubject_EVENT_SUBJECT_USER_UNFOLLOW,
				Payload: &commonenums.Event_UserUnfollow{
					UserUnfollow: &commonenums.UserUnfollowPayload{
						SenderId:   req.ActorID,
						FollowedId: req.TargetID,
					},
				},
			},
		})
		return err
	})
	if err != nil {
		return err
	}
	if outboxEvent != nil {
		if _, err := d.outboxUsecase.Publish(ctx, &PublishOutboxEventReq{
			ID: outboxEvent.ID,
		}); err != nil {
			d.logger.WarnContext(ctx, "publish outbox event best effort failed", constant.LogFieldEventID, outboxEvent.EventID, constant.LogFieldErr, err)
		}
	}
	return nil
}

type BlockRelationReq struct {
	ActorID  int64
	TargetID int64
}

func (d *RelationUsecase) Block(ctx context.Context, req *BlockRelationReq) error {
	if req.ActorID <= 0 || req.TargetID <= 0 {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_RELATION_INVALID)
	}
	if req.ActorID == req.TargetID {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_SELF_OPERATION_NOT_ALLOWED)
	}
	var outboxEvent *model.OutboxEvent
	err := d.tx(ctx, func(ctx context.Context) error {
		if err := d.ensureNormalAccounts(ctx, req.ActorID, req.TargetID); err != nil {
			return err
		}
		exists, err := d.relationRepo.Exists(ctx, &repo.RelationGetReq{
			ActorId:  &req.ActorID,
			TargetId: &req.TargetID,
			Type:     new(enum.RelationTypeBlock),
		})
		if err != nil {
			return err
		}
		if exists {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_RELATION_ALREADY_EXISTS)
		}
		if _, err = d.relationRepo.Create(ctx, &model.Relation{
			ActorID:  req.ActorID,
			TargetID: req.TargetID,
			Type:     enum.RelationTypeBlock,
		}); err != nil {
			return err
		}
		if err = d.removeFollowForBlock(ctx, req.ActorID, req.TargetID); err != nil {
			return err
		}
		outboxEvent, err = d.outboxRepo.Save(ctx, &repo.OutboxEventSave{
			Event: &commonenums.Event{
				Type:    commonenums.EventType_EVENT_TYPE_USER_BLOCK,
				Subject: commonenums.EventSubject_EVENT_SUBJECT_USER_BLOCK,
				Payload: &commonenums.Event_UserBlock{
					UserBlock: &commonenums.UserBlockPayload{
						SenderId:  req.ActorID,
						BlockedId: req.TargetID,
					},
				},
			},
		})
		return err
	})
	if err != nil {
		return err
	}
	if outboxEvent != nil {
		if _, err := d.outboxUsecase.Publish(ctx, &PublishOutboxEventReq{
			ID: outboxEvent.ID,
		}); err != nil {
			d.logger.WarnContext(ctx, "publish outbox event best effort failed", constant.LogFieldEventID, outboxEvent.EventID, constant.LogFieldErr, err)
		}
	}
	return nil
}

func (d *RelationUsecase) ensureNormalAccounts(ctx context.Context, actorID int64, targetID int64) error {
	actor, err := d.accountRepo.Get(ctx, &repo.AccountGetReq{UserID: &actorID})
	if err != nil {
		return err
	}
	target, err := d.accountRepo.Get(ctx, &repo.AccountGetReq{UserID: &targetID})
	if err != nil {
		return err
	}
	if actor == nil || target == nil || actor.Status == nil || target.Status == nil ||
		*actor.Status != enum.AccountStatusNormal || *target.Status != enum.AccountStatusNormal {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_ACCOUNT_NOT_FOUND)
	}
	return nil
}

func (d *RelationUsecase) removeFollowForBlock(ctx context.Context, actorID int64, targetID int64) error {
	follow := enum.RelationTypeFollow
	for _, pair := range [][2]int64{{actorID, targetID}, {targetID, actorID}} {
		deleted, err := d.relationRepo.Delete(ctx, &repo.RelationDeleteReq{
			ActorID:  pair[0],
			TargetID: pair[1],
			Type:     follow,
		})
		if err != nil {
			return err
		}
		if deleted == 0 {
			continue
		}
		if _, err = d.accountRepo.AddStat(ctx, &repo.AccountAddStatReq{UserID: pair[0], StatType: enum.AccountStatTypeFollow, Num: -1}); err != nil {
			return err
		}
		if _, err = d.accountRepo.AddStat(ctx, &repo.AccountAddStatReq{UserID: pair[1], StatType: enum.AccountStatTypeFollower, Num: -1}); err != nil {
			return err
		}
	}
	return nil
}

type UnblockRelationReq struct {
	ActorID  int64
	TargetID int64
}

func (d *RelationUsecase) Unblock(ctx context.Context, req *UnblockRelationReq) error {
	if req.ActorID <= 0 || req.TargetID <= 0 {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_RELATION_INVALID)
	}
	if req.ActorID == req.TargetID {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_SELF_OPERATION_NOT_ALLOWED)
	}
	var outboxEvent *model.OutboxEvent
	err := d.tx(ctx, func(ctx context.Context) error {
		if _, err := d.accountRepo.Get(ctx, &repo.AccountGetReq{UserID: &req.ActorID}); err != nil {
			return err
		}
		if _, err := d.accountRepo.Get(ctx, &repo.AccountGetReq{UserID: &req.TargetID}); err != nil {
			return err
		}
		exists, err := d.relationRepo.Exists(ctx, &repo.RelationGetReq{
			ActorId:  &req.ActorID,
			TargetId: &req.TargetID,
			Type:     new(enum.RelationTypeBlock),
		})
		if err != nil {
			return err
		}
		if !exists {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_RELATION_NOT_FOUND)
		}
		deleted, err := d.relationRepo.Delete(ctx, &repo.RelationDeleteReq{
			ActorID:  req.ActorID,
			TargetID: req.TargetID,
			Type:     enum.RelationTypeBlock,
		})
		if err != nil {
			return err
		}
		if deleted == 0 {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_RELATION_NOT_FOUND)
		}
		outboxEvent, err = d.outboxRepo.Save(ctx, &repo.OutboxEventSave{
			Event: &commonenums.Event{
				Type:    commonenums.EventType_EVENT_TYPE_USER_UNBLOCK,
				Subject: commonenums.EventSubject_EVENT_SUBJECT_USER_UNBLOCK,
				Payload: &commonenums.Event_UserUnblock{
					UserUnblock: &commonenums.UserUnblockPayload{
						SenderId:    req.ActorID,
						UnblockedId: req.TargetID,
					},
				},
			},
		})
		return err
	})
	if err != nil {
		return err
	}
	if outboxEvent != nil {
		if _, err := d.outboxUsecase.Publish(ctx, &PublishOutboxEventReq{
			ID: outboxEvent.ID,
		}); err != nil {
			d.logger.WarnContext(ctx, "publish outbox event best effort failed", constant.LogFieldEventID, outboxEvent.EventID, constant.LogFieldErr, err)
		}
	}
	return nil
}

type ListFollowingRelationsReq struct {
	Page            *commonmodel.PageReq
	ActorID         int64
	ExcludedUserIDs []int64
}

type ListFollowingRelationsResp struct {
	Rows []*model.Relation
	Page *commonmodel.PageResp
}

func (d *RelationUsecase) ListFollowing(ctx context.Context, req *ListFollowingRelationsReq) (*ListFollowingRelationsResp, error) {
	pageResp, err := d.relationRepo.Page(ctx, &repo.RelationPageReq{
		Page: *commonmodel.NormalizePage(req.Page),
		Query: repo.RelationGetReq{
			ActorId:           &req.ActorID,
			Type:              new(enum.RelationTypeFollow),
			ExcludedTargetIDs: req.ExcludedUserIDs,
		},
	})
	if err != nil {
		return nil, err
	}
	return &ListFollowingRelationsResp{
		Rows: pageResp.Rows,
		Page: &pageResp.Page,
	}, nil
}

type ListFollowersRelationsReq struct {
	Page            *commonmodel.PageReq
	TargetID        int64
	ExcludedUserIDs []int64
}

type ListFollowersRelationsResp struct {
	Rows []*model.Relation
	Page *commonmodel.PageResp
}

func (d *RelationUsecase) ListFollowers(ctx context.Context, req *ListFollowersRelationsReq) (*ListFollowersRelationsResp, error) {
	pageResp, err := d.relationRepo.Page(ctx, &repo.RelationPageReq{
		Page: *commonmodel.NormalizePage(req.Page),
		Query: repo.RelationGetReq{
			TargetId:         &req.TargetID,
			Type:             new(enum.RelationTypeFollow),
			ExcludedActorIDs: req.ExcludedUserIDs,
		},
	})
	if err != nil {
		return nil, err
	}
	return &ListFollowersRelationsResp{
		Rows: pageResp.Rows,
		Page: &pageResp.Page,
	}, nil
}

type ListBlockedRelationsReq struct {
	Page    *commonmodel.PageReq
	ActorID int64
}

type ListBlockedRelationsResp struct {
	Rows []*model.Relation
	Page *commonmodel.PageResp
}

func (d *RelationUsecase) ListBlocked(ctx context.Context, req *ListBlockedRelationsReq) (*ListBlockedRelationsResp, error) {
	pageResp, err := d.relationRepo.Page(ctx, &repo.RelationPageReq{
		Page: *commonmodel.NormalizePage(req.Page),
		Query: repo.RelationGetReq{
			ActorId: &req.ActorID,
			Type:    new(enum.RelationTypeBlock),
		},
	})
	if err != nil {
		return nil, err
	}
	return &ListBlockedRelationsResp{
		Rows: pageResp.Rows,
		Page: &pageResp.Page,
	}, nil
}

type MapRelationStatusReq struct {
	ActorID   int64
	TargetIDs []int64
}

func (d *RelationUsecase) MapStatus(ctx context.Context, req *MapRelationStatusReq) (map[int64]*model.RelationStatus, error) {
	statuses := make(map[int64]*model.RelationStatus, len(req.TargetIDs))
	for _, targetID := range req.TargetIDs {
		statuses[targetID] = &model.RelationStatus{
			TargetID: targetID,
		}
	}
	if req.ActorID <= 0 || len(req.TargetIDs) == 0 {
		return statuses, nil
	}
	listResp, err := d.relationRepo.List(ctx, &repo.RelationGetReq{ActorId: &req.ActorID, TargetIDs: req.TargetIDs})
	if err != nil {
		return nil, err
	}
	for _, row := range listResp {
		if status, ok := statuses[row.TargetID]; ok {
			switch row.Type {
			case enum.RelationTypeFollow:
				status.Following = true
			case enum.RelationTypeBlock:
				status.Blocking = true
			}
		}
	}
	listResp, err = d.relationRepo.List(ctx, &repo.RelationGetReq{TargetId: &req.ActorID, ActorIDs: req.TargetIDs})
	if err != nil {
		return nil, err
	}
	for _, row := range listResp {
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

func (d *RelationUsecase) ListBlockedCounterparts(ctx context.Context, userID int64) ([]int64, error) {
	if userID <= 0 {
		return []int64{}, nil
	}
	block := enum.RelationTypeBlock
	rows, err := d.relationRepo.List(ctx, &repo.RelationGetReq{ActorOrTargetId: &userID, Type: &block})
	if err != nil {
		return nil, err
	}
	result := make([]int64, 0, len(rows))
	seen := make(map[int64]struct{}, len(rows))
	for _, row := range rows {
		if row == nil {
			continue
		}
		counterpart := row.TargetID
		if row.TargetID == userID {
			counterpart = row.ActorID
		}
		if counterpart > 0 {
			if _, ok := seen[counterpart]; !ok {
				seen[counterpart] = struct{}{}
				result = append(result, counterpart)
			}
		}
	}
	return result, nil
}
