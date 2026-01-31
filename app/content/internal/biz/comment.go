package biz

import (
	cv1 "common/api/common/v1"
	v1 "common/api/content/v1"
	notifyv1 "common/api/notify/v1"
	userv1 "common/api/user/v1"
	"common/pkg/client"
	"common/pkg/constant"
	"common/pkg/cutil/base"
	"common/pkg/cutil/collections/set"
	commonModel "common/pkg/model"
	"common/pkg/util"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/ent"
	"content/internal/data/ent/gen"
	"context"

	"github.com/google/uuid"
)

type CommentDomain struct {
	*BaseDomain
	commentRepo             repo.CommentRepo
	commentActionRecordRepo repo.CommentActionRecordRepo
	articleRepo             repo.ArticleRepo
}

func NewCommentDomain(baseDomain *BaseDomain, commentRepo repo.CommentRepo, commentActionRecordRepo repo.CommentActionRecordRepo, articleRepo repo.ArticleRepo) *CommentDomain {
	return &CommentDomain{
		BaseDomain:              baseDomain,
		commentRepo:             commentRepo,
		commentActionRecordRepo: commentActionRecordRepo,
		articleRepo:             articleRepo,
	}
}

func (d *CommentDomain) Add(ctx context.Context, comment *model.Comment) (c *model.Comment, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cv1.ErrorUnauthorized("user not login")
	}
	err = ent.WithTx(ctx, d.db, func(tx *gen.Client) error {
		// 回复文章
		exist, err := d.articleRepo.GetOne(ctx, tx, &repo.ArticleGetReq{
			ArticleId: base.Ptr(comment.ArticleID),
			Status:    base.Ptr(v1.ArticleStatus_ARTICLE_STATUS_NORMAL),
		})
		if err != nil {
			return err
		}
		if !exist.Commentable {
			return cv1.ErrorBadRequest("article not commentable")
		}

		// 回复评论
		replyComment := &model.Comment{Comment: &gen.Comment{}}
		if comment.ReplyID != nil {
			replyComment, err = d.commentRepo.GetOne(ctx, tx, &repo.CommentGetReq{
				CommentId: comment.ReplyID,
				ArticleId: base.Ptr(comment.ArticleID),
				Status:    base.Ptr(v1.CommentStatus_COMMENT_STATUS_NORMAL),
			})
			if err != nil {
				return err
			}

			err = d.commentRepo.UpdateStat(ctx, tx, replyComment.ID, v1.CommentAction_COMMENT_ACTION_REPLY, 1)
			if err != nil {
				return err
			}
		}

		_, err = d.articleRepo.UpdateStat(ctx, tx, exist.ID, v1.ArticleAction_ARTICLE_ACTION_REPLY, 1)
		if err != nil {
			return err
		}

		save := &model.Comment{Comment: &gen.Comment{ArticleID: comment.ArticleID,
			Content:  comment.Content,
			Level:    replyComment.Level + 1,
			ParentID: base.If(comment.ReplyID == nil, nil, base.If(replyComment.ParentID == nil, &replyComment.ID, replyComment.ParentID)),
			ReplyID:  comment.ReplyID,
		}}
		save.FormatContent()
		c, err = d.commentRepo.Save(ctx, tx, save)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	err = d.eventPool.Submit(func() {
		atUserNames := c.ParseContent()
		err = d.rabbitmq.Publish(constant.ExchangeContent.String(), constant.RoutingKeyContentArticleAt.String(), &commonModel.Notification{
			UUID:       uuid.New().String(),
			Type:       base.Ptr(notifyv1.NotificationType_NotificationTypeArticleAt),
			SenderId:   user.ID,
			SenderName: user.Name,
			Channels:   []*notifyv1.NotificationChannel{base.Ptr(notifyv1.NotificationChannel_NotificationChannelWebSite)},
			Meta: commonModel.Meta{
				AtUsernames: atUserNames.ToSlice(),
				Comment: &commonModel.CommentMeta{
					CommentId:     c.ID,
					ArticleId:     c.ArticleID,
					Content:       c.Content,
					ReplyId:       c.ReplyID,
					CreatedBy:     *c.CreatedBy,
					CreatedByName: *c.CreatedByName,
				},
			},
		})
		if err != nil {
			d.log.Errorf("publish a at event error: %v", err)
			return
		}
	})
	if err != nil {
		return nil, err
	}
	return c, err
}

func (d *CommentDomain) Page(ctx context.Context, page *cv1.PageRequest, req *repo.CommentGetReq) (*cv1.PageReply, []*model.Comment, error) {
	var (
		pageReply *cv1.PageReply
		reply     []*model.Comment
		err       error
	)
	err = ent.WithTx(ctx, d.db, func(tx *gen.Client) error {
		reply, pageReply, err = d.commentRepo.GetPage(ctx, tx, page, req)
		if err != nil {
			return err
		}
		userIds := set.New[int64](0)
		for _, item := range reply {
			userIds.Add(*item.CreatedBy)
			if item.Edges.Reply != nil {
				userIds.Add(*item.Edges.Reply.CreatedBy)
			}
		}

		userService, err := client.GetServiceClient(d.etcd, constant.UserServiceName.String(), userv1.NewUserUserServiceClient)
		if err != nil {
			return err
		}
		userMap, err := userService.GetMap(ctx, &userv1.GetMapRequest{Query: &userv1.UserQueryParams{UserIds: userIds.ToSlice()}})
		if err != nil {
			return err
		}
		users := userMap.Users

		for i := range reply {
			reply[i].User = users[*reply[i].CreatedBy]
			if reply[i].Edges.Reply != nil {
				reply[i].ReplyUser = users[*reply[i].Edges.Reply.CreatedBy]
			}
		}
		return nil
	})
	return pageReply, reply, err
}

func (d *CommentDomain) UpdateStatus(ctx context.Context, commentId int64, status v1.CommentStatus) error {
	err := ent.WithTx(ctx, d.db, func(tx *gen.Client) error {
		return d.commentRepo.UpdateStatus(ctx, tx, commentId, status)
	})
	return err
}

func (d *CommentDomain) UpdateStat(ctx context.Context, commentId int64, userId int64, action v1.CommentAction, active bool) error {
	var err error
	err = ent.WithTx(ctx, d.db, func(tx *gen.Client) error {
		if active {
			err = d.commentRepo.UpdateStat(ctx, tx, commentId, action, 1)
			if err != nil {
				return err
			}
			_, err = d.commentActionRecordRepo.Save(ctx, tx, &model.CommentActionRecord{
				CommentID: commentId,
				UserID:    userId,
				Type:      int32(action),
			})
			if err != nil {
				return err
			}
			return nil
		}

		err = d.commentRepo.UpdateStat(ctx, tx, commentId, action, -1)
		if err != nil {
			return err
		}
		err = d.commentActionRecordRepo.Delete(ctx, tx, commentId, userId, action)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}
