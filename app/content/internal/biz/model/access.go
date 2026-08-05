package model

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	"content/internal/enum"
)

type ContentAccess struct {
	Scope       enum.ContentAccessScope
	ActorUserID int64
}

func (a *ContentAccess) Normalize(defaultScope enum.ContentAccessScope) (*ContentAccess, error) {
	if a == nil || a.Scope == "" {
		a = &ContentAccess{Scope: defaultScope}
	}
	if a.Scope == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	switch a.Scope {
	case enum.ContentAccessScopeGuest, enum.ContentAccessScopeInternalTask:
		return a, nil
	case enum.ContentAccessScopeUser, enum.ContentAccessScopeAuthor, enum.ContentAccessScopeAdmin:
		if a.ActorUserID <= 0 {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		return a, nil
	default:
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
}
