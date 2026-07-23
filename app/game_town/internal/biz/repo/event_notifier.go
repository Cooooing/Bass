package repo

import "game_town/internal/biz/model"

type EventNotifier interface {
	Notify(worldID int64)
	SubscribeAll() (<-chan int64, func())
	Publish(event *model.Event)
	Watch(worldID int64) (<-chan *model.Event, func())
}
