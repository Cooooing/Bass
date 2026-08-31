package schema

import (
	"common/pkg/constant"
	utilent "common/pkg/util/ent"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// CharacterActionQueue 定义角色行动队列冷备。
type CharacterActionQueue struct {
	ent.Schema
}

func (CharacterActionQueue) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: constant.TablePrefixGameIdle.String() + "character_action_queues",
		},
		entsql.WithComments(true),
	}
}

func (CharacterActionQueue) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("character_id").Comment("角色 ID"),
		field.String("queue_item_id").Comment("队列项编码"),
		field.String("action_id").Comment("行动编码"),
		field.Int64("times").Comment("剩余执行次数，-1 表示无限执行"),
		field.Int32("position").Comment("队列位置"),
		field.Time("queued_at").Comment("队列项创建时间"),
	}
}

func (CharacterActionQueue) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (CharacterActionQueue) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("character_id", "position").
			Unique().
			StorageKey("game_idle_character_action_queues_character_position_unique"),
		index.Fields("character_id", "queue_item_id").
			Unique().
			StorageKey("game_idle_character_action_queues_character_item_unique"),
		index.Fields("action_id").
			StorageKey("game_idle_character_action_queues_action_idx"),
	}
}

func (CharacterActionQueue) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("character", Character.Type).
			Ref("action_queues").
			Field("character_id").
			Unique().
			Required(),
		edge.From("action", Action.Type).
			Ref("character_action_queues").
			Field("action_id").
			Unique().
			Required(),
	}
}
