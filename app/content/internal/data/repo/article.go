package repo

import (
	"common/api/gen/common"
	cerrors "common/api/gen/common/errors"
	v1 "common/api/gen/content/v1"
	commonClient "common/pkg/client"
	"common/pkg/constant"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/conf"
	"content/internal/data/gen"
	"content/internal/data/gen/article"
	"content/internal/data/gen/articleactionrecord"
	"content/internal/data/gen/articlepostscript"
	"content/internal/data/gen/tag"
	"content/internal/enum"
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/samber/lo"
)

var _ repo.ArticleRepo = (*ArticleRepo)(nil)

type ArticleRepo struct {
	conf           *conf.Bootstrap
	log            *log.Helper
	consul         *commonClient.ConsulClient
	redis          *commonClient.RedisClient
	nats           *commonClient.NatsClient
	postscriptRepo repo.ArticlePostscriptRepo
	commentRepo    repo.CommentRepo
	domainRepo     repo.DomainRepo
	tagRepo        repo.TagRepo
}

func NewArticleRepo(
	conf *conf.Bootstrap,
	logger log.Logger,
	consul *commonClient.ConsulClient,
	redis *commonClient.RedisClient,
	nats *commonClient.NatsClient,
	postscriptRepo repo.ArticlePostscriptRepo,
	commentRepo repo.CommentRepo,
	domainRepo repo.DomainRepo,
	tagRepo repo.TagRepo,
) repo.ArticleRepo {
	return &ArticleRepo{
		conf:           conf,
		log:            log.NewHelper(logger),
		consul:         consul,
		redis:          redis,
		nats:           nats,
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
		tags = lo.Uniq(tags)
		tagNames := lo.SliceToMap[*model.Tag, string, struct{}](tags, func(m *model.Tag) (string, struct{}) { return m.Name, struct{}{} })
		constantTags, err := r.tagRepo.GetList(ctx, tx, &repo.TagGetReq{Names: lo.Keys(tagNames)})
		if err != nil {
			return nil, err
		}
		constantTags = lo.Uniq(constantTags)
		constantTagNameSet := lo.SliceToMap[*model.Tag, string, struct{}](constantTags, func(m *model.Tag) (string, struct{}) { return m.Name, struct{}{} })
		for _, i := range tags {
			if _, ok := constantTagNameSet[i.Name]; !ok {
				saveTags = append(saveTags, i)
			}
		}
		saveTags, err = r.tagRepo.Saves(ctx, tx, saveTags)
		if err != nil {
			return nil, err
		}
		for _, t := range constantTags {
			bindTagIds = append(bindTagIds, t.ID)
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
	return &model.Article{Article: save}, nil
}

func (r *ArticleRepo) Update(ctx context.Context, tx *gen.Client, updateArticle *model.Article, tags []*model.Tag) (*model.Article, error) {

	// 处理标签，去除重复
	bindTagIds := make([]int64, 0)
	if len(tags) > 0 {
		saveTags := make([]*model.Tag, 0)
		tags = lo.Uniq(tags)
		tagNames := lo.SliceToMap[*model.Tag, string, struct{}](tags, func(m *model.Tag) (string, struct{}) { return m.Name, struct{}{} })
		constantTags, err := r.tagRepo.GetList(ctx, tx, &repo.TagGetReq{Names: lo.Keys(tagNames)})
		if err != nil {
			return nil, err
		}
		constantTags = lo.Uniq(constantTags)
		constantTagNameSet := lo.SliceToMap[*model.Tag, string, struct{}](constantTags, func(m *model.Tag) (string, struct{}) { return m.Name, struct{}{} })
		for _, i := range tags {
			if _, ok := constantTagNameSet[i.Name]; !ok {
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
	return &model.Article{Article: save}, nil
}

func (r *ArticleRepo) UpdateContent(ctx context.Context, tx *gen.Client, articleId int64, content string) error {
	return tx.Article.UpdateOneID(articleId).
		SetContent(content).
		Exec(ctx)
}
func (r *ArticleRepo) UpdateStatus(ctx context.Context, tx *gen.Client, articleId int64, status v1.ArticleStatus) error {
	dbStatus, _ := enum.ArticleStatusMap.ToEnum(status)
	return tx.Article.UpdateOneID(articleId).
		SetStatus(article.Status(dbStatus)).
		Exec(ctx)
}

func (r *ArticleRepo) UpdateHasPostscript(ctx context.Context, tx *gen.Client, articleId int64, hasPostscript bool) error {
	return tx.Article.UpdateOneID(articleId).
		SetHasPostscript(hasPostscript).
		Exec(ctx)
}

func (r *ArticleRepo) UpdateStat(ctx context.Context, tx *gen.Client, articleId int64, action v1.ArticleAction, num int32) (*model.Article, error) {
	updateOne := tx.Article.UpdateOneID(articleId)
	switch action {
	case v1.ArticleAction_ARTICLE_ACTION_LIKE:
		updateOne.AddLikeCount(num)
	case v1.ArticleAction_ARTICLE_ACTION_THANK:
		updateOne.AddThankCount(num)
	case v1.ArticleAction_ARTICLE_ACTION_COLLECT:
		updateOne.AddCollectCount(num)
	case v1.ArticleAction_ARTICLE_ACTION_WATCH:
		updateOne.AddWatchCount(num)
	case v1.ArticleAction_ARTICLE_ACTION_REPLY:
		updateOne.AddReplyCount(num)
	case v1.ArticleAction_ARTICLE_ACTION_VOTE:
		updateOne.AddVoteTotal(num)
	case v1.ArticleAction_ARTICLE_ACTION_LOTTERY:
		updateOne.AddLotteryParticipantCount(num)
	case v1.ArticleAction_ARTICLE_ACTION_LOTTERY_WINNER:
		updateOne.AddLotteryWinnerCount(num)
	default:
		return nil, fmt.Errorf("unknown action")
	}
	save, err := updateOne.Save(ctx)
	return &model.Article{Article: save}, err
}

func (r *ArticleRepo) UpdateAcceptAnswer(ctx context.Context, tx *gen.Client, articleId int64, commentId int64) (*model.Article, error) {
	exist, err := r.commentRepo.Exist(ctx, tx, &repo.CommentGetReq{CommentId: new(commentId), ArticleId: new(articleId)})
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, fmt.Errorf("comment not exist")
	}
	a, err := tx.Article.UpdateOneID(articleId).
		SetAcceptedAnswerID(commentId).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.Article{Article: a}, nil
}

func (r *ArticleRepo) Publish(ctx context.Context, tx *gen.Client, articleId int64) (*model.Article, error) {
	first, err := r.GetOne(ctx, tx, &repo.ArticleGetReq{ArticleId: new(articleId)})
	if err != nil {
		return nil, err
	}

	if first.Status != article.StatusDrafts {
		return nil, cerrors.ErrorBadRequest("only update draft")
	}

	err = r.UpdateStatus(ctx, tx, articleId, v1.ArticleStatus_ARTICLE_STATUS_NORMAL)
	if err != nil {
		return nil, err
	}
	return first, nil
}

func (r *ArticleRepo) Delete(ctx context.Context, tx *gen.Client, articleId int64) error {
	return tx.Article.UpdateOneID(articleId).SetStatus(article.StatusDeleted).Exec(ctx)
}

func (r *ArticleRepo) Exist(ctx context.Context, tx *gen.Client, req *repo.ArticleGetReq) (bool, error) {
	query := tx.Article.Query()
	query = r.getQuery(query, req)
	return query.Exist(ctx)
}

func (r *ArticleRepo) GetOne(ctx context.Context, tx *gen.Client, req *repo.ArticleGetReq) (*model.Article, error) {
	query := tx.Article.Query().
		WithPostscripts(func(q *gen.ArticlePostscriptQuery) {
			q.Where(articlepostscript.StatusEQ(articlepostscript.StatusNormal)).
				Order(gen.Asc(articlepostscript.FieldCreatedAt))
		}).
		WithTags()
	query = r.getQuery(query, req)
	a, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, cerrors.ErrorBadRequest("article is not found")
	}
	return &model.Article{Article: a, IsSummary: req.IsSummary}, err
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
		articles = append(articles, &model.Article{Article: list[i], IsSummary: req.IsSummary})
	}
	return articles, nil
}

func (r *ArticleRepo) GetPage(ctx context.Context, tx *gen.Client, page *common.PageRequest, req *repo.ArticleGetReq) ([]*model.Article, *common.PageReply, error) {
	var (
		articles []*model.Article
		err      error
		total    int
	)
	page = constant.PageValid(page)
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
		articles = append(articles, &model.Article{Article: list[i], IsSummary: req.IsSummary})
	}
	return articles, &common.PageReply{
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
		dbStatus, _ := enum.ArticleStatusMap.ToEnum(*req.Status)
		query = query.Where(article.StatusEQ(article.Status(dbStatus)))
	}
	if req.AuthorId != nil {
		query = query.Where(article.CreatedByEQ(*req.AuthorId))
	}
	if req.Type != nil {
		dbType, _ := enum.ArticleTypeMap.ToEnum(*req.Type)
		query = query.Where(article.TypeEQ(article.Type(dbType)))
	}
	if req.QueryUserId != nil {
		query = query.WithActionRecords(func(query *gen.ArticleActionRecordQuery) {
			query.Where(articleactionrecord.UserIDEQ(*req.QueryUserId))
		})
	}
	if req.Keyword != nil {
		// TODO: 后续考虑使用 zhparser 全文搜索扩展实现。
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
		case v1.ArticleOrder_ARTICLE_ORDER_NEWEST:
			query = query.Order(gen.Desc(article.FieldCreatedAt))
		case v1.ArticleOrder_ARTICLE_ORDER_HOTTEST:
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
