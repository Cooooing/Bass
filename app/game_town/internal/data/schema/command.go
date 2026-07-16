package schema

import (
	"common/pkg/constant"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Command 保存文字命令解析和执行结果。
type Command struct{ ent.Schema }

func (Command) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: constant.TablePrefixGameTown.String() + "commands"}, entsql.WithComments(true)}
}
func (Command) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("world_id").Comment("世界 ID").Nillable().Optional(),
		field.Int64("session_id").Comment("会话 ID"),
		field.Int64("player_id").Comment("玩家 ID").Nillable().Optional(),
		field.Text("raw_text").Comment("原始命令"),
		field.String("type").Comment("命令类型").MaxLen(64).NotEmpty(),
		field.JSON("parsed_payload", map[string]any{}).Comment("解析结果").Optional(),
		field.String("status").Comment("命令状态").MaxLen(32).NotEmpty(),
		field.Int32("error_code").Comment("错误码").Nillable().Optional(),
		field.String("result_summary").Comment("执行摘要").MaxLen(512).Default(""),
		field.Time("created_at").Comment("创建时间"),
		field.Time("handled_at").Comment("处理时间").Nillable().Optional(),
	}
}
func (Command) Indexes() []ent.Index {
	return []ent.Index{index.Fields("session_id", "created_at").StorageKey("game_town_commands_session_time_idx"), index.Fields("world_id", "player_id", "created_at").StorageKey("game_town_commands_world_player_time_idx")}
}
func (Command) Edges() []ent.Edge { return nil }
