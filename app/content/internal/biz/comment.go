package biz

import (
	cv1 "common/api/common/v1"
	v1 "common/api/content/v1"
	"common/pkg/util"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/ent"
	"content/internal/data/ent/gen"
	"context"
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
		exist, err := d.articleRepo.GetArticleById(ctx, d.db, comment.ArticleID)
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
			replyComment, err = d.commentRepo.GetCommentById(ctx, d.db, *comment.ReplyID)
			if err != nil {
				return err
			}
			if replyComment == nil {
				return cv1.ErrorBadRequest("reply comment not exist")
			}
			if replyComment.ArticleID != comment.ArticleID {
				return cv1.ErrorBadRequest("reply comment not belong to this article")
			}
		}

		save := &model.Comment{
			ArticleID: comment.ArticleID,
			UserID:    comment.UserID,
			Content:   comment.Content,
			Level:     replyComment.Level + 1,
			ParentID:  util.If(comment.ReplyID == nil, nil, util.If(replyComment.ParentID == nil, &replyComment.ID, replyComment.ParentID)),
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

func (d *CommentDomain) Get(ctx context.Context, req *v1.GetCommentRequest) (rsp *v1.GetCommentReply, err error) {
	err = ent.WithTx(ctx, d.db, func(tx *gen.Client) error {
		rsp, err = d.commentRepo.GetCommentList(ctx, tx, req)
		if err != nil {
			return err
		}
		return nil
	})
	return
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
