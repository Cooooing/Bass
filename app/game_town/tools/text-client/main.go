package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"

	commonClient "common/pkg/client"
	"common/pkg/client/rpc"
	"common/pkg/constant"
	"common/proto/gen/common"
	v1 "common/proto/gen/game_town/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/durationpb"
)

func main() {
	addr := flag.String("addr", "", "direct game_town gRPC address; empty means Consul discovery")
	consulAddr := flag.String("consul-addr", envDefault("CONSUL_ADDRESS", envDefault("CONSUL_HOST", "127.0.0.1")+":"+envDefault("CONSUL_PORT", "8500")), "Consul address")
	consulDatacenter := flag.String("consul-datacenter", envDefault("CONSUL_DATACENTER", "dc1"), "Consul datacenter")
	consulToken := flag.String("consul-token", envDefault("CONSUL_ACL_APP_TOKEN", ""), "Consul ACL token")
	consulTimeout := flag.Duration("consul-timeout", 5*time.Second, "Consul dial timeout")
	flag.Parse()

	gameTownClient, cleanup, target, err := newGameTownClient(*addr, *consulAddr, *consulDatacenter, *consulToken, *consulTimeout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "connect game_town gRPC server failed: target=%s err=%v\n", target, err)
		fmt.Fprintln(os.Stderr, "start app/game_town server first, or pass -addr for direct gRPC connection")
		os.Exit(1)
	}
	defer cleanup()

	ctx := context.Background()
	sessionReply, err := gameTownClient.Session.Start(ctx, &v1.StartGameTownSession_Request{
		PlayerId:   new(int64(1)),
		ClientType: v1.GameTownSessionClientType_GAME_TOWN_SESSION_CLIENT_TYPE_TEXT_CLIENT,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "connect game_town gRPC server failed: target=%s err=%v\n", target, err)
		fmt.Fprintln(os.Stderr, "start app/game_town server first, or pass -addr for direct gRPC connection")
		os.Exit(1)
	}
	sessionID := sessionReply.GetRow().GetId()
	var playerID int64
	fmt.Println("game_town text client started. use /register <name>, /world create ..., /quit")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		if line == "/quit" || line == "/exit" {
			break
		}
		if strings.HasPrefix(line, "/register ") {
			name := strings.TrimSpace(strings.TrimPrefix(line, "/register "))
			reply, err := gameTownClient.Player.Register(ctx, &v1.RegisterGameTownPlayer_Request{Name: name, DisplayName: name})
			if err != nil {
				fmt.Println("error:", err)
				continue
			}
			playerID = reply.GetRow().GetId()
			fmt.Printf("registered: %s (%d)\n", reply.GetRow().GetDisplayName(), playerID)
			continue
		}
		if playerID == 0 {
			fmt.Println("please register first: /register <name>")
			continue
		}
		reply, err := gameTownClient.Command.Execute(ctx, &v1.ExecuteGameTownCommand_Request{SessionId: sessionID, PlayerId: playerID, RawText: line})
		if err != nil {
			fmt.Println("error:", err)
			continue
		}
		for _, item := range reply.GetFeedbackLines() {
			fmt.Println(item)
		}
		if world := reply.GetCurrentWorld(); world != nil && world.GetCode() != "" {
			fmt.Printf("world: %s (%s)\n", world.GetName(), world.GetCode())
		}
		if loc := reply.GetCurrentLocation(); loc != nil && loc.GetCode() != "" {
			fmt.Printf("location: %s (%s)\n", loc.GetName(), loc.GetCode())
		}
		for _, npc := range reply.GetVisibleNpcs() {
			fmt.Printf("npc: %s (%s)\n", npc.GetName(), npc.GetCode())
		}
	}
	_, _ = gameTownClient.Session.End(ctx, &v1.EndGameTownSession_Request{Id: sessionID})
}

func newGameTownClient(addr string, consulAddr string, datacenter string, token string, timeout time.Duration) (*rpc.GameTownClient, func(), string, error) {
	if addr != "" {
		conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return nil, nil, addr, err
		}
		return rpc.NewGameTownClient(conn), func() { _ = conn.Close() }, addr, nil
	}

	target := "consul://" + consulAddr + "/" + constant.GameTownServiceName.String()
	consulClient, cleanup, err := commonClient.NewConsulClient(slog.Default(), &common.Consul{
		Address:     consulAddr,
		Datacenter:  datacenter,
		Token:       token,
		DialTimeout: durationpb.New(timeout),
	}, nil)
	if err != nil {
		return nil, nil, target, err
	}
	gameTownClient, err := rpc.ProvideGameTownClient(consulClient)
	if err != nil {
		cleanup()
		return nil, nil, target, err
	}
	return gameTownClient, cleanup, target, nil
}

func envDefault(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
