package repo

import (
	"context"
	"fmt"

	"common/api/gen/common"
	cerrors "common/api/gen/common/errors"
	v1 "common/api/gen/content/v1"
	"common/pkg/constant"
	utilent "common/pkg/util/ent"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/gen"
	commentent "content/internal/data/gen/comment"
	"content/internal/enum"

	"entgo.io/ent/dialect/sql"
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
		SetStatus(commentent.StatusNormal).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.Comment{
		ID:         save.ID,
		ArticleID:  save.ArticleID,
		Content:    save.Content,
		Level:      save.Level,
		ParentID:   save.ParentID,
		ReplyID:    save.ReplyID,
		Status:     enum.CommentStatus(save.Status),
		ThankCount: save.ThankCount,
		LikeCount:  save.LikeCount,
		ReplyCount: save.ReplyCount,
		CreatedAt:  save.CreatedAt,
		UpdatedAt:  save.UpdatedAt,
		CreatedBy:  save.CreatedBy,
		UpdatedBy:  save.UpdatedBy,
	}, nil
}

func (r *CommentRepo) UpdateStatus(ctx context.Context, commentId int64, status v1.CommentStatus) error {
	dbStatus, _ := enum.CommentStatusMap.ToEnum(status)
	_, err := r.getClient(ctx).Comment.UpdateOneID(commentId).
		SetStatus(commentent.Status(dbStatus)).
		Save(ctx)
	return err
}

func (r *CommentRepo) UpdateStat(ctx context.Context, commentId int64, action v1.CommentAction, num int32) error {
	updateOne := r.getClient(ctx).Comment.UpdateOneID(commentId)
	switch action {
	case v1.CommentAction_COMMENT_ACTION_LIKE:
		updateOne.AddLikeCount(num)
	case v1.CommentAction_COMMENT_ACTION_THANK:
		updateOne.AddThankCount(num)
	case v1.CommentAction_COMMENT_ACTION_REPLY:
		updateOne.AddReplyCount(num)
	default:
		return fmt.Errorf("unknown action")
	}
	_, err := updateOne.Save(ctx)
	return err
}

func (r *CommentRepo) Exist(ctx context.Context, req *repo.CommentGetReq) (bool, error) {
	query := r.getClient(ctx).Comment.Query()
	query = r.getQuery(query, req)
	return query.Exist(ctx)
}

func (r *CommentRepo) Get(ctx context.Context, req *repo.CommentGetReq) (*model.Comment, error) {
	query := r.getClient(ctx).Comment.Query()
	query = r.getQuery(query, req)
	c, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, cerrors.ErrorBadRequest("comment is not found")
	}
	if err != nil {
		return nil, err
	}
	reply := &model.Comment{
		ID:          c.ID,
		ArticleID:   c.ArticleID,
		Content:     c.Content,
		Level:       c.Level,
		ParentID:    c.ParentID,
		ReplyID:     c.ReplyID,
		Status:      enum.CommentStatus(c.Status),
		ThankCount:  c.ThankCount,
		LikeCount:   c.LikeCount,
		ReplyCount:  c.ReplyCount,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
		CreatedBy:   c.CreatedBy,
		UpdatedBy:   c.UpdatedBy,
		WithArticle: req.WithArticle,
	}
	if c.Edges.Article != nil {
		reply.Article = &model.Article{
			ID:        c.Edges.Article.ID,
			Title:     c.Edges.Article.Title,
			CreatedBy: c.Edges.Article.CreatedBy,
		}
	}
	return reply, nil
}

func (r *CommentRepo) GetList(ctx context.Context, req *repo.CommentGetReq) ([]*model.Comment, error) {
	query := r.getClient(ctx).Comment.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	comments := make([]*model.Comment, 0, len(list))
	for i := range list {
		item := &model.Comment{
			ID:          list[i].ID,
			ArticleID:   list[i].ArticleID,
			Content:     list[i].Content,
			Level:       list[i].Level,
			ParentID:    list[i].ParentID,
			ReplyID:     list[i].ReplyID,
			Status:      enum.CommentStatus(list[i].Status),
			ThankCount:  list[i].ThankCount,
			LikeCount:   list[i].LikeCount,
			ReplyCount:  list[i].ReplyCount,
			CreatedAt:   list[i].CreatedAt,
			UpdatedAt:   list[i].UpdatedAt,
			CreatedBy:   list[i].CreatedBy,
			UpdatedBy:   list[i].UpdatedBy,
			WithArticle: req.WithArticle,
		}
		if list[i].Edges.Article != nil {
			item.Article = &model.Article{
				ID:        list[i].Edges.Article.ID,
				Title:     list[i].Edges.Article.Title,
				CreatedBy: list[i].Edges.Article.CreatedBy,
			}
		}
		comments = append(comments, item)
	}
	return comments, nil
}

