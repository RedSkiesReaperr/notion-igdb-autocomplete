package main

import (
	"notion-igdb-autocomplete/config"
	"notion-igdb-autocomplete/tui"
	tuiApp "notion-igdb-autocomplete/tui/app"

	tea "github.com/charmbracelet/bubbletea"
)

type App struct {
	config *config.Config
	model  tea.Model
}

func NewApp(conf *config.Config) *App {
	return &App{
		config: conf,
		model:  tuiApp.NewModel(),
	}
}

// Init implements tea.Model interface
func (m App) Init() tea.Cmd {
	return m.model.Init()
}

// Update implements tea.Model interface
func (m App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	newModel, newCmd := m.model.Update(msg)

	m.model = newModel
	cmd = newCmd

	switch msg.(type) {
	case tui.SaveConfigMsg:
		casted := msg.(tui.SaveConfigMsg)
		m.config.NotionAPISecret = casted.NotionApiSecret
		m.config.NotionPageID = casted.NotionPageId
		m.config.IGDBClientID = casted.IgdbClientId
		m.config.IGDBSecret = casted.IgdbSecret
		m.config.RefreshDelay = casted.RefreshDelay
		m.config.Save()
	}

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

// View implements tea.Model interface
func (m App) View() string {
	return m.model.View()
}
