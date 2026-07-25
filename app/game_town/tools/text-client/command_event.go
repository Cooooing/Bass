package main

import (
	"context"
	"fmt"

	"common/pkg/client/rpc"
	"common/proto/gen/common"
	v1 "common/proto/gen/game_town/v1"

	"github.com/samber/lo"
)

func listEvents(
	ctx context.Context,
	client *rpc.GameTownClient,
	playerID int64,
	worldID int64,
) commandResult {
	if err := requireWorldContext(playerID, worldID); err != nil {
		return commandResult{
			err: err,
		}
	}
	reply, err := client.Event.Page(ctx, &v1.PageGameTownEvents_Request{
		Page: &common.PageReq{
			Page: 1,
			Size: 20,
		},
		WorldId:  worldID,
		PlayerId: playerID,
	})
	if err != nil {
		return commandResult{
			err: err,
		}
	}
	if len(reply.GetRows()) == 0 {
		return commandResult{
			lines: []string{"暂无世界事件"},
		}
	}
	return commandResult{
		lines: lo.Map(reply.GetRows(), func(row *v1.PageGameTownEvents_Resp_Row, _ int) string {
			return fmt.Sprintf("[%d] %s %s", row.GetSequence(), eventName(row.GetType()), row.GetSummary())
		}),
	}
}
