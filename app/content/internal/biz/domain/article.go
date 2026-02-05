package domain

import (
	cv1 "common/api/common/v1"
	v1 "common/api/content/v1"
	notifyv1 "common/api/notify/v1"
	userv1 "common/api/user/v1"
	"common/pkg/client"
	"common/pkg/constant"
	"common/pkg/cutil/base"
	"common/pkg/cutil/base/str"
	"common/pkg/cutil/collections/dict"
	"common/pkg/cutil/collections/set"
	commonModel "common/pkg/model"
	"common/pkg/util"
	domainbase "content/internal/biz/base"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/ent"
	"content/internal/data/ent/gen"
	"context"

	"github.com/google/uuid"
	"github.com/sony/sonyflake/v2"
)

type ArticleDomain struct {
	*domainbase.BaseDomain
	articleRepo      repo.ArticleRepo
	postscriptRepo   repo.ArticlePostscriptRepo
	actionRecordRepo repo.ArticleActionRecordRepo
	commentRepo      repo.CommentRepo
	tagRepo          repo.TagRepo
	domainRepo       repo.DomainRepo
	sf               *sonyflake.Sonyflake
}

func NewArticleDomain(base *domainbase.BaseDomain, articleRepo repo.ArticleRepo, postscriptRepo repo.ArticlePostscriptRepo, actionRecordRepo repo.ArticleActionRecordRepo, commentRepo repo.CommentRepo, tagRepo repo.TagRepo, domainRepo repo.DomainRepo) (*ArticleDomain, error) {
	sf, err := str.NewSonyflake()
	if err != nil {
		return nil, err
	}
	return &ArticleDomain{
		BaseDomain:       base,
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

func (d *ArticleDomain) Add(ctx context.Context, article *model.Article, tags []*model.Tag) (*model.Article, error) {
	var (
		save *model.Article
		err  error
	)
	status := article.Status
	article.Status = int32(v1.ArticleStatus_ARTICLE_STATUS_DRAFTS) // 默认均为草稿
	err = ent.WithTx(ctx, d.Db, func(tx *gen.Client) error {
		save, err = d.articleRepo.Save(ctx, tx, article, tags)
		if err != nil {
			return err
		}
		// 正常文章，进行发布
		if status == int32(v1.ArticleStatus_ARTICLE_STATUS_NORMAL) {
			err = d.Publish(ctx, tx, save.ID)
			if err != nil {
				return err
			}
		}
		return err
	})
	return save, err
}

func (d *ArticleDomain) AddPostscript(ctx context.Context, articleId int64, content string) (*model.ArticlePostscript, error) {
	var save *model.ArticlePostscript
	err := ent.WithTx(ctx, d.Db, func(tx *gen.Client) error {
		var err error
		save, err = d.postscriptRepo.Save(ctx, tx, &model.ArticlePostscript{ArticlePostscript: &gen.ArticlePostscript{
			ArticleID: articleId,
			Content:   content,
			Status:    int32(v1.ArticlePostscriptStatus_ARTICLE_POSTSCRIPT_STATUS_NORMAL),
		}})
		if err != nil {
			return err
		}
		err = d.articleRepo.UpdateHasPostscript(ctx, tx, articleId, true)
		if err != nil {
			return err
		}
		return err
	})
	return save, err
}

// --- 更新 ---

func (d *ArticleDomain) UpdateDraft(ctx context.Context, article *model.Article, tags []*model.Tag) (*model.Article, error) {
	var (
		save *model.Article
		err  error
	)
	status := article.Status
	article.Status = int32(v1.ArticleStatus_ARTICLE_STATUS_DRAFTS) // 默认均为草稿
	err = ent.WithTx(ctx, d.Db, func(tx *gen.Client) error {
		exist, err := d.articleRepo.Exist(ctx, tx, &repo.ArticleGetReq{
			ArticleId: base.Ptr(article.ID),
			Status:    base.Ptr(v1.ArticleStatus_ARTICLE_STATUS_DRAFTS),
			CreatedBy: article.CreatedBy,
		})
		if err != nil {
			return err
		}
		if !exist {
			return cv1.ErrorBadRequest("article not exist")
		}

		save, err = d.articleRepo.Update(ctx, tx, article, tags)
		if err != nil {
			return err
		}
		// 正常文章，进行发布
		if status == int32(v1.ArticleStatus_ARTICLE_STATUS_NORMAL) {
			err = d.Publish(ctx, tx, save.ID)
			if err != nil {
				return err
			}
		}
		return err
	})
	return save, err
}

func (d *ArticleDomain) Action(ctx context.Context, articleId int64, userId int64, action v1.ArticleAction, active bool) error {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return cv1.ErrorUnauthorized("user not login")
	}

	var a *model.Article
	err := ent.WithTx(ctx, d.Db, func(tx *gen.Client) error {
		var err error
		if active {
			a, err = d.articleRepo.UpdateStat(ctx, tx, articleId, action, 1)
			if err != nil {
				return err
			}
			_, err = d.actionRecordRepo.Save(ctx, tx, &model.ArticleActionRecord{ArticleActionRecord: &gen.ArticleActionRecord{
				ArticleID: articleId,
				UserID:    userId,
				Type:      int32(action),
			}})
			if err != nil {
				return err
			}
		} else {
			a, err = d.articleRepo.UpdateStat(ctx, tx, articleId, action, -1)
			if err != nil {
				return err
			}
			err = d.actionRecordRepo.Delete(ctx, tx, articleId, userId, action)
			if err != nil {
				return err
			}
		}
		return err
	})
	if active {
		err = d.EventPool.Submit(func() {
			switch action {
			case v1.ArticleAction_ARTICLE_ACTION_LIKE:
				err = d.Rabbitmq.Publish(constant.ExchangeContent.String(), constant.RoutingKeyContentArticleLike.String(), &commonModel.Notification{
					UUID:       uuid.New().String(),
					Type:       base.Ptr(notifyv1.NotificationType_NOTIFICATION_TYPE_ARTICLE_LIKE),
					SenderId:   user.ID,
					SenderName: user.Name,
					Channels:   []*notifyv1.NotificationChannel{base.Ptr(notifyv1.NotificationChannel_NOTIFICATION_CHANNEL_WEBSITE)},
					Meta: commonModel.Meta{
						Article: &commonModel.ArticleMeta{ArticleId: a.ID, Title: a.Title, CreatedBy: *a.CreatedBy, CreatedByName: *a.CreatedByName},
					},
					Status: notifyv1.NotificationStatus_NOTIFICATION_STATUS_NORMAL,
				})
				if err != nil {
					d.Log.Errorf("publish article like event error: %v", err)
					return
				}
			case v1.ArticleAction_ARTICLE_ACTION_THANK:
				err = d.Rabbitmq.Publish(constant.ExchangeContent.String(), constant.RoutingKeyContentArticleThank.String(), &commonModel.Notification{
					UUID:       uuid.New().String(),
					Type:       base.Ptr(notifyv1.NotificationType_NOTIFICATION_TYPE_ARTICLE_THANK),
					SenderId:   user.ID,
					SenderName: user.Name,
					Channels:   []*notifyv1.NotificationChannel{base.Ptr(notifyv1.NotificationChannel_NOTIFICATION_CHANNEL_WEBSITE)},
					Meta: commonModel.Meta{
						Article: &commonModel.ArticleMeta{ArticleId: a.ID, Title: a.Title, CreatedBy: *a.CreatedBy, CreatedByName: *a.CreatedByName},
					},
					Status: notifyv1.NotificationStatus_NOTIFICATION_STATUS_NORMAL,
				})
				if err != nil {
					d.Log.Errorf("publish article thank event error: %v", err)
					return
				}
			case v1.ArticleAction_ARTICLE_ACTION_COLLECT:
				err = d.Rabbitmq.Publish(constant.ExchangeContent.String(), constant.RoutingKeyContentArticleCollect.String(), &commonModel.Notification{
					UUID:       uuid.New().String(),
					Type:       base.Ptr(notifyv1.NotificationType_NOTIFICATION_TYPE_ARTICLE_COLLECT),
					SenderId:   user.ID,
					SenderName: user.Name,
					Channels:   []*notifyv1.NotificationChannel{base.Ptr(notifyv1.NotificationChannel_NOTIFICATION_CHANNEL_WEBSITE)},
					Meta: commonModel.Meta{
						Article: &commonModel.ArticleMeta{ArticleId: a.ID, Title: a.Title, CreatedBy: *a.CreatedBy, CreatedByName: *a.CreatedByName},
					},
					Status: notifyv1.NotificationStatus_NOTIFICATION_STATUS_NORMAL,
				})
				if err != nil {
					d.Log.Errorf("publish article collect event error: %v", err)
					return
				}
			case v1.ArticleAction_ARTICLE_ACTION_WATCH:
				err = d.Rabbitmq.Publish(constant.ExchangeContent.String(), constant.RoutingKeyContentArticleWatch.String(), &commonModel.Notification{
					UUID:       uuid.New().String(),
					Type:       base.Ptr(notifyv1.NotificationType_NOTIFICATION_TYPE_ARTICLE_WATCH),
					SenderId:   user.ID,
					SenderName: user.Name,
					Channels:   []*notifyv1.NotificationChannel{base.Ptr(notifyv1.NotificationChannel_NOTIFICATION_CHANNEL_WEBSITE)},
					Meta: commonModel.Meta{
						Article: &commonModel.ArticleMeta{ArticleId: a.ID, Title: a.Title, CreatedBy: *a.CreatedBy, CreatedByName: *a.CreatedByName},
					},
					Status: notifyv1.NotificationStatus_NOTIFICATION_STATUS_NORMAL,
				})
				if err != nil {
					d.Log.Errorf("publish article watch event error: %v", err)
					return
				}
			case v1.ArticleAction_ARTICLE_ACTION_REWARD:
				err = d.Rabbitmq.Publish(constant.ExchangeContent.String(), constant.RoutingKeyContentArticleWatch.String(), &commonModel.Notification{
					UUID:       uuid.New().String(),
					Type:       base.Ptr(notifyv1.NotificationType_NOTIFICATION_TYPE_ARTICLE_REWARD),
					SenderId:   user.ID,
					SenderName: user.Name,
					Channels:   []*notifyv1.NotificationChannel{base.Ptr(notifyv1.NotificationChannel_NOTIFICATION_CHANNEL_WEBSITE)},
					Meta: commonModel.Meta{
						Article: &commonModel.ArticleMeta{ArticleId: a.ID, Title: a.Title, CreatedBy: *a.CreatedBy, CreatedByName: *a.CreatedByName},
					},
					Status: notifyv1.NotificationStatus_NOTIFICATION_STATUS_NORMAL,
				})
				if err != nil {
					d.Log.Errorf("publish article watch event error: %v", err)
					return
				}
			default:
				return
			}
		})
	}
	return err
}

func (d *ArticleDomain) Publish(ctx context.Context, tx *gen.Client, articleId int64) error {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return cv1.ErrorUnauthorized("user not login")
	}

	var err error
	var a *model.Article
	err = ent.WithTx(ctx, tx, func(tx *gen.Client) error {
		a, err = d.articleRepo.Publish(ctx, tx, articleId)
		if err != nil {
			return err
		}
		return err
	})
	if err != nil {
		return err
	}
	err = d.EventPool.Submit(func() {

		// 广播发布文章事件
		err = d.Rabbitmq.Publish(constant.ExchangeContent.String(), constant.RoutingKeyContentArticlePublish.String(), &commonModel.Notification{
			UUID:       uuid.New().String(),
			Type:       base.Ptr(notifyv1.NotificationType_NOTIFICATION_TYPE_ARTICLE_PUBLISH),
			SenderId:   user.ID,
			SenderName: user.Name,
			Channels:   []*notifyv1.NotificationChannel{base.Ptr(notifyv1.NotificationChannel_NOTIFICATION_CHANNEL_WEBSITE)},
			Meta: commonModel.Meta{
				Article: &commonModel.ArticleMeta{ArticleId: a.ID, Title: a.Title, CreatedBy: *a.CreatedBy, CreatedByName: *a.CreatedByName},
			},
			Status: notifyv1.NotificationStatus_NOTIFICATION_STATUS_NORMAL,
		})
		if err != nil {
			d.Log.Errorf("publish a publish event error: %v", err)
			return
		}

		// 广播@用户通知
		atUserNames := a.ParseContent()
		if atUserNames.Len() > 0 {
			err = d.Rabbitmq.Publish(constant.ExchangeContent.String(), constant.RoutingKeyContentArticleAt.String(), &commonModel.Notification{
				UUID:       uuid.New().String(),
				Type:       base.Ptr(notifyv1.NotificationType_NOTIFICATION_TYPE_ARTICLE_AT),
				SenderId:   user.ID,
				SenderName: user.Name,
				Channels:   []*notifyv1.NotificationChannel{base.Ptr(notifyv1.NotificationChannel_NOTIFICATION_CHANNEL_WEBSITE)},
				Meta: commonModel.Meta{
					AtUsernames: atUserNames.ToSlice(),
					Article:     &commonModel.ArticleMeta{ArticleId: a.ID, Title: a.Title, CreatedBy: *a.CreatedBy, CreatedByName: *a.CreatedByName},
				},
			})
			if err != nil {
				d.Log.Errorf("publish a at event error: %v", err)
				return
			}
		}
	})
	return err
}

