package model

type User struct {
	// ID 是账号 ID。
	ID int64 `json:"id,omitempty"`
	// Name 是账号名。
	Name string `json:"name,omitempty"`
	// Nickname 是展示昵称。
	Nickname string `json:"nickname,omitempty"`
	// Language 是首选语言。
	Language string `json:"language,omitempty"`
	// Timezone 是首选时区。
	Timezone string `json:"timezone,omitempty"`
}
