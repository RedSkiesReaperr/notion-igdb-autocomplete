package tui

import (
	"github.com/charmbracelet/lipgloss"
)

func (m configurationModel) render() string {
	width, height := m.root.width, m.root.height
	mainContentHeight := height - headerStyle.GetHeight() - subtitleStyle.GetHeight() - 1 // Reduce by 1 (1=>Help height)

	title := renderTitle(width, "Notion IGDB Autocomplete")
	mainContent := m.renderMainContent(width, mainContentHeight)
	helpContent := renderHelp(m.help, m.binds)

	return lipgloss.JoinVertical(lipgloss.Top, title, mainContent, helpContent)
}

func (m configurationModel) renderMainContent(width, height int) string {
	intro := configurationIntroStyle.Width(width).Render("To complete the configuration below, please follow the setup guide available at => https://github.com/RedSkiesReaperr/notion-igdb-autocomplete?tab=readme-ov-file#configuration")
	inputs := m.renderInputs(width, height-configurationIntroStyle.GetVerticalFrameSize())

	return lipgloss.JoinVertical(lipgloss.Top, intro, inputs)
}

func (m configurationModel) renderInputs(width, height int) string {
	configItems := []string{}

	for i, item := range m.items {
		itemStyle := configurationItemStyle.Copy()

		if i == m.cursor {
			item.input.PromptStyle = configurationInputFocusStyle
			item.input.TextStyle = configurationInputFocusStyle
			itemStyle.BorderForeground(colorFocus).Foreground(colorFocus)
		} else {
			itemStyle.BorderForeground(colorBlur).Foreground(colorBlur)
		}

		input := item.input.View()
		title := configurationInputTitleStyle.Render(item.title)
		desc := configurationInputDescStyle.Render(item.desc)
		input = configurationInputStyle.Render(input)
		configItems = append(configItems, itemStyle.Render(lipgloss.JoinVertical(lipgloss.Top, title, desc, input)))
	}

	return lipgloss.NewStyle().Width(width).Height(height).Render(lipgloss.JoinVertical(lipgloss.Top, configItems...))
}
