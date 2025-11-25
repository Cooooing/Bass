package biz

import (
	cv1 "common/api/common/v1"
	v1 "common/api/content/v1"
	userv1 "common/api/user/v1"
	"common/pkg/client"
	"common/pkg/constant"
	"common/pkg/util"
	"common/pkg/util/base"
	"common/pkg/util/collections/dict"
	"common/pkg/util/collections/set"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/ent"
	"content/internal/data/ent/gen"
	"context"
	"time"

	"github.com/sony/sonyflake/v2"
)

type ArticleDomain struct {
	*BaseDomain
	articleRepo      repo.ArticleRepo
	postscriptRepo   repo.ArticlePostscriptRepo
	actionRecordRepo repo.ArticleActionRecordRepo
	commentRepo      repo.CommentRepo
	tagRepo          repo.TagRepo
	domainRepo       repo.DomainRepo
	sf               *sonyflake.Sonyflake
}

func NewArticleDomain(base *BaseDomain, articleRepo repo.ArticleRepo, postscriptRepo repo.ArticlePostscriptRepo, actionRecordRepo repo.ArticleActionRecordRepo, commentRepo repo.CommentRepo, tagRepo repo.TagRepo, domainRepo repo.DomainRepo) (*ArticleDomain, error) {
	sf, err := util.NewSonyflake()
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
	article.Status = int32(cv1.ArticleStatus_ArticleDrafts) // 默认均为草稿
	err = ent.WithTx(ctx, d.db, func(tx *gen.Client) error {
		save, err = d.articleRepo.Save(ctx, tx, article, tags)
		if err != nil {
			return err
		}
		// 正常文章，进行发布
		if status == int32(cv1.ArticleStatus_ArticleNormal) {
			err = d.articleRepo.Publish(ctx, tx, save.ID)
			if err != nil {
				return err
			}
		}
		return err
	})
	return save, err
}

func (d *ArticleDomain) AddPostscript(ctx context.Context, articleId int64, content string) error {
	err := ent.WithTx(ctx, d.db, func(tx *gen.Client) error {
		var err error
		err = d.postscriptRepo.AddPostscript(ctx, tx, articleId, content)
		if err != nil {
			return err
		}
		err = d.articleRepo.UpdateHasPostscript(ctx, tx, articleId, true)
		if err != nil {
			return err
		}
		return err
	})
	// Todo 广播添加事件
	return err
}

// --- 更新 ---

func (d *ArticleDomain) UpdateDraft(ctx context.Context, article *model.Article, tags []*model.Tag) (*model.Article, error) {
	var (
		save *model.Article
		err  error
	)
	status := article.Status
	article.Status = int32(cv1.ArticleStatus_ArticleDrafts) // 默认均为草稿
	err = ent.WithTx(ctx, d.db, func(tx *gen.Client) error {
		articleById, err := d.articleRepo.GetById(ctx, tx, article.ID)
		if err != nil {
			return err
		}
		if articleById.Status != int32(cv1.ArticleStatus_ArticleDrafts) {
			return cv1.ErrorBadRequest("only update draft")
		}

		save, err = d.articleRepo.Update(ctx, tx, article, tags)
		if err != nil {
			return err
		}
		// 正常文章，进行发布
		if status == int32(cv1.ArticleStatus_ArticleNormal) {
			err = d.articleRepo.Publish(ctx, tx, save.ID)
			if err != nil {
				return err
			}
		}
		return err
	})
	return save, err
}

