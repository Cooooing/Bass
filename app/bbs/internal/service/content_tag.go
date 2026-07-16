package service

import (
	"bbs/internal/biz/usecase"
	bbscontentv1 "common/proto/gen/bbs/v1/content"
	"common/proto/gen/common"
	"context"

	"github.com/go-kratos/kratos/v3/transport/http"
)

type ContentTagService struct {
	bbscontentv1.UnimplementedTagServiceServer
	contentTagUsecase *usecase.ContentTagUsecase
}

func NewContentTagService(contentTagUsecase *usecase.ContentTagUsecase) *ContentTagService {
	return &ContentTagService{contentTagUsecase: contentTagUsecase}
}
func (s *ContentTagService) RegisterHttp(hs *http.Server) {
	bbscontentv1.RegisterTagServiceHTTPServer(hs, s)
}
func (s *ContentTagService) Create(ctx context.Context, req *bbscontentv1.CreateTag_Request) (*bbscontentv1.CreateTag_Response, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	var tagSave *usecase.ContentTagSave
	if tag := req.GetTag(); tag != nil {
		tagSave = &usecase.ContentTagSave{Name: tag.GetName(), Description: tag.Description, DomainID: tag.DomainId, Status: tag.Status}
	}
	response, err := s.contentTagUsecase.CreateTag(ctx, &usecase.CreateTagReq{UserID: userID, Tag: tagSave})
	if err != nil {
		return nil, err
	}
	var tag *bbscontentv1.CreateTag_Response_Tag
	if response.Tag != nil {
		tag = &bbscontentv1.CreateTag_Response_Tag{Id: response.Tag.ID, Name: response.Tag.Name, Description: response.Tag.Description, DomainId: response.Tag.DomainID, CreatedBy: response.Tag.CreatedBy, UpdatedBy: response.Tag.UpdatedBy, CreatedAt: response.Tag.CreatedAt, UpdatedAt: response.Tag.UpdatedAt}
		if response.Tag.Status != nil {
			status := bbscontentv1.TagStatus(*response.Tag.Status)
			tag.Status = &status
		}
	}
	return &bbscontentv1.CreateTag_Response{Tag: tag}, nil
}
func (s *ContentTagService) Update(ctx context.Context, req *bbscontentv1.UpdateTag_Request) (*bbscontentv1.UpdateTag_Response, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	var tagSave *usecase.ContentTagSave
	if tag := req.GetTag(); tag != nil {
		tagSave = &usecase.ContentTagSave{Name: tag.GetName(), Description: tag.Description, DomainID: tag.DomainId, Status: tag.Status}
	}
	response, err := s.contentTagUsecase.UpdateTag(ctx, &usecase.UpdateTagReq{UserID: userID, TagID: req.GetTagId(), Tag: tagSave})
	if err != nil {
		return nil, err
	}
	var tag *bbscontentv1.UpdateTag_Response_Tag
	if response.Tag != nil {
		tag = &bbscontentv1.UpdateTag_Response_Tag{Id: response.Tag.ID, Name: response.Tag.Name, Description: response.Tag.Description, DomainId: response.Tag.DomainID, CreatedBy: response.Tag.CreatedBy, UpdatedBy: response.Tag.UpdatedBy, CreatedAt: response.Tag.CreatedAt, UpdatedAt: response.Tag.UpdatedAt}
		if response.Tag.Status != nil {
			status := bbscontentv1.TagStatus(*response.Tag.Status)
			tag.Status = &status
		}
	}
	return &bbscontentv1.UpdateTag_Response{Tag: tag}, nil
}
func (s *ContentTagService) List(ctx context.Context, req *bbscontentv1.ListTags_Request) (*bbscontentv1.ListTags_Response, error) {
	response, err := s.contentTagUsecase.ListTags(ctx, &usecase.ListTagsReq{Page: req.GetPage(), Query: req.GetQuery()})
	if err != nil {
		return nil, err
	}
	var page *common.PageResponse
	if response.Page != nil {
		page = &common.PageResponse{Page: response.Page.Page, Size: response.Page.Size, Total: response.Page.Total}
	}
	rows := make([]*bbscontentv1.ListTags_Response_Tag, 0, len(response.Rows))
	for _, row := range response.Rows {
		if row == nil {
			rows = append(rows, nil)
			continue
		}
		tag := &bbscontentv1.ListTags_Response_Tag{Id: row.ID, Name: row.Name, Description: row.Description, DomainId: row.DomainID, CreatedBy: row.CreatedBy, UpdatedBy: row.UpdatedBy, CreatedAt: row.CreatedAt, UpdatedAt: row.UpdatedAt}
		if row.Status != nil {
			status := bbscontentv1.TagStatus(*row.Status)
			tag.Status = &status
		}
		rows = append(rows, tag)
	}
	return &bbscontentv1.ListTags_Response{Page: page, Rows: rows}, nil
}
