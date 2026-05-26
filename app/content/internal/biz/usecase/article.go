package usecase

import (
	"context"

	"common/api/gen/common"
	commonenums "common/api/gen/common/enums"
	cerrors "common/api/gen/common/errors"
	v1 "common/api/gen/content/v1"
	userv1 "common/api/gen/user/v1"
	commonenum "common/pkg/enum"
	"common/pkg/util"
	base "content/internal/biz/base"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/enum"

	"github.com/samber/lo"
)

type ArticleUsecase struct {
	tx         base.Tx
	userClient repo.UserClient

	articleRepo      repo.ArticleRepo
	postscriptRepo   repo.ArticlePostscriptRepo
	actionRecordRepo repo.ArticleActionRecordRepo
	commentRepo      repo.CommentRepo
	outboxRepo       repo.OutboxEventRepo
}

func NewArticleUsecase(
	tx base.Tx,
	userClient repo.UserClient,
	articleRepo repo.ArticleRepo,
	postscriptRepo repo.ArticlePostscriptRepo,
	actionRecordRepo repo.ArticleActionRecordRepo,
	commentRepo repo.CommentRepo,
	outboxRepo repo.OutboxEventRepo,
) *ArticleUsecase {
	return &ArticleUsecase{
		tx:               tx,
		userClient:       userClient,
		articleRepo:      articleRepo,
		postscriptRepo:   postscriptRepo,
		actionRecordRepo: actionRecordRepo,
		commentRepo:      commentRepo,
		outboxRepo:       outboxRepo,
	}
}

func validateArticleSave(article *model.Article) error {
	if article == nil {
		return cerrors.ErrorBadRequest("article is required")
	}
	switch article.Status {
	case enum.ArticleStatusNormal, enum.ArticleStatusDrafts:
	default:
		return cerrors.ErrorBadRequest("status only be normal or drafts")
	}
	switch article.Type {
	case enum.ArticleTypeNormal, enum.ArticleTypeQA:
	default:
		return cerrors.ErrorBadRequest("invalid article type")
	}
	if article.Type != enum.ArticleTypeQA && article.BountyPoints != nil {
		return cerrors.ErrorBadRequest("bounty points only be set when type is qa")
	}
	return nil
}

func (d *ArticleUsecase) Add(ctx context.Context, article *model.Article, tags []*model.Tag) (*model.Article, error) {
	var (
		save *model.Article
		err  error
	)
	if err := validateArticleSave(article); err != nil {
		return nil, err
	}
	status := article.Status
	article.Status = enum.ArticleStatusDrafts
	err = d.tx(ctx, func(ctx context.Context) error {
		save, err = d.articleRepo.Save(ctx, article, tags)
		if err != nil {
			return err
		}
		if status == enum.ArticleStatusNormal {
			senderID := int64(0)
			if save.CreatedBy != nil {
				senderID = *save.CreatedBy
			}
			return d.Publish(ctx, save.ID, senderID)
		}
		return nil
	})
	return save, err
}

func (d *ArticleUsecase) AddPostscript(ctx context.Context, articleId int64, content string) (*model.ArticlePostscript, error) {
	var save *model.ArticlePostscript
	err := d.tx(ctx, func(ctx context.Context) error {
		exist, err := d.articleRepo.Exist(ctx, &repo.ArticleGetReq{
			ArticleId: &articleId,
			Status:    new(v1.ArticleStatus_ARTICLE_STATUS_NORMAL),
		})
		if err != nil {
			return err
		}
		if !exist {
			return cerrors.ErrorBadRequest("article not exist")
		}
		save, err = d.postscriptRepo.Save(ctx, &model.ArticlePostscript{
			ArticleID: articleId,
			Content:   content,
			Status:    enum.ArticlePostscriptStatusNormal,
		})
		if err != nil {
			return err
		}
		return d.articleRepo.UpdateHasPostscript(ctx, articleId, true)
	})
	return save, err
}

