package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User 实体定义
type User struct {
	ent.Schema
}

// Fields 定义表字段
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.Int("age").Positive(),
		field.String("email").Unique(),
		field.String("email1").Unique(),
		field.String("email2").Unique(),
	}
}
