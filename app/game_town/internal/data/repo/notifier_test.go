package repo

import (
	"testing"
	"time"

	"game_town/internal/biz/model"
)

func TestEventNotifierPublishesWakeAndEvent(
	t *testing.T,
) {
	n := NewEventNotifier().(*EventNotifier)
	all, stopAll := n.SubscribeAll()
	defer stopAll()
	events, stopEvents := n.Watch(42)
	defer stopEvents()
	event := &model.Event{
		ID:       1,
		WorldID:  42,
		Sequence: 7,
	}
	n.Publish(event)
	select {
	case worldID := <-all:
		if worldID != 42 {
			t.Fatalf("world id %d", worldID)
		}
	case <-time.After(time.Second):
		t.Fatal("wake timeout")
	}
	select {
	case got := <-events:
		if got != event {
			t.Fatal("unexpected event")
		}
	case <-time.After(time.Second):
		t.Fatal("event timeout")
	}
}
