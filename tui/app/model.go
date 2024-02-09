package app

import (
	"fmt"
	"notion-igdb-autocomplete/config"
	"notion-igdb-autocomplete/tui"
	tuiConfiguration "notion-igdb-autocomplete/tui/configuration"
	tuiDashboard "notion-igdb-autocomplete/tui/dashboard"
	tuiDialog "notion-igdb-autocomplete/tui/dialog"
	tuiVerifyConfig "notion-igdb-autocomplete/tui/verifyconfiguration"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	config              *config.Config
	choices             []tui.Choice
	width               int
	height              int
	cursor              int // manage choices selector
	viewState           tui.ViewState
	dashboard           tea.Model              // dashboard view model
	configuration       tuiConfiguration.Model // configuration view model
	verifyConfiguration tea.Model              // verify configuration view model
	dialog              tuiDialog.Model
	binds               bindings // keymap of available controls
	help                help.Model
}

func NewModel(config *config.Config) *Model {
	return &Model{
		config: config,
		choices: []tui.Choice{
			{Id: tui.DashboardView, Label: "Dashboard", Description: "Launch program & show monitoring dashboard"},
			{Id: tui.ConfigurationView, Label: "Configure", Description: "Edit your configuration"},
			{Id: tui.VerifyConfigurationView, Label: "Verify configuration", Description: "Check if your configuration is correct"},
		},
		width:               0,
		height:              0,
		cursor:              0,
		viewState:           tui.MainView,
		dashboard:           tuiDashboard.NewModel(),
		configuration:       tuiConfiguration.NewModel(*config),
		verifyConfiguration: tuiVerifyConfig.NewModel(),
		dialog:              tuiDialog.NewModel(),
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
	case tui.DialogView:
		return m.dialog
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

	switch msg := msg.(type) {
	case tui.BackMsg:
		m.width, m.height = msg.Width, msg.Height
		m.viewState = tui.MainView
	case tui.SaveConfigMsg:
		if err := m.saveConfig(msg); err != nil {
			m.dialog.Title = "Saving configuration"
			m.dialog.Message = fmt.Sprintf("An error happened while saving your configuration settings:\n\n%s", err.Error())
			m.dialog.Type = tuiDialog.ErrorDialog
			m.viewState = tui.DialogView
		} else {
			m.dialog.Title = "Saving configuration"
			m.dialog.Message = "All your configuration settings have been successfully saved !"
			m.dialog.Type = tuiDialog.SuccessDialog
			m.viewState = tui.DialogView
		}
	}

	switch m.viewState {
	case tui.DashboardView:
		newModel, newCmd := m.dashboard.Update(msg)
		m.dashboard = newModel
		cmd = newCmd
	case tui.ConfigurationView:
		newModel, newCmd := m.configuration.Update(msg)
		newConfig, _ := newModel.(tuiConfiguration.Model)
		m.configuration = newConfig
		cmd = newCmd
	case tui.VerifyConfigurationView:
		newModel, newCmd := m.verifyConfiguration.Update(msg)
		m.dashboard = newModel
		cmd = newCmd
	case tui.DialogView:
		newModel, newCmd := m.dialog.Update(msg)
		newDialog, _ := newModel.(tuiDialog.Model)
		m.dialog = newDialog
		cmd = newCmd
	case tui.MainView:
		switch msg := msg.(type) {
		case tea.WindowSizeMsg:
			m.width, m.height = msg.Width, msg.Height
			m.configuration.Width, m.configuration.Height = msg.Width, msg.Height
			m.dialog.Width, m.dialog.Height = msg.Width, msg.Height
			return m, nil
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

		prompt = choicePromptStyle.Render(prompt)
		label := choiceLabelStyle.Render(c.Label)
		desc := choiceDescStyle.Render(c.Description)

		view += fmt.Sprintf("%s %s\n %s\n\n\n", prompt, label, desc)
	}

	headerContent := headerStyle.Copy().Width(m.width).Render("Notion IGDB autocomplete")
	mainContent := mainStyle.Copy().Width(m.width).Height(m.height - headerStyle.GetHeight() - helpStyle.GetHeight()).Render(view)
	helpContent := helpStyle.Copy().Width(m.width).Render(m.help.View(m.binds))

	return lipgloss.JoinVertical(lipgloss.Top, headerContent, mainContent, helpContent)
}

func (m Model) saveConfig(newValues tui.SaveConfigMsg) error {
	//TODO: If save error happened, reset config to previous values
	m.config.NotionAPISecret = newValues.NotionApiSecret
	m.config.NotionPageID = newValues.NotionPageId
	m.config.IGDBClientID = newValues.IgdbClientId
	m.config.IGDBSecret = newValues.IgdbSecret
	m.config.RefreshDelay = newValues.RefreshDelay

	return m.config.Save()
}
