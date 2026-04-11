package model

import (
	v1 "common/api/gen/user/v1"
	"user/internal/data/ent/gen"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserRelation struct {
	*gen.UserRelation
}

func (u *UserRelation) ConvertToRpc() *v1.UserRelation {
	return &v1.UserRelation{
		CreatedAt: timestamppb.New(*u.CreatedAt),
		UpdatedAt: timestamppb.New(*u.UpdatedAt),
		Id:        u.ID,
		Type:      v1.UserRelationType(u.Type),
		ActorId:   u.ActorID,
		TargetId:  u.TargetID,
	}
}
