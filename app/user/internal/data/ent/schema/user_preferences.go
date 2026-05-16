package schema

import (
	"common/pkg/constant"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserPreferences 用户偏好设置
type UserPreferences struct {
	ent.Schema
}

func (UserPreferences) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixUser.String() + "user_preferences"},
		entsql.WithComments(true),
	}
}

func (UserPreferences) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("user_id").Comment("用户ID").Unique(),
		field.String("language").Comment("用户语言").Default("zh-CN").Nillable(),
		field.String("timezone").Comment("时区").Default("Asia/Shanghai").Nillable(),
		field.String("theme").Comment("桌面皮肤主题").Default("default").Nillable(),
		field.String("mobile_theme").Comment("移动端皮肤主题").Default("default").Nillable(),
		field.Bool("enable_web_notify").Comment("启用Web通知").Default(true).Nillable(),
		field.Bool("enable_email_subscribe").Comment("启用邮件订阅").Default(true).Nillable(),
	}
}

func (UserPreferences) Indexes() []ent.Index {
	return nil
}

func (UserPreferences) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("preferences").Field("user_id").Required().Unique(),
	}
}
