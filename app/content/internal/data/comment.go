package data

import (
	cv1 "common/api/common/v1"
	v1 "common/api/content/v1"
	userv1 "common/api/user/v1"
	"common/pkg/client"
	"common/pkg/constant"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/ent/gen"
	"content/internal/data/ent/gen/comment"
	"context"

	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/types/known/timestamppb"
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
		SetUserID(comment.UserID).
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
	}
	_, err := updateOne.Save(ctx)
	return err
}

func (r CommentRepo) GetCommentById(ctx context.Context, tx *gen.Client, id int64) (*model.Comment, error) {
	query, err := tx.Comment.Query().Where(comment.IDEQ(id)).First(ctx)
	return (*model.Comment)(query), err
}

func (r CommentRepo) GetCommentList(ctx context.Context, tx *gen.Client, req *v1.GetCommentRequest) (*v1.GetCommentReply, error) {
	query := tx.Comment.Query()
	if req.ArticleId != 0 {
		query = query.Where(comment.ArticleIDEQ(req.ArticleId))
	}
	if req.Id != 0 {
		query = query.Where(comment.IDEQ(req.Id))
	}
	if req.UserId != 0 {
		query = query.Where(comment.UserIDEQ(req.UserId))
	}
	countQuery := query.Clone()
	count, err := countQuery.Count(ctx)
	if err != nil {
		return nil, err
	}
	list, err := query.Limit(int(req.Page.Size)).Offset(int((req.Page.Page - 1) * req.Page.Size)).All(ctx)
	if err != nil {
		return nil, err
	}

	userIds := make([]int64, 0)
	for _, item := range list {
		userIds = append(userIds, item.UserID)
	}

	userService, err := client.GetServiceClient(ctx, r.etcd, constant.UserServiceName.String(), userv1.NewUserUserServiceClient)
	if err != nil {
		return nil, err
	}
	userMap, err := userService.GetMap(ctx, &userv1.GetMapRequest{Ids: userIds})
	if err != nil {
		return nil, err
	}
	users := userMap.Users

	comments := make([]*v1.Comment, 0)
	for _, item := range list {
		elems := &v1.Comment{
			CreatedAt: timestamppb.New(*item.CreatedAt),
			UpdatedAt: timestamppb.New(*item.UpdatedAt),
		}
		err = copier.Copy(elems, item)
		if err != nil {
			return nil, err
		}
		elems.User = users[item.UserID]
		comments = append(comments, elems)
	}

	rsp := &v1.GetCommentReply{
		Page: &cv1.PageReply{
			Total: uint32(count),
			Page:  req.Page.Page,
			Size:  req.Page.Size,
		},
		Comments: comments,
	}
	return rsp, err
}
