package usecase

import (
	"bbs/internal/biz/repo"
	"common/pkg/apperror"
	bbscontentv1 "common/proto/gen/bbs/v1/content"
	bbscontentv1enum "common/proto/gen/bbs/v1/content/enum"
	"common/proto/gen/common"
	cerrors "common/proto/gen/common/errors"
	"context"
)

type ContentTagUsecase struct {
	contentTagClient repo.ContentTagClient
}

func NewContentTagUsecase(
	contentTagClient repo.ContentTagClient,
) *ContentTagUsecase {
	return &ContentTagUsecase{
		contentTagClient: contentTagClient,
	}
}

type ContentTagSave struct {
	Name        string
	Description *string
	DomainID    *int64
	Status      *bbscontentv1enum.TagStatus
}

type CreateTagReq struct {
	UserID int64
	Tag    *ContentTagSave
}

func (u *ContentTagUsecase) CreateTag(
	ctx context.Context,
	req *CreateTagReq,
) (*repo.Tag, error) {
	if req == nil || req.Tag == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	var status *int32
	if req.Tag.Status != nil {
		value := int32(*req.Tag.Status)
		status = &value
	}
	resp, err := u.contentTagClient.CreateTag(ctx, &repo.CreateTagReq{
		UserID: req.UserID,
		Tag: &repo.TagSave{
			Name:        req.Tag.Name,
			Description: req.Tag.Description,
			DomainID:    req.Tag.DomainID,
			Status:      status,
		},
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type UpdateTagReq struct {
	UserID int64
	TagID  int64
	Tag    *ContentTagSave
}

func (u *ContentTagUsecase) UpdateTag(
	ctx context.Context,
	req *UpdateTagReq,
) (*repo.Tag, error) {
	if req == nil || req.Tag == nil || req.TagID <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	var status *int32
	if req.Tag.Status != nil {
		value := int32(*req.Tag.Status)
		status = &value
	}
	resp, err := u.contentTagClient.UpdateTag(ctx, &repo.UpdateTagReq{
		UserID: req.UserID,
		TagID:  req.TagID,
		Tag: &repo.TagSave{
			Name:        req.Tag.Name,
			Description: req.Tag.Description,
			DomainID:    req.Tag.DomainID,
			Status:      status,
		},
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ListTagsReq struct {
	Page  *common.PageReq
	Query *bbscontentv1.ListTags_Req_TagQuery
}

type ListTagsResp struct {
	Page *repo.PageResp
	Rows []*repo.Tag
}

func (u *ContentTagUsecase) ListTags(
	ctx context.Context,
	req *ListTagsReq,
) (*ListTagsResp, error) {
	if req == nil {
		req = &ListTagsReq{}
	}
	var page *repo.PageReq
	if req.Page != nil {
		page = &repo.PageReq{
			Page: req.Page.GetPage(),
			Size: req.Page.GetSize(),
		}
	}
	query := &repo.TagQuery{}
	if req.Query != nil {
		query.IDs = req.Query.GetIds()
		query.Name = req.Query.Name
		query.Names = req.Query.GetNames()
		query.Description = req.Query.Description
		query.DomainID = req.Query.DomainId
		if req.Query.Status != nil {
			value := int32(*req.Query.Status)
			query.Status = &value
		}
	}
	if query.Status == nil {
		value := int32(bbscontentv1enum.TagStatus_TAG_STATUS_ENABLED)
		query.Status = &value
	}
	resp, err := u.contentTagClient.ListTags(ctx, &repo.ListTagsReq{
		Page:  page,
		Query: query,
	})
	if err != nil {
		return nil, err
	}
	return &ListTagsResp{
		Page: resp.Page,
		Rows: resp.Rows,
	}, nil
}
