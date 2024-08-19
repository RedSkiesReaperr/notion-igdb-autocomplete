package tui

import (
	"notion-igdb-autocomplete/core"

	tea "github.com/charmbracelet/bubbletea"
)

type RootModel struct {
	Model  tea.Model
	core   *core.Core
	width  int
	height int
}

// Init implements tea.Model interface
func (m RootModel) Init() tea.Cmd {
	return m.Model.Init()
}

// Update implements tea.Model interface
func (m RootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m.Model.Update(msg)
}

// View implements tea.Model interface
func (m RootModel) View() string {
	return m.Model.View()
}

func (m RootModel) switchScreen(screen screenType) (tea.Model, tea.Cmd) {
	switch screen {
	case home:
		m.Model = homeScreen(&m)
	case dashboard:
		m.Model = dashboardScreen(&m)
	case configuration:
		m.Model = configurationScreen(&m)
	}

	return m.Model, m.Model.Init()
}

func (m RootModel) switchScreenModel(model tea.Model) (tea.Model, tea.Cmd) {
	m.Model = model

	return m.Model, m.Model.Init()
}
