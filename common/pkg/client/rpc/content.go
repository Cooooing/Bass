package rpc

import (
	"common/pkg/client/localrpc"
	contentv1 "common/proto/gen/content/v1"

	"google.golang.org/grpc"
)

type ContentClient struct {
	Article    contentv1.ContentArticleServiceClient
	Comment    contentv1.ContentCommentServiceClient
	Domain     contentv1.ContentDomainServiceClient
	Outbox     contentv1.OutboxServiceClient
	Postscript contentv1.ContentPostscriptServiceClient
	Tag        contentv1.ContentTagServiceClient
}

func NewContentClient(
	conn grpc.ClientConnInterface,
) *ContentClient {
	return &ContentClient{
		Article:    contentv1.NewContentArticleServiceClient(conn),
		Comment:    contentv1.NewContentCommentServiceClient(conn),
		Domain:     contentv1.NewContentDomainServiceClient(conn),
		Outbox:     contentv1.NewOutboxServiceClient(conn),
		Postscript: contentv1.NewContentPostscriptServiceClient(conn),
		Tag:        contentv1.NewContentTagServiceClient(conn),
	}
}

func MountContentServices[T any](conn *localrpc.Conn, services []T) {
	for _, service := range services {
		conn.RegisterMatching(&contentv1.ContentArticleService_ServiceDesc, service)
		conn.RegisterMatching(&contentv1.ContentCommentService_ServiceDesc, service)
		conn.RegisterMatching(&contentv1.ContentDomainService_ServiceDesc, service)
		conn.RegisterMatching(&contentv1.OutboxService_ServiceDesc, service)
		conn.RegisterMatching(&contentv1.ContentPostscriptService_ServiceDesc, service)
		conn.RegisterMatching(&contentv1.ContentTagService_ServiceDesc, service)
	}
}
