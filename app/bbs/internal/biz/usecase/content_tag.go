package usecase

import (
	"bbs/internal/biz/repo"
	"common/pkg/apperror"
	bbscontentv1 "common/proto/gen/bbs/v1/content"
	cerrors "common/proto/gen/common/errors"
	"context"
)

type ContentTagUsecase struct {
	contentTagClient repo.ContentTagClient
}

func NewContentTagUsecase(contentTagClient repo.ContentTagClient) *ContentTagUsecase {
	return &ContentTagUsecase{contentTagClient: contentTagClient}
}

func (u *ContentTagUsecase) CreateTag(ctx context.Context, req *bbscontentv1.CreateTag_Request) (*bbscontentv1.CreateTag_Reply, error) {
	if req == nil || req.GetTag() == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	return u.contentTagClient.CreateTag(ctx, req)
}

func (u *ContentTagUsecase) UpdateTag(ctx context.Context, req *bbscontentv1.UpdateTag_Request) (*bbscontentv1.UpdateTag_Reply, error) {
	if req == nil || req.GetTag() == nil || req.GetTagId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	return u.contentTagClient.UpdateTag(ctx, req)
}

func (u *ContentTagUsecase) ListTags(ctx context.Context, req *bbscontentv1.ListTags_Request) (*bbscontentv1.ListTags_Reply, error) {
	if req == nil {
		req = &bbscontentv1.ListTags_Request{}
	}
	if req.Query == nil {
		req.Query = &bbscontentv1.TagQuery{}
	}
	if req.Query.Status == nil {
		req.Query.Status = new(bbscontentv1.TagStatus_TAG_STATUS_ENABLED)
	}
	return u.contentTagClient.ListTags(ctx, req)
}
