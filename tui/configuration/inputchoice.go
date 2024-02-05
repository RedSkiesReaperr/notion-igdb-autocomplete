package configuration

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

func newInputChoice(bind *string, name string, validatorFunc textinput.ValidateFunc) inputChoice {
	input := textinput.New()
	input.Placeholder = fmt.Sprintf("Type your %s here", name)
	input.Prompt = ""
	input.Width = 55
	input.PlaceholderStyle.Width(55)
	input.PlaceholderStyle = blurredInputStyle
	input.Validate = validatorFunc

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
