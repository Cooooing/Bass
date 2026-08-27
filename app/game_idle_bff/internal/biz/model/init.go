package model

type WebSocketInit struct {
	ActionQueue   *ActionQueue                 `json:"action_queue"`
	Abilities     map[string]*CharacterAbility `json:"abilities"`
	BackpackItems map[string]*CharacterItem    `json:"backpack_items"`
	ChatMessages  []*WebSocketChatMessage      `json:"chat_messages"`
	ServerTime    int64                        `json:"server_time"`
}
