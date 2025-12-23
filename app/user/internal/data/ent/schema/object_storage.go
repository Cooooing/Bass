package schema

import (
	"common/pkg"
	"common/pkg/constant"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// ObjectStorage 用户实体定义
type ObjectStorage struct {
	ent.Schema
}

func (ObjectStorage) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixUser.String() + "object_storages"},
		entsql.WithComments(true),
	}
}

// Fields 定义表字段
func (ObjectStorage) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.String("provider").Comment("OSS 提供商，如 minio/qiniu/aliyun/tencent"),
		field.String("bucket").Comment("bucket 名称"),
		field.String("key").Comment("对象 key"),
		field.String("mime_type").Comment("文件 MIME 类型"),
		field.Int64("size").Comment("文件大小（字节）"),
		field.String("hash").Comment("文件 Hash"),

		field.Int64("upload_by").Comment("上传者ID"),
		field.String("upload_by_name").Comment("上传者名称"),

		field.String("audit_callback_reply").Comment("审核回调响应").Optional().Nillable(),
		field.Bool("blocked").Comment("是否违规").Default(false),
		field.String("blocked_reason").Comment("违规原因").Optional().Nillable(),
		field.Time("blocked_at").Comment("违规时间").Optional().Nillable(),
		field.Int64("blocked_by").Comment("违规处理人").Optional().Nillable(),
		field.String("blocked_by_name").Comment("违规处理人名称").Optional().Nillable(),
	}
	fields = append(fields, pkg.TimeAuditFields()...)
	return fields
}

func (ObjectStorage) Indexes() []ent.Index {
	return []ent.Index{}
}

func (ObjectStorage) Edges() []ent.Edge {
	return []ent.Edge{}
}
