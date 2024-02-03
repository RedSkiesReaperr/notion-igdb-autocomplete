package dashboard

import (
	"fmt"
	"notion-igdb-autocomplete/tui"

	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	binds bindings
	help  help.Model
}

func NewModel() *Model {
	return &Model{
		binds: newBindings(),
		help:  help.New(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "q":
			return m, func() tea.Msg { return tui.BackMsg{} }
		}
	}

	return m, nil
}

func (m Model) View() string {
	view := fmt.Sprintf("Dashboard view \n\n%s", m.help.View(m.binds))

	return view
}
