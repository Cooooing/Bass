package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strings"
	"sync"

	"common/pkg/client/rpc"
)

type consolePrinter struct {
	mu     sync.Mutex
	writer io.Writer
}

func (p *consolePrinter) Println(value string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	fmt.Fprintln(p.writer, value)
}

type consoleSession struct {
	mu          sync.Mutex
	dialogNpcID int64
	suggestions []suggestedChoice
}

func (s *consoleSession) Dialog() (int64, []suggestedChoice) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.dialogNpcID, append([]suggestedChoice(nil), s.suggestions...)
}

func (s *consoleSession) SetDialog(npcID int64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.dialogNpcID = npcID
	s.suggestions = nil
}

func (s *consoleSession) ClearDialog() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.dialogNpcID = 0
	s.suggestions = nil
}

func (s *consoleSession) ClearSuggestions() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.suggestions = nil
}

func (s *consoleSession) SetSuggestions(values []suggestedChoice) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.suggestions = append([]suggestedChoice(nil), values...)
}

func (s *consoleSession) SetDialogSuggestions(npcID int64, values []suggestedChoice) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if npcID > 0 {
		s.dialogNpcID = npcID
	}
	s.suggestions = append([]suggestedChoice(nil), values...)
}

func runConsole(ctx context.Context, client *rpc.GameTownClient, target string, input io.Reader, output io.Writer) error {
	printer := &consolePrinter{
		writer: output,
	}
	session := &consoleSession{}
	printer.Println("Game Town Console connected: " + target)
	printer.Println("Console line mode enabled. Type /help for commands. Non-command text is sent as free action after joining a world.")

	scanner := bufio.NewScanner(input)
	scanner.Buffer(make([]byte, 4096), 1024*1024)
	var playerID int64
	var worldID int64
	var watchCancel context.CancelFunc
	defer func() {
		if watchCancel != nil {
			watchCancel()
		}
	}()

	for scanner.Scan() {
		raw := strings.TrimSpace(scanner.Text())
		if raw == "" {
			continue
		}
		if raw == "/quit" || raw == "/exit" {
			return nil
		}
		dialogNpcID, suggestions := session.Dialog()
		result := executeCommand(ctx, client, playerID, worldID, dialogNpcID, suggestions, raw)
		printCommandResult(printer, result)
		if result.playerID > 0 {
			playerID = result.playerID
		}
		if result.dialogNpcID > 0 {
			session.SetDialog(result.dialogNpcID)
		}
		if result.clearDialog {
			session.ClearDialog()
		}
		if result.clearSuggestions {
			session.ClearSuggestions()
		}
		if result.worldID > 0 && result.worldID != worldID {
			worldID = result.worldID
			if watchCancel != nil {
				watchCancel()
			}
			var events <-chan eventResult
			watchCancel, events = startWatcher(ctx, client, playerID, worldID, 0)
			go printConsoleEvents(ctx, printer, session, events)
		}
	}
	return scanner.Err()
}

func printCommandResult(printer *consolePrinter, result commandResult) {
	if result.err != nil {
		printer.Println("error: " + result.err.Error())
		return
	}
	for _, line := range result.lines {
		printer.Println(line)
	}
}

func printConsoleEvents(ctx context.Context, printer *consolePrinter, session *consoleSession, events <-chan eventResult) {
	for {
		select {
		case <-ctx.Done():
			return
		case result, ok := <-events:
			if !ok {
				return
			}
			if result.err != nil {
				printer.Println("event stream: " + result.err.Error())
				continue
			}
			if result.event == nil {
				continue
			}
			choices := eventSuggestedChoices(result.event)
			if len(choices) > 0 {
				session.SetDialogSuggestions(result.event.GetNpcId(), choices)
			}
			printer.Println(formatEventLine(result.event))
		}
	}
}
