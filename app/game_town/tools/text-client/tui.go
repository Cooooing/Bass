package main

import (
	"context"
	"fmt"
	"strings"

	"common/pkg/client/rpc"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	ctx          context.Context
	client       *rpc.GameTownClient
	target       string
	input        textinput.Model
	viewport     viewport.Model
	spinner      spinner.Model
	lines        []string
	history      []string
	historyIndex int
	busy         bool
	playerID     int64
	worldID      int64
	lastSequence uint64
	dialogNpcID  int64
	suggestions  []suggestedChoice
	watchCancel  context.CancelFunc
	eventCh      <-chan eventResult
}

var titleStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("205"))

var statusStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("241"))

var eventStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("86"))

var errorStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("196"))

func newModel(
	ctx context.Context,
	client *rpc.GameTownClient,
	target string,
) model {
	input := textinput.New()
	input.Placeholder = "输入 /help 查看命令"
	input.Focus()
	input.CharLimit = 2000
	input.ShowSuggestions = true
	input.SetSuggestions(commandSuggestions())

	spin := spinner.New()
	spin.Spinner = spinner.Dot
	result := model{
		ctx:          ctx,
		client:       client,
		target:       target,
		input:        input,
		viewport:     viewport.New(80, 20),
		spinner:      spin,
		historyIndex: -1,
		lines: []string{
			"Game Town TUI connected: " + target,
			"Type /help for commands.",
		},
	}
	result.refresh()
	return result
}

func (m model) Init() tea.Cmd {
	return tea.Batch(textinput.Blink, m.spinner.Tick)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.viewport.Width = max(1, msg.Width-2)
		m.viewport.Height = max(5, msg.Height-5)
		m.refresh()
		return m, nil
	case tea.KeyMsg:
		return m.updateKey(msg)
	case commandResult:
		return m.updateCommandResult(msg)
	case eventResult:
		return m.updateEventResult(msg)
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func (m model) updateKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c":
		m.stopWatch()
		return m, tea.Quit
	case "enter":
		raw := strings.TrimSpace(m.input.Value())
		if raw == "" {
			return m, nil
		}
		if raw == "/quit" || raw == "/exit" {
			m.stopWatch()
			return m, tea.Quit
		}
		if m.busy {
			return m, nil
		}
		m.history = append(m.history, raw)
		m.historyIndex = -1
		m.lines = append(m.lines, "> "+raw)
		m.input.SetValue("")
		m.busy = true
		m.refresh()
		return m, executeCommandCmd(
			m.ctx,
			m.client,
			m.playerID,
			m.worldID,
			m.dialogNpcID,
			m.suggestions,
			raw,
		)
	case "up":
		m.previousHistory()
		return m, nil
	case "down":
		m.nextHistory()
		return m, nil
	default:
		var cmd tea.Cmd
		m.input, cmd = m.input.Update(msg)
		return m, cmd
	}
}

func (m model) updateCommandResult(
	result commandResult,
) (tea.Model, tea.Cmd) {
	m.busy = false
	if result.err != nil {
		m.lines = append(
			m.lines,
			errorStyle.Render("error: "+result.err.Error()),
		)
	} else {
		m.lines = append(m.lines, result.lines...)
		if result.playerID > 0 {
			m.playerID = result.playerID
		}
		if result.dialogNpcID > 0 {
			m.dialogNpcID = result.dialogNpcID
			m.suggestions = nil
		}
		if result.clearDialog {
			m.dialogNpcID = 0
			m.suggestions = nil
		}
		if result.clearSuggestions {
			m.suggestions = nil
		}
		if result.worldID > 0 && result.worldID != m.worldID {
			m.worldID = result.worldID
			m.lastSequence = 0
			m.stopWatch()
			m.watchCancel, m.eventCh = startWatcher(
				m.ctx,
				m.client,
				m.playerID,
				m.worldID,
				m.lastSequence,
			)
			m.refresh()
			return m, waitEvent(m.eventCh)
		}
	}
	m.refresh()
	return m, nil
}

func (m model) updateEventResult(
	result eventResult,
) (tea.Model, tea.Cmd) {
	if result.err != nil {
		m.lines = append(
			m.lines,
			errorStyle.Render("event stream: "+result.err.Error()),
		)
	} else if result.event != nil &&
		result.event.GetSequence() > m.lastSequence {
		m.lastSequence = result.event.GetSequence()
		choices := eventSuggestedChoices(result.event)
		if len(choices) > 0 {
			m.suggestions = choices
			if result.event.GetNpcId() > 0 {
				m.dialogNpcID = result.event.GetNpcId()
			}
		}
		m.lines = append(m.lines, eventStyle.Render(formatEventLine(result.event)))
	}
	m.refresh()
	if m.eventCh != nil {
		return m, waitEvent(m.eventCh)
	}
	return m, nil
}

func (m model) View() string {
	status := fmt.Sprintf(
		"player=%d world=%d seq=%d dialog_npc=%d suggestions=%d",
		m.playerID,
		m.worldID,
		m.lastSequence,
		m.dialogNpcID,
		len(m.suggestions),
	)
	if m.busy {
		status += " working" + m.spinner.View()
	}
	return titleStyle.Render(" Game Town ") +
		"\n" + m.viewport.View() +
		"\n" + statusStyle.Render(status) +
		"\n" + m.input.View()
}

func (m *model) previousHistory() {
	if len(m.history) == 0 {
		return
	}
	if m.historyIndex < 0 {
		m.historyIndex = len(m.history) - 1
	} else if m.historyIndex > 0 {
		m.historyIndex--
	}
	m.input.SetValue(m.history[m.historyIndex])
	m.input.CursorEnd()
}

func (m *model) nextHistory() {
	if m.historyIndex >= 0 &&
		m.historyIndex < len(m.history)-1 {
		m.historyIndex++
		m.input.SetValue(m.history[m.historyIndex])
		m.input.CursorEnd()
		return
	}
	m.historyIndex = -1
	m.input.SetValue("")
}

func (m *model) refresh() {
	m.viewport.SetContent(strings.Join(m.lines, "\n"))
	m.viewport.GotoBottom()
}

func (m *model) stopWatch() {
	if m.watchCancel == nil {
		return
	}
	m.watchCancel()
	m.watchCancel = nil
	m.eventCh = nil
}

func waitEvent(events <-chan eventResult) tea.Cmd {
	return func() tea.Msg {
		result, ok := <-events
		if !ok {
			return eventResult{err: fmt.Errorf("event stream closed")}
		}
		return result
	}
}

func executeCommandCmd(
	ctx context.Context,
	client *rpc.GameTownClient,
	playerID int64,
	worldID int64,
	dialogNpcID int64,
	suggestions []suggestedChoice,
	raw string,
) tea.Cmd {
	return func() tea.Msg {
		return executeCommand(
			ctx,
			client,
			playerID,
			worldID,
			dialogNpcID,
			suggestions,
			raw,
		)
	}
}

func commandSuggestions() []string {
	return []string{
		"/register ",
		"/player use ",
		"/config add ",
		"/config list",
		"/world create ",
		"/world list",
		"/world use ",
		"/world join ",
		"/look",
		"/targets",
		"/nearby",
		"/who",
		"/npcs",
		"/factions",
		"/move ",
		"/talk ",
		"/act ",
		"/events",
		"/status",
		"/help",
		"/quit",
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
