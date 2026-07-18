package schema

import (
	"common/pkg/constant"
	utilent "common/pkg/util/ent"
	"platform/internal/enum"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type ObjectStorage struct {
	ent.Schema
}

func (ObjectStorage) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixPlatform.String() + "object_storages"},
		entsql.WithComments(true),
	}
}

func (ObjectStorage) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Enum("provider").Values(enum.ObjectStorageProviderMap.EnumValues()...).Comment("对象存储提供商"),
		field.String("bucket").Comment("存储桶名称"),
		field.String("key").Comment("对象键"),
		field.String("mime_type").Comment("文件 MIME 类型"),
		field.Int64("size").Comment("文件大小（字节）"),
		field.String("hash").Comment("文件哈希"),
		field.Int64("upload_by").Comment("上传者ID"),
		field.String("audit_callback_reply").Comment("审核回调响应").Optional().Nillable(),
		field.Bool("blocked").Comment("是否违规").Default(false),
		field.String("blocked_reason").Comment("违规原因").Optional().Nillable(),
		field.Time("blocked_at").Comment("违规时间").Optional().Nillable(),
		field.Int64("blocked_by").Comment("违规处理人").Optional().Nillable(),
	}
	return fields
}

func (ObjectStorage) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (ObjectStorage) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("provider", "bucket", "key").Unique(),
	}
}
