package rpc

import (
	contentv1 "common/proto/gen/content/v1"

	"google.golang.org/grpc"
)

type ContentClient struct {
	Article contentv1.ContentArticleServiceClient
	Comment contentv1.ContentCommentServiceClient
	Domain  contentv1.ContentDomainServiceClient
	Outbox  contentv1.OutboxServiceClient
	Tag     contentv1.ContentTagServiceClient
}

func NewContentClient(
	conn *grpc.ClientConn,
) *ContentClient {
	return &ContentClient{
		Article: contentv1.NewContentArticleServiceClient(conn),
		Comment: contentv1.NewContentCommentServiceClient(conn),
		Domain:  contentv1.NewContentDomainServiceClient(conn),
		Outbox:  contentv1.NewOutboxServiceClient(conn),
		Tag:     contentv1.NewContentTagServiceClient(conn),
	}
}
