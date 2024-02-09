package dashboard

import "github.com/charmbracelet/bubbles/key"

type bindings struct {
	SwitchPanel key.Binding
	Back        key.Binding
}

func newBindings() bindings {
	return bindings{
		SwitchPanel: key.NewBinding(
			key.WithKeys("tab"),
			key.WithHelp("⇆", "Switch panel"),
		),
		Back: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "Back to main menu"),
		),
	}
}

// ShortHelp returns keybindings to be shown in the mini help view. It's part
// of the key.Map interface.
func (b bindings) ShortHelp() []key.Binding {
	return []key.Binding{b.SwitchPanel, b.Back}
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// key.Map interface.
func (b bindings) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{b.SwitchPanel, b.Back}, // first column
		{},                      // second column
	}
}
