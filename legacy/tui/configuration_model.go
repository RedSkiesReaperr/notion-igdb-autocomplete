package tui

import (
	"notion-igdb-autocomplete/config"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type configurationItem struct {
	title  string
	desc   string
	input  textinput.Model
	target config.ConfigKey
}

type configurationModel struct {
	root   *RootModel
	binds  configurationBindings
	help   help.Model
	cursor int
	items  []configurationItem
}

// Init implements tea.Model interface
func (m configurationModel) Init() tea.Cmd {
	for i, item := range m.items {
		m.items[i].input.SetValue(m.root.core.Config.Get(item.target))

		if i == m.cursor {
			m.items[i].input.Focus()
		} else {
			m.items[i].input.Blur()
		}
	}

	return nil
}

// Update implements tea.Model interface
func (m configurationModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.root.width = msg.Width
		m.root.height = msg.Height
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.binds.MoveUp):
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(m.items) - 1
			}
		case key.Matches(msg, m.binds.MoveDown):
			m.cursor++
			if m.cursor > len(m.items)-1 {
				m.cursor = 0
			}
		case key.Matches(msg, m.binds.Save):
			for _, item := range m.items {
				m.root.core.Config.Update(item.target, item.input.Value())
			}
			return m.root.switchScreen(home)
		case key.Matches(msg, m.binds.Back):
			return m.root.switchScreen(home)
		}
	}

	for i, _ := range m.items {
		if i == m.cursor {
			m.items[i].input.Focus()
		} else {
			m.items[i].input.Blur()
		}
	}

	m.items[m.cursor].input, cmd = m.items[m.cursor].input.Update(msg)

	return m, cmd
}

// View implements tea.Model interface
func (m configurationModel) View() string {
	return m.render()
}