func (d *ArticleUsecase) UpdateDraft(ctx context.Context, article *model.Article, tags []*model.Tag) (*model.Article, error) {
	var (
		save *model.Article
		err  error
	)
	if err := validateArticleSave(article); err != nil {
		return nil, err
	}
	status := article.Status
	article.Status = enum.ArticleStatusDrafts
	err = d.tx(ctx, func(ctx context.Context) error {
		exist, err := d.articleRepo.Exist(ctx, &repo.ArticleGetReq{
			ArticleId: new(article.ID),
			Status:    new(v1.ArticleStatus_ARTICLE_STATUS_DRAFTS),
		})
		if err != nil {
			return err
		}
		if !exist {
			return cerrors.ErrorBadRequest("article not exist")
		}

		save, err = d.articleRepo.Update(ctx, article, tags)
		if err != nil {
			return err
		}
		if status == enum.ArticleStatusNormal {
			senderID := int64(0)
			if save.CreatedBy != nil {
				senderID = *save.CreatedBy
			}
			return d.Publish(ctx, save.ID, senderID)
		}
		return nil
	})
	return save, err
}

func (d *ArticleUsecase) Action(ctx context.Context, articleId int64, userId int64, action v1.ArticleAction, active bool) error {
	dbAction, ok := enum.ArticleActionMap.ToEnum(action)
	if !ok {
		return cerrors.ErrorBadRequest("invalid article action")
	}
	switch action {
	case v1.ArticleAction_ARTICLE_ACTION_LIKE,
		v1.ArticleAction_ARTICLE_ACTION_THANK,
		v1.ArticleAction_ARTICLE_ACTION_COLLECT,
		v1.ArticleAction_ARTICLE_ACTION_WATCH:
	case v1.ArticleAction_ARTICLE_ACTION_REWARD:
		return cerrors.ErrorNotImplemented("article reward not implemented")
	default:
		return cerrors.ErrorBadRequest("unsupported article action")
	}
	senderName, err := accountName(ctx, d.userClient, userId)
	if err != nil {
		return err
	}
	return d.tx(ctx, func(ctx context.Context) error {
		article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{
			ArticleId: &articleId,
			Status:    new(v1.ArticleStatus_ARTICLE_STATUS_NORMAL),
		})
		if err != nil {
			return err
		}
		authorID := int64(0)
		if article.CreatedBy != nil {
			authorID = *article.CreatedBy
		}
		if active {
			existRecord, err := d.actionRecordRepo.Exist(ctx, &repo.ArticleActionRecordReq{
				ArticleId: &articleId,
				UserId:    &userId,
				Type:      &action,
			})
			if err != nil {
				return err
			}
			if existRecord {
				return nil
			}
			_, err = d.actionRecordRepo.Save(ctx, &model.ArticleActionRecord{
				ArticleID: articleId,
				UserID:    userId,
				Type:      dbAction,
			})
			if err != nil {
				return err
			}
			_, err = d.articleRepo.UpdateStat(ctx, articleId, action, 1)
			if err != nil {
				return err
			}
			var eventType commonenum.EventType
			var subject commonenum.EventSubject
			var event *commonenums.Event
			switch action {
			case v1.ArticleAction_ARTICLE_ACTION_LIKE:
				eventType = commonenum.EventTypeContentArticleLike
				subject = commonenum.EventSubjectContentArticleLike
				event = &commonenums.Event{
					Payload: &commonenums.Event_ArticleLiked{
						ArticleLiked: &commonenums.ArticleLikedPayload{
							SenderId:   userId,
							SenderName: senderName,
							ArticleId:  articleId,
							AuthorId:   authorID,
							Title:      article.Title,
						},
					},
				}
			case v1.ArticleAction_ARTICLE_ACTION_THANK:
				eventType = commonenum.EventTypeContentArticleThank
				subject = commonenum.EventSubjectContentArticleThank
				event = &commonenums.Event{
					Payload: &commonenums.Event_ArticleThanked{
						ArticleThanked: &commonenums.ArticleThankedPayload{
							SenderId:   userId,
							SenderName: senderName,
							ArticleId:  articleId,
							AuthorId:   authorID,
							Title:      article.Title,
						},
					},
				}
			case v1.ArticleAction_ARTICLE_ACTION_COLLECT:
				eventType = commonenum.EventTypeContentArticleCollect
				subject = commonenum.EventSubjectContentArticleCollect
				event = &commonenums.Event{
					Payload: &commonenums.Event_ArticleCollected{
						ArticleCollected: &commonenums.ArticleCollectedPayload{
							SenderId:   userId,
							SenderName: senderName,
							ArticleId:  articleId,
							AuthorId:   authorID,
							Title:      article.Title,
						},
					},
				}
			case v1.ArticleAction_ARTICLE_ACTION_WATCH:
				eventType = commonenum.EventTypeContentArticleWatch
				subject = commonenum.EventSubjectContentArticleWatch
				event = &commonenums.Event{
					Payload: &commonenums.Event_ArticleWatched{
						ArticleWatched: &commonenums.ArticleWatchedPayload{
							SenderId:   userId,
							SenderName: senderName,
							ArticleId:  articleId,
							AuthorId:   authorID,
							Title:      article.Title,
						},
					},
				}
			}
			if event == nil {
				return nil
			}
			return d.outboxRepo.Save(ctx, &repo.OutboxEventSave{
				EventType: eventType,
				Subject:   subject,
				Event:     event,
			})
		}
		deleted, err := d.actionRecordRepo.Delete(ctx, articleId, userId, action)
		if err != nil {
			return err
		}
		if deleted == 0 {
			return nil
		}
		_, err = d.articleRepo.UpdateStat(ctx, articleId, action, -1)
		return err
	})
}

