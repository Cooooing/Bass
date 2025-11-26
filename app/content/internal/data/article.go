package data

import (
	cv1 "common/api/common/v1"
	v1 "common/api/content/v1"
	userv1 "common/api/user/v1"
	"common/pkg/client"
	"common/pkg/constant"
	"common/pkg/util/base"
	"common/pkg/util/collections/set"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/ent/gen"
	"content/internal/data/ent/gen/article"
	"content/internal/data/ent/gen/articlepostscript"
	"content/internal/data/ent/gen/tag"
	"context"
	"encoding/json"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/jinzhu/copier"
)

type ArticleRepo struct {
	*BaseRepo
	client         *gen.Client
	postscriptRepo repo.ArticlePostscriptRepo
	commentRepo    repo.CommentRepo
	domainRepo     repo.DomainRepo
	tagRepo        repo.TagRepo
}

func NewArticleRepo(baseRepo *BaseRepo, client *gen.Client, postscriptRepo repo.ArticlePostscriptRepo, commentRepo repo.CommentRepo, domainRepo repo.DomainRepo, tagRepo repo.TagRepo) repo.ArticleRepo {
	return &ArticleRepo{
		BaseRepo:       baseRepo,
		client:         client,
		postscriptRepo: postscriptRepo,
		commentRepo:    commentRepo,
		domainRepo:     domainRepo,
		tagRepo:        tagRepo,
	}
}

func (r *ArticleRepo) Save(ctx context.Context, tx *gen.Client, article *model.Article, tags []*model.Tag) (*model.Article, error) {

	// 处理标签，去除重复
	bindTagIds := make([]int64, 0)
	if len(tags) > 0 {
		saveTags := make([]*model.Tag, 0)
		tagNames := set.NewFromSlice[*model.Tag, string](tags, func(m *model.Tag) string { return m.Name })
		constantTags, err := r.tagRepo.GetList(ctx, tx, &repo.TagGetReq{Names: tagNames.ToSlice()})
		if err != nil {
			return nil, err
		}
		constantTagNameSet := set.NewFromSlice[*model.Tag, string](constantTags, func(m *model.Tag) string { return m.Name })
		for _, i := range tags {
			if !constantTagNameSet.Contains(i.Name) {
				saveTags = append(saveTags, i)
			}
		}
		saveTags, err = r.tagRepo.Saves(ctx, tx, saveTags)
		if err != nil {
			return nil, err
		}
		for _, i := range saveTags {
			bindTagIds = append(bindTagIds, i.ID)
		}
	}

	article.FormatContent()
	create := tx.Article.Create().
		SetTitle(article.Title).
		SetContent(article.Content).
		SetNillableRewardContent(article.RewardContent).
		SetNillableRewardPoints(article.RewardPoints).
		SetStatus(article.Status).
		SetType(article.Type).
		SetNillableBountyPoints(article.BountyPoints).
		SetNillableStatement(article.Statement).
		SetCommentable(article.Commentable).
		SetAnonymous(article.Anonymous).
		SetListable(article.Listable)
	if len(bindTagIds) > 0 {
		create.AddTagIDs(bindTagIds...)
	}
	save, err := create.Save(ctx)
	if err != nil {
		return nil, err
	}
	return (*model.Article)(save), nil
}

func (r *ArticleRepo) Update(ctx context.Context, tx *gen.Client, updateArticle *model.Article, tags []*model.Tag) (*model.Article, error) {

	// 处理标签，去除重复
	bindTagIds := make([]int64, 0)
	if len(tags) > 0 {
		saveTags := make([]*model.Tag, 0)
		tagNames := set.NewFromSlice[*model.Tag, string](tags, func(m *model.Tag) string { return m.Name })
		constantTags, err := r.tagRepo.GetList(ctx, tx, &repo.TagGetReq{Names: tagNames.ToSlice()})
		if err != nil {
			return nil, err
		}
		constantTagNameSet := set.NewFromSlice[*model.Tag, string](constantTags, func(m *model.Tag) string { return m.Name })
		for _, i := range tags {
			if !constantTagNameSet.Contains(i.Name) {
				saveTags = append(saveTags, i)
			}
		}
		saveTags, err = r.tagRepo.Saves(ctx, tx, saveTags)
		if err != nil {
			return nil, err
		}
		for _, i := range saveTags {
			bindTagIds = append(bindTagIds, i.ID)
		}
	}

	updateArticle.FormatContent()
	update := tx.Article.UpdateOneID(updateArticle.ID).
		SetTitle(updateArticle.Title).
		SetContent(updateArticle.Content).
		SetNillableRewardContent(updateArticle.RewardContent).
		SetNillableRewardPoints(updateArticle.RewardPoints).
		SetStatus(updateArticle.Status).
		SetType(updateArticle.Type).
		SetNillableBountyPoints(updateArticle.BountyPoints).
		SetNillableStatement(updateArticle.Statement).
		SetCommentable(updateArticle.Commentable).
		SetAnonymous(updateArticle.Anonymous).
		SetListable(updateArticle.Listable)
	if len(bindTagIds) > 0 {
		update.ClearTags().AddTagIDs(bindTagIds...)
	}
	save, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return (*model.Article)(save), nil
}

