package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type inputChoice struct {
	name  string
	input textinput.Model
	bind  *string
}

func newInputChoice(name string, bind *string, validate func(string) error) inputChoice {
	input := textinput.New()
	input.Placeholder = fmt.Sprintf("Type your %s here", name)
	input.Validate = validate
	input.Cursor.Style = focusedStyle
	input.PlaceholderStyle = blurredStyle.Copy().Italic(true)

	input.SetValue(*bind)

	return inputChoice{
		name:  name,
		bind:  bind,
		input: input,
	}
}

// Implements saver interface
func (ic *inputChoice) Update(msg tea.Msg) (textinput.Model, tea.Cmd) {
	txt, cmd := ic.input.Update(msg)

	*ic.bind = txt.Value()

	return txt, cmd
}
