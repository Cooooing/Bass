package usecase

import (
	"context"
	"time"

	"game_town/internal/biz/base"
	"game_town/internal/biz/model"
	"game_town/internal/biz/repo"
	"game_town/internal/enum"
)

type EventUsecase struct {
	eventRepo       repo.EventRepo
	observationRepo repo.ObservationRepo
	worldStateRepo  repo.WorldStateRepo
	worldMemberRepo repo.WorldMemberRepo
	npcRepo         repo.NpcRepo
	eventNotifier   repo.EventNotifier
}

func NewEventUsecase(
	eventRepo repo.EventRepo,
	observationRepo repo.ObservationRepo,
	worldStateRepo repo.WorldStateRepo,
	worldMemberRepo repo.WorldMemberRepo,
	npcRepo repo.NpcRepo,
	eventNotifier repo.EventNotifier,
) *EventUsecase {
	return &EventUsecase{
		eventRepo:       eventRepo,
		observationRepo: observationRepo,
		worldStateRepo:  worldStateRepo,
		worldMemberRepo: worldMemberRepo,
		npcRepo:         npcRepo,
		eventNotifier:   eventNotifier,
	}
}

type AppendEventReq struct {
	WorldID          int64
	Type             enum.EventType
	ActorPlayerID    *int64
	NpcID            *int64
	LocationID       *int64
	CausationEventID *int64
	Summary          string
	Content          string
	Payload          map[string]any
}

func (u *EventUsecase) AppendInTx(ctx context.Context, req *AppendEventReq) (*model.Event, error) {
	sequence, err := u.worldStateRepo.NextEventSequence(ctx, req.WorldID)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	state, err := u.worldStateRepo.Get(ctx, &repo.WorldStateQuery{
		WorldID: new(req.WorldID),
	})
	if err != nil {
		return nil, err
	}

	scale := state.TimeScale
	if scale == 0 {
		scale = 24
	}
	worldTime := state.WorldTime.Add(time.Duration(float64(now.Sub(state.TimeAnchor)) * float64(scale)))

	event, err := u.eventRepo.Save(ctx, &model.Event{
		WorldID:          req.WorldID,
		Sequence:         sequence,
		Type:             req.Type,
		ActorPlayerID:    req.ActorPlayerID,
		NpcID:            req.NpcID,
		LocationID:       req.LocationID,
		CausationEventID: req.CausationEventID,
		Summary:          req.Summary,
		Content:          req.Content,
		Payload:          req.Payload,
		WorldTime:        worldTime,
		OccurredAt:       now,
		CreatedAt:        now,
	})
	if err != nil {
		return nil, err
	}
	if err = u.projectObservations(ctx, event); err != nil {
		return nil, err
	}
	return event, nil
}

