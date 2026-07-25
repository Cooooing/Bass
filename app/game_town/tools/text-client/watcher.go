package main

import (
	"context"
	"time"

	"common/pkg/client/rpc"
	v1 "common/proto/gen/game_town/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type eventResult struct {
	event *v1.WatchGameTownEvents_Resp
	err   error
}

func startWatcher(parent context.Context, client *rpc.GameTownClient, playerID int64, worldID int64, after uint64) (context.CancelFunc, <-chan eventResult) {
	ctx, cancel := context.WithCancel(parent)
	out := make(chan eventResult, 128)
	go func() {
		defer close(out)
		lastSequence := after
		for {
			if ctx.Err() != nil {
				return
			}
			stream, err := client.Event.Watch(ctx, &v1.WatchGameTownEvents_Request{
				WorldId:       worldID,
				PlayerId:      playerID,
				AfterSequence: lastSequence,
			})
			if err != nil {
				if retryWatcher(ctx, err) {
					continue
				}
				publishWatcherResult(ctx, out, eventResult{
					err: err,
				})
				return
			}
			for {
				event, err := stream.Recv()
				if err != nil {
					if ctx.Err() != nil {
						return
					}
					if retryWatcher(ctx, err) {
						break
					}
					publishWatcherResult(ctx, out, eventResult{
						err: err,
					})
					return
				}
				if event.GetSequence() > lastSequence {
					lastSequence = event.GetSequence()
				}
				if !publishWatcherResult(ctx, out, eventResult{
					event: event,
				}) {
					return
				}
			}
		}
	}()
	return cancel, out
}

func retryWatcher(ctx context.Context, err error) bool {
	if !isRetryableWatcherError(err) {
		return false
	}
	timer := time.NewTimer(time.Second)
	defer timer.Stop()
	select {
	case <-timer.C:
		return true
	case <-ctx.Done():
		return false
	}
}

func isRetryableWatcherError(err error) bool {
	switch status.Code(err) {
	case codes.Unavailable,
		codes.DeadlineExceeded,
		codes.ResourceExhausted,
		codes.Aborted,
		codes.Internal:
		return true
	default:
		return false
	}
}

func publishWatcherResult(ctx context.Context, out chan<- eventResult, result eventResult) bool {
	select {
	case out <- result:
		return true
	case <-ctx.Done():
		return false
	}
}
