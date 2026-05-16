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

// UserRelation 用户实体定义
type UserRelation struct {
	ent.Schema
}

func (UserRelation) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixUser.String() + "follows"},
		entsql.WithComments(true),
	}
}

// Fields 定义表字段
func (UserRelation) Fields() []ent.Field {
	fields := []ent.Field{
		// --- 基础信息 ---
		field.Int64("id").Immutable().Unique(),
		field.Int64("actor_id").Comment("关系发起者"),
		field.Int64("target_id").Comment("关系目标用户"),
		field.Int32("type").Comment("关系类型 0-follow 1-block"),
	}
	return fields
}

func (UserRelation) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (UserRelation) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("actor_id", "target_id", "type").Unique(),
		index.Fields("actor_id", "type"),
		index.Fields("target_id", "type"),
	}
}

func (UserRelation) Edges() []ent.Edge {
	return []ent.Edge{
		// 关联用户关系发起者 多对一
		edge.From("actor", User.Type).Ref("relations_as_actor").Field("actor_id").Required().Unique(),
		// 关联用户关系目标 多对一
		edge.From("target", User.Type).Ref("relations_as_target").Field("target_id").Required().Unique(),
	}
}
