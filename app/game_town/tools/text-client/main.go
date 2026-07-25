package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/x/term"
)

const (
	clientModeAuto    = "auto"
	clientModeTUI     = "tui"
	clientModeConsole = "console"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	addr := flag.String("addr", "", "direct game_town gRPC address")
	defaultConsulAddress := envDefault("CONSUL_HOST", "127.0.0.1") + ":" + envDefault("CONSUL_PORT", "8500")
	consulAddr := flag.String(
		"consul-addr",
		envDefault("CONSUL_ADDRESS", defaultConsulAddress),
		"Consul address",
	)
	datacenter := flag.String(
		"consul-datacenter",
		envDefault("CONSUL_DATACENTER", "dc1"),
		"Consul datacenter",
	)
	token := flag.String("consul-token", envDefault("CONSUL_ACL_APP_TOKEN", ""), "Consul token")
	mode := flag.String("mode", clientModeAuto, "client mode: auto, tui, console")
	flag.Parse()

	resolvedMode, err := resolveClientMode(*mode, term.IsTerminal(os.Stdin.Fd()))
	if err != nil {
		return err
	}
	client, cleanup, target, err := newGameTownClient(
		*addr,
		*consulAddr,
		*datacenter,
		*token,
		5*time.Second,
	)
	if err != nil {
		return err
	}
	defer cleanup()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if resolvedMode == clientModeConsole {
		return runConsole(ctx, client, target, os.Stdin, os.Stdout)
	}
	_, err = tea.NewProgram(
		newModel(ctx, client, target),
		tea.WithAltScreen(),
	).Run()
	return err
}

func resolveClientMode(
	mode string,
	terminal bool,
) (string, error) {
	switch mode {
	case clientModeAuto:
		if terminal {
			return clientModeTUI, nil
		}
		return clientModeConsole, nil
	case clientModeTUI, clientModeConsole:
		return mode, nil
	default:
		return "", fmt.Errorf("unsupported client mode %q", mode)
	}
}
