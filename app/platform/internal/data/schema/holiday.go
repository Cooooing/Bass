package schema

import (
	"common/pkg/constant"
	utilent "common/pkg/util/ent"

	"platform/internal/enum"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Holiday 保存按自然日缓存的节假日状态。
type Holiday struct {
	ent.Schema
}

func (Holiday) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixPlatform.String() + "holidays"},
		entsql.WithComments(true),
	}
}

func (Holiday) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Time("date").Unique().Comment("日期"),
		field.Enum("type").Values(enum.HolidayTypeMap.EnumValues()...).Comment("日期类型"),
		field.String("day_name").Comment("当天名称"),
		field.String("holiday_name").Optional().Nillable().Comment("假期名称"),
		field.Int32("days_to_holiday").Default(0).Comment("距离下一个假期的天数"),
		field.Int32("days_remaining").Default(0).Comment("当前假期剩余天数"),
	}
}

func (Holiday) Mixin() []ent.Mixin {
	return []ent.Mixin{utilent.TimeAuditMixin{}}
}
