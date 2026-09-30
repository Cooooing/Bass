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

// Asset is a completed, content-addressed storage object. It deliberately has
// no business purpose, owner or access level: those facts belong to references
// owned by user, content, IM and other domains.
type Asset struct {
	ent.Schema
}

func (Asset) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixPlatform.String() + "assets"},
		entsql.WithComments(true),
	}
}

func (Asset) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Enum("provider").Values(enum.AssetProviderMap.EnumValues()...).Comment("对象存储提供商"),
		field.String("bucket").Comment("存储桶名称"),
		field.String("object_key").Comment("对象键"),
		field.String("hash").Optional().Nillable().Comment("SHA-256 内容摘要"),
		field.Int64("upload_by").Optional().Nillable().Comment("上传者 ID"),
		field.String("provider_etag").Optional().Nillable().Comment("存储提供商对象标识"),
		field.String("mime_type").Comment("对象 MIME 类型"),
		field.Int64("size").Comment("对象大小"),
		field.Enum("status").Values(enum.AssetStatusMap.EnumValues()...).Default(enum.AssetStatusAvailable.String()).Comment("资源状态"),
		field.String("blocked_reason").Optional().Nillable().Comment("封禁原因"),
		field.Time("blocked_at").Optional().Nillable().Comment("封禁时间"),
		field.Int64("blocked_by").Optional().Nillable().Comment("封禁处理人"),
	}
}

func (Asset) Mixin() []ent.Mixin {
	// Assets must retain ordinary creation and modification audit facts even
	// though their future storage-cleanup lifecycle is modeled explicitly.
	return []ent.Mixin{utilent.TimeAuditMixin{}}
}

func (Asset) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("provider", "bucket", "object_key").Unique(),
		// PostgreSQL unique indexes permit multiple NULL values. A normal unique
		// index therefore keeps legacy rows without a SHA-256 while letting Ent's
		// ON CONFLICT (hash) safely implement content-addressed creation.
		index.Fields("hash").Unique(),
		index.Fields("status"),
	}
}
