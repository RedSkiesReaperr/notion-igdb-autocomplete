package tui

import "github.com/charmbracelet/bubbles/key"

type taskBindings struct {
	Back    key.Binding
	Refresh key.Binding
}

func newTaskBindings() taskBindings {
	return taskBindings{
		Back:    key.NewBinding(key.WithKeys("esc", "enter"), key.WithHelp("esc/↵", "Back to home")),
		Refresh: key.NewBinding(key.WithKeys("r"), key.WithHelp("r", "Refresh informations")),
	}
}

func (d taskBindings) ShortHelp() []key.Binding {
	return []key.Binding{d.Back, d.Refresh}
}

func (d taskBindings) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{d.Back},
		{d.Refresh},
	}
}
