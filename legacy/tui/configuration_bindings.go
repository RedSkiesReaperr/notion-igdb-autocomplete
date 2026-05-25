package tui

import "github.com/charmbracelet/bubbles/key"

type configurationBindings struct {
	MoveUp   key.Binding
	MoveDown key.Binding
	Back     key.Binding
	Save     key.Binding
}

func newConfigurationBindings() configurationBindings {
	return configurationBindings{
		Back:     key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "Back to home")),
		MoveUp:   key.NewBinding(key.WithKeys("up"), key.WithHelp("↑", "Select previous input")),
		MoveDown: key.NewBinding(key.WithKeys("down"), key.WithHelp("↓", "Select next input")),
		Save:     key.NewBinding(key.WithKeys("enter"), key.WithHelp("↵", "Save")),
	}
}

func (d configurationBindings) ShortHelp() []key.Binding {
	return []key.Binding{d.MoveUp, d.MoveDown, d.Save, d.Back}
}

func (d configurationBindings) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{d.MoveUp, d.MoveDown},
		{d.Back, d.Save},
	}
}
