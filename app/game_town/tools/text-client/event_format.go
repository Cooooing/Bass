package main

import (
	v1enum "common/proto/gen/game_town/v1/enum"
	"fmt"
	"strings"

	v1 "common/proto/gen/game_town/v1"
)

func formatEventLine(event *v1.WatchGameTownEvents_Resp) string {
	line := fmt.Sprintf("[%d] %s · %s", event.GetSequence(), eventName(event.GetType()), event.GetSummary())
	if event.GetContent() != "" {
		line += "\n  " + event.GetContent()
	}
	if len(event.GetSuggestedActions()) > 0 {
		line += "\n  可选回答："
		for index, action := range event.GetSuggestedActions() {
			label := strings.TrimSpace(action.GetLabel())
			content := strings.TrimSpace(action.GetContent())
			if label == "" {
				label = content
			}
			line += fmt.Sprintf("\n  %d. %s", index+1, label)
		}
		line += "\n  也可以直接输入任何内容。"
	}
	return line
}

func eventSuggestedChoices(event *v1.WatchGameTownEvents_Resp) []suggestedChoice {
	choices := make([]suggestedChoice, 0, len(event.GetSuggestedActions()))
	for _, action := range event.GetSuggestedActions() {
		content := strings.TrimSpace(action.GetContent())
		if content == "" {
			continue
		}
		targets := make([]*v1.SubmitGameTownAction_Request_EntityRef, 0, len(action.GetTargets()))
		for _, target := range action.GetTargets() {
			if target.GetId() <= 0 || target.GetType() == v1enum.GameTownEntityType_GAME_TOWN_ENTITY_TYPE_UNSPECIFIED {
				continue
			}
			targets = append(targets, &v1.SubmitGameTownAction_Request_EntityRef{
				Type: target.GetType(),
				Id:   target.GetId(),
			})
		}
		if len(targets) == 0 && event.GetNpcId() > 0 {
			targets = append(targets, npcTarget(event.GetNpcId()))
		}
		choices = append(choices, suggestedChoice{
			label:   strings.TrimSpace(action.GetLabel()),
			content: content,
			targets: targets,
		})
	}
	return choices
}
