package usecase

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	"content/internal/biz/model"
	"content/internal/enum"
)

type CommentAccessUsecase struct{}

func NewCommentAccessUsecase() *CommentAccessUsecase {
	return &CommentAccessUsecase{}
}

func (u *CommentAccessUsecase) BuildScope(access *model.ContentAccess) (*model.CommentScopeFilter, error) {
	access, err := access.Normalize(enum.ContentAccessScopeGuest)
	if err != nil {
		return nil, err
	}
	scope := &model.CommentScopeFilter{}
	switch access.Scope {
	case enum.ContentAccessScopeGuest, enum.ContentAccessScopeUser:
		none := enum.ContentRestrictionNone
		locked := enum.ContentRestrictionLocked
		scope.Restrictions = []enum.ContentRestriction{none, locked}
		scope.ArticlePublicVisible = true
	case enum.ContentAccessScopeAdmin, enum.ContentAccessScopeInternalTask:
	default:
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	return scope, nil
}

func (u *CommentAccessUsecase) CanCreate(access *model.ContentAccess, article *model.Article) error {
	access, err := access.Normalize("")
	if err != nil {
		return err
	}
	if access.Scope != enum.ContentAccessScopeUser {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
	}
	return article.CanComment()
}

func (u *CommentAccessUsecase) CanManage(access *model.ContentAccess) error {
	access, err := access.Normalize("")
	if err != nil {
		return err
	}
	if access.Scope != enum.ContentAccessScopeAdmin {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
	}
	return nil
}

func (u *CommentAccessUsecase) CanInteract(access *model.ContentAccess, comment *model.Comment, article *model.Article) error {
	access, err := access.Normalize("")
	if err != nil {
		return err
	}
	if access.Scope != enum.ContentAccessScopeUser {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
	}
	if comment == nil || comment.DeletedAt != nil || comment.Restriction != enum.ContentRestrictionNone {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_COMMENT_NOT_FOUND)
	}
	return article.CanInteract()
}
