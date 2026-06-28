# Ent Schema 模板

schema 字段必须来自明确业务需求或实际查询、投递、审计需求。

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

func (Article) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}
```

## 必须

- 字段使用明确中文注释。
- 业务固定取值使用枚举。
- 唯一约束对应真实业务幂等或唯一性。
- PostgreSQL 软删除唯一约束使用 `WHERE deleted_at IS NULL` 的部分唯一索引。

## 禁止

- 意义不明、无读写路径、可推导或重复的信息落库。
- 业务表保留 `deleted_by`。
- 用 `deleted` 枚举状态表达逻辑删除。
- 无设计确认就新增或删除字段。
