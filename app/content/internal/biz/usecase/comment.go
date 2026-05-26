package usecase

import (
	"context"

	"common/api/gen/common"
	commonenums "common/api/gen/common/enums"
	cerrors "common/api/gen/common/errors"
	v1 "common/api/gen/content/v1"
	commonenum "common/pkg/enum"
	"common/pkg/util"
	base "content/internal/biz/base"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/enum"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/samber/lo"
)

type CommentUsecase struct {
	log        *log.Helper
	tx         base.Tx
	userClient repo.UserClient

	commentRepo             repo.CommentRepo
	commentActionRecordRepo repo.CommentActionRecordRepo
	articleRepo             repo.ArticleRepo
	outboxRepo              repo.OutboxEventRepo
}

func NewCommentUsecase(
	logger log.Logger,
	tx base.Tx,
	userClient repo.UserClient,
	commentRepo repo.CommentRepo,
	commentActionRecordRepo repo.CommentActionRecordRepo,
	articleRepo repo.ArticleRepo,
	outboxRepo repo.OutboxEventRepo,
) *CommentUsecase {
	return &CommentUsecase{
		log:                     log.NewHelper(logger),
		tx:                      tx,
		userClient:              userClient,
		commentRepo:             commentRepo,
		commentActionRecordRepo: commentActionRecordRepo,
		articleRepo:             articleRepo,
		outboxRepo:              outboxRepo,
	}
}

func (d *CommentUsecase) Add(ctx context.Context, userId int64, comment *model.Comment) (c *model.Comment, err error) {
	senderName, err := accountName(ctx, d.userClient, userId)
	if err != nil {
		return nil, err
	}
	err = d.tx(ctx, func(ctx context.Context) error {
		article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{
			ArticleId: new(comment.ArticleID),
			Status:    new(v1.ArticleStatus_ARTICLE_STATUS_NORMAL),
		})
		if err != nil {
			return err
		}
		if !article.Commentable {
			return cerrors.ErrorBadRequest("article not commentable")
		}

		replyComment := &model.Comment{}
		if comment.ReplyID != nil {
			replyComment, err = d.commentRepo.Get(ctx, &repo.CommentGetReq{
				CommentId: comment.ReplyID,
				ArticleId: new(comment.ArticleID),
				Status:    new(v1.CommentStatus_COMMENT_STATUS_NORMAL),
			})
			if err != nil {
				return err
			}

			err = d.commentRepo.UpdateStat(ctx, replyComment.ID, v1.CommentAction_COMMENT_ACTION_REPLY, 1)
			if err != nil {
				return err
			}
		}

		_, err = d.articleRepo.UpdateStat(ctx, article.ID, v1.ArticleAction_ARTICLE_ACTION_REPLY, 1)
		if err != nil {
			return err
		}

		save := &model.Comment{
			ArticleID: comment.ArticleID,
			Content:   comment.Content,
			Level:     replyComment.Level + 1,
			ParentID:  util.If(comment.ReplyID == nil, nil, util.If(replyComment.ParentID == nil, &replyComment.ID, replyComment.ParentID)),
			ReplyID:   comment.ReplyID,
		}
		save.FormatContent()
		c, err = d.commentRepo.Save(ctx, save)
		if err != nil {
			return err
		}
		articleAuthorID := int64(0)
		if article.CreatedBy != nil {
			articleAuthorID = *article.CreatedBy
		}
		return d.outboxRepo.Save(ctx, &repo.OutboxEventSave{
			EventType: commonenum.EventTypeContentCommentPublish,
			Subject:   commonenum.EventSubjectContentCommentPublish,
			Event: &commonenums.Event{
				Payload: &commonenums.Event_CommentPublished{
					CommentPublished: &commonenums.CommentPublishedPayload{
						SenderId:   userId,
						SenderName: senderName,
						CommentId:  c.ID,
						ArticleId:  comment.ArticleID,
						AuthorId:   articleAuthorID,
						Content:    c.Content,
						Title:      article.Title,
					},
				},
			},
		})
	})
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (d *CommentUsecase) Page(ctx context.Context, page *common.PageRequest, req *repo.CommentGetReq) (*common.PageReply, []*model.Comment, error) {
	var (
		pageReply *common.PageReply
		reply     []*model.Comment
		err       error
	)
	err = d.tx(ctx, func(ctx context.Context) error {
		reply, pageReply, err = d.commentRepo.GetPage(ctx, page, req)
		if err != nil {
			return err
		}
		userIDs := make(map[int64]struct{})
		for _, item := range reply {
			if item.CreatedBy != nil {
				userIDs[*item.CreatedBy] = struct{}{}
			}
			if item.Reply != nil && item.Reply.CreatedBy != nil {
				userIDs[*item.Reply.CreatedBy] = struct{}{}
			}
		}

		users, err := d.userClient.BatchGetBasicAccounts(ctx, lo.Keys(userIDs))
		if err != nil {
			return err
		}

		for i := range reply {
			if reply[i].CreatedBy != nil {
				reply[i].User = users[*reply[i].CreatedBy]
			}
			if reply[i].Reply != nil && reply[i].Reply.CreatedBy != nil {
				reply[i].ReplyUser = users[*reply[i].Reply.CreatedBy]
			}
		}
		return nil
	})
	return pageReply, reply, err
}

