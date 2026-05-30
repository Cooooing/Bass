package schema

import (
	"common/pkg/constant"
	userenum "user/internal/enum"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Preferences 存储账号可选偏好设置。
type Preferences struct {
	ent.Schema
}

func (Preferences) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixUser.String() + "preferences"},
		entsql.WithComments(true),
	}
}

func (Preferences) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("user_id").Comment("账号 ID").Unique(),
		field.Enum("language").Values(userenum.LanguageMap.EnumValues()...).Default(string(userenum.LanguageZhCN)).Comment("语言"),
		field.String("timezone").Comment("时区").Default("Asia/Shanghai"),
		field.String("theme").Comment("桌面端主题").Default("default"),
		field.String("mobile_theme").Comment("移动端主题").Default("default"),
	}
}

func (Preferences) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", Account.Type).Ref("preferences").Field("user_id").Required().Unique(),
	}
}
