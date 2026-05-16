package doamin

import (
	"common/api/gen/common"
	cerrors "common/api/gen/common/errors"
	v1 "common/api/gen/user/v1"
	"context"
	domainbase "user/internal/biz/base"
	"user/internal/biz/model"
	"user/internal/biz/repo"
)

type UserRelationDomain struct {
	*domainbase.BaseDomain
	userRelationRepo repo.UserRelationRepo
	userRepo         repo.UserRepo
}

func NewUserRelationDomain(
	base *domainbase.BaseDomain,
	userRelationRepo repo.UserRelationRepo,
	userRepo repo.UserRepo,
) (*UserRelationDomain, error) {
	return &UserRelationDomain{
		BaseDomain:       base,
		userRelationRepo: userRelationRepo,
		userRepo:         userRepo,
	}, nil
}

func (d *UserRelationDomain) UpdateUserRelation(ctx context.Context, relationType v1.UserRelationType, isAdd bool, actorId int64, targetId int64) error {
	err := d.TxRunner(ctx, func(ctx context.Context) error {
		var err error
		var num int32

		exist, err := d.userRelationRepo.Exist(ctx, &repo.UserRelationGetReq{
			ActorId:  new(actorId),
			TargetId: new(targetId),
			Type:     new(relationType),
		})
		if err != nil {
			return err
		}
		if exist {
			return cerrors.ErrorBadRequest("relation already exists")
		}

		if isAdd {
			num = 1
			_, err = d.userRelationRepo.Save(ctx, &model.UserRelation{
				ActorID:  actorId,
				TargetID: targetId,
				Type:     relationType,
			})
			if err != nil {
				return err
			}
		} else {
			num = -1
			_, err = d.userRelationRepo.Delete(ctx, &model.UserRelation{
				ActorID:  actorId,
				TargetID: targetId,
				Type:     relationType,
			})
		}
		switch relationType {
		case v1.UserRelationType_USER_RELATION_TYPE_FOLLOW:
			_, err = d.userRepo.UpdateStat(ctx, actorId, v1.UserStatType_USER_STAT_TYPE_FOLLOW, num)
			if err != nil {
				return err
			}
			_, err = d.userRepo.UpdateStat(ctx, targetId, v1.UserStatType_USER_STAT_TYPE_FOLLOWER, num)
			if err != nil {
				return err
			}
		case v1.UserRelationType_USER_RELATION_TYPE_BLOCK:
			_, err = d.userRepo.UpdateStat(ctx, actorId, v1.UserStatType_USER_STAT_TYPE_BLOCK, num)
			if err != nil {
				return err
			}
			_, err = d.userRepo.UpdateStat(ctx, targetId, v1.UserStatType_USER_STAT_TYPE_BLOCKED, num)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	//u, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	//if !ok {
	//	return cerrors.ErrorUnauthorized("user not login")
	//}
	// 发送通知，仅关注通知
	if relationType == v1.UserRelationType_USER_RELATION_TYPE_FOLLOW {
		err = d.EventPool.Submit(func() {
			//err := d.Rabbitmq.Publish(constant.ExchangeUser.String(), util.If[string](isAdd, constant.RoutingKeyUserFollow.String(), constant.RoutingKeyUserUnfollow.String()),
			//	&commonModel.Notification{
			//		UUID:       uuid.New().String(),
			//		Type:       new(notifyv1.NotificationType_NOTIFICATION_TYPE_USER_FOLLOW),
			//		SenderId:   u.ID,
			//		SenderName: u.Name,
			//		Channels:   []*notifyv1.NotificationChannel{new(notifyv1.NotificationChannel_NOTIFICATION_CHANNEL_STATION)},
			//		Meta: commonModel.Meta{
			//			User: &commonModel.UserMeta{UserId: targetId, UserName: u.Name},
			//		},
			//	},
			//)
			//if err != nil {
			//	d.Log.Errorf("publish user follow event error: %v", err)
			//}
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *UserRelationDomain) Page(ctx context.Context, page *common.PageRequest, req *repo.UserRelationGetReq) ([]*model.UserRelation, *common.PageReply, error) {
	return d.userRelationRepo.GetPage(ctx, page, req)
}
