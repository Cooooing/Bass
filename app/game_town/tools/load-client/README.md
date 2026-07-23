# Game Town Load Client

`load-client` is a real gRPC load and acceptance client for Game Town.
It does not call internal usecases and does not write the database directly. Every step goes through public gRPC APIs.

The flow is:

1. Create or reuse an Ollama `AgentConfig`.
2. Register at least two players.
3. Create a new world and wait for `world_ready`.
4. Join all players and wait for `player_character_ready`.
5. Start `EventService.Watch` for each player and keep reconnecting with the last sequence.
6. Submit many `SubmitAction` calls, mixing free actions, movement, NPC talk, and major events.
7. Use `EventService.Page` as a durable backfill path and collect acceptance counters.

## Smoke test

```powershell
go run ./tools/load-client -addr 192.168.100.1:9105 -rounds 10 -big-event-every 5 -timeout 10m
```

## Strict smoke test

```powershell
go run ./tools/load-client -addr 192.168.100.1:9105 -rounds 20 -big-event-every 5 -timeout 20m -strict
```

## Full 1000-round acceptance run

```powershell
go run ./tools/load-client -addr 192.168.100.1:9105 -players 2 -rounds 1000 -big-event-every 100 -timeout 6h -poll-interval 2s -strict
```

With Consul service discovery:

```powershell
go run ./tools/load-client -consul-addr consul.dev.bass.local:80 -players 2 -rounds 1000 -big-event-every 100 -timeout 6h -strict
```

## Useful flags

- `-addr`: direct game_town gRPC address.
- `-consul-addr`: Consul address, used when `-addr` is empty.
- `-players`: number of players; minimum is 2.
- `-rounds`: submitted action rounds; use 1000 for acceptance.
- `-big-event-every`: inject one major event every N rounds.
- `-ollama-url`: Ollama base URL, default `http://192.168.100.10:31434`.
- `-model`: chat model, default `qwen3:1.7b`.
- `-world-description`: custom world description.
- `-timeout`: whole run timeout.
- `-strict`: fail the run when key acceptance counters are missing.

## Acceptance counters

Check the final report:

- `submitted_actions` reaches the requested rounds.
- `big_events` matches the injection cadence.
- `stream_events` is greater than zero, proving `EventService.Watch` delivered real events.
- `stream_reconnects` and `stream_errors` stay explainable.
- `npc_replied` grows during NPC talk rounds.
- `world_evolved` or `npc_planned` appears during long runs, proving autonomous world evolution.
- `agent_job_failed` remains explainable and visible.
- `page_events` continues to grow, proving durable history backfill works.