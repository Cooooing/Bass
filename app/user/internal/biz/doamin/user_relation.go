package doamin

import (
	cv1 "common/api/common/v1"
	notifyv1 "common/api/notify/v1"
	v1 "common/api/user/v1"
	"common/pkg/constant"
	commonModel "common/pkg/model"
	"common/pkg/util"

	"context"
	domainbase "user/internal/biz/base"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/data/ent"
	"user/internal/data/ent/gen"

	"github.com/google/uuid"
)

type UserRelationDomain struct {
	*domainbase.BaseDomain
	userRelationRepo repo.UserRelationRepo
	userRepo         repo.UserRepo
}

func NewUserRelationDomain(base *domainbase.BaseDomain, userRelationRepo repo.UserRelationRepo, userRepo repo.UserRepo) (*UserRelationDomain, error) {
	return &UserRelationDomain{
		BaseDomain:       base,
		userRelationRepo: userRelationRepo,
		userRepo:         userRepo,
	}, nil
}

func (d *UserRelationDomain) UpdateUserRelation(ctx context.Context, relationType v1.UserRelationType, isAdd bool, actorId int64, targetId int64) error {
	err := ent.WithTx(ctx, d.Db, func(tx *gen.Client) error {
		var err error
		var num int32

		exist, err := d.userRelationRepo.Exist(ctx, tx, &repo.UserRelationGetReq{
			ActorId:  util.Ptr(actorId),
			TargetId: util.Ptr(targetId),
			Type:     util.Ptr(relationType),
		})
		if err != nil {
			return err
		}
		if exist {
			return cv1.ErrorBadRequest("relation already exists")
		}

		if isAdd {
			num = 1
			_, err = d.userRelationRepo.Save(ctx, tx, &model.UserRelation{UserRelation: &gen.UserRelation{
				ActorID:  actorId,
				TargetID: targetId,
				Type:     int32(relationType),
			}})
			if err != nil {
				return err
			}
		} else {
			num = -1
			_, err = d.userRelationRepo.Delete(ctx, tx, &model.UserRelation{UserRelation: &gen.UserRelation{
				ActorID:  actorId,
				TargetID: targetId,
				Type:     int32(relationType),
			}})
		}
		switch relationType {
		case v1.UserRelationType_USER_RELATION_TYPE_FOLLOW:
			_, err = d.userRepo.UpdateStat(ctx, tx, actorId, v1.UserStatType_USER_STAT_TYPE_FOLLOW, num)
			if err != nil {
				return err
			}
			_, err = d.userRepo.UpdateStat(ctx, tx, targetId, v1.UserStatType_USER_STAT_TYPE_FOLLOWER, num)
			if err != nil {
				return err
			}
		case v1.UserRelationType_USER_RELATION_TYPE_BLOCK:
			_, err = d.userRepo.UpdateStat(ctx, tx, actorId, v1.UserStatType_USER_STAT_TYPE_BLOCK, num)
			if err != nil {
				return err
			}
			_, err = d.userRepo.UpdateStat(ctx, tx, targetId, v1.UserStatType_USER_STAT_TYPE_BLOCKED, num)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	u, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return cv1.ErrorUnauthorized("user not login")
	}
	// 发送通知，仅关注通知
	if relationType == v1.UserRelationType_USER_RELATION_TYPE_FOLLOW {
		err = d.EventPool.Submit(func() {
			err := d.Rabbitmq.Publish(constant.ExchangeUser.String(), util.If[string](isAdd, constant.RoutingKeyUserFollow.String(), constant.RoutingKeyUserUnfollow.String()),
				&commonModel.Notification{
					UUID:       uuid.New().String(),
					Type:       util.Ptr(notifyv1.NotificationType_NOTIFICATION_TYPE_USER_FOLLOW),
					SenderId:   u.ID,
					SenderName: u.Name,
					Channels:   []*notifyv1.NotificationChannel{util.Ptr(notifyv1.NotificationChannel_NOTIFICATION_CHANNEL_WEBSITE)},
					Meta: commonModel.Meta{
						User: &commonModel.UserMeta{UserId: targetId, UserName: u.Name},
					},
				},
			)
			if err != nil {
				d.Log.Errorf("publish user follow event error: %v", err)
			}
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *UserRelationDomain) Page(ctx context.Context, page *cv1.PageRequest, req *repo.UserRelationGetReq) ([]*model.UserRelation, *cv1.PageReply, error) {
	return d.userRelationRepo.GetPage(ctx, d.Db, page, req)
}
