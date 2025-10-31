package data

import (
	"common/api/common/v1"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/ent/gen"
	"context"
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

func (c CommentRepo) Save(ctx context.Context, client *gen.Client, comment *model.Comment) (*model.Comment, error) {
	save, err := client.Comment.Create().
		SetArticleID(comment.ArticleID).
		SetUserID(comment.UserID).
		SetContent(comment.Content).
		SetLevel(comment.Level).
		SetParentID(comment.ParentID).
		SetStatus(0).
		Save(ctx)
	return (*model.Comment)(save), err
}

func (c CommentRepo) UpdateStatus(ctx context.Context, client *gen.Client, commentId int, status v1.CommentStatus) error {
	_, err := client.Comment.UpdateOneID(commentId).
		SetStatus(int(status)).
		Save(ctx)
	return err
}

func (c CommentRepo) UpdateStat(ctx context.Context, client *gen.Client, commentId int, action v1.CommentAction, num int) error {
	updateOne := client.Comment.UpdateOneID(commentId)
	switch action {
	case v1.CommentAction_CommentActionLike:
		updateOne.AddLikeCount(num)
	case v1.CommentAction_CommentActionCollect:
		updateOne.AddCollectCount(num)
	}
	_, err := updateOne.Save(ctx)
	return err
}
