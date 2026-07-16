package repo

import (
	cerrors "common/proto/gen/common/errors"
	"content/internal/biz/base"
	"context"
	"time"

	"common/pkg/apperror"
	utilent "common/pkg/util/ent"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/gen"
	articleent "content/internal/data/gen/article"
	tagent "content/internal/data/gen/tag"
	"content/internal/enum"

	"entgo.io/ent/dialect/sql"
	"github.com/samber/lo"
)

var _ repo.ArticleRepo = (*ArticleRepo)(nil)

type ArticleRepo struct {
	db *gen.Client
}

func NewArticleRepo(db *gen.Client) repo.ArticleRepo {
	return &ArticleRepo{db: db}
}

func (r *ArticleRepo) getClient(ctx context.Context) *gen.Client {
	if tx, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return tx
	}
	return r.db
}

func (r *ArticleRepo) Save(ctx context.Context, req *repo.ArticleSaveReq) (*repo.ArticleSaveResponse, error) {
	article := req.Article
	save, err := r.getClient(ctx).Article.Create().
		SetTitle(article.Title).
		SetContent(article.Content).
		SetNillableRewardContent(article.RewardContent).
		SetNillableRewardPoints(article.RewardPoints).
		SetPublishStatus(articleent.PublishStatus(article.PublishStatus)).
		SetVisibility(articleent.Visibility(article.Visibility)).
		SetRestriction(articleent.Restriction(article.Restriction)).
		SetType(articleent.Type(article.Type)).
		SetNillableBountyPoints(article.BountyPoints).
		SetNillableStatement(article.Statement).
		SetCommentable(article.Commentable).
		SetAnonymous(article.Anonymous).
		SetNillablePublishedAt(article.PublishedAt).
		SetNillableEditedAt(article.EditedAt).
		SetNillableCreatedBy(article.CreatedBy).
		SetNillableUpdatedBy(article.UpdatedBy).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &repo.ArticleSaveResponse{Article: &model.Article{
		ID:               save.ID,
		Title:            save.Title,
		Content:          save.Content,
		HasPostscript:    save.HasPostscript,
		RewardContent:    save.RewardContent,
		RewardPoints:     save.RewardPoints,
		PublishStatus:    enum.ArticlePublishStatus(save.PublishStatus),
		Visibility:       enum.ArticleVisibility(save.Visibility),
		Restriction:      enum.ContentRestriction(save.Restriction),
		Type:             enum.ArticleType(save.Type),
		Statement:        save.Statement,
		Commentable:      save.Commentable,
		Anonymous:        save.Anonymous,
		PublishedAt:      save.PublishedAt,
		EditedAt:         save.EditedAt,
		ViewCount:        save.ViewCount,
		ThankCount:       save.ThankCount,
		LikeCount:        save.LikeCount,
		CollectCount:     save.CollectCount,
		WatchCount:       save.WatchCount,
		ReplyCount:       save.ReplyCount,
		BountyPoints:     save.BountyPoints,
		AcceptedAnswerID: save.AcceptedAnswerID,
		CreatedAt:        save.CreatedAt,
		UpdatedAt:        save.UpdatedAt,
		CreatedBy:        save.CreatedBy,
		UpdatedBy:        save.UpdatedBy,
		DeletedAt:        save.DeletedAt,
	}}, nil
}

