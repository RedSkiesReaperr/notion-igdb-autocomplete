package tui

import (
	"notion-igdb-autocomplete/core"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type taskModel struct {
	root  *RootModel
	task  *core.Task
	binds taskBindings
	help  help.Model
}

// Init implements tea.Model interface
func (m taskModel) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model interface
func (m taskModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.root.width = msg.Width
		m.root.height = msg.Height
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.binds.Refresh):
			// Just re-render
		case key.Matches(msg, m.binds.Back):
			return m.root.switchScreen(dashboard)
		}
	}

	return m, nil
}

// View implements tea.Model interface
func (m taskModel) View() string {
	return m.render()
}
