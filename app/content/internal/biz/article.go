package biz

import (
	"common/api/common/v1"
	"common/pkg/util"
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
	domainRepo       repo.DomainRepo
	sf               *sonyflake.Sonyflake
}

func NewArticleDomain(base *BaseDomain, articleRepo repo.ArticleRepo, postscriptRepo repo.ArticlePostscriptRepo, actionRecordRepo repo.ArticleActionRecordRepo, domainRepo repo.DomainRepo) (*ArticleDomain, error) {
	sf, err := util.NewSonyflake()
	if err != nil {
		return nil, err
	}
	return &ArticleDomain{
		BaseDomain:       base,
		articleRepo:      articleRepo,
		postscriptRepo:   postscriptRepo,
		actionRecordRepo: actionRecordRepo,
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
		if save.Status != int32(v1.ArticleStatus_ArticleDrafts) {
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

func (d *ArticleDomain) Action(ctx context.Context, articleId int64, userId int64, action v1.ArticleAction, active bool) error {
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
		err := d.articleRepo.UpdateStatus(ctx, client, articleId, v1.ArticleStatus_ArticleNormal)
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
