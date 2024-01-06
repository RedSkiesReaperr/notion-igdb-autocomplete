package main

import (
	"fmt"
	"notion-igdb-autocomplete/config"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type homeModel struct {
	config  *config.Config
	choices []inputChoice
	cursor  int
	err     error
}

func newHomeModel(config *config.Config) homeModel {
	return homeModel{
		config: config,
		choices: []inputChoice{
			newInputChoice("Notion API secret", &config.NotionAPISecret, validateString),
			newInputChoice("Notion page ID", &config.NotionPageID, validateString),
			newInputChoice("IGDB secret", &config.IGDBSecret, validateString),
			newInputChoice("IGDB client ID", &config.IGDBClientID, validateString),
			newInputChoice("Refresh delay", &config.RefreshDelay, validateInteger),
		},
		cursor: 0,
		err:    nil,
	}
}

func (m homeModel) Init() tea.Cmd {
	return nil
}

func (m homeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.err = nil

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			if err := m.config.Save(); err != nil { // If something went wrong, report error
				m.err = err
			} else {
				return m, tea.Quit
			}
		case tea.KeyUp:
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(m.choices) - 1
			}
		case tea.KeyDown:
			m.cursor++
			if m.cursor >= len(m.choices) {
				m.cursor = 0
			}
		}
	case error:
		m.err = msg
		cmd = nil
	}

	// Only activate current element input
	for i := range m.choices {
		if i == m.cursor {
			m.choices[i].input.Focus()
		} else {
			m.choices[i].input.Blur()
		}
	}

	m.choices[m.cursor].input, cmd = m.choices[m.cursor].Update(msg)
	return m, cmd
}

func (m homeModel) View() string {
	var b strings.Builder

	header := m.renderHeader()
	inputs := m.renderBody()
	errors := m.renderErrors()
	man := m.renderHelp()

	template := fmt.Sprintf(`%s
%s

%s

%s
`, &header, &inputs, &errors, &man)

	b.WriteString(template)

	return b.String()
}

func (m homeModel) renderHeader() strings.Builder {
	var b strings.Builder

	b.WriteString(headerStyle.Render(`==================================================================
===== Welcome to the notion-igdb-autocomplete configurator ! =====
==================================================================
`))

	b.WriteString(headerStyle.Render(`
This program helps you to setup all mandatory configurations to
smoothly run notion-igdb-autocomplete tool.
`))

	b.WriteString(headerStyle.Render(`
You can find more informations about them here => 
github.com/RedSkiesReaperr/notion-igdb-autocomplete
`))

	return b
}

func (m homeModel) renderBody() strings.Builder {
	var b strings.Builder
	choiceTemplate := `
 %s  %s
    %s
`

	for i, choice := range m.choices {
		cursor := " "
		style := blurredStyle
		if i == m.cursor {
			cursor = "$"
			style = focusedStyle
		}

		choice.input.PromptStyle = style
		choice.input.TextStyle = style

		cursor = style.Copy().Blink(true).Bold(true).Render(cursor)
		name := style.Copy().Bold(true).Underline(true).Italic(true).Render(choice.name)
		input := choice.input.View()

		b.WriteString(fmt.Sprintf(choiceTemplate, cursor, name, input))
	}

	return b
}

func (m homeModel) renderErrors() strings.Builder {
	var b strings.Builder

	if m.err != nil {
		b.WriteString(errorStyle.Render(fmt.Sprintf(`%s`, m.err)))
	}

	return b
}

func (m homeModel) renderHelp() strings.Builder {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("%s %s %s", helpPrimaryStyle.Render("Press"), helpAccentStyle.Render("up/down"), helpPrimaryStyle.Render(fmt.Sprintf("to %s", "move between values"))))
	b.WriteRune('\n')
	b.WriteString(fmt.Sprintf("%s %s %s", helpPrimaryStyle.Render("Press"), helpAccentStyle.Render("enter"), helpPrimaryStyle.Render(fmt.Sprintf("to %s", "save & quit"))))
	b.WriteRune('\n')
	b.WriteString(fmt.Sprintf("%s %s %s", helpPrimaryStyle.Render("Press"), helpAccentStyle.Render("esc"), helpPrimaryStyle.Render(fmt.Sprintf("to %s", "quit without saving"))))

	return b
}