func (r *ArticleRepo) UpdateContent(ctx context.Context, tx *gen.Client, articleId int64, content string) error {
	return tx.Article.UpdateOneID(articleId).
		SetContent(content).
		Exec(ctx)
}
func (r *ArticleRepo) UpdateStatus(ctx context.Context, tx *gen.Client, articleId int64, status cv1.ArticleStatus) error {
	return tx.Article.UpdateOneID(articleId).
		SetStatus(int32(status)).
		Exec(ctx)
}

func (r *ArticleRepo) UpdateHasPostscript(ctx context.Context, tx *gen.Client, articleId int64, hasPostscript bool) error {
	return tx.Article.UpdateOneID(articleId).
		SetHasPostscript(hasPostscript).
		Exec(ctx)
}

func (r *ArticleRepo) UpdateStat(ctx context.Context, tx *gen.Client, articleId int64, action cv1.ArticleAction, num int32) error {
	updateOne := tx.Article.UpdateOneID(articleId)
	switch action {
	case cv1.ArticleAction_ArticleActionLike:
		updateOne.AddLikeCount(num)
	case cv1.ArticleAction_ArticleActionThank:
		updateOne.AddThankCount(num)
	case cv1.ArticleAction_ArticleActionCollect:
		updateOne.AddCollectCount(num)
	case cv1.ArticleAction_ArticleActionWatch:
		updateOne.AddWatchCount(num)
	case cv1.ArticleAction_ArticleActionReply:
		updateOne.AddReplyCount(num)
	case cv1.ArticleAction_ArticleActionVote:
		updateOne.AddVoteTotal(num)
	case cv1.ArticleAction_ArticleActionLottery:
		updateOne.AddLotteryParticipantCount(num)
	case cv1.ArticleAction_ArticleActionLotteryWinner:
		updateOne.AddLotteryWinnerCount(num)
	default:
		return nil
	}
	return updateOne.Exec(ctx)
}

func (r *ArticleRepo) Publish(ctx context.Context, tx *gen.Client, articleId int64) error {
	first, err := r.GetById(ctx, tx, &repo.ArticleGetReq{ArticleId: base.Ptr(articleId)})
	if err != nil {
		return err
	}

	if first.Status != int32(cv1.ArticleStatus_ArticleDrafts) {
		return cv1.ErrorBadRequest("only update draft")
	}

	err = r.UpdateStatus(ctx, tx, articleId, cv1.ArticleStatus_ArticleNormal)
	if err != nil {
		return err
	}

	return nil

	// Todo 广播添加文章事件

	publish := &v1.ArticleEventPublish{}
	err = copier.Copy(&publish, first)
	if err != nil {
		return err
	}
	marshal, err := json.Marshal(publish)
	if err != nil {
		return err
	}
	err = r.rabbitmq.Publish(constant.ExchangeContent.String(), constant.RoutingKeyArticleCreate.String(), marshal)
	if err != nil {
		return err
	}
	// Todo 广播@用户通知
	_, atUserNames := first.ParseContent()

	userServiceClient, err := client.GetServiceClient(r.etcd, constant.UserServiceName.String(), userv1.NewUserUserServiceClient)
	atUserList, err := userServiceClient.GetList(ctx, &userv1.GetListRequest{Query: &userv1.UserQueryParams{Names: atUserNames}})
	if err != nil {
		return err
	}
	_ = atUserList

	return nil
}

func (r *ArticleRepo) Delete(ctx context.Context, tx *gen.Client, articleId int64) error {
	return tx.Article.UpdateOneID(articleId).SetStatus(int32(cv1.ArticleStatus_ArticleDeleted)).Exec(ctx)
}

