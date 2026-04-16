package domain

import (
	"common/api/gen/common"
	v1 "common/api/gen/content/v1"
	notifyv1 "common/api/gen/notify/v1"
	userv1 "common/api/gen/user/v1"
	"common/pkg/constant"
	commonModel "common/pkg/model"
	"common/pkg/util"

	domainbase "content/internal/biz/base"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/ent"
	"content/internal/data/ent/gen"
	"context"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type CommentDomain struct {
	*domainbase.BaseDomain
	commentRepo             repo.CommentRepo
	commentActionRecordRepo repo.CommentActionRecordRepo
	articleRepo             repo.ArticleRepo
}

func NewCommentDomain(baseDomain *domainbase.BaseDomain, commentRepo repo.CommentRepo, commentActionRecordRepo repo.CommentActionRecordRepo, articleRepo repo.ArticleRepo) *CommentDomain {
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
		return nil, common.ErrorUnauthorized("user not login")
	}
	err = ent.WithTx(ctx, d.Db, func(tx *gen.Client) error {
		// 回复文章
		exist, err := d.articleRepo.GetOne(ctx, tx, &repo.ArticleGetReq{
			ArticleId: new(comment.ArticleID),
			Status:    new(v1.ArticleStatus_ARTICLE_STATUS_NORMAL),
		})
		if err != nil {
			return err
		}
		if !exist.Commentable {
			return common.ErrorBadRequest("article not commentable")
		}

		// 回复评论
		replyComment := &model.Comment{Comment: &gen.Comment{}}
		if comment.ReplyID != nil {
			replyComment, err = d.commentRepo.GetOne(ctx, tx, &repo.CommentGetReq{
				CommentId: comment.ReplyID,
				ArticleId: new(comment.ArticleID),
				Status:    new(v1.CommentStatus_COMMENT_STATUS_NORMAL),
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
			ParentID: util.If(comment.ReplyID == nil, nil, util.If(replyComment.ParentID == nil, &replyComment.ID, replyComment.ParentID)),
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
	err = d.EventPool.Submit(func() {
		atUserNames := c.ParseContent()
		err = d.Rabbitmq.Publish(constant.ExchangeContent.String(), constant.RoutingKeyContentArticleAt.String(), &commonModel.Notification{
			UUID:       uuid.New().String(),
			Type:       new(notifyv1.NotificationType_NOTIFICATION_TYPE_ARTICLE_AT),
			SenderId:   user.ID,
			SenderName: user.Name,
			Channels:   []*notifyv1.NotificationChannel{new(notifyv1.NotificationChannel_NOTIFICATION_CHANNEL_WEBSITE)},
			Meta: commonModel.Meta{
				AtUsernames: lo.Keys(atUserNames),
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
			d.Log.Errorf("publish a at event error: %v", err)
			return
		}
	})
	if err != nil {
		return nil, err
	}
	return c, err
}

func (d *CommentDomain) Page(ctx context.Context, page *common.PageRequest, req *repo.CommentGetReq) (*common.PageReply, []*model.Comment, error) {
	var (
		pageReply *common.PageReply
		reply     []*model.Comment
		err       error
	)
	err = ent.WithTx(ctx, d.Db, func(tx *gen.Client) error {
		reply, pageReply, err = d.commentRepo.GetPage(ctx, tx, page, req)
		if err != nil {
			return err
		}
		userIds := make(map[int64]struct{})
		for _, item := range reply {
			userIds[*item.CreatedBy] = struct{}{}
			if item.Edges.Reply != nil {
				userIds[*item.Edges.Reply.CreatedBy] = struct{}{}
			}
		}

		userMap, err := d.UserClient.User.GetMap(ctx, &userv1.GetMapUser_Request{Query: &userv1.UserQueryParams{UserIds: lo.Keys(userIds)}})
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
	err := ent.WithTx(ctx, d.Db, func(tx *gen.Client) error {
		return d.commentRepo.UpdateStatus(ctx, tx, commentId, status)
	})
	return err
}

func (d *CommentDomain) UpdateStat(ctx context.Context, commentId int64, userId int64, action v1.CommentAction, active bool) error {
	var err error
	err = ent.WithTx(ctx, d.Db, func(tx *gen.Client) error {
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
