package schema

import (
	"common/pkg/constant"
	utilent "common/pkg/util/ent"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type ChatGroupMember struct {
	ent.Schema
}

func (ChatGroupMember) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixIM.String() + "chat_group_members"},
	}
}

func (ChatGroupMember) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("group_id").Comment("群id"),
		field.Int64("user_id").Comment("成员id"),
		field.String("nickname").Comment("群内昵称").Optional().Nillable(),
		field.Int32("role").Comment("角色: 1-成员, 2-管理员, 3-群主").Default(1),
		field.Time("mute_end_at").Comment("禁言结束时间").Optional().Nillable(),
	}
	return fields
}

func (ChatGroupMember) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.UserAuditMixin{},
	}
}

func (ChatGroupMember) Edges() []ent.Edge {
	return []ent.Edge{
		// 关联群组 多对一
		edge.From("group", ChatGroup.Type).Ref("members").Unique().Required().Field("group_id"),
	}
}

func (ChatGroupMember) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id").Unique(),
		index.Fields("group_id").Unique(),
	}
}