func (r *ArticleRepo) Exist(ctx context.Context, tx *gen.Client, req *repo.ArticleGetReq) (bool, error) {
	if req.ArticleId == nil {
		return false, cv1.ErrorBadRequest("articleId is required")
	}
	query := tx.Article.Query()
	query = r.getQuery(query, req)
	return query.Exist(ctx)
}

func (r *ArticleRepo) GetById(ctx context.Context, tx *gen.Client, req *repo.ArticleGetReq) (*model.Article, error) {
	if req.ArticleId == nil {
		return nil, cv1.ErrorBadRequest("articleId is required")
	}
	query := tx.Article.Query().
		WithPostscripts(func(q *gen.ArticlePostscriptQuery) {
			q.Where(articlepostscript.StatusEQ(int32(cv1.ArticlePostscriptStatus_ArticlePostscriptNormal))).
				Order(gen.Asc(articlepostscript.FieldCreatedAt))
		}).
		WithTags()
	query = r.getQuery(query, req)
	a, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, cv1.ErrorBadRequest("article is not found")
	}
	return (*model.Article)(a), err
}

func (r *ArticleRepo) GetList(ctx context.Context, tx *gen.Client, req *repo.ArticleGetReq) ([]*model.Article, error) {
	var (
		articles []*model.Article
		err      error
	)
	query := tx.Article.Query().WithTags()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	for i := range list {
		articles = append(articles, (*model.Article)(list[i]))
	}
	return articles, nil
}

func (r *ArticleRepo) GetPage(ctx context.Context, tx *gen.Client, page *cv1.PageRequest, req *repo.ArticleGetReq) ([]*model.Article, *cv1.PageReply, error) {
	var (
		articles []*model.Article
		err      error
		total    int
	)
	page = base.OrDefault(page, constant.GetPageDefault())
	query := tx.Article.Query().WithTags()
	query = r.getQuery(query, req)
	countQuery := query.Clone()
	total, err = countQuery.Count(ctx)
	if err != nil {
		return nil, nil, err
	}
	list, err := query.Limit(int(page.Size)).Offset(int((page.Page - 1) * page.Size)).All(ctx)
	if err != nil {
		return nil, nil, err
	}
	for i := range list {
		articles = append(articles, (*model.Article)(list[i]))
	}
	return articles, &cv1.PageReply{
		Total: uint32(total),
		Page:  page.Page,
		Size:  page.Size,
	}, nil
}

func (r *ArticleRepo) getQuery(query *gen.ArticleQuery, req *repo.ArticleGetReq) *gen.ArticleQuery {
	if req.ArticleId != nil {
		query = query.Where(article.IDEQ(*req.ArticleId))
	}
	if req.CreatedBy != nil {
		query = query.Where(article.CreatedByEQ(*req.CreatedBy))
	}
	if req.TagId != nil {
		query = query.Where(article.HasTagsWith(tag.IDEQ(*req.TagId)))
	}
	if req.DomainId != nil {
		query = query.Where(article.HasTagsWith(tag.DomainIDEQ(*req.DomainId)))
	}
	if req.Status != nil {
		query = query.Where(article.StatusEQ(int32(*req.Status)))
	}
	if req.AuthorId != nil {
		query = query.Where(article.CreatedByEQ(*req.AuthorId))
	}
	if req.Type != nil {
		query = query.Where(article.TypeEQ(int32(*req.Type)))
	}
	if req.Keyword != nil {
		// Todo 后续考虑使用 zhparser 全文搜索拓展实现
		query = query.Where(
			article.Or(
				article.TitleContains(*req.Keyword),
				article.ContentContains(*req.Keyword),
			),
		)
	}
	if req.Listable != nil {
		query = query.Where(article.ListableEQ(*req.Listable))
	}
	if req.Order != nil {
		switch *req.Order {
		case cv1.ArticleOrder_ArticleOrderNewest:
			query = query.Order(gen.Desc(article.FieldCreatedAt))
		case cv1.ArticleOrder_ArticleOrderHottest:
			query = query.Where(article.CreatedAtGTE(time.Now().Add(-30 * 24 * time.Hour))).
				Order(func(s *sql.Selector) {
					s.OrderExpr(sql.Expr(`
        (
            (reply_count * 8 + like_count * 4 + collect_count * 6 + thank_count * 2 + watch_count * 1)
            /
            pow((extract(epoch from (now() - created_at)) / 3600) + 2 , 1.3)
        ) DESC`))
				})
		}
	} else {
		query = query.Order(gen.Desc(article.FieldCreatedAt))
	}
	return query
}
