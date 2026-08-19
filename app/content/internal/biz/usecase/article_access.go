package usecase

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	"content/internal/biz/model"
	"content/internal/enum"
	"time"
)

type ArticleAccessUsecase struct{}

func NewArticleAccessUsecase() *ArticleAccessUsecase {
	return &ArticleAccessUsecase{}
}

func (u *ArticleAccessUsecase) BuildScope(access *model.ContentAccess, _ *model.ArticleFilter) (*model.ArticleScopeFilter, error) {
	access, err := access.Normalize(enum.ContentAccessScopeGuest)
	if err != nil {
		return nil, err
	}
	scope := &model.ArticleScopeFilter{}
	switch access.Scope {
	case enum.ContentAccessScopeGuest:
		published := enum.ArticlePublishStatusPublished
		public := enum.ArticleVisibilityPublic
		none := enum.ContentRestrictionNone
		locked := enum.ContentRestrictionLocked
		scope.PublishStatus = new(published)
		scope.Visibility = new(public)
		scope.Restrictions = []enum.ContentRestriction{none, locked}
		scope.PublicVisibleOnly = true
	case enum.ContentAccessScopeUser:
		published := enum.ArticlePublishStatusPublished
		public := enum.ArticleVisibilityPublic
		none := enum.ContentRestrictionNone
		locked := enum.ContentRestrictionLocked
		scope.PublishStatus = new(published)
		scope.Visibility = new(public)
		scope.Restrictions = []enum.ContentRestriction{none, locked}
		scope.PublicVisibleOnly = true
	case enum.ContentAccessScopeAuthor:
		scope.AuthorID = new(access.ActorUserID)
	case enum.ContentAccessScopeAdmin, enum.ContentAccessScopeInternalTask:
	default:
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	return scope, nil
}

func (u *ArticleAccessUsecase) CanGet(access *model.ContentAccess, article *model.Article) error {
	access, err := access.Normalize(enum.ContentAccessScopeGuest)
	if err != nil {
		return err
	}
	if article == nil || article.DeletedAt != nil {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_NOT_FOUND)
	}
	return article.CanView(access)
}

func (u *ArticleAccessUsecase) CanCreateDraft(access *model.ContentAccess) error {
	access, err := access.Normalize("")
	if err != nil {
		return err
	}
	if access.Scope != enum.ContentAccessScopeAuthor {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
	}
	return nil
}

func (u *ArticleAccessUsecase) CanUpdateDraft(access *model.ContentAccess, article *model.Article) error {
	access, err := access.Normalize("")
	if err != nil {
		return err
	}
	if access.Scope != enum.ContentAccessScopeAuthor {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
	}
	return article.CanEditDraft(access.ActorUserID)
}

func (u *ArticleAccessUsecase) CanPublish(access *model.ContentAccess, article *model.Article, now time.Time) error {
	access, err := access.Normalize(enum.ContentAccessScopeInternalTask)
	if err != nil {
		return err
	}
	switch access.Scope {
	case enum.ContentAccessScopeAuthor:
		return article.CanPublish(access, nil, now)
	case enum.ContentAccessScopeInternalTask:
		return article.CanPublish(access, nil, now)
	default:
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
	}
}

func (u *ArticleAccessUsecase) CanCancelPublish(access *model.ContentAccess, article *model.Article) error {
	access, err := access.Normalize("")
	if err != nil {
		return err
	}
	if access.Scope != enum.ContentAccessScopeAuthor {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
	}
	return article.CanCancelSchedule(access.ActorUserID)
}

func (u *ArticleAccessUsecase) CanArchive(access *model.ContentAccess, article *model.Article) error {
	access, err := access.Normalize("")
	if err != nil {
		return err
	}
	if access.Scope != enum.ContentAccessScopeAuthor {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
	}
	return article.CanArchive(access.ActorUserID)
}

func (u *ArticleAccessUsecase) CanUnarchive(access *model.ContentAccess, _ *model.Article) error {
	return u.CanManage(access)
}

func (u *ArticleAccessUsecase) CanManage(access *model.ContentAccess) error {
	access, err := access.Normalize("")
	if err != nil {
		return err
	}
	if access.Scope != enum.ContentAccessScopeAdmin {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
	}
	return nil
}

func (u *ArticleAccessUsecase) CanInteract(access *model.ContentAccess, article *model.Article) error {
	access, err := access.Normalize("")
	if err != nil {
		return err
	}
	if access.Scope != enum.ContentAccessScopeAuthor {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
	}
	if article.IsAuthor(access.ActorUserID) {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
	}
	return article.CanInteract()
}

func (u *ArticleAccessUsecase) CanBindTags(access *model.ContentAccess, article *model.Article) error {
	access, err := access.Normalize("")
	if err != nil {
		return err
	}
	if access.Scope == enum.ContentAccessScopeAdmin {
		return nil
	}
	if access.Scope != enum.ContentAccessScopeAuthor {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
	}
	return article.CanEditDraft(access.ActorUserID)
}

func (u *ArticleAccessUsecase) CanAddPostscript(access *model.ContentAccess, article *model.Article) error {
	access, err := access.Normalize("")
	if err != nil {
		return err
	}
	if access.Scope != enum.ContentAccessScopeUser {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
	}
	return article.CanAddPostscript(access.ActorUserID)
}

func (u *ArticleAccessUsecase) CanViewRewardContent(access *model.ContentAccess, article *model.Article, rewarded bool) error {
	access, err := access.Normalize(enum.ContentAccessScopeGuest)
	if err != nil {
		return err
	}
	if access.Scope == enum.ContentAccessScopeAdmin || article.IsAuthor(access.ActorUserID) || rewarded {
		return nil
	}
	return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
}