func (d *CommentUsecase) UpdateStatus(ctx context.Context, commentId int64, status v1.CommentStatus) error {
	if _, ok := enum.CommentStatusMap.ToEnum(status); !ok {
		return cerrors.ErrorBadRequest("invalid comment status")
	}
	return d.tx(ctx, func(ctx context.Context) error {
		return d.commentRepo.UpdateStatus(ctx, commentId, status)
	})
}

func (d *CommentUsecase) UpdateStat(ctx context.Context, commentId int64, userId int64, action v1.CommentAction, active bool) error {
	dbAction, ok := enum.CommentActionMap.ToEnum(action)
	if !ok {
		return cerrors.ErrorBadRequest("invalid comment action")
	}
	switch action {
	case v1.CommentAction_COMMENT_ACTION_LIKE, v1.CommentAction_COMMENT_ACTION_THANK:
	default:
		return cerrors.ErrorBadRequest("unsupported comment action")
	}
	senderName, err := accountName(ctx, d.userClient, userId)
	if err != nil {
		return err
	}
	return d.tx(ctx, func(ctx context.Context) error {
		comment, err := d.commentRepo.Get(ctx, &repo.CommentGetReq{
			CommentId: &commentId,
			Status:    new(v1.CommentStatus_COMMENT_STATUS_NORMAL),
		})
		if err != nil {
			return err
		}
		commentAuthorID := int64(0)
		if comment.CreatedBy != nil {
			commentAuthorID = *comment.CreatedBy
		}
		if active {
			existRecord, err := d.commentActionRecordRepo.Exist(ctx, commentId, userId, action)
			if err != nil {
				return err
			}
			if existRecord {
				return nil
			}
			_, err = d.commentActionRecordRepo.Save(ctx, &model.CommentActionRecord{
				CommentID: commentId,
				UserID:    userId,
				Type:      dbAction,
			})
			if err != nil {
				return err
			}
			err = d.commentRepo.UpdateStat(ctx, commentId, action, 1)
			if err != nil {
				return err
			}
			if action != v1.CommentAction_COMMENT_ACTION_LIKE {
				return nil
			}
			return d.outboxRepo.Save(ctx, &repo.OutboxEventSave{
				EventType: commonenum.EventTypeContentCommentLike,
				Subject:   commonenum.EventSubjectContentCommentLike,
				Event: &commonenums.Event{
					Payload: &commonenums.Event_CommentLiked{
						CommentLiked: &commonenums.CommentLikedPayload{
							SenderId:        userId,
							SenderName:      senderName,
							CommentId:       commentId,
							ArticleId:       comment.ArticleID,
							CommentAuthorId: commentAuthorID,
						},
					},
				},
			})
		}

		deleted, err := d.commentActionRecordRepo.Delete(ctx, commentId, userId, action)
		if err != nil {
			return err
		}
		if deleted == 0 {
			return nil
		}
		return d.commentRepo.UpdateStat(ctx, commentId, action, -1)
	})
}
