package model

import schedulerenum "scheduler/internal/enum"

type AvailableTask struct {
	HandlerName schedulerenum.TaskHandlerName
	Title       string
	Description string
}