func (d *ArticleDomain) AcceptAnswer(ctx context.Context, articleId int64, commentId int64) error {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return cv1.ErrorUnauthorized("user not login")
	}
	err := ent.WithTx(ctx, d.Db, func(tx *gen.Client) error {
		a, err := d.articleRepo.GetOne(ctx, tx, &repo.ArticleGetReq{ArticleId: base.Ptr(articleId)})
		if err != nil {
			return err
		}
		if *a.CreatedBy != user.ID {
			return cv1.ErrorForbidden("you are not the author of this article")
		}
		if a.AcceptedAnswerID != nil {
			return cv1.ErrorBadRequest("article already accepted answer")
		}
		_, err = d.articleRepo.UpdateAcceptAnswer(ctx, tx, articleId, commentId)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	// Todo 发送通知
	return nil
}

// --- 查询 ---

func (d *ArticleDomain) GetOne(ctx context.Context, articleId int64) (*model.Article, error) {
	var (
		reply *model.Article
		err   error
	)
	reply, err = d.articleRepo.GetOne(ctx, d.Db, &repo.ArticleGetReq{
		ArticleId: base.Ptr(articleId),
	})
	if err != nil {
		return nil, err
	}

	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)

	/*
	 * 正常状态的只能查看公开
	 * 草稿状态的只能查看自己
	 */

	if reply.Status == int32(v1.ArticleStatus_ARTICLE_STATUS_DRAFTS) && !ok && *reply.CreatedBy != user.ID {
		return nil, cv1.ErrorUnauthorized("login required to view drafts")
	}

	lastReplyComment, err := d.commentRepo.GetArticleLastComment(ctx, d.Db, &repo.CommentGetReq{ArticleId: base.Ptr(reply.ID)})
	if err != nil {
		return nil, err
	}

	userServiceClient, err := client.GetConsulServiceClient(d.Consul, constant.UserServiceName.String(), userv1.NewUserUserServiceClient)
	if err != nil {
		return nil, err
	}
	userIds := []int64{*reply.CreatedBy}
	if lastReplyComment != nil {
		userIds = append(userIds, *lastReplyComment.CreatedBy)
	}
	userAuthorsMap, err := userServiceClient.GetMap(ctx, &userv1.GetMapRequest{Query: &userv1.UserQueryParams{UserIds: userIds}})
	if err != nil {
		return nil, err
	}

	if lastReplyComment != nil {
		reply.LastReplyCommentAt = lastReplyComment.CreatedAt
		reply.LastReplyCommentUser = userAuthorsMap.Users[*lastReplyComment.CreatedBy]
	}
	reply.AuthorUser = base.If(reply.Anonymous, nil, userAuthorsMap.Users[*reply.CreatedBy])
	return reply, err
}

