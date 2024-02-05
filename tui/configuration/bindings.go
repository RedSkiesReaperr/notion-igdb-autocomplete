package configuration

import "github.com/charmbracelet/bubbles/key"

type bindings struct {
	MoveUp   key.Binding
	MoveDown key.Binding
	Back     key.Binding
	Save     key.Binding
}

func newBindings() bindings {
	return bindings{
		MoveUp: key.NewBinding(
			key.WithKeys("up"),
			key.WithHelp("↑", "Move up"),
		),
		MoveDown: key.NewBinding(
			key.WithKeys("down"),
			key.WithHelp("↓", "Move down"),
		),
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
	return []key.Binding{b.MoveUp, b.MoveDown, b.Save, b.Back}
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// key.Map interface.
func (b bindings) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{b.MoveUp, b.MoveDown}, // first column
		{b.Save, b.Back},       // second column
	}
}
