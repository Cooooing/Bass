package usecase

import (
	"common/api/gen/common"
	cerrors "common/api/gen/common/errors"
	v1 "common/api/gen/content/v1"
	userv1 "common/api/gen/user/v1"
	"common/pkg/client/rpc"
	"common/pkg/constant"
	commonModel "common/pkg/model"
	"common/pkg/util"
	utilent "common/pkg/util/ent"

	"common/pkg/util/str"
	base "content/internal/biz/base"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/gen"
	articleent "content/internal/data/gen/article"
	"content/internal/data/gen/articlepostscript"
	"content/internal/enum"
	"context"
	"errors"

	"github.com/samber/lo"
	"github.com/sony/sonyflake/v2"
)

type ArticleUsecase struct {
	tx         base.Tx
	userClient *rpc.UserClient

	articleRepo      repo.ArticleRepo
	postscriptRepo   repo.ArticlePostscriptRepo
	actionRecordRepo repo.ArticleActionRecordRepo
	commentRepo      repo.CommentRepo
	tagRepo          repo.TagRepo
	domainRepo       repo.DomainRepo
	sf               *sonyflake.Sonyflake
}

func NewArticleUsecase(
	tx base.Tx,
	userClient *rpc.UserClient,
	articleRepo repo.ArticleRepo,
	postscriptRepo repo.ArticlePostscriptRepo,
	actionRecordRepo repo.ArticleActionRecordRepo,
	commentRepo repo.CommentRepo,
	tagRepo repo.TagRepo,
	domainRepo repo.DomainRepo,
) (*ArticleUsecase, error) {
	sf, err := str.NewSonyflake()
	if err != nil {
		return nil, err
	}
	return &ArticleUsecase{
		tx:               tx,
		userClient:       userClient,
		articleRepo:      articleRepo,
		postscriptRepo:   postscriptRepo,
		actionRecordRepo: actionRecordRepo,
		commentRepo:      commentRepo,
		tagRepo:          tagRepo,
		domainRepo:       domainRepo,
		sf:               sf,
	}, nil
}

// --- 新增 ---

func (d *ArticleUsecase) Add(ctx context.Context, article *model.Article, tags []*model.Tag) (*model.Article, error) {
	var (
		save *model.Article
		err  error
	)
	status := article.Status
	article.Status = articleent.StatusDrafts // 默认均为草稿
	err = d.tx(ctx, func(ctx context.Context) error {
		c, ok := utilent.ClientFromCtx[*gen.Client](ctx)
		if !ok {
			return errors.New("no transaction in context")
		}
		save, err = d.articleRepo.Save(ctx, c, article, tags)
		if err != nil {
			return err
		}
		// 正常文章，进行发布
		if enum.ArticleStatus(status) == enum.ArticleStatusNormal {
			err = d.Publish(ctx, save.ID, *save.CreatedBy)
			if err != nil {
				return err
			}
		}
		return err
	})
	return save, err
}

func (d *ArticleUsecase) AddPostscript(ctx context.Context, articleId int64, userId int64, content string) (*model.ArticlePostscript, error) {
	var save *model.ArticlePostscript
	err := d.tx(ctx, func(ctx context.Context) error {
		c, ok := utilent.ClientFromCtx[*gen.Client](ctx)
		if !ok {
			return errors.New("no transaction in context")
		}
		exist, err := d.articleRepo.Exist(ctx, c, &repo.ArticleGetReq{
			ArticleId: &articleId,
			Status:    new(v1.ArticleStatus_ARTICLE_STATUS_NORMAL),
			CreatedBy: &userId,
		})
		if err != nil {
			return err
		}
		if !exist {
			return cerrors.ErrorBadRequest("article not exist")
		}
		save, err = d.postscriptRepo.Save(ctx, c, &model.ArticlePostscript{ArticlePostscript: &gen.ArticlePostscript{
			ArticleID: articleId,
			Content:   content,
			Status:    articlepostscript.StatusNormal,
		}})
		if err != nil {
			return err
		}
		err = d.articleRepo.UpdateHasPostscript(ctx, c, articleId, true)
		if err != nil {
			return err
		}
		return err
	})
	return save, err
}

// --- 更新 ---