func (d *ArticleDomain) Page(ctx context.Context, page *cv1.PageRequest, req *repo.ArticleGetReq) ([]*model.Article, *cv1.PageReply, error) {
	var (
		list      []*model.Article
		pageReply *cv1.PageReply
		err       error
	)
	req.IsSummary = true
	list, pageReply, err = d.articleRepo.GetPage(ctx, d.Db, page, req)
	if err != nil {
		return nil, nil, err
	}
	articleIds := set.New[int64](0)
	userIds := set.New[int64](0)
	for _, item := range list {
		articleIds.Add(item.ID)
		userIds.Add(*item.CreatedBy)
	}

	lastCommentMap := dict.New[int64, *model.Comment](0)
	if articleIds.Len() > 0 {
		lastCommentMap, err = d.commentRepo.GetArticleLastComments(ctx, d.Db, &repo.CommentGetReq{ArticleIds: articleIds.ToSlice()})
		if err != nil {
			return nil, nil, err
		}
		lastCommentMap.Foreach(func(e *dict.Entry[int64, *model.Comment]) bool {
			userIds.Add(*e.Value.CreatedBy)
			return true
		})
	}

	userAuthorsMap := &userv1.GetMapReply{}
	if userIds.Len() > 0 {
		userServiceClient, err := client.GetConsulServiceClient(d.Consul, constant.UserServiceName.String(), userv1.NewUserUserServiceClient)
		if err != nil {
			return nil, nil, err
		}
		userAuthorsMap, err = userServiceClient.GetMap(ctx, &userv1.GetMapRequest{Query: &userv1.UserQueryParams{UserIds: userIds.ToSlice()}})
		if err != nil {
			return nil, nil, err
		}
	}

	for i := range list {
		if lastReplyComment, ok := lastCommentMap.Get(list[i].ID); ok {
			list[i].LastReplyCommentAt = lastReplyComment.CreatedAt
			list[i].LastReplyCommentUser = userAuthorsMap.Users[*lastReplyComment.CreatedBy]
		}
		list[i].AuthorUser = base.If(list[i].Anonymous, nil, userAuthorsMap.Users[*list[i].CreatedBy])
	}
	return list, pageReply, err
}
