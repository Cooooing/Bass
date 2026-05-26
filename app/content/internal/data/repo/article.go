package repo

import (
	"context"
	"fmt"
	"time"

	"common/api/gen/common"
	cerrors "common/api/gen/common/errors"
	v1 "common/api/gen/content/v1"
	"common/pkg/constant"
	utilent "common/pkg/util/ent"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/gen"
	articleent "content/internal/data/gen/article"
	"content/internal/data/gen/articleactionrecord"
	"content/internal/data/gen/articlepostscript"
	tagent "content/internal/data/gen/tag"
	"content/internal/enum"

	"entgo.io/ent/dialect/sql"
	"github.com/samber/lo"
)

var _ repo.ArticleRepo = (*ArticleRepo)(nil)

type ArticleRepo struct {
	db          *gen.Client
	commentRepo repo.CommentRepo
	tagRepo     repo.TagRepo
}

func NewArticleRepo(
	db *gen.Client,
	commentRepo repo.CommentRepo,
	tagRepo repo.TagRepo,
) repo.ArticleRepo {
	return &ArticleRepo{
		db:          db,
		commentRepo: commentRepo,
		tagRepo:     tagRepo,
	}
}

func (r *ArticleRepo) getClient(ctx context.Context) *gen.Client {
	if tx, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return tx
	}
	return r.db
}

func (r *ArticleRepo) resolveArticleTagIDs(ctx context.Context, tags []*model.Tag) ([]int64, error) {
	tagIDSet := make(map[int64]struct{})
	tagByName := make(map[string]*model.Tag)
	for _, item := range tags {
		if item == nil {
			continue
		}
		if item.ID > 0 {
			tagIDSet[item.ID] = struct{}{}
			continue
		}
		if item.Name != "" {
			tagByName[item.Name] = item
		}
	}

	bindTagIDSet := make(map[int64]struct{})
	if len(tagIDSet) > 0 {
		existingTags, err := r.tagRepo.GetList(ctx, &repo.TagGetReq{TagIds: lo.Keys(tagIDSet)})
		if err != nil {
			return nil, err
		}
		if len(existingTags) != len(tagIDSet) {
			return nil, cerrors.ErrorBadRequest("tag not exist")
		}
		for _, item := range existingTags {
			bindTagIDSet[item.ID] = struct{}{}
		}
	}

	if len(tagByName) > 0 {
		existingTags, err := r.tagRepo.GetList(ctx, &repo.TagGetReq{Names: lo.Keys(tagByName)})
		if err != nil {
			return nil, err
		}
		for _, item := range existingTags {
			bindTagIDSet[item.ID] = struct{}{}
			delete(tagByName, item.Name)
		}

		saveTags := make([]*model.Tag, 0, len(tagByName))
		for _, item := range tagByName {
			saveTags = append(saveTags, item)
		}
		if len(saveTags) > 0 {
			savedTags, err := r.tagRepo.Saves(ctx, saveTags)
			if err != nil {
				return nil, err
			}
			for _, item := range savedTags {
				bindTagIDSet[item.ID] = struct{}{}
			}
		}
	}

	return lo.Keys(bindTagIDSet), nil
}

func (r *ArticleRepo) Save(ctx context.Context, article *model.Article, tags []*model.Tag) (*model.Article, error) {
	var bindTagIds []int64
	if len(tags) > 0 {
		var err error
		bindTagIds, err = r.resolveArticleTagIDs(ctx, tags)
		if err != nil {
			return nil, err
		}
	}

	article.FormatContent()
	create := r.getClient(ctx).Article.Create().
		SetTitle(article.Title).
		SetContent(article.Content).
		SetNillableRewardContent(article.RewardContent).
		SetNillableRewardPoints(article.RewardPoints).
		SetStatus(articleent.Status(article.Status)).
		SetType(articleent.Type(article.Type)).
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
	return &model.Article{
		ID:               save.ID,
		Title:            save.Title,
		Content:          save.Content,
		HasPostscript:    save.HasPostscript,
		RewardContent:    save.RewardContent,
		RewardPoints:     save.RewardPoints,
		Status:           enum.ArticleStatus(save.Status),
		Type:             enum.ArticleType(save.Type),
		Statement:        save.Statement,
		Commentable:      save.Commentable,
		Anonymous:        save.Anonymous,
		Listable:         save.Listable,
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
	}, nil
}

