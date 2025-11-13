package biz

import (
	cv1 "common/api/common/v1"
	v1 "common/api/content/v1"
	userv1 "common/api/user/v1"
	"common/pkg/client"
	"common/pkg/constant"
	"common/pkg/util/base"
	"common/pkg/util/collections/set"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/ent"
	"content/internal/data/ent/gen"
	"context"

	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CommentDomain struct {
	*BaseDomain
	commentRepo repo.CommentRepo
	articleRepo repo.ArticleRepo
}

func NewCommentDomain(baseDomain *BaseDomain, commentRepo repo.CommentRepo, articleRepo repo.ArticleRepo) *CommentDomain {
	return &CommentDomain{
		BaseDomain:  baseDomain,
		commentRepo: commentRepo,
		articleRepo: articleRepo,
	}
}

func (d *CommentDomain) Add(ctx context.Context, comment *model.Comment) (res *model.Comment, err error) {
	err = ent.WithTx(ctx, d.db, func(tx *gen.Client) error {
		// 回复文章
		exist, err := d.articleRepo.GetArticleById(ctx, tx, comment.ArticleID)
		if err != nil {
			return err
		}
		if exist == nil || exist.Status != int32(cv1.ArticleStatus_ArticleNormal) {
			return cv1.ErrorBadRequest("article not exist")
		}
		if !exist.Commentable {
			return cv1.ErrorBadRequest("article not commentable")
		}

		// 回复评论
		replyComment := &model.Comment{}
		if comment.ReplyID != nil {
			replyComment, err = d.commentRepo.GetById(ctx, tx, *comment.ReplyID)
			if err != nil {
				return err
			}
			if replyComment == nil {
				return cv1.ErrorBadRequest("reply comment not exist")
			}
			if replyComment.ArticleID != comment.ArticleID {
				return cv1.ErrorBadRequest("reply comment not belong to this article")
			}

			err = d.commentRepo.UpdateStat(ctx, tx, replyComment.ID, cv1.CommentAction_CommentActionReply, 1)
			if err != nil {
				return err
			}
		}

		err = d.articleRepo.UpdateStat(ctx, tx, exist.ID, cv1.ArticleAction_ArticleActionReply, 1)
		if err != nil {
			return err
		}

		save := &model.Comment{
			ArticleID: comment.ArticleID,
			UserID:    comment.UserID,
			Content:   comment.Content,
			Level:     replyComment.Level + 1,
			ParentID:  base.If(comment.ReplyID == nil, nil, base.If(replyComment.ParentID == nil, &replyComment.ID, replyComment.ParentID)),
			ReplyID:   comment.ReplyID,
		}

		_, err = d.commentRepo.Save(ctx, tx, save)
		if err != nil {
			return err
		}

		return nil
	})
	return res, err
}

func (d *CommentDomain) Page(ctx context.Context, page *cv1.PageRequest, req *repo.CommentGetReq) (*v1.PageCommentReply, error) {
	var (
		reply *v1.PageCommentReply
		err   error
	)
	err = ent.WithTx(ctx, d.db, func(tx *gen.Client) error {
		list, pageReply, err := d.commentRepo.GetPage(ctx, tx, page, req)
		if err != nil {
			return err
		}
		userIds := set.New[int64](0)
		for _, item := range list {
			userIds.Add(item.UserID)
		}

		userService, err := client.GetServiceClient(ctx, d.etcd, constant.UserServiceName.String(), userv1.NewUserUserServiceClient)
		if err != nil {
			return err
		}
		userMap, err := userService.GetMap(ctx, &userv1.GetMapRequest{Ids: userIds.ToSlice()})
		if err != nil {
			return err
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
				return err
			}
			elems.User = users[item.UserID]
			comments = append(comments, elems)
		}

		reply = &v1.PageCommentReply{
			Page:     pageReply,
			Comments: comments,
		}
		return nil
	})
	return reply, err
}

func (d *CommentDomain) UpdateStatus(ctx context.Context, commentId int64, status cv1.CommentStatus) error {
	err := ent.WithTx(ctx, d.db, func(tx *gen.Client) error {
		return d.commentRepo.UpdateStatus(ctx, tx, commentId, status)
	})
	return err
}

func (d *CommentDomain) UpdateStat(ctx context.Context, commentId int64, action cv1.CommentAction, active bool) error {
	err := ent.WithTx(ctx, d.db, func(tx *gen.Client) error {
		if active {
			return d.commentRepo.UpdateStat(ctx, tx, commentId, action, 1)
		} else {
			return d.commentRepo.UpdateStat(ctx, tx, commentId, action, -1)
		}
	})
	return err
}
