package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// AAA 实体定义
type AAA struct {
	ent.Schema
}

// Fields 定义表字段
func (AAA) Fields() []ent.Field {
	fields := []ent.Field{
		// --- 基础信息 ---
		field.String("name").Comment("用户名").NotEmpty(),
		field.String("nickname").Comment("昵称").Optional(),
	}
	return fields
}
