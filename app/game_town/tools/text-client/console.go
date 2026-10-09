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
	errors chan<- error
}

func (p *consolePrinter) Println(value string) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	_, err := fmt.Fprintln(p.writer, value)
	return err
}

func (p *consolePrinter) reportError(err error) {
	select {
	case p.errors <- err:
	default:
	}
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
	consoleCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	printErrors := make(chan error, 1)
	printer := &consolePrinter{
		writer: output,
		errors: printErrors,
	}
	session := &consoleSession{}
	if err := printer.Println("Game Town Console connected: " + target); err != nil {
		return err
	}
	if err := printer.Println("Console line mode enabled. Type /help for commands. Non-command text is sent as free action after joining a world."); err != nil {
		return err
	}

	scanner := bufio.NewScanner(input)
	scanner.Buffer(make([]byte, 4096), 1024*1024)
	lines := make(chan string)
	scanErrors := make(chan error, 1)
	go scanConsoleLines(consoleCtx, scanner, lines, scanErrors)
	var playerID int64
	var worldID int64
	var watchCancel context.CancelFunc
	defer func() {
		if watchCancel != nil {
			watchCancel()
		}
	}()

	for {
		select {
		case err := <-printErrors:
			return err
		case err := <-scanErrors:
			return err
		case raw, ok := <-lines:
			if !ok {
				return nil
			}
			raw = strings.TrimSpace(raw)
			if raw == "" {
				continue
			}
			if raw == "/quit" || raw == "/exit" {
				return nil
			}
			dialogNpcID, suggestions := session.Dialog()
			result := executeCommand(ctx, client, playerID, worldID, dialogNpcID, suggestions, raw)
			if err := printCommandResult(printer, result); err != nil {
				return err
			}
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
				watchCancel, events = startWatcher(consoleCtx, client, playerID, worldID, 0)
				go printConsoleEvents(consoleCtx, printer, session, events)
			}
		}
	}
}

func scanConsoleLines(ctx context.Context, scanner *bufio.Scanner, lines chan<- string, scanErrors chan<- error) {
	defer close(lines)
	for scanner.Scan() {
		select {
		case lines <- scanner.Text():
		case <-ctx.Done():
			return
		}
	}
	if err := scanner.Err(); err != nil {
		scanErrors <- err
	}
}

func printCommandResult(printer *consolePrinter, result commandResult) error {
	if result.err != nil {
		return printer.Println("error: " + result.err.Error())
	}
	for _, line := range result.lines {
		if err := printer.Println(line); err != nil {
			return err
		}
	}
	return nil
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
				if err := printer.Println("event stream: " + result.err.Error()); err != nil {
					printer.reportError(err)
					return
				}
				continue
			}
			if result.event == nil {
				continue
			}
			choices := eventSuggestedChoices(result.event)
			if len(choices) > 0 {
				session.SetDialogSuggestions(result.event.GetNpcId(), choices)
			}
			if err := printer.Println(formatEventLine(result.event)); err != nil {
				printer.reportError(err)
				return
			}
		}
	}
}