func (d *ArticleUsecase) UpdateDraft(ctx context.Context, article *model.Article, tags []*model.Tag) (*model.Article, error) {
	var (
		save *model.Article
		err  error
	)
	status := article.Status
	article.Status = articleent.StatusDrafts // 默认均为草稿
	err = d.tx(ctx, func(ctx context.Context) error {
		c, ok := utilent.ClientFromCtx[*gen.Client](ctx)
		if !ok {
			return errors.New("no transaction in context")
		}
		exist, err := d.articleRepo.Exist(ctx, c, &repo.ArticleGetReq{
			ArticleId: new(article.ID),
			Status:    new(v1.ArticleStatus_ARTICLE_STATUS_DRAFTS),
			CreatedBy: article.CreatedBy,
		})
		if err != nil {
			return err
		}
		if !exist {
			return cerrors.ErrorBadRequest("article not exist")
		}

		save, err = d.articleRepo.Update(ctx, c, article, tags)
		if err != nil {
			return err
		}
		// 正常文章，进行发布
		if enum.ArticleStatus(status) == enum.ArticleStatusNormal {
			err = d.Publish(ctx, save.ID, *save.CreatedBy)
			if err != nil {
				return err
			}
		}
		return err
	})
	return save, err
}

func (d *ArticleUsecase) Action(ctx context.Context, articleId int64, userId int64, action v1.ArticleAction, active bool) error {
	var err error
	//user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	//if !ok {
	//	return cerrors.ErrorUnauthorized("user not login")
	//}

	//var a *model.Article
	//err := ent.WithTx(ctx, d.db, func(tx *gen.Client) error {
	//	var err error
	//	if active {
	//		a, err = d.articleRepo.UpdateStat(ctx, tx, articleId, action, 1)
	//		if err != nil {
	//			return err
	//		}
	//		_, err = d.actionRecordRepo.Save(ctx, tx, &model.ArticleActionRecord{ArticleActionRecord: &gen.ArticleActionRecord{
	//			ArticleID: articleId,
	//			UserID:    userId,
	//			Type:      int32(action),
	//		}})
	//		if err != nil {
	//			return err
	//		}
	//	} else {
	//		a, err = d.articleRepo.UpdateStat(ctx, tx, articleId, action, -1)
	//		if err != nil {
	//			return err
	//		}
	//		err = d.actionRecordRepo.Delete(ctx, tx, articleId, userId, action)
	//		if err != nil {
	//			return err
	//		}
	//	}
	//	return err
	//})
	//if active {
	//err = d.eventPool.Submit(func() {
	//	switch action {
	//	case v1.ArticleAction_ARTICLE_ACTION_LIKE:
	//		err = d.Rabbitmq.Publish(constant.ExchangeContent.String(), constant.RoutingKeyContentArticleLike.String(), &commonModel.Notification{
	//			UUID:       uuid.New().String(),
	//			Type:       new(notifyv1.NotificationType_NOTIFICATION_TYPE_ARTICLE_LIKE),
	//			SenderId:   user.ID,
	//			SenderName: user.Name,
	//			Channels:   []*notifyv1.NotificationChannel{new(notifyv1.NotificationChannel_NOTIFICATION_CHANNEL_STATION)},
	//			Meta: commonModel.Meta{
	//				Article: &commonModel.ArticleMeta{ArticleId: a.ID, Title: a.Title, CreatedBy: *a.CreatedBy, CreatedByName: *a.CreatedByName},
	//			},
	//			Status: notifyv1.NotificationStatus_NOTIFICATION_STATUS_NORMAL,
	//		})
	//		if err != nil {
	//			d.log.Errorf("publish article like event error: %v", err)
	//			return
	//		}
	//	case v1.ArticleAction_ARTICLE_ACTION_THANK:
	//		err = d.Rabbitmq.Publish(constant.ExchangeContent.String(), constant.RoutingKeyContentArticleThank.String(), &commonModel.Notification{
	//			UUID:       uuid.New().String(),
	//			Type:       new(notifyv1.NotificationType_NOTIFICATION_TYPE_ARTICLE_THANK),
	//			SenderId:   user.ID,
	//			SenderName: user.Name,
	//			Channels:   []*notifyv1.NotificationChannel{new(notifyv1.NotificationChannel_NOTIFICATION_CHANNEL_STATION)},
	//			Meta: commonModel.Meta{
	//				Article: &commonModel.ArticleMeta{ArticleId: a.ID, Title: a.Title, CreatedBy: *a.CreatedBy, CreatedByName: *a.CreatedByName},
	//			},
	//			Status: notifyv1.NotificationStatus_NOTIFICATION_STATUS_NORMAL,
	//		})
	//		if err != nil {
	//			d.log.Errorf("publish article thank event error: %v", err)
	//			return
	//		}
	//	case v1.ArticleAction_ARTICLE_ACTION_COLLECT:
	//		err = d.Rabbitmq.Publish(constant.ExchangeContent.String(), constant.RoutingKeyContentArticleCollect.String(), &commonModel.Notification{
	//			UUID:       uuid.New().String(),
	//			Type:       new(notifyv1.NotificationType_NOTIFICATION_TYPE_ARTICLE_COLLECT),
	//			SenderId:   user.ID,
	//			SenderName: user.Name,
	//			Channels:   []*notifyv1.NotificationChannel{new(notifyv1.NotificationChannel_NOTIFICATION_CHANNEL_STATION)},
	//			Meta: commonModel.Meta{
	//				Article: &commonModel.ArticleMeta{ArticleId: a.ID, Title: a.Title, CreatedBy: *a.CreatedBy, CreatedByName: *a.CreatedByName},
	//			},
	//			Status: notifyv1.NotificationStatus_NOTIFICATION_STATUS_NORMAL,
	//		})
	//		if err != nil {
	//			d.log.Errorf("publish article collect event error: %v", err)
	//			return
	//		}
	//	case v1.ArticleAction_ARTICLE_ACTION_WATCH:
	//		err = d.Rabbitmq.Publish(constant.ExchangeContent.String(), constant.RoutingKeyContentArticleWatch.String(), &commonModel.Notification{
	//			UUID:       uuid.New().String(),
	//			Type:       new(notifyv1.NotificationType_NOTIFICATION_TYPE_ARTICLE_WATCH),
	//			SenderId:   user.ID,
	//			SenderName: user.Name,
	//			Channels:   []*notifyv1.NotificationChannel{new(notifyv1.NotificationChannel_NOTIFICATION_CHANNEL_STATION)},
	//			Meta: commonModel.Meta{
	//				Article: &commonModel.ArticleMeta{ArticleId: a.ID, Title: a.Title, CreatedBy: *a.CreatedBy, CreatedByName: *a.CreatedByName},
	//			},
	//			Status: notifyv1.NotificationStatus_NOTIFICATION_STATUS_NORMAL,
	//		})
	//		if err != nil {
	//			d.log.Errorf("publish article watch event error: %v", err)
	//			return
	//		}
	//	case v1.ArticleAction_ARTICLE_ACTION_REWARD:
	//		err = d.Rabbitmq.Publish(constant.ExchangeContent.String(), constant.RoutingKeyContentArticleWatch.String(), &commonModel.Notification{
	//			UUID:       uuid.New().String(),
	//			Type:       new(notifyv1.NotificationType_NOTIFICATION_TYPE_ARTICLE_REWARD),
	//			SenderId:   user.ID,
	//			SenderName: user.Name,
	//			Channels:   []*notifyv1.NotificationChannel{new(notifyv1.NotificationChannel_NOTIFICATION_CHANNEL_STATION)},
	//			Meta: commonModel.Meta{
	//				Article: &commonModel.ArticleMeta{ArticleId: a.ID, Title: a.Title, CreatedBy: *a.CreatedBy, CreatedByName: *a.CreatedByName},
	//			},
	//			Status: notifyv1.NotificationStatus_NOTIFICATION_STATUS_NORMAL,
	//		})
	//		if err != nil {
	//			d.log.Errorf("publish article watch event error: %v", err)
	//			return
	//		}
	//	default:
	//		return
	//	}
	//})
	//}
	return err
}

