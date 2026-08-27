package model

import "game_idle/internal/enum"

// Region 是前端展示区域配置，不参与行动调度。
type Region struct {
	ID          string
	Name        string
	Description string
	ActionKind  enum.ActionKind
	Enabled     bool
	Sort        int32
}
