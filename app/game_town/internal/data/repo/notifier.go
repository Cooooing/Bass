package repo

import (
	"sync"

	"game_town/internal/biz/model"
	bizrepo "game_town/internal/biz/repo"

	"github.com/samber/lo"
)

type EventNotifier struct {
	mu     sync.Mutex
	next   uint64
	all    map[uint64]chan int64
	events map[int64]map[uint64]chan *model.Event
}

func NewEventNotifier() bizrepo.EventNotifier {
	return &EventNotifier{
		all:    map[uint64]chan int64{},
		events: map[int64]map[uint64]chan *model.Event{},
	}
}

func (n *EventNotifier) Notify(
	worldID int64,
) {
	n.mu.Lock()
	subs := lo.Values(n.all)
	n.mu.Unlock()
	for _, ch := range subs {
		select {
		case ch <- worldID:
		default:
		}
	}
}

func (n *EventNotifier) SubscribeAll() (<-chan int64, func()) {
	n.mu.Lock()
	n.next++
	id := n.next
	ch := make(chan int64, 256)
	n.all[id] = ch
	n.mu.Unlock()
	cancel := func() {
		n.mu.Lock()
		delete(n.all, id)
		n.mu.Unlock()
	}
	return ch, cancel
}

func (n *EventNotifier) Publish(
	event *model.Event,
) {
	if event == nil {
		return
	}
	n.Notify(event.WorldID)
	n.mu.Lock()
	subs := lo.Values(n.events[event.WorldID])
	n.mu.Unlock()
	for _, ch := range subs {
		select {
		case ch <- event:
		default:
		}
	}
}

func (n *EventNotifier) Watch(
	worldID int64,
) (<-chan *model.Event, func()) {
	n.mu.Lock()
	n.next++
	id := n.next
	ch := make(chan *model.Event, 128)
	if n.events[worldID] == nil {
		n.events[worldID] = map[uint64]chan *model.Event{}
	}
	n.events[worldID][id] = ch
	n.mu.Unlock()
	return ch, func() {
		n.mu.Lock()
		if subs := n.events[worldID]; subs != nil {
			delete(subs, id)
		}
		n.mu.Unlock()
	}
}