func (r *ArticleRepo) Update(ctx context.Context, updateArticle *model.Article, tags []*model.Tag) (*model.Article, error) {
	var bindTagIds []int64
	if len(tags) > 0 {
		var err error
		bindTagIds, err = r.resolveArticleTagIDs(ctx, tags)
		if err != nil {
			return nil, err
		}
	}

	updateArticle.FormatContent()
	update := r.getClient(ctx).Article.UpdateOneID(updateArticle.ID).
		SetTitle(updateArticle.Title).
		SetContent(updateArticle.Content).
		SetNillableRewardContent(updateArticle.RewardContent).
		SetNillableRewardPoints(updateArticle.RewardPoints).
		SetStatus(articleent.Status(updateArticle.Status)).
		SetType(articleent.Type(updateArticle.Type)).
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
	return &model.Article{
		ID:               save.ID,
		Title:            save.Title,
		Content:          save.Content,
		HasPostscript:    save.HasPostscript,
		RewardContent:    save.RewardContent,
		RewardPoints:     save.RewardPoints,
		Status:           enum.ArticleStatus(save.Status),
		Type:             enum.ArticleType(save.Type),
		Statement:        save.Statement,
		Commentable:      save.Commentable,
		Anonymous:        save.Anonymous,
		Listable:         save.Listable,
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
	}, nil
}

func (r *ArticleRepo) UpdateContent(ctx context.Context, articleId int64, content string) error {
	return r.getClient(ctx).Article.UpdateOneID(articleId).
		SetContent(content).
		Exec(ctx)
}

func (r *ArticleRepo) UpdateStatus(ctx context.Context, articleId int64, status v1.ArticleStatus) error {
	dbStatus, _ := enum.ArticleStatusMap.ToEnum(status)
	return r.getClient(ctx).Article.UpdateOneID(articleId).
		SetStatus(articleent.Status(dbStatus)).
		Exec(ctx)
}

func (r *ArticleRepo) UpdateControlFields(ctx context.Context, articleId int64, status v1.ArticleStatus, commentable bool, anonymous bool, listable *bool) error {
	dbStatus, _ := enum.ArticleStatusMap.ToEnum(status)
	update := r.getClient(ctx).Article.UpdateOneID(articleId).
		SetStatus(articleent.Status(dbStatus)).
		SetCommentable(commentable).
		SetAnonymous(anonymous)
	if listable != nil {
		update.SetListable(*listable)
	}
	return update.Exec(ctx)
}

func (r *ArticleRepo) UpdateHasPostscript(ctx context.Context, articleId int64, hasPostscript bool) error {
	return r.getClient(ctx).Article.UpdateOneID(articleId).
		SetHasPostscript(hasPostscript).
		Exec(ctx)
}

func (r *ArticleRepo) UpdateStat(ctx context.Context, articleId int64, action v1.ArticleAction, num int32) (*model.Article, error) {
	updateOne := r.getClient(ctx).Article.UpdateOneID(articleId)
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
	default:
		return nil, fmt.Errorf("unknown action")
	}
	save, err := updateOne.Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.Article{
		ID:               save.ID,
		Title:            save.Title,
		Content:          save.Content,
		HasPostscript:    save.HasPostscript,
		RewardContent:    save.RewardContent,
		RewardPoints:     save.RewardPoints,
		Status:           enum.ArticleStatus(save.Status),
		Type:             enum.ArticleType(save.Type),
		Statement:        save.Statement,
		Commentable:      save.Commentable,
		Anonymous:        save.Anonymous,
		Listable:         save.Listable,
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
	}, nil
}

func (r *ArticleRepo) UpdateAcceptAnswer(ctx context.Context, articleId int64, commentId int64) (*model.Article, error) {
	exist, err := r.commentRepo.Exist(ctx, &repo.CommentGetReq{CommentId: new(commentId), ArticleId: new(articleId)})
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, fmt.Errorf("comment not exist")
	}
	a, err := r.getClient(ctx).Article.UpdateOneID(articleId).
		SetAcceptedAnswerID(commentId).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.Article{
		ID:               a.ID,
		Title:            a.Title,
		Content:          a.Content,
		HasPostscript:    a.HasPostscript,
		RewardContent:    a.RewardContent,
		RewardPoints:     a.RewardPoints,
		Status:           enum.ArticleStatus(a.Status),
		Type:             enum.ArticleType(a.Type),
		Statement:        a.Statement,
		Commentable:      a.Commentable,
		Anonymous:        a.Anonymous,
		Listable:         a.Listable,
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
	}, nil
}

func (r *ArticleRepo) Publish(ctx context.Context, articleId int64) (*model.Article, error) {
	first, err := r.Get(ctx, &repo.ArticleGetReq{ArticleId: new(articleId)})
	if err != nil {
		return nil, err
	}

	if first.Status != enum.ArticleStatusDrafts {
		return nil, cerrors.ErrorBadRequest("only update draft")
	}

	err = r.UpdateStatus(ctx, articleId, v1.ArticleStatus_ARTICLE_STATUS_NORMAL)
	if err != nil {
		return nil, err
	}
	return first, nil
}

