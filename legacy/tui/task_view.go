package tui

import (
	"github.com/charmbracelet/lipgloss"
)

func (m taskModel) render() string {
	width, height := m.root.width, m.root.height
	mainContentHeight := height - headerStyle.GetHeight() - subtitleStyle.GetHeight() - 1 // Reduce by 1 (1=>Help height)

	title := renderTitle(width, "Notion IGDB Autocomplete")
	subtitle := renderSubtitle(width, height, "Task details")
	mainContent := m.renderMainContent(width, mainContentHeight)
	help := renderHelp(m.help, m.binds)

	return lipgloss.JoinVertical(lipgloss.Top, title, subtitle, mainContent, help)
}

func (m taskModel) renderMainContent(width, height int) string {
	id := m.renderItem("ID", m.task.Id.String())
	ttype := m.renderItem("Type", m.task.TypeString())
	status := m.renderItem("Status", m.task.StatusString())
	query := m.renderItem("Query", m.task.Query)
	notion := m.renderItem("Notion Page ID", m.task.NotionId)
	createdAt := m.renderItem("Created At", timeToHumanDate(m.task.CreatedAt))
	queuedAt := m.renderItem("Queued At", timeToHumanDate(m.task.QueuedAt))
	startedAt := m.renderItem("Started At", timeToHumanDate(m.task.StartedAt))
	endedAt := m.renderItem("Ended At", timeToHumanDate(m.task.EndedAt))
	elapsed := m.renderItem("Time elapsed", m.task.Elapsed().String())
	terror := m.renderItem("Error", "No error")

	if m.task.Error != nil {
		terror = m.renderItem("Error", m.task.Error.Error())
	}

	content := lipgloss.JoinVertical(lipgloss.Top, id, ttype, status, query, notion, createdAt, queuedAt, startedAt, endedAt, elapsed, terror)

	return taskMainContentStyle.Height(height).Render(content)
}

func (m taskModel) renderItem(title, value string) string {
	styledTitle := taskItemTitleStyle.Render(title)
	styledValue := taskItemValueStyle.Render(value)
	itemContent := lipgloss.JoinVertical(lipgloss.Top, styledTitle, styledValue)

	return taskItemStyle.Render(itemContent)
}
