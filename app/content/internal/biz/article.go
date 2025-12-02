package biz

import (
	cv1 "common/api/common/v1"
	v1 "common/api/content/v1"
	userv1 "common/api/user/v1"
	"common/pkg/client"
	"common/pkg/constant"
	commonModel "common/pkg/model"
	"common/pkg/util"
	"common/pkg/util/base"
	"common/pkg/util/collections/dict"
	"common/pkg/util/collections/set"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/ent"
	"content/internal/data/ent/gen"
	"context"

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
	article.Status = int32(v1.ArticleStatus_ArticleDrafts) // 默认均为草稿
	err = ent.WithTx(ctx, d.db, func(tx *gen.Client) error {
		save, err = d.articleRepo.Save(ctx, tx, article, tags)
		if err != nil {
			return err
		}
		// 正常文章，进行发布
		if status == int32(v1.ArticleStatus_ArticleNormal) {
			err = d.articleRepo.Publish(ctx, tx, save.ID)
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
	err := ent.WithTx(ctx, d.db, func(tx *gen.Client) error {
		var err error
		save, err = d.postscriptRepo.Save(ctx, tx, &model.ArticlePostscript{ArticlePostscript: &gen.ArticlePostscript{
			ArticleID: articleId,
			Content:   content,
			Status:    int32(v1.ArticlePostscriptStatus_ArticlePostscriptNormal),
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
	// Todo 广播添加事件
	return save, err
}

// --- 更新 ---

func (d *ArticleDomain) UpdateDraft(ctx context.Context, article *model.Article, tags []*model.Tag) (*model.Article, error) {
	var (
		save *model.Article
		err  error
	)
	status := article.Status
	article.Status = int32(v1.ArticleStatus_ArticleDrafts) // 默认均为草稿
	err = ent.WithTx(ctx, d.db, func(tx *gen.Client) error {
		exist, err := d.articleRepo.Exist(ctx, tx, &repo.ArticleGetReq{
			ArticleId: base.Ptr(article.ID),
			Status:    base.Ptr(v1.ArticleStatus_ArticleDrafts),
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
		if status == int32(v1.ArticleStatus_ArticleNormal) {
			err = d.articleRepo.Publish(ctx, tx, save.ID)
			if err != nil {
				return err
			}
		}
		return err
	})
	return save, err
}

func (d *ArticleDomain) Action(ctx context.Context, articleId int64, userId int64, action v1.ArticleAction, active bool) error {
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

func (d *ArticleDomain) GetOne(ctx context.Context, articleId int64) (*model.Article, error) {
	var (
		reply *model.Article
		err   error
	)
	reply, err = d.articleRepo.GetOne(ctx, d.db, &repo.ArticleGetReq{
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

	if reply.Status == int32(v1.ArticleStatus_ArticleDrafts) && !ok && *reply.CreatedBy != user.ID {
		return nil, cv1.ErrorUnauthorized("login required to view drafts")
	}

	lastReplyComment, err := d.commentRepo.GetArticleLastComment(ctx, d.db, &repo.CommentGetReq{ArticleId: base.Ptr(reply.ID)})
	if err != nil {
		return nil, err
	}

	userServiceClient, err := client.GetServiceClient(d.etcd, constant.UserServiceName.String(), userv1.NewUserUserServiceClient)
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
	list, pageReply, err = d.articleRepo.GetPage(ctx, d.db, page, req)
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
		lastCommentMap, err = d.commentRepo.GetArticleLastComments(ctx, d.db, &repo.CommentGetReq{ArticleIds: articleIds.ToSlice()})
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
		userServiceClient, err := client.GetServiceClient(d.etcd, constant.UserServiceName.String(), userv1.NewUserUserServiceClient)
		if err != nil {
			return nil, nil, err
		}
		userAuthorsMap, err = userServiceClient.GetMap(ctx, &userv1.GetMapRequest{Query: &userv1.UserQueryParams{UserIds: userIds.ToSlice()}})
		if err != nil {
			return nil, nil, err
		}
	}

	for i := range list {
		list[i].Summary()
		if lastReplyComment, ok := lastCommentMap.Get(list[i].ID); ok {
			list[i].LastReplyCommentAt = lastReplyComment.CreatedAt
			list[i].LastReplyCommentUser = userAuthorsMap.Users[*lastReplyComment.CreatedBy]
		}
		list[i].AuthorUser = base.If(list[i].Anonymous, nil, userAuthorsMap.Users[*list[i].CreatedBy])
	}
	return list, pageReply, err
}
