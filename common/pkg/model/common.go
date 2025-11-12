package model

type Range[T comparable] struct {
	Start *T `json:"start"`
	End   *T `json:"end"`
}