func (r *ArticleRepo) Update(ctx context.Context, req *repo.ArticleUpdateReq) (*repo.ArticleUpdateResponse, error) {
	updateArticle := req.Article
	save, err := r.getClient(ctx).Article.UpdateOneID(updateArticle.ID).
		SetTitle(updateArticle.Title).
		SetContent(updateArticle.Content).
		SetNillableRewardContent(updateArticle.RewardContent).
		SetNillableRewardPoints(updateArticle.RewardPoints).
		SetPublishStatus(articleent.PublishStatus(updateArticle.PublishStatus)).
		SetVisibility(articleent.Visibility(updateArticle.Visibility)).
		SetRestriction(articleent.Restriction(updateArticle.Restriction)).
		SetType(articleent.Type(updateArticle.Type)).
		SetNillableBountyPoints(updateArticle.BountyPoints).
		SetNillableStatement(updateArticle.Statement).
		SetCommentable(updateArticle.Commentable).
		SetAnonymous(updateArticle.Anonymous).
		SetNillablePublishedAt(updateArticle.PublishedAt).
		SetNillableEditedAt(updateArticle.EditedAt).
		SetNillableUpdatedBy(updateArticle.UpdatedBy).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &repo.ArticleUpdateResponse{Article: &model.Article{
		ID:               save.ID,
		Title:            save.Title,
		Content:          save.Content,
		HasPostscript:    save.HasPostscript,
		RewardContent:    save.RewardContent,
		RewardPoints:     save.RewardPoints,
		PublishStatus:    enum.ArticlePublishStatus(save.PublishStatus),
		Visibility:       enum.ArticleVisibility(save.Visibility),
		Restriction:      enum.ContentRestriction(save.Restriction),
		Type:             enum.ArticleType(save.Type),
		Statement:        save.Statement,
		Commentable:      save.Commentable,
		Anonymous:        save.Anonymous,
		PublishedAt:      save.PublishedAt,
		EditedAt:         save.EditedAt,
		ViewCount:        save.ViewCount,
		ThankCount:       save.ThankCount,
		LikeCount:        save.LikeCount,
		CollectCount:     save.CollectCount,
		WatchCount:       save.WatchCount,
		ReplyCount:       save.ReplyCount,
		BountyPoints:     save.BountyPoints,
		AcceptedAnswerID: save.AcceptedAnswerID,
		CreatedAt:        save.CreatedAt,
		UpdatedAt:        save.UpdatedAt,
		CreatedBy:        save.CreatedBy,
		UpdatedBy:        save.UpdatedBy,
		DeletedAt:        save.DeletedAt,
	}}, nil
}

func (r *ArticleRepo) UpdatePublishStatus(ctx context.Context, req *repo.ArticleUpdatePublishStatusReq) (*repo.ArticleUpdatePublishStatusResponse, error) {
	articleId := req.ArticleID
	publishStatus := req.PublishStatus
	visibility := req.Visibility
	publishedAt := req.PublishedAt
	updatedBy := req.UpdatedBy
	_, err := r.getClient(ctx).Article.Get(ctx, articleId)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	update := r.getClient(ctx).Article.UpdateOneID(articleId).
		SetPublishStatus(articleent.PublishStatus(publishStatus)).
		SetVisibility(articleent.Visibility(visibility)).
		SetUpdatedBy(updatedBy)
	if publishedAt != nil {
		update.SetPublishedAt(*publishedAt)
	}
	if err := update.Exec(ctx); err != nil {
		return nil, err
	}
	return &repo.ArticleUpdatePublishStatusResponse{}, nil
}

func (r *ArticleRepo) UpdateVisibility(ctx context.Context, req *repo.ArticleUpdateVisibilityReq) (*repo.ArticleUpdateVisibilityResponse, error) {
	articleId := req.ArticleID
	visibility := req.Visibility
	updatedBy := req.UpdatedBy
	if err := r.getClient(ctx).Article.UpdateOneID(articleId).
		SetVisibility(articleent.Visibility(visibility)).
		SetUpdatedBy(updatedBy).
		Exec(ctx); err != nil {
		return nil, err
	}
	return &repo.ArticleUpdateVisibilityResponse{}, nil
}

func (r *ArticleRepo) UpdateRestriction(ctx context.Context, req *repo.ArticleUpdateRestrictionReq) (*repo.ArticleUpdateRestrictionResponse, error) {
	articleId := req.ArticleID
	restriction := req.Restriction
	updatedBy := req.UpdatedBy
	_, err := r.getClient(ctx).Article.Get(ctx, articleId)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	if err := r.getClient(ctx).Article.UpdateOneID(articleId).
		SetRestriction(articleent.Restriction(restriction)).
		SetUpdatedBy(updatedBy).
		Exec(ctx); err != nil {
		return nil, err
	}
	return &repo.ArticleUpdateRestrictionResponse{}, nil
}

func (r *ArticleRepo) DiscardDraft(ctx context.Context, req *repo.ArticleDiscardDraftReq) (*repo.ArticleDiscardDraftResponse, error) {
	articleId := req.ArticleID
	if err := r.getClient(ctx).Article.DeleteOneID(articleId).Exec(ctx); err != nil {
		return nil, err
	}
	return &repo.ArticleDiscardDraftResponse{}, nil
}

func (r *ArticleRepo) UpdateHasPostscript(ctx context.Context, req *repo.ArticleUpdateHasPostscriptReq) (*repo.ArticleUpdateHasPostscriptResponse, error) {
	articleId := req.ArticleID
	hasPostscript := req.HasPostscript
	updatedBy := req.UpdatedBy
	if err := r.getClient(ctx).Article.UpdateOneID(articleId).
		SetHasPostscript(hasPostscript).
		SetUpdatedBy(updatedBy).
		Exec(ctx); err != nil {
		return nil, err
	}
	return &repo.ArticleUpdateHasPostscriptResponse{}, nil
}

