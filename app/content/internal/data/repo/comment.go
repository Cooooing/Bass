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
	commentent "content/internal/data/gen/comment"
	"content/internal/enum"
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
)

type CommentRepo struct {
	conf   *conf.Bootstrap
	log    *log.Helper
	consul *commonClient.ConsulClient
	redis  *commonClient.RedisClient
	nats   *commonClient.NatsClient
}

func NewCommentRepo(
	conf *conf.Bootstrap,
	logger log.Logger,
	consul *commonClient.ConsulClient,
	redis *commonClient.RedisClient,
	nats *commonClient.NatsClient,
) repo.CommentRepo {
	return &CommentRepo{
		conf:   conf,
		log:    log.NewHelper(logger),
		consul: consul,
		redis:  redis,
		nats:   nats,
	}
}

func (r *CommentRepo) Save(ctx context.Context, client *gen.Client, comment *model.Comment) (*model.Comment, error) {
	save, err := client.Comment.Create().
		SetArticleID(comment.ArticleID).
		SetContent(comment.Content).
		SetLevel(comment.Level).
		SetNillableParentID(comment.ParentID).
		SetNillableReplyID(comment.ReplyID).
		SetStatus(commentent.StatusNormal).
		Save(ctx)
	return &model.Comment{Comment: save}, err
}

func (r *CommentRepo) UpdateStatus(ctx context.Context, client *gen.Client, commentId int64, status v1.CommentStatus) error {
	dbStatus, _ := enum.CommentStatusMap.ToEnum(status)
	_, err := client.Comment.UpdateOneID(commentId).
		SetStatus(commentent.Status(dbStatus)).
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
		return nil, cerrors.ErrorBadRequest("comment is not found")
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

func (r *CommentRepo) GetPage(ctx context.Context, tx *gen.Client, page *common.PageRequest, req *repo.CommentGetReq) ([]*model.Comment, *common.PageReply, error) {
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
		case int32(v1.CommentOrder_COMMENT_ORDER_NEWEST):
			query = query.Order(gen.Desc(commentent.FieldCreatedAt))
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
		query = query.Order(gen.Desc(commentent.FieldCreatedAt))
	}
	return query
}

func (r *CommentRepo) GetArticleLastComment(ctx context.Context, client *gen.Client, req *repo.CommentGetReq) (*model.Comment, error) {
	if req.ArticleId == nil {
		return nil, cerrors.ErrorBadRequest("articleId is required")
	}
	query := client.Comment.Query()
	query = r.getQuery(query, req)
	c, err := query.Order(gen.Desc(commentent.FieldCreatedAt)).First(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	return &model.Comment{Comment: c, WithArticle: req.WithArticle}, err
}

func (r *CommentRepo) GetArticleLastComments(ctx context.Context, tx *gen.Client, req *repo.CommentGetReq) (map[int64]*model.Comment, error) {
	if len(req.ArticleIds) == 0 {
		return nil, cerrors.ErrorBadRequest("articleIds is required")
	}
	articleIdsAny := make([]any, len(req.ArticleIds))
	for i, v := range req.ArticleIds {
		articleIdsAny[i] = v
	}
	comments, err := tx.Comment.Query().
		Where(func(s *sql.Selector) {
			// 子查询 SELECT article_id, MAX(created_at)
			sub := sql.Select(
				commentent.FieldArticleID,
				sql.As(sql.Max(commentent.FieldCreatedAt), "latest_time"),
			).
				From(sql.Table(commentent.Table)).
				Where(sql.EQ(commentent.FieldStatus, string(commentent.StatusNormal))).
				Where(sql.In(commentent.FieldArticleID, articleIdsAny...)).
				GroupBy(commentent.FieldArticleID)

			// JOIN 子查询
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
		commentMap[item.ArticleID] = &model.Comment{Comment: item, WithArticle: req.WithArticle}
	}
	return commentMap, err
}
