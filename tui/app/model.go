package app

import (
	"fmt"
	"notion-igdb-autocomplete/tui"
	tuiConfiguration "notion-igdb-autocomplete/tui/configuration"
	tuiDashboard "notion-igdb-autocomplete/tui/dashboard"
	tuiVerifyConfig "notion-igdb-autocomplete/tui/verifyconfiguration"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	choices             []tui.Choice
	cursor              int // manage choices selector
	viewState           tui.ViewState
	dashboard           tea.Model // dashboard view model
	configuration       tea.Model // configuration view model
	verifyConfiguration tea.Model // verify configuration view model
	binds               bindings  // keymap of available controls
	help                help.Model
}

func NewModel() *Model {
	return &Model{
		choices: []tui.Choice{
			{Id: tui.DashboardView, Label: "Dashboard", Description: "Launch program & show monitoring dashboard"},
			{Id: tui.ConfigurationView, Label: "Configure", Description: "Edit your configuration"},
			{Id: tui.VerifyConfigurationView, Label: "Verify configuration", Description: "Check if your configuration is correct"},
		},
		cursor:              0,
		viewState:           tui.MainView,
		dashboard:           tuiDashboard.NewModel(),
		configuration:       tuiConfiguration.NewModel(),
		verifyConfiguration: tuiVerifyConfig.NewModel(),
		binds:               newBindings(),
		help:                help.New(),
	}
}
func (m Model) currentView() tea.Model {
	switch m.viewState {
	case tui.MainView:
		return m
	case tui.DashboardView:
		return m.dashboard
	case tui.ConfigurationView:
		return m.configuration
	case tui.VerifyConfigurationView:
		return m.verifyConfiguration
	default:
		return m
	}
}

// Init implements tea.Model interface
func (m Model) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model interface
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg.(type) {
	case tui.BackMsg:
		m.viewState = tui.MainView
	}

	switch m.viewState {
	case tui.DashboardView:
		newModel, newCmd := m.dashboard.Update(msg)
		m.dashboard = newModel
		cmd = newCmd
	case tui.ConfigurationView:
		newModel, newCmd := m.configuration.Update(msg)
		m.configuration = newModel
		cmd = newCmd
	case tui.VerifyConfigurationView:
		newModel, newCmd := m.verifyConfiguration.Update(msg)
		m.dashboard = newModel
		cmd = newCmd
	case tui.MainView:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, m.binds.MoveUp):
				m.cursor--
				if m.cursor < 0 {
					m.cursor = len(m.choices) - 1
				}
			case key.Matches(msg, m.binds.MoveDown):
				m.cursor++
				if m.cursor >= len(m.choices) {
					m.cursor = 0
				}
			case key.Matches(msg, m.binds.Select):
				m.viewState = m.choices[m.cursor].Id
			case key.Matches(msg, m.binds.Quit):
				return m, tea.Quit
			}
		}
	}

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

// View implements tea.Model interface
func (m Model) View() string {
	if m.viewState != tui.MainView {
		return m.currentView().View()
	}

	view := ""
	for i, c := range m.choices {
		prompt := " "

		if i == m.cursor {
			prompt = ">"
		}

		view += fmt.Sprintf("%s %s\n  %s\n\n", prompt, c.Label, c.Description)
	}

	view += m.help.View(m.binds)
	view += "\n"

	return view
}
