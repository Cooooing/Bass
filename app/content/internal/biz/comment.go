package biz

import (
	v1 "common/api/common/v1"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/ent"
	"content/internal/data/ent/gen"
	"context"
)

type CommentDomain struct {
	*BaseDomain
	commentRepo repo.CommentRepo
}

func NewCommentDomain(baseDomain *BaseDomain, commentRepo repo.CommentRepo) *CommentDomain {
	return &CommentDomain{
		BaseDomain:  baseDomain,
		commentRepo: commentRepo,
	}
}

func (d *CommentDomain) Add(ctx context.Context, client *gen.Client, comment *model.Comment) (*model.Comment, error) {
	var (
		save *model.Comment
		err  error
	)
	err = ent.WithTx(ctx, client, func(tx *gen.Client) error {
		save, err = d.commentRepo.Save(ctx, tx, comment)
		if err != nil {
			return err
		}
		return nil
	})
	return save, err
}

func (d *CommentDomain) UpdateStatus(ctx context.Context, client *gen.Client, commentId int, status v1.CommentStatus) error {
	err := ent.WithTx(ctx, client, func(tx *gen.Client) error {
		return d.commentRepo.UpdateStatus(ctx, tx, commentId, status)
	})
	return err
}

func (d *CommentDomain) UpdateStat(ctx context.Context, client *gen.Client, commentId int, action v1.CommentAction, active bool) error {
	err := ent.WithTx(ctx, client, func(tx *gen.Client) error {
		if active {
			return d.commentRepo.UpdateStat(ctx, tx, commentId, action, 1)
		} else {
			return d.commentRepo.UpdateStat(ctx, tx, commentId, action, -1)
		}
	})
	return err
}
