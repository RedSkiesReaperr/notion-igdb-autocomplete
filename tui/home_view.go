package tui

import (
	"github.com/charmbracelet/lipgloss"
)

func (m homeModel) render() string {
	width, height := m.root.width, m.root.height
	mainContentHeight := height - headerStyle.GetHeight() - 1 // Reduce by 1 (1=>Help height)

	title := renderTitle(width, "Notion IGDB Autocomplete")
	content := m.renderContent(width, mainContentHeight)
	help := renderHelp(m.help, m.binds)

	return lipgloss.JoinVertical(lipgloss.Top, title, content, help)
}

func (m homeModel) renderContent(width, height int) string {
	list := m.renderMenus()

	return lipgloss.NewStyle().Width(width).Height(height).Align(lipgloss.Center).Render(list)
}

func (m homeModel) renderMenus() string {
	var menus []string

	for i, menu := range m.menus {
		itemStyle := menuItemStyle.Copy()
		titleStyle := menuItemTitleStyle.Copy()
		descStyle := menuItemDescStyle.Copy()

		if i == m.cursor {
			itemStyle.Foreground(colorFocus).BorderForeground(colorFocus)
		} else {
			itemStyle.Foreground(colorBlur).BorderLeft(false).PaddingLeft(1)
		}

		title := titleStyle.Render(menu.title)
		desc := descStyle.Render(menu.desc)
		item := lipgloss.JoinVertical(lipgloss.Top, title, desc)

		menus = append(menus, itemStyle.Render(item))
	}

	return lipgloss.JoinVertical(lipgloss.Top, menus...)
}
