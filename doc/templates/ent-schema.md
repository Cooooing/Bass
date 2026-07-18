# Ent Schema 模板

schema 字段必须来自明确业务需求、查询路径、投递需求或审计需求。

```go
type Article struct {
	ent.Schema
}

func (Article) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixContent.String() + "articles"},
		entsql.WithComments(true),
	}
}

func (Article) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.String("title").Comment("标题").NotEmpty(),
		field.Int64("created_by").Comment("创建人 ID").Nillable().Optional(),
		field.Int64("updated_by").Comment("更新人 ID").Nillable().Optional(),
	}
}

func (Article) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("title").StorageKey("content_articles_title_idx"),
	}
}

func (Article) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}
```

## 必须

- 表名使用归属服务表前缀。
- 字段使用明确中文注释。
- 业务固定取值使用枚举。`field.Enum` 必须通过对应 `xxxMap.EnumValues()` 填充 `Values`，例如 `field.Enum("provider").Values(enum.ObjectStorageProviderMap.EnumValues()...).Comment("对象存储提供商")`，禁止手写枚举值列表或只使用 `GoType` 承载数据库枚举值；需要绑定内部类型时可以同时使用 `GoType`，但 `Values` 仍然必须来自 `EnumValues()`。
- 唯一约束对应真实业务幂等或唯一性。
- 索引只服务实际查询、唯一约束、补偿扫描或审计路径。
- 逻辑删除统一使用 `deleted_at`；软删除唯一约束使用 `WHERE deleted_at IS NULL` 的部分唯一索引。
- schema 变更前说明设计、影响和回滚方式。

## 禁止

- 意义不明、无读写路径、可推导或重复的信息落库。
- 为了未来可能查询提前增加索引。
- 业务表保留 `deleted_by`。
- 用 `deleted` 枚举状态表达逻辑删除。
- 无设计确认就新增、删除或改名字段。