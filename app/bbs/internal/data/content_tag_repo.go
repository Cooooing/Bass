package data

import (
	"bbs/internal/biz/repo"
	"common/pkg/client/rpc"
	"common/proto/gen/common"
	contentv1 "common/proto/gen/content/v1"
	contentv1enum "common/proto/gen/content/v1/enum"
	"context"
)

var _ repo.ContentTagClient = (*ContentTagClient)(nil)

type ContentTagClient struct {
	contentClient *rpc.ContentClient
}

func NewContentTagClient(
	contentClient *rpc.ContentClient,
) repo.ContentTagClient {
	return &ContentTagClient{
		contentClient: contentClient,
	}
}

func (r *ContentTagClient) CreateTag(ctx context.Context, req *repo.CreateTagReq) (*repo.Tag, error) {
	tag := req.Tag
	var status *contentv1enum.TagStatus
	if tag.Status != nil {
		status = new(contentv1enum.TagStatus(*tag.Status))
	}
	reply, err := r.contentClient.Tag.BatchCreate(ctx, &contentv1.BatchCreateTags_Req{
		UserId: req.UserID,
		Tags: []*contentv1.BatchCreateTags_Req_Tag{
			{
				Code:        tag.Code,
				Name:        tag.Name,
				Description: tag.Description,
				DomainId:    tag.DomainID,
				Status:      status,
				Icon:        tag.Icon,
				Sort:        tag.Sort,
			},
		},
	})
	if err != nil {
		return nil, err
	}
	rows := reply.GetRows()
	if len(rows) == 0 {
		return nil, nil
	}
	item := rows[0]
	return &repo.Tag{
		ID:           item.GetId(),
		Code:         item.GetCode(),
		Name:         item.GetName(),
		Description:  item.Description,
		DomainID:     item.DomainId,
		Status:       new(int32(item.GetStatus())),
		Icon:         item.Icon,
		Sort:         item.GetSort(),
		ArticleCount: item.GetArticleCount(),
		CreatedBy:    item.CreatedBy,
		UpdatedBy:    item.UpdatedBy,
		CreatedAt:    new(item.GetCreatedAt().AsTime()),
		UpdatedAt:    new(item.GetUpdatedAt().AsTime()),
	}, nil
}

func (r *ContentTagClient) UpdateTag(ctx context.Context, req *repo.UpdateTagReq) (*repo.Tag, error) {
	tag := req.Tag
	var status *contentv1enum.TagStatus
	if tag.Status != nil {
		status = new(contentv1enum.TagStatus(*tag.Status))
	}
	reply, err := r.contentClient.Tag.Update(ctx, &contentv1.UpdateTag_Req{
		TagId:  req.TagID,
		UserId: req.UserID,
		Tag: &contentv1.UpdateTag_Req_Tag{
			Code:        tag.Code,
			Name:        tag.Name,
			Description: tag.Description,
			DomainId:    tag.DomainID,
			Status:      status,
			Icon:        tag.Icon,
			Sort:        tag.Sort,
		},
	})
	if err != nil {
		return nil, err
	}
	item := reply.GetTag()
	return &repo.Tag{
		ID:           item.GetId(),
		Code:         item.GetCode(),
		Name:         item.GetName(),
		Description:  item.Description,
		DomainID:     item.DomainId,
		Status:       new(int32(item.GetStatus())),
		Icon:         item.Icon,
		Sort:         item.GetSort(),
		ArticleCount: item.GetArticleCount(),
		CreatedBy:    item.CreatedBy,
		UpdatedBy:    item.UpdatedBy,
		CreatedAt:    new(item.GetCreatedAt().AsTime()),
		UpdatedAt:    new(item.GetUpdatedAt().AsTime()),
	}, nil
}

