package configuration

import (
	"fmt"
	"notion-igdb-autocomplete/tui"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
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
		switch {
		case key.Matches(msg, m.binds.Save):
			return m, func() tea.Msg {
				return tui.SaveConfigMsg{
					NotionApiSecret: "SECRET1",     //TODO: Replace by real value
					NotionPageId:    "PAGE1345678", //TODO: Replace by real value
					IgdbClientId:    "CLIENT",      //TODO: Replace by real value
					IgdbSecret:      "SECRET2",     //TODO: Replace by real value
					RefreshDelay:    "10",          //TODO: Replace by real value
				}
			}
		case key.Matches(msg, m.binds.Back):
			return m, func() tea.Msg { return tui.BackMsg{} }
		}
	}

	return m, nil
}

func (m Model) View() string {
	view := fmt.Sprintf("Configuration view \n\n%s", m.help.View(m.binds))

	return view
}
