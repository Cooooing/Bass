package task

import "context"

type Task interface {
	Name() string
	Title() string
	Description() string
	Execute(ctx context.Context, payload string) error
}
