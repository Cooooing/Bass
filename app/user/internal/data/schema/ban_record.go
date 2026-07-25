package schema

import (
	"common/pkg/constant"
	commonenum "common/pkg/enum"
	utilent "common/pkg/util/ent"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// BanRecord 追加记录账号封禁事实，用于审计和过期解封校验。
type BanRecord struct {
	ent.Schema
}

func (BanRecord) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: constant.TablePrefixUser.String() + "ban_records",
		},
		entsql.WithComments(true),
	}
}

func (BanRecord) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("user_id").Comment("被封禁账号 ID"),
		field.Int64("operator_id").Comment("操作人账号 ID"),
		field.Enum("operator_realm").Values(commonenum.LoginRealmMap.EnumValues()...).Comment("操作人登录域"),
		field.String("reason").Comment("封禁原因").MaxLen(128).NotEmpty(),
		field.Text("remark").Comment("操作备注").Default(""),
		field.Time("started_at").Comment("封禁开始时间"),
		field.Time("banned_until").Comment("封禁截止时间，空表示永久封禁").Optional().Nillable(),
	}
}

func (BanRecord) Mixin() []ent.Mixin {
	return []ent.Mixin{utilent.TimeAuditMixin{}}
}

func (BanRecord) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "created_at").StorageKey("user_ban_records_user_created_idx"),
		index.Fields("user_id", "banned_until").StorageKey("user_ban_records_user_until_idx"),
		index.Fields("operator_id", "created_at").StorageKey("user_ban_records_operator_created_idx"),
	}
}

func (BanRecord) Edges() []ent.Edge {
	return nil
}
