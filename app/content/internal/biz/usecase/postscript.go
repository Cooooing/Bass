package usecase

import (
	utilent "common/pkg/util/ent"
	"context"
	"log/slog"

	"common/pkg/apperror"
	commonevent "common/proto/gen/common/enums/event"
	cerrors "common/proto/gen/common/errors"

	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/enum"
)

type PostscriptUsecase struct {
	log *slog.Logger
	tx  utilent.TxFunc

	articleRepo    repo.ArticleRepo
	postscriptRepo repo.PostscriptRepo
	outboxRepo     repo.OutboxEventRepo
	outboxUsecase  *OutboxUsecase
}

func NewPostscriptUsecase(
	logger *slog.Logger,
	tx utilent.TxFunc,
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
	articleID := req.ArticleID
	access, err := req.Access.Normalize("")
	if err != nil {
		return nil, err
	}
	if access.Scope != enum.ContentAccessScopeAuthor {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
	}
	userID := access.ActorUserID
	postscript := &model.Postscript{
		ArticleID:   articleID,
		Content:     req.Content,
		Restriction: enum.ContentRestrictionNone,
		CreatedBy:   new(userID),
		UpdatedBy:   new(userID),
	}
	postscript.FormatContent()

	var (
		save        *model.Postscript
		outboxEvent *repo.OutboxEvent
	)
	err = d.tx(ctx, func(ctx context.Context) error {
		article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{
			Filter: &model.ArticleFilter{ArticleID: new(articleID)},
		})
		if err != nil {
			return err
		}
		if err = article.CanAddPostscript(userID); err != nil {
			return err
		}
		save, err = d.postscriptRepo.Save(ctx, postscript)
		if err != nil {
			return err
		}
		if err = d.articleRepo.UpdateHasPostscript(ctx, &repo.ArticleUpdateHasPostscriptReq{
			ArticleID:     articleID,
			HasPostscript: true,
			UpdatedBy:     userID,
		}); err != nil {
			return err
		}
		outboxEvent, err = d.outboxRepo.Save(ctx, &commonevent.Event{
			Subject: commonevent.EventSubject_EVENT_SUBJECT_ARTICLE_POSTSCRIPT_ADDED,
			Payload: &commonevent.Event_ArticlePostscriptAdded{
				ArticlePostscriptAdded: &commonevent.ArticlePostscriptAddedPayload{
					SenderId:     userID,
					ArticleId:    articleID,
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
	articleID := req.ArticleID
	article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{Filter: &model.ArticleFilter{ArticleID: new(articleID)}})
	if err != nil {
		return nil, err
	}
	if err = article.CanView(access); err != nil {
		return nil, err
	}
	postscriptReq := &repo.PostscriptGetReq{ArticleID: new(articleID)}
	if access.Scope != enum.ContentAccessScopeAdmin && access.Scope != enum.ContentAccessScopeInternalTask {
		postscriptReq.Restrictions = []enum.ContentRestriction{enum.ContentRestrictionNone, enum.ContentRestrictionLocked}
	}
	rows, err := d.postscriptRepo.List(ctx, postscriptReq)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
