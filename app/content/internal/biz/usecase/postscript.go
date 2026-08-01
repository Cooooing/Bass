package usecase

import (
	"context"
	"log/slog"

	commonenums "common/proto/gen/common/enums"
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
	UserID    int64
}

func (d *PostscriptUsecase) Add(ctx context.Context, req *PostscriptAddReq) (*model.Postscript, error) {
	articleId := req.ArticleID
	userId := req.UserID
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
		err         error
	)
	err = d.tx(ctx, func(ctx context.Context) error {
		article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{
			ArticleId: new(articleId),
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

func (d *PostscriptUsecase) List(ctx context.Context, articleID int64) ([]*model.Postscript, error) {
	articleId := articleID
	if _, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{ArticleId: new(articleId)}); err != nil {
		return nil, err
	}
	rows, err := d.postscriptRepo.List(ctx, &repo.PostscriptGetReq{
		ArticleID:   new(articleId),
		Restriction: new(enum.ContentRestrictionNone),
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}
