package service

import (
	"bbs/internal/biz/usecase"
	bbscontentv1 "common/proto/gen/bbs/v1/content"
	"common/proto/gen/common"
	"context"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

type ContentTagService struct {
	bbscontentv1.UnimplementedTagServiceServer
	contentTagUsecase *usecase.ContentTagUsecase
}

func NewContentTagService(contentTagUsecase *usecase.ContentTagUsecase) *ContentTagService {
	return &ContentTagService{contentTagUsecase: contentTagUsecase}
}

func (s *ContentTagService) RegisterGrpc(gs *grpc.Server) {}

func (s *ContentTagService) RegisterHttp(hs *http.Server) {
	bbscontentv1.RegisterTagServiceHTTPServer(hs, s)
}

func (s *ContentTagService) Create(ctx context.Context, req *bbscontentv1.CreateTag_Req) (*bbscontentv1.CreateTag_Resp, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	var tagSave *usecase.ContentTagSave
	if tag := req.GetTag(); tag != nil {
		tagSave = &usecase.ContentTagSave{Name: tag.GetName(), Description: tag.Description, DomainID: tag.DomainId, Status: tag.Status}
	}
	resp, err := s.contentTagUsecase.CreateTag(ctx, &usecase.CreateTagReq{UserID: userID, Tag: tagSave})
	if err != nil {
		return nil, err
	}
	var tag *bbscontentv1.CreateTag_Resp_Tag
	if resp != nil {
		tag = &bbscontentv1.CreateTag_Resp_Tag{Id: resp.ID, Name: resp.Name, Description: resp.Description, DomainId: resp.DomainID, CreatedBy: resp.CreatedBy, UpdatedBy: resp.UpdatedBy, CreatedAt: resp.CreatedAt, UpdatedAt: resp.UpdatedAt}
		if resp.Status != nil {
			status := bbscontentv1.TagStatus(*resp.Status)
			tag.Status = &status
		}
	}
	return &bbscontentv1.CreateTag_Resp{Tag: tag}, nil
}
func (s *ContentTagService) Update(ctx context.Context, req *bbscontentv1.UpdateTag_Req) (*bbscontentv1.UpdateTag_Resp, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	var tagSave *usecase.ContentTagSave
	if tag := req.GetTag(); tag != nil {
		tagSave = &usecase.ContentTagSave{Name: tag.GetName(), Description: tag.Description, DomainID: tag.DomainId, Status: tag.Status}
	}
	resp, err := s.contentTagUsecase.UpdateTag(ctx, &usecase.UpdateTagReq{UserID: userID, TagID: req.GetTagId(), Tag: tagSave})
	if err != nil {
		return nil, err
	}
	var tag *bbscontentv1.UpdateTag_Resp_Tag
	if resp != nil {
		tag = &bbscontentv1.UpdateTag_Resp_Tag{Id: resp.ID, Name: resp.Name, Description: resp.Description, DomainId: resp.DomainID, CreatedBy: resp.CreatedBy, UpdatedBy: resp.UpdatedBy, CreatedAt: resp.CreatedAt, UpdatedAt: resp.UpdatedAt}
		if resp.Status != nil {
			status := bbscontentv1.TagStatus(*resp.Status)
			tag.Status = &status
		}
	}
	return &bbscontentv1.UpdateTag_Resp{Tag: tag}, nil
}
func (s *ContentTagService) List(ctx context.Context, req *bbscontentv1.ListTags_Req) (*bbscontentv1.ListTags_Resp, error) {
	resp, err := s.contentTagUsecase.ListTags(ctx, &usecase.ListTagsReq{Page: req.GetPage(), Query: req.GetQuery()})
	if err != nil {
		return nil, err
	}
	var page *common.PageResp
	if resp.Page != nil {
		page = &common.PageResp{Page: resp.Page.Page, Size: resp.Page.Size, Total: resp.Page.Total}
	}
	rows := make([]*bbscontentv1.ListTags_Resp_Tag, 0, len(resp.Rows))
	for _, row := range resp.Rows {
		if row == nil {
			rows = append(rows, nil)
			continue
		}
		tag := &bbscontentv1.ListTags_Resp_Tag{Id: row.ID, Name: row.Name, Description: row.Description, DomainId: row.DomainID, CreatedBy: row.CreatedBy, UpdatedBy: row.UpdatedBy, CreatedAt: row.CreatedAt, UpdatedAt: row.UpdatedAt}
		if row.Status != nil {
			status := bbscontentv1.TagStatus(*row.Status)
			tag.Status = &status
		}
		rows = append(rows, tag)
	}
	return &bbscontentv1.ListTags_Resp{Page: page, Rows: rows}, nil
}
