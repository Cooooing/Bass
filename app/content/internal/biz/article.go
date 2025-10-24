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
	save, err := d.articleRepo.Save(ctx, d.db, article)
	if err != nil {
		return nil, err
	}
	// Todo 广播发帖事件
	return save, err
}

func (d *ArticleDomain) AddPostscript(ctx context.Context, articleId int, content string) error {
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

func (d *ArticleDomain) Action(ctx context.Context, action v1.ArticleAction, articleId int, userId int, active bool) error {
	err := ent.WithTx(ctx, d.db, func(client *gen.Client) error {
		var err error
		if active {
			err = d.articleRepo.UpdateStat(ctx, client, articleId, action, 1)
			if err != nil {
				return err
			}
			_, err = d.actionRecordRepo.Save(ctx, client, &model.ActionRecord{
				ArticleID: articleId,
				UserID:    userId,
				Type:      int(action),
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
