package schema

import (
	"common/pkg"
	"common/pkg/constant"

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
		field.String("nickname").Comment("群内昵称").Default(""),
		field.Int32("role").Comment("角色: 0-成员, 1-管理员, 2-群主").Default(0),
	}
	fields = append(fields, pkg.UserAuditFields()...)
	fields = append(fields, pkg.TimeAuditFields()...)
	fields = append(fields, pkg.UsernameAuditFields()...)
	return fields
}

func (ChatGroupMember) Edges() []ent.Edge {
	return []ent.Edge{
		// 指向所属群组 (M2O)
		edge.From("group", ChatGroup.Type).Ref("members").Unique().Required().Field("group_id"),
	}
}

func (ChatGroupMember) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("group_id", "user_id").Unique(),
	}
}