func (d *ArticleUsecase) Publish(ctx context.Context, articleId int64, userId int64) error {
	senderName, err := accountName(ctx, d.userClient, userId)
	if err != nil {
		return err
	}
	return d.tx(ctx, func(ctx context.Context) error {
		exist, err := d.articleRepo.Exist(ctx, &repo.ArticleGetReq{
			ArticleId: &articleId,
			Status:    new(v1.ArticleStatus_ARTICLE_STATUS_DRAFTS),
		})
		if err != nil {
			return err
		}
		if !exist {
			return cerrors.ErrorBadRequest("article not exist")
		}
		article, err := d.articleRepo.Publish(ctx, articleId)
		if err != nil {
			return err
		}
		return d.outboxRepo.Save(ctx, &repo.OutboxEventSave{
			EventType: commonenum.EventTypeContentArticlePublish,
			Subject:   commonenum.EventSubjectContentArticlePublish,
			Event: &commonenums.Event{
				Payload: &commonenums.Event_ArticlePublished{
					ArticlePublished: &commonenums.ArticlePublishedPayload{
						SenderId:   userId,
						SenderName: senderName,
						ArticleId:  articleId,
						Title:      article.Title,
					},
				},
			},
		})
	})
}

func (d *ArticleUsecase) UpdateArticle(ctx context.Context, articleId int64, status v1.ArticleStatus, commentable bool, anonymous bool, listable *bool) error {
	dbStatus, ok := enum.ArticleStatusMap.ToEnum(status)
	if !ok {
		return cerrors.ErrorBadRequest("invalid article status")
	}
	switch dbStatus {
	case enum.ArticleStatusNormal, enum.ArticleStatusHidden, enum.ArticleStatusLocked:
	default:
		return cerrors.ErrorBadRequest("invalid article status flow")
	}
	return d.tx(ctx, func(ctx context.Context) error {
		a, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{ArticleId: &articleId})
		if err != nil {
			return err
		}
		switch a.Status {
		case enum.ArticleStatusNormal, enum.ArticleStatusHidden, enum.ArticleStatusLocked:
		default:
			return cerrors.ErrorBadRequest("invalid article status flow")
		}
		return d.articleRepo.UpdateControlFields(ctx, articleId, status, commentable, anonymous, listable)
	})
}

func (d *ArticleUsecase) AcceptAnswer(ctx context.Context, articleId int64, commentId int64) error {
	return d.tx(ctx, func(ctx context.Context) error {
		a, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{
			ArticleId: new(articleId),
			Status:    new(v1.ArticleStatus_ARTICLE_STATUS_NORMAL),
		})
		if err != nil {
			return err
		}
		if a.Type != enum.ArticleTypeQA {
			return cerrors.ErrorBadRequest("only qa article can accept answer")
		}
		if a.AcceptedAnswerID != nil {
			return cerrors.ErrorBadRequest("article already accepted answer")
		}
		exist, err := d.commentRepo.Exist(ctx, &repo.CommentGetReq{
			CommentId: &commentId,
			ArticleId: &articleId,
			Status:    new(v1.CommentStatus_COMMENT_STATUS_NORMAL),
		})
		if err != nil {
			return err
		}
		if !exist {
			return cerrors.ErrorBadRequest("comment not exist")
		}
		_, err = d.articleRepo.UpdateAcceptAnswer(ctx, articleId, commentId)
		return err
	})
}

