package usecase

import (
	"context"
	"log/slog"

	"common/pkg/apperror"
	commonenums "common/proto/gen/common/enums"
	cerrors "common/proto/gen/common/errors"
	"content/internal/biz/base"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/enum"
)

type PostscriptUsecase struct {
	log *slog.Logger
	tx  base.Tx

	articleRepo    repo.ArticleRepo
	postscriptRepo repo.PostscriptRepo
	outboxRepo     repo.OutboxEventRepo
	outboxUsecase  *OutboxUsecase
}

func NewPostscriptUsecase(
	logger *slog.Logger,
	tx base.Tx,
	articleRepo repo.ArticleRepo,
	postscriptRepo repo.PostscriptRepo,
	outboxRepo repo.OutboxEventRepo,
	outboxUsecase *OutboxUsecase,
) *PostscriptUsecase {
	return &PostscriptUsecase{
		log:            logger,
		tx:             tx,
		articleRepo:    articleRepo,
		postscriptRepo: postscriptRepo,
		outboxRepo:     outboxRepo,
		outboxUsecase:  outboxUsecase,
	}
}

type PostscriptAddReq struct {
	ArticleID int64
	Content   string
	Access    *model.ContentAccess
}

func (d *PostscriptUsecase) Add(ctx context.Context, req *PostscriptAddReq) (*model.Postscript, error) {
	articleId := req.ArticleID
	access, err := req.Access.Normalize("")
	if err != nil {
		return nil, err
	}
	if access.Scope != enum.ContentAccessScopeAuthor {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
	}
	userId := access.ActorUserID
	postscript := &model.Postscript{
		ArticleID:   articleId,
		Content:     req.Content,
		Restriction: enum.ContentRestrictionNone,
		CreatedBy:   new(userId),
		UpdatedBy:   new(userId),
	}
	postscript.FormatContent()

	var (
		save        *model.Postscript
		outboxEvent *repo.OutboxEvent
	)
	err = d.tx(ctx, func(ctx context.Context) error {
		article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{
			Filter: &model.ArticleFilter{ArticleID: new(articleId)},
		})
		if err != nil {
			return err
		}
		if err = article.CanAddPostscript(userId); err != nil {
			return err
		}
		save, err = d.postscriptRepo.Save(ctx, postscript)
		if err != nil {
			return err
		}
		if err = d.articleRepo.UpdateHasPostscript(ctx, &repo.ArticleUpdateHasPostscriptReq{
			ArticleID:     articleId,
			HasPostscript: true,
			UpdatedBy:     userId,
		}); err != nil {
			return err
		}
		outboxEvent, err = d.outboxRepo.Save(ctx, &commonenums.Event{
			Type:    commonenums.EventType_EVENT_TYPE_ARTICLE_POSTSCRIPT_ADDED,
			Subject: commonenums.EventSubject_EVENT_SUBJECT_ARTICLE_POSTSCRIPT_ADDED,
			Payload: &commonenums.Event_ArticlePostscriptAdded{
				ArticlePostscriptAdded: &commonenums.ArticlePostscriptAddedPayload{
					SenderId:     userId,
					ArticleId:    articleId,
					PostscriptId: save.ID,
				},
			},
		})
		return err
	})
	if err != nil {
		return nil, err
	}
	if outboxEvent != nil {
		if _, publishErr := d.outboxUsecase.Publish(ctx, &PublishOutboxEventReq{ID: outboxEvent.ID}); publishErr != nil {
			d.log.WarnContext(ctx, "publish content outbox event failed", slog.Int64("outbox_id", outboxEvent.ID), slog.Any("err", publishErr))
		}
	}
	return save, nil
}

type PostscriptListReq struct {
	ArticleID int64
	Access    *model.ContentAccess
}

func (d *PostscriptUsecase) List(ctx context.Context, req *PostscriptListReq) ([]*model.Postscript, error) {
	access, err := req.Access.Normalize(enum.ContentAccessScopeGuest)
	if err != nil {
		return nil, err
	}
	articleId := req.ArticleID
	article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{Filter: &model.ArticleFilter{ArticleID: new(articleId)}})
	if err != nil {
		return nil, err
	}
	if err = article.CanView(access); err != nil {
		return nil, err
	}
	postscriptReq := &repo.PostscriptGetReq{ArticleID: new(articleId)}
	if access.Scope != enum.ContentAccessScopeAdmin && access.Scope != enum.ContentAccessScopeInternalTask {
		postscriptReq.Restrictions = []enum.ContentRestriction{enum.ContentRestrictionNone, enum.ContentRestrictionLocked}
	}
	rows, err := d.postscriptRepo.List(ctx, postscriptReq)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