func (r *ArticleRepo) AddStats(ctx context.Context, req *repo.ArticleAddStatsReq) (*repo.ArticleAddStatsResponse, error) {
	articleId := req.ArticleID
	stats := req.Stats
	updateOne := r.getClient(ctx).Article.UpdateOneID(articleId)
	if stats.ViewCount != 0 {
		updateOne.AddViewCount(stats.ViewCount)
	}
	if stats.ThankCount != 0 {
		updateOne.AddThankCount(stats.ThankCount)
	}
	if stats.LikeCount != 0 {
		updateOne.AddLikeCount(stats.LikeCount)
	}
	if stats.CollectCount != 0 {
		updateOne.AddCollectCount(stats.CollectCount)
	}
	if stats.WatchCount != 0 {
		updateOne.AddWatchCount(stats.WatchCount)
	}
	if stats.ReplyCount != 0 {
		updateOne.AddReplyCount(stats.ReplyCount)
	}
	if err := updateOne.Exec(ctx); err != nil {
		return nil, err
	}
	return &repo.ArticleAddStatsResponse{}, nil
}

func (r *ArticleRepo) UpdateAcceptedAnswerID(ctx context.Context, req *repo.ArticleUpdateAcceptedAnswerIDReq) (*repo.ArticleUpdateAcceptedAnswerIDResponse, error) {
	articleId := req.ArticleID
	commentId := req.CommentID
	updatedBy := req.UpdatedBy
	a, err := r.getClient(ctx).Article.UpdateOneID(articleId).
		SetAcceptedAnswerID(commentId).
		SetUpdatedBy(updatedBy).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &repo.ArticleUpdateAcceptedAnswerIDResponse{Article: &model.Article{
		ID:               a.ID,
		Title:            a.Title,
		Content:          a.Content,
		HasPostscript:    a.HasPostscript,
		RewardContent:    a.RewardContent,
		RewardPoints:     a.RewardPoints,
		PublishStatus:    enum.ArticlePublishStatus(a.PublishStatus),
		Visibility:       enum.ArticleVisibility(a.Visibility),
		Restriction:      enum.ContentRestriction(a.Restriction),
		Type:             enum.ArticleType(a.Type),
		Statement:        a.Statement,
		Commentable:      a.Commentable,
		Anonymous:        a.Anonymous,
		PublishedAt:      a.PublishedAt,
		EditedAt:         a.EditedAt,
		ViewCount:        a.ViewCount,
		ThankCount:       a.ThankCount,
		LikeCount:        a.LikeCount,
		CollectCount:     a.CollectCount,
		WatchCount:       a.WatchCount,
		ReplyCount:       a.ReplyCount,
		BountyPoints:     a.BountyPoints,
		AcceptedAnswerID: a.AcceptedAnswerID,
		CreatedAt:        a.CreatedAt,
		UpdatedAt:        a.UpdatedAt,
		CreatedBy:        a.CreatedBy,
		UpdatedBy:        a.UpdatedBy,
		DeletedAt:        a.DeletedAt,
	}}, nil
}

func (r *ArticleRepo) ReplaceTags(ctx context.Context, req *repo.ArticleReplaceTagsReq) (*repo.ArticleReplaceTagsResponse, error) {
	articleId := req.ArticleID
	tagIds := req.TagIDs
	update := r.getClient(ctx).Article.UpdateOneID(articleId).ClearTags()
	if len(tagIds) > 0 {
		update.AddTagIDs(tagIds...)
	}
	if err := update.Exec(ctx); err != nil {
		return nil, err
	}
	return &repo.ArticleReplaceTagsResponse{}, nil
}

func (r *ArticleRepo) Exist(ctx context.Context, req *repo.ArticleGetReq) (*repo.ArticleExistResponse, error) {
	query := r.getClient(ctx).Article.Query()
	query = r.getQuery(query, req)
	exist, err := query.Exist(ctx)
	if err != nil {
		return nil, err
	}
	return &repo.ArticleExistResponse{Exist: exist}, nil
}

