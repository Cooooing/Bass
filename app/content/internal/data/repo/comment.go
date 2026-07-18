package repo

import (
	cerrors "common/proto/gen/common/errors"
	"content/internal/biz/base"
	"context"

	"common/pkg/apperror"
	utilent "common/pkg/util/ent"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/gen"
	commentent "content/internal/data/gen/comment"
	"content/internal/enum"

	"entgo.io/ent/dialect/sql"
	"github.com/samber/lo"
)

var _ repo.CommentRepo = (*CommentRepo)(nil)

type CommentRepo struct {
	db *gen.Client
}

func NewCommentRepo(db *gen.Client) repo.CommentRepo {
	return &CommentRepo{db: db}
}

func (r *CommentRepo) getClient(ctx context.Context) *gen.Client {
	if tx, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return tx
	}
	return r.db
}

func (r *CommentRepo) Save(ctx context.Context, comment *model.Comment) (*model.Comment, error) {
	save, err := r.getClient(ctx).Comment.Create().
		SetArticleID(comment.ArticleID).
		SetContent(comment.Content).
		SetLevel(comment.Level).
		SetNillableParentID(comment.ParentID).
		SetNillableReplyID(comment.ReplyID).
		SetRestriction(commentent.Restriction(comment.Restriction)).
		SetNillableCreatedBy(comment.CreatedBy).
		SetNillableUpdatedBy(comment.UpdatedBy).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.Comment{
		ID:          save.ID,
		ArticleID:   save.ArticleID,
		Content:     save.Content,
		Level:       save.Level,
		ParentID:    save.ParentID,
		ReplyID:     save.ReplyID,
		Restriction: enum.ContentRestriction(save.Restriction),
		ThankCount:  save.ThankCount,
		LikeCount:   save.LikeCount,
		ReplyCount:  save.ReplyCount,
		CreatedAt:   save.CreatedAt,
		UpdatedAt:   save.UpdatedAt,
		CreatedBy:   save.CreatedBy,
		UpdatedBy:   save.UpdatedBy,
		DeletedAt:   save.DeletedAt,
	}, nil
}

