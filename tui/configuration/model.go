package configuration

import (
	"fmt"
	"notion-igdb-autocomplete/config"
	"notion-igdb-autocomplete/tui"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type configValues struct {
	notionApiSecret string
	notionPageId    string
	igdbClientId    string
	igdbSecret      string
	refreshDelay    string
}

type Model struct {
	cursor  int
	choices []inputChoice
	values  *configValues
	binds   bindings
	help    help.Model
	Width   int
	Height  int
}

func NewModel(initialConf config.Config) *Model {
	values := configValues{
		notionApiSecret: initialConf.NotionAPISecret,
		notionPageId:    initialConf.NotionPageID,
		igdbClientId:    initialConf.IGDBClientID,
		igdbSecret:      initialConf.IGDBSecret,
		refreshDelay:    initialConf.RefreshDelay,
	}

	choices := []inputChoice{
		newInputChoice(&values.notionApiSecret, "Notion API secret", tui.ValidateString),
		newInputChoice(&values.notionPageId, "Notion page ID", tui.ValidateString),
		newInputChoice(&values.igdbSecret, "IGDB secret", tui.ValidateString),
		newInputChoice(&values.igdbClientId, "IGDB client ID", tui.ValidateString),
		newInputChoice(&values.refreshDelay, "Refresh delay", tui.ValidateInteger),
	}

	return &Model{
		cursor:  0,
		values:  &values,
		choices: choices,
		binds:   newBindings(),
		help:    help.New(),
		Width:   0,
		Height:  0,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width, m.Height = msg.Width, msg.Height
		return m, nil
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.binds.Save):
			return m, func() tea.Msg {
				return tui.SaveConfigMsg{
					NotionApiSecret: m.values.notionApiSecret,
					NotionPageId:    m.values.notionPageId,
					IgdbClientId:    m.values.igdbClientId,
					IgdbSecret:      m.values.igdbSecret,
					RefreshDelay:    m.values.refreshDelay,
				}
			}
		case key.Matches(msg, m.binds.MoveUp):
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(m.choices) - 1
			}
		case key.Matches(msg, m.binds.MoveDown):
			m.cursor++
			if m.cursor >= len(m.choices) {
				m.cursor = 0
			}
		case key.Matches(msg, m.binds.Back):
			return m, func() tea.Msg { return tui.BackMsg{Width: m.Width, Height: m.Height} }
		}
	}

	// Only activate current element input
	for i := range m.choices {
		if i == m.cursor {
			m.choices[i].input.Focus()
		} else {
			m.choices[i].input.Blur()
		}
	}

	var cmd tea.Cmd
	m.choices[m.cursor].input, cmd = m.choices[m.cursor].Update(msg)
	return m, cmd
}

func (m Model) View() string {
	intro := introductionStyle.Render("This view helps you to setup all mandatory configuration parameters to smoothly run notion-igdb-autocomplete tool.\n\nFirst, you need to follow the configuration guide here => https://github.com/RedSkiesReaperr/notion-igdb-autocomplete?tab=readme-ov-file#configuration")
	view := ""

	for i, choice := range m.choices {
		cursor, name, input := " ", "", choice.input.View()

		if i == m.cursor {
			cursor = ">"
			choice.input.PromptStyle = focusedInputStyle
			choice.input.TextStyle = focusedInputStyle
			cursor = focusedChoiceCursorStyle.Render(cursor)
			name = focusedChoiceLabelStyle.Render(choice.name)
			input = focusedInputStyle.Render(input)
			choice.input.PromptStyle.BorderForeground(focusedColor)
		} else {
			choice.input.PromptStyle = blurredInputStyle
			choice.input.TextStyle = blurredInputStyle
			cursor = blurredChoiceCursorStyle.Render(cursor)
			name = blurredChoiceLabelStyle.Render(choice.name)
			input = blurredInputStyle.Render(input)
		}

		view += fmt.Sprintf("%s %s\n%s\n\n", cursor, name, input)
	}

	headerContent := headerStyle.Copy().Width(m.Width).Render("Notion IGDB autocomplete - Configuration")
	mainContent := mainStyle.Copy().Width(m.Width).Height(m.Height - headerStyle.GetHeight() - introductionStyle.GetHeight() - helpStyle.GetHeight()).Render(view)
	helpContent := helpStyle.Copy().Width(m.Width).Render(m.help.View(m.binds))

	return lipgloss.JoinVertical(lipgloss.Top, headerContent, intro, mainContent, helpContent)
}