func (d *ArticleUsecase) Publish(ctx context.Context, articleId int64, userId int64) error {
	return d.tx(ctx, func(ctx context.Context) error {
		c, ok := utilent.ClientFromCtx[*gen.Client](ctx)
		if !ok {
			return errors.New("no transaction in context")
		}
		exist, err := d.articleRepo.Exist(ctx, c, &repo.ArticleGetReq{
			ArticleId: &articleId,
			Status:    new(v1.ArticleStatus_ARTICLE_STATUS_DRAFTS),
			CreatedBy: &userId,
		})
		if err != nil {
			return err
		}
		if !exist {
			return cerrors.ErrorBadRequest("article not exist")
		}
		// TODO: 实现发布逻辑
		return nil
	})
}

func (d *ArticleUsecase) AcceptAnswer(ctx context.Context, articleId int64, commentId int64) error {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return cerrors.ErrorUnauthorized("user not login")
	}
	err := d.tx(ctx, func(ctx context.Context) error {
		c, ok := utilent.ClientFromCtx[*gen.Client](ctx)
		if !ok {
			return errors.New("no transaction in context")
		}
		a, err := d.articleRepo.GetOne(ctx, c, &repo.ArticleGetReq{ArticleId: new(articleId)})
		if err != nil {
			return err
		}
		if *a.CreatedBy != user.ID {
			return cerrors.ErrorForbidden("you are not the author of this article")
		}
		if a.AcceptedAnswerID != nil {
			return cerrors.ErrorBadRequest("article already accepted answer")
		}
		_, err = d.articleRepo.UpdateAcceptAnswer(ctx, c, articleId, commentId)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	// TODO: 发送通知。
	return nil
}

// --- 查询 ---

func (d *ArticleUsecase) GetOne(ctx context.Context, articleId int64) (*model.Article, error) {
	var (
		reply *model.Article
		err   error
	)
	c, ok := utilent.ClientFromCtx[*gen.Client](ctx)
	if !ok {
		return nil, errors.New("no client in context")
	}
	reply, err = d.articleRepo.GetOne(ctx, c, &repo.ArticleGetReq{
		ArticleId: new(articleId),
	})
	if err != nil {
		return nil, err
	}

	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)

	/*
	 * 正常状态的只能查看公开
	 * 草稿状态的只能查看自己
	 */

	if enum.ArticleStatus(reply.Status) == enum.ArticleStatusDrafts && !ok && *reply.CreatedBy != user.ID {
		return nil, cerrors.ErrorUnauthorized("login required to view drafts")
	}

	lastReplyComment, err := d.commentRepo.GetArticleLastComment(ctx, c, &repo.CommentGetReq{ArticleId: new(reply.ID)})
	if err != nil {
		return nil, err
	}

	userIds := []int64{*reply.CreatedBy}
	if lastReplyComment != nil {
		userIds = append(userIds, *lastReplyComment.CreatedBy)
	}
	userAuthorsMap, err := d.userClient.Account.BatchGetBasic(ctx, &userv1.BatchGetBasicAccount_Request{UserIds: userIds})
	if err != nil {
		return nil, err
	}

	if lastReplyComment != nil {
		reply.LastReplyCommentAt = lastReplyComment.CreatedAt
		reply.LastReplyCommentUser = userAuthorsMap.Accounts[*lastReplyComment.CreatedBy]
	}
	reply.AuthorUser = util.If(reply.Anonymous, nil, userAuthorsMap.Accounts[*reply.CreatedBy])
	return reply, err
}