func (d *ArticleUsecase) Get(ctx context.Context, articleId int64) (*model.Article, error) {
	reply, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{
		ArticleId: new(articleId),
	})
	if err != nil {
		return nil, err
	}

	lastReplyComment, err := d.commentRepo.GetArticleLastComment(ctx, &repo.CommentGetReq{ArticleId: new(reply.ID)})
	if err != nil {
		return nil, err
	}

	userIDs := make([]int64, 0, 2)
	if reply.CreatedBy != nil {
		userIDs = append(userIDs, *reply.CreatedBy)
	}
	if lastReplyComment != nil && lastReplyComment.CreatedBy != nil {
		userIDs = append(userIDs, *lastReplyComment.CreatedBy)
	}
	userAuthorsMap := map[int64]*userv1.AccountBasic{}
	if len(userIDs) > 0 {
		userAuthorsMap, err = d.userClient.BatchGetBasicAccounts(ctx, userIDs)
		if err != nil {
			return nil, err
		}
	}

	if lastReplyComment != nil {
		reply.LastReplyCommentAt = lastReplyComment.CreatedAt
		if lastReplyComment.CreatedBy != nil {
			reply.LastReplyCommentUser = userAuthorsMap[*lastReplyComment.CreatedBy]
		}
	}
	if reply.CreatedBy != nil {
		reply.AuthorUser = util.If(reply.Anonymous, nil, userAuthorsMap[*reply.CreatedBy])
	}
	return reply, nil
}

func (d *ArticleUsecase) Page(ctx context.Context, page *common.PageRequest, req *repo.ArticleGetReq) ([]*model.Article, *common.PageReply, error) {
	req.IsSummary = true
	list, pageReply, err := d.articleRepo.GetPage(ctx, page, req)
	if err != nil {
		return nil, nil, err
	}
	articleIDs := make(map[int64]struct{})
	userIDs := make(map[int64]struct{})
	for _, item := range list {
		articleIDs[item.ID] = struct{}{}
		if item.CreatedBy != nil {
			userIDs[*item.CreatedBy] = struct{}{}
		}
	}

	lastCommentMap := make(map[int64]*model.Comment)
	if len(articleIDs) > 0 {
		lastCommentMap, err = d.commentRepo.GetArticleLastComments(ctx, &repo.CommentGetReq{ArticleIds: lo.Keys(articleIDs)})
		if err != nil {
			return nil, nil, err
		}
		for _, v := range lastCommentMap {
			if v.CreatedBy != nil {
				userIDs[*v.CreatedBy] = struct{}{}
			}
		}
	}

	userAuthorsMap := map[int64]*userv1.AccountBasic{}
	if len(userIDs) > 0 {
		userAuthorsMap, err = d.userClient.BatchGetBasicAccounts(ctx, lo.Keys(userIDs))
		if err != nil {
			return nil, nil, err
		}
	}

	for i := range list {
		if lastReplyComment, ok := lastCommentMap[list[i].ID]; ok {
			list[i].LastReplyCommentAt = lastReplyComment.CreatedAt
			if lastReplyComment.CreatedBy != nil {
				list[i].LastReplyCommentUser = userAuthorsMap[*lastReplyComment.CreatedBy]
			}
		}
		if list[i].CreatedBy != nil {
			list[i].AuthorUser = util.If(list[i].Anonymous, nil, userAuthorsMap[*list[i].CreatedBy])
		}
	}
	return list, pageReply, nil
}

func (d *ArticleUsecase) Delete(ctx context.Context, articleId int64) error {
	return d.tx(ctx, func(ctx context.Context) error {
		exist, err := d.articleRepo.Exist(ctx, &repo.ArticleGetReq{
			ArticleId: &articleId,
			Status:    new(v1.ArticleStatus_ARTICLE_STATUS_DRAFTS),
		})
		if err != nil {
			return err
		}
		if !exist {
			return cerrors.ErrorBadRequest("article not exist")
		}
		return d.articleRepo.Delete(ctx, articleId)
	})
}
