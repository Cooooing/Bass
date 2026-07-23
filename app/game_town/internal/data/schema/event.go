package schema

import (
	"time"

	"common/pkg/constant"
	gameenum "game_town/internal/enum"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Event 定义不可变的世界领域事件。
type Event struct{ ent.Schema }

func (Event) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixGameTown.String() + "events"},
		entsql.WithComments(true),
	}
}

func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("world_id").Comment("世界 ID"),
		field.Uint64("sequence").Comment("世界内事件序号"),
		field.Enum("type").Values(gameenum.EventTypeMap.EnumValues()...).Comment("事件类型"),
		field.Int64("actor_player_id").Comment("事件玩家 ID").Nillable().Optional(),
		field.Int64("npc_id").Comment("关联 NPC ID").Nillable().Optional(),
		field.Int64("location_id").Comment("关联地点 ID").Nillable().Optional(),
		field.Int64("causation_event_id").Comment("原因事件 ID").Nillable().Optional(),
		field.String("summary").Comment("事件摘要").MaxRuneLen(512).Default(""),
		field.Text("content").Comment("事件内容").Default(""),
		field.JSON("payload", map[string]any{}).Comment("最小事件载荷").Optional(),
		field.Time("world_time").Comment("事件发生的世界时间").Default(time.Now),
		field.Time("occurred_at").Comment("事件发生的现实时间"),
		field.Time("created_at").Comment("记录创建时间"),
	}
}

func (Event) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("world_id", "sequence").
			Unique().
			StorageKey("game_town_events_world_sequence_unique"),
		index.Fields("world_id", "type", "sequence").
			StorageKey("game_town_events_world_type_sequence_idx"),
		index.Fields("causation_event_id", "type").
			Unique().
			StorageKey("game_town_events_causation_type_unique").
			Annotations(entsql.IndexWhere("causation_event_id IS NOT NULL")),
	}
}
