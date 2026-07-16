package schema

import (
	"common/pkg/constant"
	utilent "common/pkg/util/ent"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Location 保存随机世界内的地点。
type Location struct{ ent.Schema }

func (Location) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: constant.TablePrefixGameTown.String() + "locations"}, entsql.WithComments(true)}
}
func (Location) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("world_id").Comment("世界 ID"),
		field.String("code").Comment("地点编码").MaxLen(64).NotEmpty(),
		field.String("name").Comment("地点名称").MaxLen(128).NotEmpty(),
		field.Text("description").Comment("地点描述"),
		field.JSON("tags", map[string]any{}).Comment("标签").Optional(),
		field.Int32("sort").Comment("排序").Default(0),
		field.Bool("enabled").Comment("是否启用").Default(true),
	}
}
func (Location) Mixin() []ent.Mixin {
	return []ent.Mixin{utilent.TimeAuditMixin{}, utilent.SoftDeleteMixin{}}
}
func (Location) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("world_id", "code").Unique().StorageKey("game_town_locations_world_code_active_unique").Annotations(entsql.IndexWhere("deleted_at IS NULL")),
		index.Fields("world_id", "enabled", "sort").StorageKey("game_town_locations_world_enabled_sort_idx").Annotations(entsql.IndexWhere("deleted_at IS NULL")),
	}
}
func (Location) Edges() []ent.Edge { return nil }