func (d *ArticleDomain) Action(ctx context.Context, articleId int64, userId int64, action cv1.ArticleAction, active bool) error {
	err := ent.WithTx(ctx, d.db, func(tx *gen.Client) error {
		var err error
		if active {
			err = d.articleRepo.UpdateStat(ctx, tx, articleId, action, 1)
			if err != nil {
				return err
			}
			_, err = d.actionRecordRepo.Save(ctx, tx, &model.ArticleActionRecord{
				ArticleID: articleId,
				UserID:    userId,
				Type:      int32(action),
			})
			if err != nil {
				return err
			}
		} else {
			err = d.articleRepo.UpdateStat(ctx, tx, articleId, action, -1)
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
	// Todo 广播行为事件
	return err
}

func (d *ArticleDomain) Publish(ctx context.Context, articleId int64) error {
	err := ent.WithTx(ctx, d.db, func(tx *gen.Client) error {
		err := d.articleRepo.Publish(ctx, tx, articleId)
		if err != nil {
			return err
		}
		return err
	})
	return err
}

// --- 查询 ---

func (d *ArticleDomain) GetOne(ctx context.Context, articleId int64) (*v1.GetArticleOneReply, error) {
	var (
		reply                *v1.GetArticleOneReply
		err                  error
		authorUser           *userv1.User
		lastReplyComment     *model.Comment
		lastReplyCommentUser *userv1.User
		lastReplyCommentAt   *time.Time
	)
	query, err := d.articleRepo.GetById(ctx, d.db, articleId)
	if err != nil {
		return nil, err
	}

	lastReplyComment, err = d.commentRepo.GetArticleLastComment(ctx, d.db, query.ID)
	if err != nil {
		return nil, err
	}

	userServiceClient, err := client.GetServiceClient(d.etcd, constant.UserServiceName.String(), userv1.NewUserUserServiceClient)
	if err != nil {
		return nil, err
	}
	userIds := []int64{*query.CreatedBy}
	if lastReplyComment != nil {
		userIds = append(userIds, *lastReplyComment.CreatedBy)
	}
	userAuthorsMap, err := userServiceClient.GetMap(ctx, &userv1.GetMapRequest{Query: &userv1.UserQueryParams{UserIds: userIds}})
	if err != nil {
		return nil, err
	}

	if lastReplyComment != nil {
		lastReplyCommentAt = lastReplyComment.CreatedAt
		lastReplyCommentUser = userAuthorsMap.Users[*lastReplyComment.CreatedBy]
	}

	authorUser = base.If(query.Anonymous, nil, userAuthorsMap.Users[*query.CreatedBy])
	articleReply := query.ConvertToRpc(authorUser, lastReplyCommentUser, lastReplyCommentAt)
	articleReply.ContentRender, _ = query.ParseContent()
	reply = &v1.GetArticleOneReply{
		Article: articleReply,
	}
	return reply, err
}

func (d *ArticleDomain) Page(ctx context.Context, page *cv1.PageRequest, req *repo.ArticleGetReq) (*v1.PageArticleReply, error) {
	var (
		list                 []*model.Article
		pageReply            *cv1.PageReply
		reply                *v1.PageArticleReply
		err                  error
		authorUser           *userv1.User
		lastReplyComment     *model.Comment
		lastReplyCommentUser *userv1.User
		lastReplyCommentAt   *time.Time
		ok                   bool
	)
	list, pageReply, err = d.articleRepo.GetPage(ctx, d.db, page, req)
	if err != nil {
		return nil, err
	}
	articleIds := set.New[int64](0)
	userIds := set.New[int64](0)
	for _, item := range list {
		articleIds.Add(item.ID)
		userIds.Add(*item.CreatedBy)
	}

	lastCommentMap, err := d.commentRepo.GetArticleLastComments(ctx, d.db, articleIds.ToSlice())
	if err != nil {
		return nil, err
	}
	lastCommentMap.Foreach(func(e *dict.Entry[int64, *model.Comment]) bool {
		userIds.Add(*e.Value.CreatedBy)
		return true
	})

	userServiceClient, err := client.GetServiceClient(d.etcd, constant.UserServiceName.String(), userv1.NewUserUserServiceClient)
	if err != nil {
		return nil, err
	}
	userAuthorsMap, err := userServiceClient.GetMap(ctx, &userv1.GetMapRequest{Query: &userv1.UserQueryParams{UserIds: userIds.ToSlice()}})
	if err != nil {
		return nil, err
	}

	articles := make([]*v1.Article, 0, len(list))
	for _, item := range list {
		item.Summary()
		if lastReplyComment, ok = lastCommentMap.Get(item.ID); ok {
			lastReplyCommentAt = lastReplyComment.CreatedAt
			lastReplyCommentUser = userAuthorsMap.Users[*lastReplyComment.CreatedBy]
		}
		authorUser = base.If(item.Anonymous, nil, userAuthorsMap.Users[*item.CreatedBy])
		a := item.ConvertToRpc(authorUser, lastReplyCommentUser, lastReplyCommentAt)
		articles = append(articles, a)
	}
	reply = &v1.PageArticleReply{
		Page:     pageReply,
		Articles: articles,
	}
	return reply, err
}
