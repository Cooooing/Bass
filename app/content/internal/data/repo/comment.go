package repo

import (
	cv1 "common/gen/common/v1"
	v1 "common/gen/content/v1"
	"common/pkg/constant"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	basedata "content/internal/data/base"
	"content/internal/data/ent/gen"
	"content/internal/data/ent/gen/comment"
	"context"

	"entgo.io/ent/dialect/sql"
)

type CommentRepo struct {
	*basedata.BaseData
	client *gen.Client
}

func NewCommentRepo(BaseData *basedata.BaseData, client *gen.Client) repo.CommentRepo {
	return &CommentRepo{
		BaseData: BaseData,
		client:   client,
	}
}

func (r *CommentRepo) Save(ctx context.Context, client *gen.Client, comment *model.Comment) (*model.Comment, error) {
	save, err := client.Comment.Create().
		SetArticleID(comment.ArticleID).
		SetContent(comment.Content).
		SetLevel(comment.Level).
		SetNillableParentID(comment.ParentID).
		SetNillableReplyID(comment.ReplyID).
		SetStatus(0).
		Save(ctx)
	return &model.Comment{Comment: save}, err
}

func (r *CommentRepo) UpdateStatus(ctx context.Context, client *gen.Client, commentId int64, status v1.CommentStatus) error {
	_, err := client.Comment.UpdateOneID(commentId).
		SetStatus(int32(status)).
		Save(ctx)
	return err
}

func (r *CommentRepo) UpdateStat(ctx context.Context, client *gen.Client, commentId int64, action v1.CommentAction, num int32) error {
	updateOne := client.Comment.UpdateOneID(commentId)
	switch action {
	case v1.CommentAction_COMMENT_ACTION_LIKE:
		updateOne.AddLikeCount(num)
	case v1.CommentAction_COMMENT_ACTION_COLLECT:
		updateOne.AddCollectCount(num)
	case v1.CommentAction_COMMENT_ACTION_REPLY:
		updateOne.AddReplyCount(num)
	}
	_, err := updateOne.Save(ctx)
	return err
}

func (r *CommentRepo) Exist(ctx context.Context, tx *gen.Client, req *repo.CommentGetReq) (bool, error) {
	query := tx.Comment.Query()
	query = r.getQuery(query, req)
	exist, err := query.Exist(ctx)
	return exist, err
}

func (r *CommentRepo) GetOne(ctx context.Context, tx *gen.Client, req *repo.CommentGetReq) (*model.Comment, error) {
	query := tx.Comment.Query()
	query = r.getQuery(query, req)
	c, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, cv1.ErrorBadRequest("comment is not found")
	}
	return &model.Comment{Comment: c, WithArticle: req.WithArticle}, err
}

func (r *CommentRepo) GetList(ctx context.Context, tx *gen.Client, req *repo.CommentGetReq) ([]*model.Comment, error) {
	var (
		comments []*model.Comment
		err      error
	)
	query := tx.Comment.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	for i := range list {
		comments = append(comments, &model.Comment{Comment: list[i], WithArticle: req.WithArticle})
	}
	return comments, nil
}

func (r *CommentRepo) GetPage(ctx context.Context, tx *gen.Client, page *cv1.PageRequest, req *repo.CommentGetReq) ([]*model.Comment, *cv1.PageReply, error) {
	var (
		comments []*model.Comment
		err      error
		total    int
	)
	page = constant.PageValid(page)
	query := tx.Comment.Query().WithReply()
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
		comments = append(comments, &model.Comment{Comment: list[i], WithArticle: req.WithArticle})
	}
	return comments, &cv1.PageReply{
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
		query = query.Where(comment.ParentIDEQ(*req.ParentId))
	}
	if req.ReplyId != nil {
		query = query.Where(comment.ReplyIDEQ(*req.ReplyId))
	}
	if req.CommentId != nil {
		query = query.Where(comment.IDEQ(*req.CommentId))
	}
	if len(req.CommentIds) > 0 {
		query = query.Where(comment.IDIn(req.CommentIds...))
	}
	if req.ArticleId != nil {
		query = query.Where(comment.ArticleIDEQ(*req.ArticleId))
	}
	if len(req.ArticleIds) > 0 {
		query = query.Where(comment.ArticleIDIn(req.ArticleIds...))
	}
	if req.CreatedBy != nil {
		query = query.Where(comment.CreatedByEQ(*req.CreatedBy))
	}
	if req.Status != nil {
		query = query.Where(comment.StatusEQ(int32(*req.Status)))
	}
	if req.Level != nil {
		query = query.Where(comment.LevelEQ(*req.Level))
	}
	if req.Order != nil {
		switch *req.Order {
		case int32(v1.CommentOrder_COMMENT_ORDER_NEWEST):
			query = query.Order(gen.Desc(comment.FieldCreatedAt))
		case int32(v1.CommentOrder_COMMENT_ORDER_HOTTEST):
			query = query.
				Order(func(s *sql.Selector) {
					s.OrderExpr(sql.Expr(`
        (
            (reply_count * 6 + like_count * 4 + collect_count * 1)
            /
            pow((extract(epoch from (now() - created_at)) / 3600) + 1.5, 1.8)
        ) DESC`))
				})
		}

	} else {
		query = query.Order(gen.Desc(comment.FieldCreatedAt))
	}
	return query
}

func (r *CommentRepo) GetArticleLastComment(ctx context.Context, client *gen.Client, req *repo.CommentGetReq) (*model.Comment, error) {
	if req.ArticleId == nil {
		return nil, cv1.ErrorBadRequest("articleId is required")
	}
	query := client.Comment.Query()
	query = r.getQuery(query, req)
	c, err := query.Order(gen.Desc(comment.FieldCreatedAt)).First(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	return &model.Comment{Comment: c, WithArticle: req.WithArticle}, err
}

func (r *CommentRepo) GetArticleLastComments(ctx context.Context, tx *gen.Client, req *repo.CommentGetReq) (map[int64]*model.Comment, error) {
	if len(req.ArticleIds) == 0 {
		return nil, cv1.ErrorBadRequest("articleIds is required")
	}
	articleIdsAny := make([]any, len(req.ArticleIds))
	for i, v := range req.ArticleIds {
		articleIdsAny[i] = v
	}
	comments, err := tx.Comment.Query().
		Where(func(s *sql.Selector) {
			// 子查询 SELECT article_id, MAX(created_at)
			sub := sql.Select(
				comment.FieldArticleID,
				sql.As(sql.Max(comment.FieldCreatedAt), "latest_time"),
			).
				From(sql.Table(comment.Table)).
				Where(sql.EQ(comment.FieldStatus, int32(v1.CommentStatus_COMMENT_STATUS_NORMAL))).
				Where(sql.In(comment.FieldArticleID, articleIdsAny...)).
				GroupBy(comment.FieldArticleID)

			// JOIN 子查询
			s.Join(sub).On(
				s.C(comment.FieldArticleID), sub.C(comment.FieldArticleID),
			).On(
				s.C(comment.FieldCreatedAt), sub.C("latest_time"),
			)
		}).
		Where(comment.StatusEQ(int32(v1.CommentStatus_COMMENT_STATUS_NORMAL))).
		Where(comment.ArticleIDIn(req.ArticleIds...)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	commentMap := make(map[int64]*model.Comment)
	for _, item := range comments {
		commentMap[item.ArticleID] = &model.Comment{Comment: item, WithArticle: req.WithArticle}
	}
	return commentMap, err
}