func (r *CommentRepo) UpdateRestriction(ctx context.Context, req *repo.CommentUpdateRestrictionReq) error {
	commentId := req.CommentID
	restriction := req.Restriction
	updatedBy := req.UpdatedBy
	if err := r.getClient(ctx).Comment.UpdateOneID(commentId).
		SetRestriction(commentent.Restriction(restriction)).
		SetUpdatedBy(updatedBy).
		Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (r *CommentRepo) AddStats(ctx context.Context, req *repo.CommentAddStatsReq) error {
	commentId := req.CommentID
	stats := req.Stats
	updateOne := r.getClient(ctx).Comment.UpdateOneID(commentId)
	if stats.ThankCount != 0 {
		updateOne.AddThankCount(stats.ThankCount)
	}
	if stats.LikeCount != 0 {
		updateOne.AddLikeCount(stats.LikeCount)
	}
	if stats.ReplyCount != 0 {
		updateOne.AddReplyCount(stats.ReplyCount)
	}
	if err := updateOne.Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (r *CommentRepo) Exist(ctx context.Context, req *repo.CommentGetReq) (bool, error) {
	query := r.getClient(ctx).Comment.Query()
	query = r.getQuery(query, req)
	exist, err := query.Exist(ctx)
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (r *CommentRepo) Get(ctx context.Context, req *repo.CommentGetReq) (*model.Comment, error) {
	query := r.getClient(ctx).Comment.Query()
	query = r.getQuery(query, req)
	c, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_COMMENT_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return &model.Comment{
		ID:          c.ID,
		ArticleID:   c.ArticleID,
		Content:     c.Content,
		Level:       c.Level,
		ParentID:    c.ParentID,
		ReplyID:     c.ReplyID,
		Restriction: enum.ContentRestriction(c.Restriction),
		ThankCount:  c.ThankCount,
		LikeCount:   c.LikeCount,
		ReplyCount:  c.ReplyCount,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
		CreatedBy:   c.CreatedBy,
		UpdatedBy:   c.UpdatedBy,
		DeletedAt:   c.DeletedAt,
	}, nil
}

func (r *CommentRepo) List(ctx context.Context, req *repo.CommentGetReq) ([]*model.Comment, error) {
	query := r.getClient(ctx).Comment.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	return lo.Map(list, func(item *gen.Comment, _ int) *model.Comment {
		return &model.Comment{
			ID:          item.ID,
			ArticleID:   item.ArticleID,
			Content:     item.Content,
			Level:       item.Level,
			ParentID:    item.ParentID,
			ReplyID:     item.ReplyID,
			Restriction: enum.ContentRestriction(item.Restriction),
			ThankCount:  item.ThankCount,
			LikeCount:   item.LikeCount,
			ReplyCount:  item.ReplyCount,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
			CreatedBy:   item.CreatedBy,
			UpdatedBy:   item.UpdatedBy,
			DeletedAt:   item.DeletedAt,
		}
	}), nil
}

func (r *CommentRepo) Map(ctx context.Context, req *repo.CommentGetReq) (map[int64]*model.Comment, error) {
	list, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	return lo.SliceToMap(list, func(item *model.Comment) (int64, *model.Comment) {
		return item.ID, item
	}), nil
}

func (r *CommentRepo) Count(ctx context.Context, req *repo.CommentGetReq) (int, error) {
	query := r.getClient(ctx).Comment.Query()
	query = r.getQuery(query, req)
	count, err := query.Count(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *CommentRepo) Page(ctx context.Context, req *repo.CommentGetReq) (*repo.CommentPageResp, error) {
	page := normalizePage(req.Page)
	query := r.getClient(ctx).Comment.Query().WithReply(func(q *gen.CommentQuery) {
		q.Where(commentent.DeletedAtIsNil())
	})
	query = r.getQuery(query, req)
	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}
	list, err := query.Limit(int(page.Size)).Offset(int((page.Page - 1) * page.Size)).All(ctx)
	if err != nil {
		return nil, err
	}
	comments := lo.Map(list, func(item *gen.Comment, _ int) *model.Comment {
		out := &model.Comment{
			ID:          item.ID,
			ArticleID:   item.ArticleID,
			Content:     item.Content,
			Level:       item.Level,
			ParentID:    item.ParentID,
			ReplyID:     item.ReplyID,
			Restriction: enum.ContentRestriction(item.Restriction),
			ThankCount:  item.ThankCount,
			LikeCount:   item.LikeCount,
			ReplyCount:  item.ReplyCount,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
			CreatedBy:   item.CreatedBy,
			UpdatedBy:   item.UpdatedBy,
			DeletedAt:   item.DeletedAt,
		}
		if item.Edges.Reply != nil {
			out.ReplyUserID = item.Edges.Reply.CreatedBy
		}
		return out
	})
	return &repo.CommentPageResp{
		Rows: comments,
		Page: &base.PageResp{
			Total: int64(total),
			Page:  page.Page,
			Size:  page.Size,
		},
	}, nil
}

func (r *CommentRepo) ListReplyPreviews(ctx context.Context, req *repo.CommentReplyPreviewReq) ([]*repo.CommentReplyPreview, error) {
	parentIds := lo.Uniq(req.ParentIds)
	if len(parentIds) == 0 {
		return nil, nil
	}
	limit := int(req.LimitPerParent)
	if limit <= 0 {
		limit = 3
	}
	if limit > 5 {
		limit = 5
	}
	replyMap := make(map[int64][]*model.Comment, len(parentIds))
	for _, parentId := range parentIds {
		query := r.getClient(ctx).Comment.Query().WithReply(func(q *gen.CommentQuery) {
			q.Where(commentent.DeletedAtIsNil())
		}).Limit(limit)
		query = r.getQuery(query, &repo.CommentGetReq{
			ArticleId:    new(req.ArticleId),
			ParentId:     new(parentId),
			Restriction:  req.Restriction,
			Restrictions: req.Restrictions,
			Order:        req.Order,
		})
		list, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		replyMap[parentId] = lo.Map(list, func(item *gen.Comment, _ int) *model.Comment {
			out := &model.Comment{
				ID:          item.ID,
				ArticleID:   item.ArticleID,
				Content:     item.Content,
				Level:       item.Level,
				ParentID:    item.ParentID,
				ReplyID:     item.ReplyID,
				Restriction: enum.ContentRestriction(item.Restriction),
				ThankCount:  item.ThankCount,
				LikeCount:   item.LikeCount,
				ReplyCount:  item.ReplyCount,
				CreatedAt:   item.CreatedAt,
				UpdatedAt:   item.UpdatedAt,
				CreatedBy:   item.CreatedBy,
				UpdatedBy:   item.UpdatedBy,
				DeletedAt:   item.DeletedAt,
			}
			if item.Edges.Reply != nil {
				out.ReplyUserID = item.Edges.Reply.CreatedBy
			}
			return out
		})
	}
	return lo.Map(parentIds, func(parentId int64, _ int) *repo.CommentReplyPreview {
		return &repo.CommentReplyPreview{
			ParentId: parentId,
			Rows:     replyMap[parentId],
		}
	}), nil
}

func (r *CommentRepo) getQuery(query *gen.CommentQuery, req *repo.CommentGetReq) *gen.CommentQuery {
	query = query.Where(commentent.DeletedAtIsNil())
	if req == nil {
		req = &repo.CommentGetReq{}
	}
	if req.ParentId != nil {
		query = query.Where(commentent.ParentIDEQ(*req.ParentId))
	}
	if req.ReplyId != nil {
		query = query.Where(commentent.ReplyIDEQ(*req.ReplyId))
	}
	if req.CommentId != nil {
		query = query.Where(commentent.IDEQ(*req.CommentId))
	}
	if len(req.CommentIds) > 0 {
		query = query.Where(commentent.IDIn(req.CommentIds...))
	}
	if req.ArticleId != nil {
		query = query.Where(commentent.ArticleIDEQ(*req.ArticleId))
	}
	if len(req.ArticleIds) > 0 {
		query = query.Where(commentent.ArticleIDIn(req.ArticleIds...))
	}
	if req.CreatedBy != nil {
		query = query.Where(commentent.CreatedByEQ(*req.CreatedBy))
	}
	if req.Restriction != nil {
		query = query.Where(commentent.RestrictionEQ(commentent.Restriction(*req.Restriction)))
	}
	if len(req.Restrictions) > 0 {
		query = query.Where(commentent.RestrictionIn(lo.Map(req.Restrictions, func(item enum.ContentRestriction, _ int) commentent.Restriction {
			return commentent.Restriction(item)
		})...))
	}
	if req.Level != nil {
		query = query.Where(commentent.LevelEQ(*req.Level))
	}
	if req.Order != nil {
		switch *req.Order {
		case enum.CommentOrderNewest:
			query = query.Order(gen.Desc(commentent.FieldCreatedAt, commentent.FieldID))
		case enum.CommentOrderOldest:
			query = query.Order(gen.Asc(commentent.FieldCreatedAt, commentent.FieldID))
		case enum.CommentOrderHottest:
			query = query.
				Order(func(s *sql.Selector) {
					s.OrderExpr(sql.Expr(`
        (
            (reply_count * 6 + like_count * 4 + thank_count * 2)
            /
            pow((extract(epoch from (now() - created_at)) / 3600) + 1.5, 1.8)
        ) DESC`))
				}, gen.Desc(commentent.FieldID))
		}
	} else {
		query = query.Order(gen.Desc(commentent.FieldCreatedAt, commentent.FieldID))
	}
	return query
}

func (r *CommentRepo) GetArticleLastComment(ctx context.Context, req *repo.CommentGetReq) (*model.Comment, error) {
	if req.ArticleId == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	query := r.getClient(ctx).Comment.Query()
	query = r.getQuery(query, req)
	c, err := query.Order(gen.Desc(commentent.FieldCreatedAt)).First(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &model.Comment{
		ID:          c.ID,
		ArticleID:   c.ArticleID,
		Content:     c.Content,
		Level:       c.Level,
		ParentID:    c.ParentID,
		ReplyID:     c.ReplyID,
		Restriction: enum.ContentRestriction(c.Restriction),
		ThankCount:  c.ThankCount,
		LikeCount:   c.LikeCount,
		ReplyCount:  c.ReplyCount,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
		CreatedBy:   c.CreatedBy,
		UpdatedBy:   c.UpdatedBy,
		DeletedAt:   c.DeletedAt,
	}, nil
}

func (r *CommentRepo) MapArticleLastComments(ctx context.Context, req *repo.CommentGetReq) (map[int64]*model.Comment, error) {
	if len(req.ArticleIds) == 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	articleIdsAny := lo.Map(req.ArticleIds, func(item int64, _ int) any {
		return item
	})
	commentTable := sql.Table(commentent.Table)
	query := r.getClient(ctx).Comment.Query().
		Where(func(s *sql.Selector) {
			sub := sql.Select(
				commentTable.C(commentent.FieldArticleID),
				sql.As(sql.Max(commentTable.C(commentent.FieldCreatedAt)), "latest_time"),
			).
				From(commentTable).
				Where(sql.In(commentTable.C(commentent.FieldArticleID), articleIdsAny...)).
				Where(sql.IsNull(commentTable.C(commentent.FieldDeletedAt))).
				GroupBy(commentTable.C(commentent.FieldArticleID))
			if req.Restriction != nil {
				sub = sub.Where(sql.EQ(commentTable.C(commentent.FieldRestriction), string(commentent.Restriction(*req.Restriction))))
			}
			if len(req.Restrictions) > 0 {
				restrictions := lo.Map(req.Restrictions, func(item enum.ContentRestriction, _ int) any {
					return string(commentent.Restriction(item))
				})
				sub = sub.Where(sql.In(commentTable.C(commentent.FieldRestriction), restrictions...))
			}
			s.Join(sub).On(
				s.C(commentent.FieldArticleID), sub.C(commentent.FieldArticleID),
			).On(
				s.C(commentent.FieldCreatedAt), sub.C("latest_time"),
			)
		}).
		Where(commentent.ArticleIDIn(req.ArticleIds...), commentent.DeletedAtIsNil())
	if req.Restriction != nil {
		query = query.Where(commentent.RestrictionEQ(commentent.Restriction(*req.Restriction)))
	}
	if len(req.Restrictions) > 0 {
		query = query.Where(commentent.RestrictionIn(lo.Map(req.Restrictions, func(item enum.ContentRestriction, _ int) commentent.Restriction {
			return commentent.Restriction(item)
		})...))
	}
	comments, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	return lo.SliceToMap(comments, func(item *gen.Comment) (int64, *model.Comment) {
		return item.ArticleID, &model.Comment{
			ID:          item.ID,
			ArticleID:   item.ArticleID,
			Content:     item.Content,
			Level:       item.Level,
			ParentID:    item.ParentID,
			ReplyID:     item.ReplyID,
			Restriction: enum.ContentRestriction(item.Restriction),
			ThankCount:  item.ThankCount,
			LikeCount:   item.LikeCount,
			ReplyCount:  item.ReplyCount,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
			CreatedBy:   item.CreatedBy,
			UpdatedBy:   item.UpdatedBy,
			DeletedAt:   item.DeletedAt,
		}
	}), nil
}