func (r *ArticleRepo) Get(ctx context.Context, req *repo.ArticleGetReq) (*repo.ArticleGetResponse, error) {
	query := r.getClient(ctx).Article.Query()
	query = r.getQuery(query, req)
	a, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return &repo.ArticleGetResponse{Article: &model.Article{
		ID:               a.ID,
		Title:            a.Title,
		Content:          a.Content,
		HasPostscript:    a.HasPostscript,
		RewardContent:    a.RewardContent,
		RewardPoints:     a.RewardPoints,
		PublishStatus:    enum.ArticlePublishStatus(a.PublishStatus),
		Visibility:       enum.ArticleVisibility(a.Visibility),
		Restriction:      enum.ContentRestriction(a.Restriction),
		Type:             enum.ArticleType(a.Type),
		Statement:        a.Statement,
		Commentable:      a.Commentable,
		Anonymous:        a.Anonymous,
		PublishedAt:      a.PublishedAt,
		EditedAt:         a.EditedAt,
		ViewCount:        a.ViewCount,
		ThankCount:       a.ThankCount,
		LikeCount:        a.LikeCount,
		CollectCount:     a.CollectCount,
		WatchCount:       a.WatchCount,
		ReplyCount:       a.ReplyCount,
		BountyPoints:     a.BountyPoints,
		AcceptedAnswerID: a.AcceptedAnswerID,
		CreatedAt:        a.CreatedAt,
		UpdatedAt:        a.UpdatedAt,
		CreatedBy:        a.CreatedBy,
		UpdatedBy:        a.UpdatedBy,
		DeletedAt:        a.DeletedAt,
	}}, nil
}

func (r *ArticleRepo) List(ctx context.Context, req *repo.ArticleGetReq) (*repo.ArticleListResponse, error) {
	query := r.getClient(ctx).Article.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	return &repo.ArticleListResponse{Rows: lo.Map(list, func(item *gen.Article, _ int) *model.Article {
		return &model.Article{
			ID:               item.ID,
			Title:            item.Title,
			Content:          item.Content,
			HasPostscript:    item.HasPostscript,
			RewardContent:    item.RewardContent,
			RewardPoints:     item.RewardPoints,
			PublishStatus:    enum.ArticlePublishStatus(item.PublishStatus),
			Visibility:       enum.ArticleVisibility(item.Visibility),
			Restriction:      enum.ContentRestriction(item.Restriction),
			Type:             enum.ArticleType(item.Type),
			Statement:        item.Statement,
			Commentable:      item.Commentable,
			Anonymous:        item.Anonymous,
			PublishedAt:      item.PublishedAt,
			EditedAt:         item.EditedAt,
			ViewCount:        item.ViewCount,
			ThankCount:       item.ThankCount,
			LikeCount:        item.LikeCount,
			CollectCount:     item.CollectCount,
			WatchCount:       item.WatchCount,
			ReplyCount:       item.ReplyCount,
			BountyPoints:     item.BountyPoints,
			AcceptedAnswerID: item.AcceptedAnswerID,
			CreatedAt:        item.CreatedAt,
			UpdatedAt:        item.UpdatedAt,
			CreatedBy:        item.CreatedBy,
			UpdatedBy:        item.UpdatedBy,
			DeletedAt:        item.DeletedAt,
		}
	})}, nil
}

func (r *ArticleRepo) Map(ctx context.Context, req *repo.ArticleGetReq) (*repo.ArticleMapResponse, error) {
	listResponse, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	return &repo.ArticleMapResponse{Rows: lo.SliceToMap(listResponse.Rows, func(item *model.Article) (int64, *model.Article) {
		return item.ID, item
	})}, nil
}

func (r *ArticleRepo) Count(ctx context.Context, req *repo.ArticleGetReq) (*repo.ArticleCountResponse, error) {
	query := r.getClient(ctx).Article.Query()
	query = r.getQuery(query, req)
	count, err := query.Count(ctx)
	if err != nil {
		return nil, err
	}
	return &repo.ArticleCountResponse{Count: count}, nil
}

