package main

import (
	"notion-igdb-autocomplete/config"
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
		model:  tuiApp.NewModel(conf),
	}
}

// Init implements tea.Model interface
func (m App) Init() tea.Cmd {
	return m.model.Init()
}

// Update implements tea.Model interface
func (m App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m.model.Update(msg)
}

// View implements tea.Model interface
func (m App) View() string {
	return m.model.View()
}