func (u *EventUsecase) projectObservations(ctx context.Context, event *model.Event) error {
	players := make(map[int64]enum.ObservationSource)
	npcs := make(map[int64]enum.ObservationSource)
	if event.ActorPlayerID != nil {
		players[*event.ActorPlayerID] = enum.ObservationSourceExperienced
	}
	if event.NpcID != nil {
		npcs[*event.NpcID] = enum.ObservationSourceExperienced
	}

	public, _ := event.Payload["public"].(bool)
	if public || event.LocationID != nil {
		if err := u.projectAreaObservations(ctx, event, public, players, npcs); err != nil {
			return err
		}
	}

	summary := event.Summary
	if summary == "" {
		summary = event.Content
	}
	for playerID, source := range players {
		_, err := u.observationRepo.Save(ctx, &model.Observation{
			WorldID:    event.WorldID,
			EventID:    event.ID,
			PlayerID:   new(playerID),
			Source:     source,
			Certainty:  enum.KnowledgeCertaintyConfirmed,
			Summary:    summary,
			Salience:   u.eventSalience(event.Type),
			ObservedAt: event.OccurredAt,
			WorldTime:  event.WorldTime,
		})
		if err != nil {
			return err
		}
	}
	for npcID, source := range npcs {
		_, err := u.observationRepo.Save(ctx, &model.Observation{
			WorldID:    event.WorldID,
			EventID:    event.ID,
			NpcID:      new(npcID),
			Source:     source,
			Certainty:  enum.KnowledgeCertaintyConfirmed,
			Summary:    summary,
			Salience:   u.eventSalience(event.Type),
			ObservedAt: event.OccurredAt,
			WorldTime:  event.WorldTime,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (u *EventUsecase) projectAreaObservations(ctx context.Context, event *model.Event, public bool, players map[int64]enum.ObservationSource, npcs map[int64]enum.ObservationSource) error {
	memberQuery := &repo.WorldMemberQuery{
		WorldID: new(event.WorldID),
	}
	npcQuery := &repo.NpcQuery{
		WorldID: new(event.WorldID),
	}
	if !public {
		memberQuery.LocationID = event.LocationID
		npcQuery.LocationID = event.LocationID
	}

	members, err := u.worldMemberRepo.List(ctx, memberQuery)
	if err != nil {
		return err
	}
	for _, member := range members {
		if _, ok := players[member.PlayerID]; ok {
			continue
		}
		players[member.PlayerID] = u.observationSource(public)
	}

	rows, err := u.npcRepo.List(ctx, npcQuery)
	if err != nil {
		return err
	}
	for _, npc := range rows {
		if _, ok := npcs[npc.ID]; ok {
			continue
		}
		npcs[npc.ID] = u.observationSource(public)
	}
	return nil
}

func (u *EventUsecase) observationSource(public bool) enum.ObservationSource {
	if public {
		return enum.ObservationSourcePublic
	}
	return enum.ObservationSourceWitnessed
}

func (u *EventUsecase) eventSalience(eventType enum.EventType) float64 {
	switch eventType {
	case enum.EventTypeNpcDied, enum.EventTypeLocationChanged, enum.EventTypeFactionChanged, enum.EventTypeWorldReady:
		return 1
	case enum.EventTypeNpcReplied, enum.EventTypeActionResolved, enum.EventTypeActionRejected:
		return 0.8
	default:
		return 0.5
	}
}

func (u *EventUsecase) Publish(event *model.Event) {
	if event == nil {
		return
	}
	u.eventNotifier.Publish(event)
}

type PageEventsReq struct {
	Page          base.PageRequest
	WorldID       int64
	PlayerID      int64
	AfterSequence uint64
	Type          *enum.EventType
	SkipTotal     bool
}

type PageEventsResp struct {
	Rows []*model.PerceivedEvent
	Page base.PageResp
}

func (u *EventUsecase) Page(ctx context.Context, req *PageEventsReq) (*PageEventsResp, error) {
	pageResp, err := u.observationRepo.Page(ctx, &repo.ObservationPageReq{
		Page:      req.Page,
		SkipTotal: req.SkipTotal,
		Query: repo.ObservationQuery{
			WorldID:            new(req.WorldID),
			PlayerID:           new(req.PlayerID),
			AfterEventSequence: new(req.AfterSequence),
			EventType:          req.Type,
		},
	})
	if err != nil {
		return nil, err
	}
	observations := pageResp.Rows

	observationByEvent := make(map[int64]*model.Observation, len(observations))
	eventIDs := make([]int64, 0, len(observations))
	for _, observation := range observations {
		observationByEvent[observation.EventID] = observation
		eventIDs = append(eventIDs, observation.EventID)
	}
	if len(eventIDs) == 0 {
		return &PageEventsResp{
			Rows: nil,
			Page: pageResp.Page,
		}, nil
	}

	events, err := u.eventRepo.List(ctx, &repo.EventQuery{
		IDs:     eventIDs,
		WorldID: new(req.WorldID),
	})
	if err != nil {
		return nil, err
	}

	rows := make([]*model.PerceivedEvent, 0, len(events))
	for _, event := range events {
		rows = append(rows, &model.PerceivedEvent{
			Event:       event,
			Observation: observationByEvent[event.ID],
		})
	}

	return &PageEventsResp{
		Rows: rows,
		Page: pageResp.Page,
	}, nil
}

type ListEventsAfterReq struct {
	WorldID       int64
	PlayerID      int64
	AfterSequence uint64
}

func (u *EventUsecase) ListAfter(ctx context.Context, req *ListEventsAfterReq) ([]*model.PerceivedEvent, error) {
	resp, err := u.Page(ctx, &PageEventsReq{
		Page: base.PageRequest{
			Page: 1,
			Size: 100,
		},
		WorldID:       req.WorldID,
		PlayerID:      req.PlayerID,
		AfterSequence: req.AfterSequence,
		SkipTotal:     true,
	})
	if err != nil {
		return nil, err
	}
	return resp.Rows, nil
}

func (u *EventUsecase) VisibleToPlayer(ctx context.Context, worldID int64, playerID int64, eventID int64) (*model.Observation, bool, error) {
	rows, err := u.observationRepo.List(ctx, &repo.ObservationQuery{
		WorldID:  new(worldID),
		PlayerID: new(playerID),
		EventID:  new(eventID),
	})
	if err != nil {
		return nil, false, err
	}
	if len(rows) == 0 {
		return nil, false, nil
	}
	return rows[0], true, nil
}

func (u *EventUsecase) Watch(worldID int64) (<-chan *model.Event, func()) {
	return u.eventNotifier.Watch(worldID)
}
