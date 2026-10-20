package repo

import (
	"context"

	"realtime_hub/internal/biz/model"
	"realtime_hub/internal/enum"
)

type Registry interface {
	RegisterNode(ctx context.Context, nodeID string, group enum.NodeGroup, realms []string) error
	RenewNode(ctx context.Context, nodeID string, group enum.NodeGroup, realms []string) (bool, error)
	RemoveNode(ctx context.Context, nodeID string, realms []string) error
	BindConnection(ctx context.Context, nodeID, realm string, userID int64, connectionID string) ([]model.EvictedConnection, error)
	UnbindConnection(ctx context.Context, nodeID, realm string, userID int64, connectionID string) error
	UserNodes(ctx context.Context, realm string, userID int64) ([]string, error)
	IssueTicket(ctx context.Context, realm, ticket string, userID int64) error
	ConsumeTicket(ctx context.Context, realm, ticket string) (int64, error)
}

type ControlPublisher interface {
	PublishCloseConnection(ctx context.Context, realm, nodeID, connectionID, reason string) error
}
