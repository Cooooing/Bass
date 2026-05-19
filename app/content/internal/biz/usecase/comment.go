package usecase

import (
	"common/api/gen/common"
	cerrors "common/api/gen/common/errors"
	v1 "common/api/gen/content/v1"
	userv1 "common/api/gen/user/v1"
	"common/pkg/client/rpc"
	"common/pkg/util"
	utilent "common/pkg/util/ent"

	base "content/internal/biz/base"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/gen"
	"content/internal/data/gen/commentactionrecord"
	"content/internal/enum"
	"context"
	"errors"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/samber/lo"
)

type CommentUsecase struct {
	log        *log.Helper
	tx         base.Tx
	userClient *rpc.UserClient
	eventPool  *util.EventPool

	commentRepo             repo.CommentRepo
	commentActionRecordRepo repo.CommentActionRecordRepo
	articleRepo             repo.ArticleRepo
}

func NewCommentUsecase(
	logger log.Logger,
	tx base.Tx,
	userClient *rpc.UserClient,
	eventPool *util.EventPool,
	commentRepo repo.CommentRepo,
	commentActionRecordRepo repo.CommentActionRecordRepo,
	articleRepo repo.ArticleRepo,
) *CommentUsecase {
	return &CommentUsecase{
		log:                     log.NewHelper(logger),
		tx:                      tx,
		userClient:              userClient,
		eventPool:               eventPool,
		commentRepo:             commentRepo,
		commentActionRecordRepo: commentActionRecordRepo,
		articleRepo:             articleRepo,
	}
}

func (d *CommentUsecase) Add(ctx context.Context, comment *model.Comment) (c *model.Comment, err error) {
	//user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	//if !ok {
	//	return nil, cerrors.ErrorUnauthorized("user not login")
	//}
	err = d.tx(ctx, func(ctx context.Context) error {
		cl, ok := utilent.ClientFromCtx[*gen.Client](ctx)
		if !ok {
			return errors.New("no transaction in context")
		}
		// 回复文章
		exist, err := d.articleRepo.GetOne(ctx, cl, &repo.ArticleGetReq{
			ArticleId: new(comment.ArticleID),
			Status:    new(v1.ArticleStatus_ARTICLE_STATUS_NORMAL),
		})
		if err != nil {
			return err
		}
		if !exist.Commentable {
			return cerrors.ErrorBadRequest("article not commentable")
		}

		// 回复评论
		replyComment := &model.Comment{Comment: &gen.Comment{}}
		if comment.ReplyID != nil {
			replyComment, err = d.commentRepo.GetOne(ctx, cl, &repo.CommentGetReq{
				CommentId: comment.ReplyID,
				ArticleId: new(comment.ArticleID),
				Status:    new(v1.CommentStatus_COMMENT_STATUS_NORMAL),
			})
			if err != nil {
				return err
			}

			err = d.commentRepo.UpdateStat(ctx, cl, replyComment.ID, v1.CommentAction_COMMENT_ACTION_REPLY, 1)
			if err != nil {
				return err
			}
		}

		_, err = d.articleRepo.UpdateStat(ctx, cl, exist.ID, v1.ArticleAction_ARTICLE_ACTION_REPLY, 1)
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
		c, err = d.commentRepo.Save(ctx, cl, save)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	err = d.eventPool.Submit(func() {
		//atUserNames := c.ParseContent()
		//err = d.Rabbitmq.Publish(constant.ExchangeContent.String(), constant.RoutingKeyContentArticleAt.String(), &commonModel.Notification{
		//	UUID:       uuid.New().String(),
		//	Type:       new(notifyv1.NotificationType_NOTIFICATION_TYPE_ARTICLE_AT),
		//	SenderId:   user.ID,
		//	SenderName: user.Name,
		//	Channels:   []*notifyv1.NotificationChannel{new(notifyv1.NotificationChannel_NOTIFICATION_CHANNEL_STATION)},
		//	Meta: commonModel.Meta{
		//		AtUsernames: lo.Keys(atUserNames),
		//		Comment: &commonModel.CommentMeta{
		//			CommentId:     c.ID,
		//			ArticleId:     c.ArticleID,
		//			Content:       c.Content,
		//			ReplyId:       c.ReplyID,
		//			CreatedBy:     *c.CreatedBy,
		//			CreatedByName: *c.CreatedByName,
		//		},
		//	},
		//})
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

func (d *CommentUsecase) Page(ctx context.Context, page *common.PageRequest, req *repo.CommentGetReq) (*common.PageReply, []*model.Comment, error) {
	var (
		pageReply *common.PageReply
		reply     []*model.Comment
		err       error
	)
	err = d.tx(ctx, func(ctx context.Context) error {
		c, ok := utilent.ClientFromCtx[*gen.Client](ctx)
		if !ok {
			return errors.New("no transaction in context")
		}
		reply, pageReply, err = d.commentRepo.GetPage(ctx, c, page, req)
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

		userMap, err := d.userClient.User.GetMap(ctx, &userv1.GetMapUser_Request{Query: &userv1.UserQueryParams{UserIds: lo.Keys(userIds)}})
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

func (d *CommentUsecase) UpdateStatus(ctx context.Context, commentId int64, status v1.CommentStatus) error {
	err := d.tx(ctx, func(ctx context.Context) error {
		c, ok := utilent.ClientFromCtx[*gen.Client](ctx)
		if !ok {
			return errors.New("no transaction in context")
		}
		return d.commentRepo.UpdateStatus(ctx, c, commentId, status)
	})
	return err
}

func (d *CommentUsecase) UpdateStat(ctx context.Context, commentId int64, userId int64, action v1.CommentAction, active bool) error {
	var err error
	err = d.tx(ctx, func(ctx context.Context) error {
		c, ok := utilent.ClientFromCtx[*gen.Client](ctx)
		if !ok {
			return errors.New("no transaction in context")
		}
		if active {
			err = d.commentRepo.UpdateStat(ctx, c, commentId, action, 1)
			if err != nil {
				return err
			}
			_, err = d.commentActionRecordRepo.Save(ctx, c, &model.CommentActionRecord{
				CommentID: commentId,
				UserID:    userId,
				Type: func() commentactionrecord.Type {
					v, _ := enum.CommentActionMap.ToEnum(action)
					return commentactionrecord.Type(v)
				}(),
			})
			if err != nil {
				return err
			}
			return nil
		}

		err = d.commentRepo.UpdateStat(ctx, c, commentId, action, -1)
		if err != nil {
			return err
		}
		err = d.commentActionRecordRepo.Delete(ctx, c, commentId, userId, action)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}
