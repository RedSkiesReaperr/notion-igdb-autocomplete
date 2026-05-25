package tui

import "github.com/charmbracelet/bubbles/key"

type homeBindings struct {
	MenuUp   key.Binding
	MenuDown key.Binding
	Quit     key.Binding
	Validate key.Binding
}

func newHomeBindings() homeBindings {
	return homeBindings{
		MenuUp:   key.NewBinding(key.WithKeys("up"), key.WithHelp("↑", "Select previous menu")),
		MenuDown: key.NewBinding(key.WithKeys("down"), key.WithHelp("↓", "Select next menu")),
		Quit:     key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "Quit")),
		Validate: key.NewBinding(key.WithKeys("enter"), key.WithHelp("↵", "Validate")),
	}
}

func (d homeBindings) ShortHelp() []key.Binding {
	return []key.Binding{d.MenuUp, d.MenuDown, d.Validate, d.Quit}
}

func (d homeBindings) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{d.MenuUp, d.MenuDown},
		{d.Validate, d.Quit},
	}
}