func (r *ArticleRepo) Page(ctx context.Context, req *repo.ArticleGetReq) (*repo.ArticlePageResponse, error) {
	page := normalizePage(req.Page)
	query := r.getClient(ctx).Article.Query()
	query = r.getQuery(query, req)
	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}
	list, err := query.Limit(int(page.Size)).Offset(int((page.Page - 1) * page.Size)).All(ctx)
	if err != nil {
		return nil, err
	}
	articles := lo.Map(list, func(item *gen.Article, _ int) *model.Article {
		return &model.Article{
			ID:               item.ID,
			Title:            item.Title,
			Content:          item.Content,
			HasPostscript:    item.HasPostscript,
			RewardContent:    item.RewardContent,
			RewardPoints:     item.RewardPoints,
			PublishStatus:    enum.ArticlePublishStatus(item.PublishStatus),
			Visibility:       enum.ArticleVisibility(item.Visibility),
			Restriction:      enum.ContentRestriction(item.Restriction),
			Type:             enum.ArticleType(item.Type),
			Statement:        item.Statement,
			Commentable:      item.Commentable,
			Anonymous:        item.Anonymous,
			PublishedAt:      item.PublishedAt,
			EditedAt:         item.EditedAt,
			ViewCount:        item.ViewCount,
			ThankCount:       item.ThankCount,
			LikeCount:        item.LikeCount,
			CollectCount:     item.CollectCount,
			WatchCount:       item.WatchCount,
			ReplyCount:       item.ReplyCount,
			BountyPoints:     item.BountyPoints,
			AcceptedAnswerID: item.AcceptedAnswerID,
			CreatedAt:        item.CreatedAt,
			UpdatedAt:        item.UpdatedAt,
			CreatedBy:        item.CreatedBy,
			UpdatedBy:        item.UpdatedBy,
			DeletedAt:        item.DeletedAt,
		}
	})
	return &repo.ArticlePageResponse{
		Rows: articles,
		Page: &base.PageResponse{
			Total: int64(total),
			Page:  page.Page,
			Size:  page.Size,
		},
	}, nil
}

func (r *ArticleRepo) getQuery(query *gen.ArticleQuery, req *repo.ArticleGetReq) *gen.ArticleQuery {
	query = query.Where(articleent.DeletedAtIsNil())
	if req == nil {
		req = &repo.ArticleGetReq{}
	}
	if req.ArticleId != nil {
		query = query.Where(articleent.IDEQ(*req.ArticleId))
	}
	if len(req.ArticleIds) > 0 {
		query = query.Where(articleent.IDIn(req.ArticleIds...))
	}
	if req.CreatedBy != nil {
		query = query.Where(articleent.CreatedByEQ(*req.CreatedBy))
	}
	if req.TagId != nil {
		query = query.Where(articleent.HasTagsWith(tagent.IDEQ(*req.TagId)))
	}
	if req.DomainId != nil {
		query = query.Where(articleent.HasTagsWith(tagent.DomainIDEQ(*req.DomainId)))
	}
	if req.PublishStatus != nil {
		query = query.Where(articleent.PublishStatusEQ(articleent.PublishStatus(*req.PublishStatus)))
	}
	if len(req.PublishStatuses) > 0 {
		query = query.Where(articleent.PublishStatusIn(lo.Map(req.PublishStatuses, func(item enum.ArticlePublishStatus, _ int) articleent.PublishStatus {
			return articleent.PublishStatus(item)
		})...))
	}
	if req.Visibility != nil {
		query = query.Where(articleent.VisibilityEQ(articleent.Visibility(*req.Visibility)))
	}
	if len(req.Visibilities) > 0 {
		query = query.Where(articleent.VisibilityIn(lo.Map(req.Visibilities, func(item enum.ArticleVisibility, _ int) articleent.Visibility {
			return articleent.Visibility(item)
		})...))
	}
	if req.Restriction != nil {
		query = query.Where(articleent.RestrictionEQ(articleent.Restriction(*req.Restriction)))
	}
	if len(req.Restrictions) > 0 {
		query = query.Where(articleent.RestrictionIn(lo.Map(req.Restrictions, func(item enum.ContentRestriction, _ int) articleent.Restriction {
			return articleent.Restriction(item)
		})...))
	}
	if req.AuthorId != nil {
		query = query.Where(articleent.CreatedByEQ(*req.AuthorId))
	}
	if req.Type != nil {
		query = query.Where(articleent.TypeEQ(articleent.Type(*req.Type)))
	}
	if req.Keyword != nil {
		query = query.Where(
			articleent.Or(
				articleent.TitleContains(*req.Keyword),
				articleent.ContentContains(*req.Keyword),
			),
		)
	}
	if req.Order != nil {
		switch *req.Order {
		case enum.ArticleOrderNewest:
			query = query.Order(gen.Desc(articleent.FieldCreatedAt))
		case enum.ArticleOrderHottest:
			query = query.Where(articleent.CreatedAtGTE(time.Now().Add(-30 * 24 * time.Hour))).
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
		query = query.Order(gen.Desc(articleent.FieldCreatedAt))
	}
	return query
}
