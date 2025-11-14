package data

import (
	cv1 "common/api/common/v1"
	"common/pkg/constant"
	"common/pkg/util/base"
	"common/pkg/util/collections/dict"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/ent/gen"
	"content/internal/data/ent/gen/comment"
	"context"

	"entgo.io/ent/dialect/sql"
)

type CommentRepo struct {
	*BaseRepo
	client *gen.Client
}

func NewCommentRepo(baseRepo *BaseRepo, client *gen.Client) repo.CommentRepo {
	return &CommentRepo{
		BaseRepo: baseRepo,
		client:   client,
	}
}

func (r CommentRepo) Save(ctx context.Context, client *gen.Client, comment *model.Comment) (*model.Comment, error) {
	save, err := client.Comment.Create().
		SetArticleID(comment.ArticleID).
		SetContent(comment.Content).
		SetLevel(comment.Level).
		SetNillableParentID(comment.ParentID).
		SetNillableReplyID(comment.ReplyID).
		SetStatus(0).
		Save(ctx)
	return (*model.Comment)(save), err
}

func (r CommentRepo) UpdateStatus(ctx context.Context, client *gen.Client, commentId int64, status cv1.CommentStatus) error {
	_, err := client.Comment.UpdateOneID(commentId).
		SetStatus(int32(status)).
		Save(ctx)
	return err
}

func (r CommentRepo) UpdateStat(ctx context.Context, client *gen.Client, commentId int64, action cv1.CommentAction, num int32) error {
	updateOne := client.Comment.UpdateOneID(commentId)
	switch action {
	case cv1.CommentAction_CommentActionLike:
		updateOne.AddLikeCount(num)
	case cv1.CommentAction_CommentActionCollect:
		updateOne.AddCollectCount(num)
	case cv1.CommentAction_CommentActionReply:
		updateOne.AddReplyCount(num)
	}
	_, err := updateOne.Save(ctx)
	return err
}

func (r CommentRepo) Exist(ctx context.Context, tx *gen.Client, id int64) (bool, error) {
	exist, err := tx.Comment.Query().
		Where(comment.StatusEQ(int32(cv1.CommentStatus_CommentNormal))).
		Where(comment.IDEQ(id)).
		Exist(ctx)
	return exist, err
}

func (r CommentRepo) GetById(ctx context.Context, tx *gen.Client, id int64) (*model.Comment, error) {
	query, err := tx.Comment.Query().
		Where(comment.IDEQ(id)).
		Where(comment.StatusEQ(int32(cv1.CommentStatus_CommentNormal))).
		First(ctx)
	return (*model.Comment)(query), err
}

func (r CommentRepo) GetList(ctx context.Context, tx *gen.Client, req *repo.CommentGetReq) ([]*model.Comment, error) {
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
		comments = append(comments, (*model.Comment)(list[i]))
	}
	return comments, nil
}

func (r CommentRepo) GetPage(ctx context.Context, tx *gen.Client, page *cv1.PageRequest, req *repo.CommentGetReq) ([]*model.Comment, *cv1.PageReply, error) {
	var (
		comments []*model.Comment
		err      error
		total    int
	)
	page = base.OrDefault(page, constant.GetPageDefault())
	query := tx.Comment.Query()
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
		comments = append(comments, (*model.Comment)(list[i]))
	}
	return comments, &cv1.PageReply{
		Total: uint32(total),
		Page:  page.Page,
		Size:  page.Size,
	}, nil
}

func (r CommentRepo) getQuery(query *gen.CommentQuery, req *repo.CommentGetReq) *gen.CommentQuery {
	if req.CommentId != nil {
		query = query.Where(comment.ParentIDEQ(*req.CommentId)).
			WithReply(func(query *gen.CommentQuery) {
				query.Select(comment.FieldCreatedBy).Where(comment.LevelNEQ(1))
			}).
			Order(gen.Asc(comment.FieldCreatedAt))
	}
	if req.ArticleId != nil {
		query = query.Where(comment.ArticleIDEQ(*req.ArticleId))
	}
	if req.UserId != nil {
		query = query.Where(comment.CreatedByEQ(*req.UserId))
	}
	if req.Order != nil {
		// Todo 评论排序
	} else {
		query = query.Order(gen.Desc(comment.FieldCreatedAt))
	}
	return query
}

func (r CommentRepo) GetArticleLastComment(ctx context.Context, client *gen.Client, articleId int64) (*model.Comment, error) {
	query, err := client.Comment.Query().
		Where(comment.ArticleIDEQ(articleId)).
		Where(comment.StatusEQ(int32(cv1.CommentStatus_CommentNormal))).
		Order(gen.Desc(comment.FieldCreatedAt)).
		First(ctx)
	return (*model.Comment)(query), err
}

func (r CommentRepo) GetArticleLastComments(ctx context.Context, tx *gen.Client, articleIds []int64) (dict.Map[int64, *model.Comment], error) {
	articleIdsAny := make([]any, len(articleIds))
	for i, v := range articleIds {
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
				Where(sql.EQ(comment.FieldStatus, int32(cv1.CommentStatus_CommentNormal))).
				Where(sql.In(comment.FieldArticleID, articleIdsAny...)).
				GroupBy(comment.FieldArticleID)

			// JOIN 子查询
			s.Join(sub).On(
				s.C(comment.FieldArticleID), sub.C(comment.FieldArticleID),
			).On(
				s.C(comment.FieldCreatedAt), sub.C("latest_time"),
			)
		}).
		Where(comment.StatusEQ(int32(cv1.CommentStatus_CommentNormal))).
		Where(comment.ArticleIDIn(articleIds...)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	commentMap := dict.New[int64, *model.Comment](0)
	for _, item := range comments {
		commentMap.Set(item.ArticleID, (*model.Comment)(item))
	}
	return commentMap, err
}
