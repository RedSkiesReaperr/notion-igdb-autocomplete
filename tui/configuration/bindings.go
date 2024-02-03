package configuration

import "github.com/charmbracelet/bubbles/key"

type bindings struct {
	Back key.Binding
	Save key.Binding
}

func newBindings() bindings {
	return bindings{
		Back: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "Back to main menu"),
		),
		Save: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("↵", "Save"),
		),
	}
}

// ShortHelp returns keybindings to be shown in the mini help view. It's part
// of the key.Map interface.
func (b bindings) ShortHelp() []key.Binding {
	return []key.Binding{b.Back, b.Save}
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// key.Map interface.
func (b bindings) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{b.Back}, // first column
		{b.Save}, // second column
	}
}