func (r *ArticleRepo) Delete(ctx context.Context, articleId int64) error {
	return r.getClient(ctx).Article.UpdateOneID(articleId).SetStatus(articleent.StatusDeleted).Exec(ctx)
}

func (r *ArticleRepo) Exist(ctx context.Context, req *repo.ArticleGetReq) (bool, error) {
	query := r.getClient(ctx).Article.Query()
	query = r.getQuery(query, req)
	return query.Exist(ctx)
}

func (r *ArticleRepo) Get(ctx context.Context, req *repo.ArticleGetReq) (*model.Article, error) {
	query := r.getClient(ctx).Article.Query().
		WithPostscripts(func(q *gen.ArticlePostscriptQuery) {
			q.Where(articlepostscript.StatusEQ(articlepostscript.StatusNormal)).
				Order(gen.Asc(articlepostscript.FieldCreatedAt))
		}).
		WithTags()
	if req.QueryUserId != nil {
		query = query.WithActionRecords(func(query *gen.ArticleActionRecordQuery) {
			query.Where(articleactionrecord.UserIDEQ(*req.QueryUserId))
		})
	}
	query = r.getQuery(query, req)
	a, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, cerrors.ErrorBadRequest("article is not found")
	}
	if err != nil {
		return nil, err
	}
	result := &model.Article{
		ID:               a.ID,
		Title:            a.Title,
		Content:          a.Content,
		HasPostscript:    a.HasPostscript,
		RewardContent:    a.RewardContent,
		RewardPoints:     a.RewardPoints,
		Status:           enum.ArticleStatus(a.Status),
		Type:             enum.ArticleType(a.Type),
		Statement:        a.Statement,
		Commentable:      a.Commentable,
		Anonymous:        a.Anonymous,
		Listable:         a.Listable,
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
		IsSummary:        req.IsSummary,
	}
	for _, item := range a.Edges.Postscripts {
		result.Postscripts = append(result.Postscripts, &model.ArticlePostscript{
			ID:        item.ID,
			ArticleID: item.ArticleID,
			Content:   item.Content,
			Status:    enum.ArticlePostscriptStatus(item.Status),
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
			CreatedBy: item.CreatedBy,
			UpdatedBy: item.UpdatedBy,
		})
	}
	for _, item := range a.Edges.Tags {
		result.Tags = append(result.Tags, &model.Tag{
			ID:          item.ID,
			Name:        item.Name,
			Description: item.Description,
			DomainID:    item.DomainID,
			Status:      enum.TagStatus(item.Status),
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
			CreatedBy:   item.CreatedBy,
			UpdatedBy:   item.UpdatedBy,
		})
	}
	for _, item := range a.Edges.ActionRecords {
		result.ActionRecords = append(result.ActionRecords, &model.ArticleActionRecord{
			ID:        item.ID,
			ArticleID: item.ArticleID,
			UserID:    item.UserID,
			Type:      enum.ArticleAction(item.Type),
		})
	}
	return result, nil
}

func (r *ArticleRepo) GetList(ctx context.Context, req *repo.ArticleGetReq) ([]*model.Article, error) {
	query := r.getClient(ctx).Article.Query().WithTags()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	articles := make([]*model.Article, 0, len(list))
	for i := range list {
		item := &model.Article{
			ID:               list[i].ID,
			Title:            list[i].Title,
			Content:          list[i].Content,
			HasPostscript:    list[i].HasPostscript,
			RewardContent:    list[i].RewardContent,
			RewardPoints:     list[i].RewardPoints,
			Status:           enum.ArticleStatus(list[i].Status),
			Type:             enum.ArticleType(list[i].Type),
			Statement:        list[i].Statement,
			Commentable:      list[i].Commentable,
			Anonymous:        list[i].Anonymous,
			Listable:         list[i].Listable,
			ViewCount:        list[i].ViewCount,
			ThankCount:       list[i].ThankCount,
			LikeCount:        list[i].LikeCount,
			CollectCount:     list[i].CollectCount,
			WatchCount:       list[i].WatchCount,
			ReplyCount:       list[i].ReplyCount,
			BountyPoints:     list[i].BountyPoints,
			AcceptedAnswerID: list[i].AcceptedAnswerID,
			CreatedAt:        list[i].CreatedAt,
			UpdatedAt:        list[i].UpdatedAt,
			CreatedBy:        list[i].CreatedBy,
			UpdatedBy:        list[i].UpdatedBy,
			IsSummary:        req.IsSummary,
		}
		for _, tag := range list[i].Edges.Tags {
			item.Tags = append(item.Tags, &model.Tag{
				ID:          tag.ID,
				Name:        tag.Name,
				Description: tag.Description,
				DomainID:    tag.DomainID,
				Status:      enum.TagStatus(tag.Status),
				CreatedAt:   tag.CreatedAt,
				UpdatedAt:   tag.UpdatedAt,
				CreatedBy:   tag.CreatedBy,
				UpdatedBy:   tag.UpdatedBy,
			})
		}
		articles = append(articles, item)
	}
	return articles, nil
}

