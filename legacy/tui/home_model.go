package tui

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type homeMenuItem struct {
	title        string
	desc         string
	targetScreen screenType
}

type homeModel struct {
	root   *RootModel
	binds  homeBindings
	help   help.Model
	cursor int
	menus  []homeMenuItem
}

// Init implements tea.Model interface
func (m homeModel) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model interface
func (m homeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.root.width = msg.Width
		m.root.height = msg.Height
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.binds.MenuUp):
			if m.cursor <= 0 {
				m.cursor = len(m.menus) - 1
			} else {
				m.cursor--
			}
		case key.Matches(msg, m.binds.MenuDown):
			if m.cursor >= len(m.menus)-1 {
				m.cursor = 0
			} else {
				m.cursor++
			}
		case key.Matches(msg, m.binds.Validate):
			return m.root.switchScreen(m.menus[m.cursor].targetScreen)
		case key.Matches(msg, m.binds.Quit):
			return m, tea.Quit
		}
	}

	return m, nil
}

// View implements tea.Model interface
func (m homeModel) View() string {
	return m.render()
}
