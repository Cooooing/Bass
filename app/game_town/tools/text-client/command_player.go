package main

import (
	"context"
	"fmt"
	"strconv"

	"common/pkg/client/rpc"
	v1 "common/proto/gen/game_town/v1"
)

func usePlayer(
	ctx context.Context,
	client *rpc.GameTownClient,
	value string,
) commandResult {
	playerID, err := strconv.ParseInt(value, 10, 64)
	if err != nil || playerID <= 0 {
		return commandUsage("/player use <player_id>")
	}
	reply, err := client.Player.Get(ctx, &v1.GetGameTownPlayer_Request{Id: playerID})
	if err != nil {
		return commandResult{err: err}
	}
	return commandResult{
		playerID: playerID,
		lines: []string{fmt.Sprintf(
			"using player %s (%d)",
			reply.GetRow().GetDisplayName(),
			playerID,
		)},
	}
}
func registerPlayer(
	ctx context.Context,
	client *rpc.GameTownClient,
	name string,
) commandResult {
	reply, err := client.Player.Register(
		ctx,
		&v1.RegisterGameTownPlayer_Request{
			Name:        name,
			DisplayName: name,
		},
	)
	if err != nil {
		return commandResult{err: err}
	}
	return commandResult{
		playerID: reply.GetRow().GetId(),
		lines: []string{fmt.Sprintf(
			"registered %s (%d)",
			reply.GetRow().GetDisplayName(),
			reply.GetRow().GetId(),
		)},
	}
}