func (r *ArticleRepo) GetPage(ctx context.Context, page *common.PageRequest, req *repo.ArticleGetReq) ([]*model.Article, *common.PageReply, error) {
	page = constant.PageValid(page)
	query := r.getClient(ctx).Article.Query().WithTags()
	query = r.getQuery(query, req)
	countQuery := query.Clone()
	total, err := countQuery.Count(ctx)
	if err != nil {
		return nil, nil, err
	}
	list, err := query.Limit(int(page.Size)).Offset(int((page.Page - 1) * page.Size)).All(ctx)
	if err != nil {
		return nil, nil, err
	}
	articles := make([]*model.Article, 0, len(list))
	for i := range list {
		item := &model.Article{
			ID:               list[i].ID,
			Title:            list[i].Title,
			Content:          list[i].Content,
			HasPostscript:    list[i].HasPostscript,
			RewardContent:    list[i].RewardContent,
			RewardPoints:     list[i].RewardPoints,
			Status:           enum.ArticleStatus(list[i].Status),
			Type:             enum.ArticleType(list[i].Type),
			Statement:        list[i].Statement,
			Commentable:      list[i].Commentable,
			Anonymous:        list[i].Anonymous,
			Listable:         list[i].Listable,
			ViewCount:        list[i].ViewCount,
			ThankCount:       list[i].ThankCount,
			LikeCount:        list[i].LikeCount,
			CollectCount:     list[i].CollectCount,
			WatchCount:       list[i].WatchCount,
			ReplyCount:       list[i].ReplyCount,
			BountyPoints:     list[i].BountyPoints,
			AcceptedAnswerID: list[i].AcceptedAnswerID,
			CreatedAt:        list[i].CreatedAt,
			UpdatedAt:        list[i].UpdatedAt,
			CreatedBy:        list[i].CreatedBy,
			UpdatedBy:        list[i].UpdatedBy,
			IsSummary:        req.IsSummary,
		}
		for _, tag := range list[i].Edges.Tags {
			item.Tags = append(item.Tags, &model.Tag{
				ID:          tag.ID,
				Name:        tag.Name,
				Description: tag.Description,
				DomainID:    tag.DomainID,
				Status:      enum.TagStatus(tag.Status),
				CreatedAt:   tag.CreatedAt,
				UpdatedAt:   tag.UpdatedAt,
				CreatedBy:   tag.CreatedBy,
				UpdatedBy:   tag.UpdatedBy,
			})
		}
		articles = append(articles, item)
	}
	return articles, &common.PageReply{
		Total: uint32(total),
		Page:  page.Page,
		Size:  page.Size,
	}, nil
}

func (r *ArticleRepo) getQuery(query *gen.ArticleQuery, req *repo.ArticleGetReq) *gen.ArticleQuery {
	if req.ArticleId != nil {
		query = query.Where(articleent.IDEQ(*req.ArticleId))
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
	if req.Status != nil {
		dbStatus, _ := enum.ArticleStatusMap.ToEnum(*req.Status)
		query = query.Where(articleent.StatusEQ(articleent.Status(dbStatus)))
	}
	if req.AuthorId != nil {
		query = query.Where(articleent.CreatedByEQ(*req.AuthorId))
	}
	if req.Type != nil {
		dbType, _ := enum.ArticleTypeMap.ToEnum(*req.Type)
		query = query.Where(articleent.TypeEQ(articleent.Type(dbType)))
	}
	if req.Keyword != nil {
		query = query.Where(
			articleent.Or(
				articleent.TitleContains(*req.Keyword),
				articleent.ContentContains(*req.Keyword),
			),
		)
	}
	if req.Listable != nil {
		query = query.Where(articleent.ListableEQ(*req.Listable))
	}
	if req.Order != nil {
		switch *req.Order {
		case v1.ArticleOrder_ARTICLE_ORDER_NEWEST:
			query = query.Order(gen.Desc(articleent.FieldCreatedAt))
		case v1.ArticleOrder_ARTICLE_ORDER_HOTTEST:
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
