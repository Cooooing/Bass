package service

import (
	v1 "common/api/content/v1"
	"content/internal/biz"
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TagService struct {
	v1.UnimplementedContentTagServiceServer
	*BaseService
	domainTag *biz.TagDomain
}

func (s *TagService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterContentTagServiceServer(gs, s)
}

func (s *TagService) RegisterHttp(hs *http.Server) {
	v1.RegisterContentTagServiceHTTPServer(hs, s)
}

func NewTagService(baseService *BaseService, domainTag *biz.TagDomain) *TagService {
	return &TagService{
		BaseService: baseService,
		domainTag:   domainTag,
	}
}

func (s *TagService) Add(ctx context.Context, in *v1.AddTagRequest) (*v1.AddTagReply, error) {
	// TODO implement me
	panic("implement me")
}

func (s *TagService) Update(ctx context.Context, in *v1.UpdateTagRequest) (*v1.UpdateTagReply, error) {
	// TODO implement me
	panic("implement me")
}

func (s *TagService) Get(ctx context.Context, in *v1.GetTagRequest) (*v1.GetTagReply, error) {
	tags := make([]*v1.Tag, 0, len(list))
	for _, item := range list {
		i := &v1.Tag{}
		err = copier.Copy(i, item)
		if err != nil {
			return nil, nil, err
		}
		i.CreatedAt = timestamppb.New(*item.CreatedAt)
		i.UpdatedAt = timestamppb.New(*item.UpdatedAt)
		tags = append(tags, i)
	}
}
