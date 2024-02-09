package dialog

import "github.com/charmbracelet/bubbles/key"

type bindings struct {
	Validate key.Binding
}

func newBindings() bindings {
	return bindings{
		Validate: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("↵", "Select"),
		),
	}
}

// ShortHelp returns keybindings to be shown in the mini help view. It's part
// of the key.Map interface.
func (b bindings) ShortHelp() []key.Binding {
	return []key.Binding{b.Validate}
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// key.Map interface.
func (b bindings) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{b.Validate}, // first column
		{},           // second column
	}
}