func (d *ArticleUsecase) Page(ctx context.Context, page *common.PageRequest, req *repo.ArticleGetReq) ([]*model.Article, *common.PageReply, error) {
	var (
		list      []*model.Article
		pageReply *common.PageReply
		err       error
	)
	c, ok := utilent.ClientFromCtx[*gen.Client](ctx)
	if !ok {
		return nil, nil, errors.New("no client in context")
	}
	req.IsSummary = true
	list, pageReply, err = d.articleRepo.GetPage(ctx, c, page, req)
	if err != nil {
		return nil, nil, err
	}
	articleIds := make(map[int64]struct{})
	userIds := make(map[int64]struct{})
	for _, item := range list {
		articleIds[item.ID] = struct{}{}
		userIds[*item.CreatedBy] = struct{}{}
	}

	lastCommentMap := make(map[int64]*model.Comment)
	if len(articleIds) > 0 {
		lastCommentMap, err = d.commentRepo.GetArticleLastComments(ctx, c, &repo.CommentGetReq{ArticleIds: lo.Keys(articleIds)})
		if err != nil {
			return nil, nil, err
		}
		for _, v := range lastCommentMap {
			userIds[*v.CreatedBy] = struct{}{}
		}
	}

	userAuthorsMap := &userv1.BatchGetBasicAccount_Reply{Accounts: map[int64]*userv1.AccountBasic{}}
	if len(userIds) > 0 {
		userAuthorsMap, err = d.userClient.Account.BatchGetBasic(ctx, &userv1.BatchGetBasicAccount_Request{UserIds: lo.Keys(userIds)})
		if err != nil {
			return nil, nil, err
		}
	}

	for i := range list {
		if lastReplyComment, ok := lastCommentMap[list[i].ID]; ok {
			list[i].LastReplyCommentAt = lastReplyComment.CreatedAt
			list[i].LastReplyCommentUser = userAuthorsMap.Accounts[*lastReplyComment.CreatedBy]
		}
		list[i].AuthorUser = util.If(list[i].Anonymous, nil, userAuthorsMap.Accounts[*list[i].CreatedBy])
	}
	return list, pageReply, err
}

func (d *ArticleUsecase) Delete(ctx context.Context, articleId int64, userId int64) error {
	return d.tx(ctx, func(ctx context.Context) error {
		c, ok := utilent.ClientFromCtx[*gen.Client](ctx)
		if !ok {
			return errors.New("no transaction in context")
		}
		exist, err := d.articleRepo.Exist(ctx, c, &repo.ArticleGetReq{
			ArticleId: &articleId,
			Status:    new(v1.ArticleStatus_ARTICLE_STATUS_DRAFTS),
			CreatedBy: &userId,
		})
		if err != nil {
			return err
		}
		if !exist {
			return cerrors.ErrorBadRequest("article not exist")
		}
		return d.articleRepo.Delete(ctx, c, articleId)
	})
}
