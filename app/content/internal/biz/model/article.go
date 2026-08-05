package model

import (
	"fmt"
	"time"

	"common/pkg/apperror"
	"common/pkg/util"
	cerrors "common/proto/gen/common/errors"
	"content/internal/enum"
)

type Article struct {
	ID            int64
	Title         string
	Content       string
	HasPostscript bool
	RewardContent *string
	RewardPoints  *int32
	PublishStatus enum.ArticlePublishStatus
	Visibility    enum.ArticleVisibility
	Restriction   enum.ContentRestriction
	Type          enum.ArticleType
	Statement     *string
	Commentable   bool
	PublishedAt   *time.Time
	EditedAt      *time.Time
	ViewCount     int32
	ThankCount    int32
	LikeCount     int32
	CollectCount  int32
	RewardCount   int32
	ReplyCount    int32
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
	CreatedBy     *int64
	UpdatedBy     *int64
	DeletedAt     *time.Time
}

func (a *Article) FormatContent() {
	a.Content = util.LuteEngine.FormatStr(fmt.Sprintf("%s_%d", "article_content", a.ID), a.Content)
}

func (a *Article) IsAuthor(userID int64) bool {
	return a != nil && a.CreatedBy != nil && *a.CreatedBy == userID
}

func (a *Article) IsPublicVisible() bool {
	if a == nil || a.DeletedAt != nil || a.Visibility != enum.ArticleVisibilityPublic {
		return false
	}
	if a.PublishStatus != enum.ArticlePublishStatusPublished && a.PublishStatus != enum.ArticlePublishStatusArchived {
		return false
	}
	return a.Restriction == enum.ContentRestrictionNone || a.Restriction == enum.ContentRestrictionLocked
}

func (a *Article) CanCreateDraft() error {
	if a == nil || a.Type != enum.ArticleTypeNormal {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_TYPE)
	}
	return nil
}

func (a *Article) CanEditDraft(operatorID int64) error {
	if a == nil || !a.IsAuthor(operatorID) {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
	}
	if a.Type != enum.ArticleTypeNormal || a.PublishStatus != enum.ArticlePublishStatusDraft || a.Restriction != enum.ContentRestrictionNone {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
	}
	return nil
}

func (a *Article) CanPublish(access *ContentAccess, publishAt *time.Time, now time.Time) error {
	access, err := access.Normalize("")
	if err != nil {
		return err
	}
	if a == nil || a.Type != enum.ArticleTypeNormal || a.Restriction != enum.ContentRestrictionNone {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
	}
	switch access.Scope {
	case enum.ContentAccessScopeAuthor:
		if !a.IsAuthor(access.ActorUserID) {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
		}
		if publishAt != nil {
			if a.PublishStatus != enum.ArticlePublishStatusDraft {
				return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
			}
			if !publishAt.After(now) {
				return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_STATUS)
			}
			return nil
		}
		if a.PublishStatus != enum.ArticlePublishStatusDraft && a.PublishStatus != enum.ArticlePublishStatusScheduled {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
		}
		return nil
	case enum.ContentAccessScopeInternalTask:
		if a.PublishStatus != enum.ArticlePublishStatusScheduled {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
		}
		return nil
	default:
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
	}
}

func (a *Article) CanCancelSchedule(operatorID int64) error {
	if a == nil || !a.IsAuthor(operatorID) {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
	}
	if a.Type != enum.ArticleTypeNormal || a.PublishStatus != enum.ArticlePublishStatusScheduled || a.Restriction != enum.ContentRestrictionNone {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
	}
	return nil
}

func (a *Article) CanArchive(operatorID int64) error {
	if a == nil || !a.IsAuthor(operatorID) {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
	}
	if a.PublishStatus != enum.ArticlePublishStatusPublished || a.Restriction != enum.ContentRestrictionNone {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
	}
	return nil
}

func (a *Article) CanAddPostscript(operatorID int64) error {
	if a == nil || !a.IsAuthor(operatorID) {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
	}
	if a.PublishStatus != enum.ArticlePublishStatusPublished || a.Restriction != enum.ContentRestrictionNone {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
	}
	return nil
}

func (a *Article) CanInteract() error {
	if a == nil || a.PublishStatus != enum.ArticlePublishStatusPublished || a.Visibility != enum.ArticleVisibilityPublic || a.Restriction != enum.ContentRestrictionNone {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
	}
	return nil
}

func (a *Article) CanComment() error {
	if err := a.CanInteract(); err != nil {
		return err
	}
	if !a.Commentable {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_NOT_COMMENTABLE)
	}
	return nil
}

func (a *Article) CanView(viewerID int64) error {
	if a == nil {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_NOT_FOUND)
	}
	if a.Restriction == enum.ContentRestrictionHidden {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
	}
	if a.PublishStatus == enum.ArticlePublishStatusPublished || a.PublishStatus == enum.ArticlePublishStatusArchived {
		if a.Visibility == enum.ArticleVisibilityPublic || a.IsAuthor(viewerID) {
			return nil
		}
	}
	if a.PublishStatus == enum.ArticlePublishStatusDraft || a.PublishStatus == enum.ArticlePublishStatusScheduled {
		if a.IsAuthor(viewerID) {
			return nil
		}
	}
	return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
}
