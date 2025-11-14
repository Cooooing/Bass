package biz

import (
	cv1 "common/api/common/v1"
	v1 "common/api/content/v1"
	userv1 "common/api/user/v1"
	"common/pkg/client"
	"common/pkg/constant"
	"common/pkg/util"
	"common/pkg/util/collections/dict"
	"common/pkg/util/collections/set"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/ent"
	"content/internal/data/ent/gen"
	"context"

	"github.com/jinzhu/copier"
	"github.com/sony/sonyflake/v2"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ArticleDomain struct {
	*BaseDomain
	articleRepo      repo.ArticleRepo
	postscriptRepo   repo.ArticlePostscriptRepo
	actionRecordRepo repo.ArticleActionRecordRepo
	commentRepo      repo.CommentRepo
	domainRepo       repo.DomainRepo
	sf               *sonyflake.Sonyflake
}

func NewArticleDomain(base *BaseDomain, articleRepo repo.ArticleRepo, postscriptRepo repo.ArticlePostscriptRepo, actionRecordRepo repo.ArticleActionRecordRepo, commentRepo repo.CommentRepo, domainRepo repo.DomainRepo) (*ArticleDomain, error) {
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
		domainRepo:       domainRepo,
		sf:               sf,
	}, nil
}

func (d *ArticleDomain) Add(ctx context.Context, article *model.Article) (*model.Article, error) {
	var (
		save *model.Article
		err  error
	)
	err = ent.WithTx(ctx, d.db, func(client *gen.Client) error {
		save, err = d.articleRepo.Save(ctx, d.db, article)
		if err != nil {
			return err
		}
		// 不是草稿则进行发布
		if save.Status != int32(cv1.ArticleStatus_ArticleDrafts) {
			err = d.articleRepo.Publish(ctx, d.db, save.ID)
			if err != nil {
				return err
			}
		}
		return err
	})
	return save, err
}

func (d *ArticleDomain) AddPostscript(ctx context.Context, articleId int64, content string) error {
	err := ent.WithTx(ctx, d.db, func(client *gen.Client) error {
		var err error
		err = d.postscriptRepo.AddPostscript(ctx, client, articleId, content)
		if err != nil {
			return err
		}
		err = d.articleRepo.UpdateHasPostscript(ctx, client, articleId, true)
		if err != nil {
			return err
		}
		return err
	})
	// Todo 广播添加事件
	return err
}

func (d *ArticleDomain) Action(ctx context.Context, articleId int64, userId int64, action cv1.ArticleAction, active bool) error {
	err := ent.WithTx(ctx, d.db, func(client *gen.Client) error {
		var err error
		if active {
			err = d.articleRepo.UpdateStat(ctx, client, articleId, action, 1)
			if err != nil {
				return err
			}
			_, err = d.actionRecordRepo.Save(ctx, client, &model.ArticleActionRecord{
				ArticleID: articleId,
				UserID:    userId,
				Type:      int32(action),
			})
			if err != nil {
				return err
			}
		} else {
			err = d.articleRepo.UpdateStat(ctx, client, articleId, action, -1)
			if err != nil {
				return err
			}
			err = d.actionRecordRepo.Delete(ctx, client, articleId, userId, action)
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
	return ent.WithTx(ctx, d.db, func(client *gen.Client) error {
		err := d.articleRepo.UpdateStatus(ctx, client, articleId, cv1.ArticleStatus_ArticleNormal)
		if err != nil {
			return err
		}
		err = d.articleRepo.Publish(ctx, client, articleId)
		if err != nil {
			return err
		}
		return err
	})
}

// --- 查询 ---

func (d *ArticleDomain) Page(ctx context.Context, page *cv1.PageRequest, req *repo.ArticleGetReq) (*v1.PageArticleReply, error) {
	var (
		list      []*model.Article
		pageReply *cv1.PageReply
		reply     *v1.PageArticleReply
		err       error
	)
	err = ent.WithTx(ctx, d.db, func(tx *gen.Client) error {
		list, pageReply, err = d.articleRepo.GetPage(ctx, tx, page, req)
		if err != nil {
			return err
		}
		articleIds := set.New[int64](0)
		userIds := set.New[int64](0)
		for _, item := range list {
			articleIds.Add(item.ID)
			userIds.Add(item.UserID)
		}

		lastCommentMap, _ := d.commentRepo.GetArticleLastComments(ctx, tx, articleIds.ToSlice())
		lastCommentMap.Foreach(func(e *dict.Entry[int64, *model.Comment]) bool {
			userIds.Add(*e.Value.CreatedBy)
			return true
		})

		userServiceClient, err := client.GetServiceClient(ctx, d.etcd, constant.UserServiceName.String(), userv1.NewUserUserServiceClient)
		if err != nil {
			return err
		}
		userAuthors, err := userServiceClient.GetMap(ctx, &userv1.GetMapRequest{
			Ids: userIds.ToSlice(),
		})
		if err != nil {
			return err
		}

		articles := make([]*v1.Article, 0, len(list))
		for _, item := range list {
			item.Summary()
			a := &v1.Article{}
			err = copier.Copy(a, item)
			if err != nil {
				return err
			}
			a.CreatedAt = timestamppb.New(*item.CreatedAt)
			a.UpdatedAt = timestamppb.New(*item.UpdatedAt)
			if lastReplyComment, ok := lastCommentMap.Get(item.ID); ok {
				a.RepliedAt = timestamppb.New(*lastReplyComment.CreatedAt)
				a.ReplyUser = userAuthors.Users[*lastReplyComment.CreatedBy]
			}
			a.AuthorUser = userAuthors.Users[item.UserID]
			articles = append(articles, a)
		}
		reply = &v1.PageArticleReply{
			Page:     pageReply,
			Articles: articles,
		}
		return nil
	})
	return reply, err
}
