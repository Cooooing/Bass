package usecase

import (
	cerrors "common/proto/gen/common/errors"
	"context"

	"common/pkg/apperror"
	"content/internal/biz/base"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/enum"

	"github.com/samber/lo"
)

type TagUsecase struct {
	tx          base.Tx
	tagRepo     repo.TagRepo
	articleRepo repo.ArticleRepo
}

func NewTagUsecase(
	tx base.Tx,
	tagRepo repo.TagRepo,
	articleRepo repo.ArticleRepo,
) *TagUsecase {
	return &TagUsecase{
		tx:          tx,
		tagRepo:     tagRepo,
		articleRepo: articleRepo,
	}
}

func (t *TagUsecase) Saves(ctx context.Context, tags []*model.Tag) ([]*model.Tag, error) {
	var (
		rows []*model.Tag
		err  error
	)
	err = t.tx(ctx, func(ctx context.Context) error {
		savesResp, saveErr := t.tagRepo.Saves(ctx, tags)
		if saveErr != nil {
			return saveErr
		}
		rows = savesResp
		return err
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (t *TagUsecase) Update(ctx context.Context, tag *model.Tag) (*model.Tag, error) {
	var (
		updated *model.Tag
		err     error
	)
	err = t.tx(ctx, func(ctx context.Context) error {
		updateResp, updateErr := t.tagRepo.Update(ctx, tag)
		if updateErr != nil {
			return updateErr
		}
		updated = updateResp
		return err
	})
	if err != nil {
		return nil, err
	}
	return updated, nil
}

type TagPageReq struct {
	Page        *base.PageRequest
	TagIDs      []int64
	Code        *string
	Name        *string
	Names       []string
	Description *string
	Status      *enum.TagStatus
	DomainID    *int64
}

type TagPageResp struct {
	Rows []*model.Tag
	Page *base.PageResp
}

func (t *TagUsecase) Page(ctx context.Context, req *TagPageReq) (*TagPageResp, error) {
	if req == nil {
		req = &TagPageReq{}
	}
	pageResp, err := t.tagRepo.Page(ctx, &repo.TagGetReq{
		Page:        req.Page,
		TagIds:      req.TagIDs,
		Code:        req.Code,
		Name:        req.Name,
		Names:       req.Names,
		Description: req.Description,
		Status:      req.Status,
		DomainId:    req.DomainID,
	})
	if err != nil {
		return nil, err
	}
	return &TagPageResp{
		Rows: pageResp.Rows,
		Page: pageResp.Page,
	}, nil
}

type ArticleTagReq struct {
	ArticleID int64
	TagIDs    []int64
	UserID    int64
	Manager   bool
}

func (t *TagUsecase) BindArticle(ctx context.Context, req *ArticleTagReq) error {
	tagIDs := lo.Uniq(req.TagIDs)
	err := t.tx(ctx, func(ctx context.Context) error {
		article, err := t.articleRepo.Get(ctx, &repo.ArticleGetReq{
			ArticleId: &req.ArticleID,
		})
		if err != nil {
			return err
		}
		if !req.Manager {
			if article.CreatedBy == nil || *article.CreatedBy != req.UserID {
				return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
			}
			if article.PublishStatus != enum.ArticlePublishStatusDraft {
				return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
			}
		}
		count, err := t.tagRepo.Count(ctx, &repo.TagGetReq{
			TagIds: tagIDs,
			Status: new(enum.TagStatusEnabled),
		})
		if err != nil {
			return err
		}
		if count != len(tagIDs) {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_TAG_NOT_FOUND)
		}
		currentTags, err := t.articleRepo.ListTags(ctx, req.ArticleID)
		if err != nil {
			return err
		}
		currentTagIDs := lo.Map(currentTags, func(item *model.Tag, _ int) int64 {
			return item.ID
		})
		bindIDs, _ := lo.Difference(tagIDs, currentTagIDs)
		if len(currentTagIDs)+len(bindIDs) > 5 {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_TAG_INVALID)
		}
		actualBindIDs, err := t.articleRepo.BindTags(ctx, &repo.ArticleTagBindReq{
			ArticleID: req.ArticleID,
			TagIDs:    tagIDs,
		})
		if err != nil {
			return err
		}
		return t.tagRepo.AddArticleCount(ctx, &repo.TagAddArticleCountReq{
			TagIDs: actualBindIDs,
			Delta:  1,
		})
	})
	if err != nil {
		return err
	}
	return nil
}

func (t *TagUsecase) UnbindArticle(ctx context.Context, req *ArticleTagReq) error {
	tagIDs := lo.Uniq(req.TagIDs)
	err := t.tx(ctx, func(ctx context.Context) error {
		article, err := t.articleRepo.Get(ctx, &repo.ArticleGetReq{
			ArticleId: &req.ArticleID,
		})
		if err != nil {
			return err
		}
		if !req.Manager {
			if article.CreatedBy == nil || *article.CreatedBy != req.UserID {
				return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
			}
			if article.PublishStatus != enum.ArticlePublishStatusDraft {
				return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
			}
		}
		actualUnbindIDs, err := t.articleRepo.UnbindTags(ctx, &repo.ArticleTagBindReq{
			ArticleID: req.ArticleID,
			TagIDs:    tagIDs,
		})
		if err != nil {
			return err
		}
		return t.tagRepo.AddArticleCount(ctx, &repo.TagAddArticleCountReq{
			TagIDs: actualUnbindIDs,
			Delta:  -1,
		})
	})
	if err != nil {
		return err
	}
	return nil
}

func (t *TagUsecase) ListArticleTags(ctx context.Context, articleID int64) ([]*model.Tag, error) {
	rows, err := t.articleRepo.ListTags(ctx, articleID)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