func (r *CommentRepo) GetPage(ctx context.Context, page *common.PageRequest, req *repo.CommentGetReq) ([]*model.Comment, *common.PageReply, error) {
	page = constant.PageValid(page)
	query := r.getClient(ctx).Comment.Query().WithReply()
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
	comments := make([]*model.Comment, 0, len(list))
	for i := range list {
		item := &model.Comment{
			ID:          list[i].ID,
			ArticleID:   list[i].ArticleID,
			Content:     list[i].Content,
			Level:       list[i].Level,
			ParentID:    list[i].ParentID,
			ReplyID:     list[i].ReplyID,
			Status:      enum.CommentStatus(list[i].Status),
			ThankCount:  list[i].ThankCount,
			LikeCount:   list[i].LikeCount,
			ReplyCount:  list[i].ReplyCount,
			CreatedAt:   list[i].CreatedAt,
			UpdatedAt:   list[i].UpdatedAt,
			CreatedBy:   list[i].CreatedBy,
			UpdatedBy:   list[i].UpdatedBy,
			WithArticle: req.WithArticle,
		}
		if list[i].Edges.Reply != nil {
			item.Reply = &model.Comment{
				ID:        list[i].Edges.Reply.ID,
				CreatedBy: list[i].Edges.Reply.CreatedBy,
			}
		}
		if list[i].Edges.Article != nil {
			item.Article = &model.Article{
				ID:        list[i].Edges.Article.ID,
				Title:     list[i].Edges.Article.Title,
				CreatedBy: list[i].Edges.Article.CreatedBy,
			}
		}
		comments = append(comments, item)
	}
	return comments, &common.PageReply{
		Total: uint32(total),
		Page:  page.Page,
		Size:  page.Size,
	}, nil
}

func (r *CommentRepo) getQuery(query *gen.CommentQuery, req *repo.CommentGetReq) *gen.CommentQuery {
	if req.WithArticle {
		query = query.WithArticle()
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
	if req.Status != nil {
		dbStatus, _ := enum.CommentStatusMap.ToEnum(*req.Status)
		query = query.Where(commentent.StatusEQ(commentent.Status(dbStatus)))
	}
	if req.Level != nil {
		query = query.Where(commentent.LevelEQ(*req.Level))
	}
	if req.Order != nil {
		switch *req.Order {
		case v1.CommentOrder_COMMENT_ORDER_NEWEST:
			query = query.Order(gen.Desc(commentent.FieldCreatedAt))
		case v1.CommentOrder_COMMENT_ORDER_HOTTEST:
			query = query.
				Order(func(s *sql.Selector) {
					s.OrderExpr(sql.Expr(`
        (
            (reply_count * 6 + like_count * 4 + thank_count * 2)
            /
            pow((extract(epoch from (now() - created_at)) / 3600) + 1.5, 1.8)
        ) DESC`))
				})
		}

	} else {
		query = query.Order(gen.Desc(commentent.FieldCreatedAt))
	}
	return query
}

func (r *CommentRepo) GetArticleLastComment(ctx context.Context, req *repo.CommentGetReq) (*model.Comment, error) {
	if req.ArticleId == nil {
		return nil, cerrors.ErrorBadRequest("articleId is required")
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
		Status:      enum.CommentStatus(c.Status),
		ThankCount:  c.ThankCount,
		LikeCount:   c.LikeCount,
		ReplyCount:  c.ReplyCount,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
		CreatedBy:   c.CreatedBy,
		UpdatedBy:   c.UpdatedBy,
		WithArticle: req.WithArticle,
	}, nil
}

func (r *CommentRepo) GetArticleLastComments(ctx context.Context, req *repo.CommentGetReq) (map[int64]*model.Comment, error) {
	if len(req.ArticleIds) == 0 {
		return nil, cerrors.ErrorBadRequest("articleIds is required")
	}
	articleIdsAny := make([]any, len(req.ArticleIds))
	for i, v := range req.ArticleIds {
		articleIdsAny[i] = v
	}
	comments, err := r.getClient(ctx).Comment.Query().
		Where(func(s *sql.Selector) {
			sub := sql.Select(
				commentent.FieldArticleID,
				sql.As(sql.Max(commentent.FieldCreatedAt), "latest_time"),
			).
				From(sql.Table(commentent.Table)).
				Where(sql.EQ(commentent.FieldStatus, string(commentent.StatusNormal))).
				Where(sql.In(commentent.FieldArticleID, articleIdsAny...)).
				GroupBy(commentent.FieldArticleID)

			s.Join(sub).On(
				s.C(commentent.FieldArticleID), sub.C(commentent.FieldArticleID),
			).On(
				s.C(commentent.FieldCreatedAt), sub.C("latest_time"),
			)
		}).
		Where(commentent.StatusEQ(commentent.StatusNormal)).
		Where(commentent.ArticleIDIn(req.ArticleIds...)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	commentMap := make(map[int64]*model.Comment)
	for _, item := range comments {
		commentMap[item.ArticleID] = &model.Comment{
			ID:          item.ID,
			ArticleID:   item.ArticleID,
			Content:     item.Content,
			Level:       item.Level,
			ParentID:    item.ParentID,
			ReplyID:     item.ReplyID,
			Status:      enum.CommentStatus(item.Status),
			ThankCount:  item.ThankCount,
			LikeCount:   item.LikeCount,
			ReplyCount:  item.ReplyCount,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
			CreatedBy:   item.CreatedBy,
			UpdatedBy:   item.UpdatedBy,
			WithArticle: req.WithArticle,
		}
	}
	return commentMap, nil
}