func (r *ContentTagClient) ListTags(ctx context.Context, req *repo.ListTagsReq) (*repo.ListTagsResp, error) {
	query := req.Query
	if query == nil {
		query = &repo.TagQuery{}
	}
	contentQuery := &contentv1.PageTags_Req_Query{
		Ids:         query.IDs,
		Code:        query.Code,
		Name:        query.Name,
		Names:       query.Names,
		Description: query.Description,
		DomainId:    query.DomainID,
	}
	if query.Status != nil {
		contentQuery.Status = new(contentv1enum.TagStatus(*query.Status))
	}
	var pageReq *common.PageReq
	if req.Page != nil {
		pageReq = &common.PageReq{
			Page: req.Page.Page,
			Size: req.Page.Size,
		}
	}
	reply, err := r.contentClient.Tag.Page(ctx, &contentv1.PageTags_Req{
		Page:  pageReq,
		Query: contentQuery,
	})
	if err != nil {
		return nil, err
	}
	rows := make([]*repo.Tag, 0, len(reply.GetRows()))
	for _, item := range reply.GetRows() {
		rows = append(rows, &repo.Tag{
			ID:           item.GetId(),
			Code:         item.GetCode(),
			Name:         item.GetName(),
			Description:  item.Description,
			DomainID:     item.DomainId,
			Status:       new(int32(item.GetStatus())),
			Icon:         item.Icon,
			Sort:         item.GetSort(),
			ArticleCount: item.GetArticleCount(),
			CreatedBy:    item.CreatedBy,
			UpdatedBy:    item.UpdatedBy,
			CreatedAt:    new(item.GetCreatedAt().AsTime()),
			UpdatedAt:    new(item.GetUpdatedAt().AsTime()),
		})
	}
	var page *repo.PageResp
	if reply.GetPage() != nil {
		page = &repo.PageResp{
			Page:  reply.GetPage().GetPage(),
			Size:  reply.GetPage().GetSize(),
			Total: reply.GetPage().GetTotal(),
		}
	}
	return &repo.ListTagsResp{
		Page: page,
		Rows: rows,
	}, nil
}

func (r *ContentTagClient) BindArticleTags(ctx context.Context, req *repo.BindArticleTagsReq) error {
	_, err := r.contentClient.Tag.BindArticle(ctx, &contentv1.BindArticleTags_Req{
		ArticleId: req.ArticleID,
		TagIds:    req.TagIDs,
		UserId:    req.UserID,
	})
	return err
}

func (r *ContentTagClient) UnbindArticleTags(ctx context.Context, req *repo.UnbindArticleTagsReq) error {
	_, err := r.contentClient.Tag.UnbindArticle(ctx, &contentv1.UnbindArticleTags_Req{
		ArticleId: req.ArticleID,
		TagIds:    req.TagIDs,
		UserId:    req.UserID,
	})
	return err
}

func (r *ContentTagClient) ListArticleTags(ctx context.Context, req *repo.ListArticleTagsReq) ([]*repo.Tag, error) {
	reply, err := r.contentClient.Tag.ListArticleTags(ctx, &contentv1.ListArticleTags_Req{
		ArticleId: req.ArticleID,
	})
	if err != nil {
		return nil, err
	}
	rows := make([]*repo.Tag, 0, len(reply.GetRows()))
	for _, item := range reply.GetRows() {
		rows = append(rows, &repo.Tag{
			ID:           item.GetId(),
			Code:         item.GetCode(),
			Name:         item.GetName(),
			Description:  item.Description,
			DomainID:     item.DomainId,
			Status:       new(int32(item.GetStatus())),
			Icon:         item.Icon,
			Sort:         item.GetSort(),
			ArticleCount: item.GetArticleCount(),
			CreatedBy:    item.CreatedBy,
			UpdatedBy:    item.UpdatedBy,
			CreatedAt:    new(item.GetCreatedAt().AsTime()),
			UpdatedAt:    new(item.GetUpdatedAt().AsTime()),
		})
	}
	return rows, nil
}
