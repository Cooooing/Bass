package usecase

import (
	cerrors "common/proto/gen/common/errors"
	"context"

	"common/pkg/apperror"
	commonenums "common/proto/gen/common/enums"
	"user/internal/biz/base"
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

type RelationPageReq struct {
	Page uint32
	Size uint32
}

type RelationPageResp struct {
	Total uint32
	Page  uint32
	Size  uint32
}

type FollowRelationReq struct {
	ActorID  int64
	TargetID int64
}

func (d *RelationUsecase) Follow(ctx context.Context, req *FollowRelationReq) error {
	err := d.tx(ctx, func(ctx context.Context) error {
		exists, err := d.relationRepo.Exists(ctx, &repo.RelationGetReq{
			ActorId:  &req.ActorID,
			TargetId: &req.TargetID,
			Type:     new(enum.RelationTypeFollow),
		})
		if err != nil {
			return err
		}
		if _, err = d.accountRepo.Get(ctx, &repo.AccountGetReq{UserID: &req.ActorID}); err != nil {
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
		if _, err = d.accountRepo.AddStat(ctx, &repo.AccountAddStatReq{UserID: req.ActorID, StatType: enum.AccountStatTypeFollow, Num: 1}); err != nil {
			return err
		}
		if _, err = d.accountRepo.AddStat(ctx, &repo.AccountAddStatReq{UserID: req.TargetID, StatType: enum.AccountStatTypeFollower, Num: 1}); err != nil {
			return err
		}
		err = d.outboxRepo.Save(ctx, &repo.OutboxEventSave{
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
	return nil
}

type UnfollowRelationReq struct {
	ActorID  int64
	TargetID int64
}

func (d *RelationUsecase) Unfollow(ctx context.Context, req *UnfollowRelationReq) error {
	err := d.tx(ctx, func(ctx context.Context) error {
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
		if _, err = d.accountRepo.AddStat(ctx, &repo.AccountAddStatReq{UserID: req.ActorID, StatType: enum.AccountStatTypeFollow, Num: -1}); err != nil {
			return err
		}
		if _, err = d.accountRepo.AddStat(ctx, &repo.AccountAddStatReq{UserID: req.TargetID, StatType: enum.AccountStatTypeFollower, Num: -1}); err != nil {
			return err
		}
		err = d.outboxRepo.Save(ctx, &repo.OutboxEventSave{
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
	return nil
}

type BlockRelationReq struct {
	ActorID  int64
	TargetID int64
}

func (d *RelationUsecase) Block(ctx context.Context, req *BlockRelationReq) error {
	err := d.tx(ctx, func(ctx context.Context) error {
		exists, err := d.relationRepo.Exists(ctx, &repo.RelationGetReq{
			ActorId:  &req.ActorID,
			TargetId: &req.TargetID,
			Type:     new(enum.RelationTypeBlock),
		})
		if err != nil {
			return err
		}
		if _, err = d.accountRepo.Get(ctx, &repo.AccountGetReq{UserID: &req.ActorID}); err != nil {
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
		err = d.outboxRepo.Save(ctx, &repo.OutboxEventSave{
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
	return nil
}

type UnblockRelationReq struct {
	ActorID  int64
	TargetID int64
}

func (d *RelationUsecase) Unblock(ctx context.Context, req *UnblockRelationReq) error {
	err := d.tx(ctx, func(ctx context.Context) error {
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
		err = d.outboxRepo.Save(ctx, &repo.OutboxEventSave{
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
	return nil
}

type ListFollowingRelationsReq struct {
	Page    RelationPageReq
	ActorID int64
}

type ListFollowingRelationsResp struct {
	Rows []*model.Relation
	Page RelationPageResp
}

func (d *RelationUsecase) ListFollowing(ctx context.Context, req *ListFollowingRelationsReq) (*ListFollowingRelationsResp, error) {
	pageResp, err := d.relationRepo.Page(ctx, &repo.RelationPageReq{
		Page:  repo.PageReq{Page: req.Page.Page, Size: req.Page.Size},
		Query: repo.RelationGetReq{ActorId: &req.ActorID, Type: new(enum.RelationTypeFollow)},
	})
	if err != nil {
		return nil, err
	}
	res := RelationPageResp{Total: pageResp.Page.Total, Page: pageResp.Page.Page, Size: pageResp.Page.Size}
	return &ListFollowingRelationsResp{Rows: pageResp.Rows, Page: res}, nil
}

type ListFollowersRelationsReq struct {
	Page     RelationPageReq
	TargetID int64
}

type ListFollowersRelationsResp struct {
	Rows []*model.Relation
	Page RelationPageResp
}

func (d *RelationUsecase) ListFollowers(ctx context.Context, req *ListFollowersRelationsReq) (*ListFollowersRelationsResp, error) {
	pageResp, err := d.relationRepo.Page(ctx, &repo.RelationPageReq{
		Page:  repo.PageReq{Page: req.Page.Page, Size: req.Page.Size},
		Query: repo.RelationGetReq{TargetId: &req.TargetID, Type: new(enum.RelationTypeFollow)},
	})
	if err != nil {
		return nil, err
	}
	res := RelationPageResp{Total: pageResp.Page.Total, Page: pageResp.Page.Page, Size: pageResp.Page.Size}
	return &ListFollowersRelationsResp{Rows: pageResp.Rows, Page: res}, nil
}

type ListBlockedRelationsReq struct {
	Page    RelationPageReq
	ActorID int64
}

type ListBlockedRelationsResp struct {
	Rows []*model.Relation
	Page RelationPageResp
}

func (d *RelationUsecase) ListBlocked(ctx context.Context, req *ListBlockedRelationsReq) (*ListBlockedRelationsResp, error) {
	pageResp, err := d.relationRepo.Page(ctx, &repo.RelationPageReq{
		Page:  repo.PageReq{Page: req.Page.Page, Size: req.Page.Size},
		Query: repo.RelationGetReq{ActorId: &req.ActorID, Type: new(enum.RelationTypeBlock)},
	})
	if err != nil {
		return nil, err
	}
	res := RelationPageResp{Total: pageResp.Page.Total, Page: pageResp.Page.Page, Size: pageResp.Page.Size}
	return &ListBlockedRelationsResp{Rows: pageResp.Rows, Page: res}, nil
}

type MapRelationStatusReq struct {
	ActorID   int64
	TargetIDs []int64
}

func (d *RelationUsecase) MapStatus(ctx context.Context, req *MapRelationStatusReq) (map[int64]*model.RelationStatus, error) {
	statuses := make(map[int64]*model.RelationStatus, len(req.TargetIDs))
	for _, targetID := range req.TargetIDs {
		statuses[targetID] = &model.RelationStatus{TargetID: targetID}
	}
	listResp, err := d.relationRepo.List(ctx, &repo.RelationGetReq{ActorId: &req.ActorID})
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
	listResp, err = d.relationRepo.List(ctx, &repo.RelationGetReq{TargetId: &req.ActorID})
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
