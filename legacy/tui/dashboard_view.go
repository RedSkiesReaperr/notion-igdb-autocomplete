package tui

import (
	"fmt"
	"notion-igdb-autocomplete/core"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

func (m dashboardModel) render() string {
	width, height := m.root.width, m.root.height
	title := renderTitle(width, "Notion IGDB Autocomplete")
	subtitle := renderSubtitle(width, height, m.getSubtitle())
	tasksTables := m.renderTasks(
		width,
		height-headerStyle.GetHeight()-subtitleStyle.GetHeight())
	help := renderHelp(m.help, m.binds)

	return lipgloss.JoinVertical(lipgloss.Top, title, subtitle, tasksTables, help)
}

func (m dashboardModel) getSubtitle() string {
	waiting := m.root.core.CountTasksWithStatus(core.WaitingTask)
	running := m.root.core.CountTasksWithStatus(core.RunningTask)
	finished := m.root.core.CountTasksWithStatus(core.FinishedTask)

	return fmt.Sprintf("Waiting tasks: %d | Running tasks: %d | Finished tasks: %d", waiting, running, finished)
}

func (m dashboardModel) renderTasks(width, height int) string {
	width -= 2          // Reduce by 2 (1=>Left border + 1=>Right border)
	height = height / 3 // Evenly spread height between the 3 tables
	height -= 4         // Reduce by 4 (1=>Top border + 1=>Header line + 1=>Bottom border + 1 =>Due to constrained width)

	tables := []string{}
	tablesRenderFunctions := []func(int, int) (int, []table.Column, []table.Row){
		m.getWaitingTasksTableParams, m.getRunningTasksTableParams, m.getFinishedTasksTableParams,
	}

	for _, function := range tablesRenderFunctions {
		index, cols, rows := function(width, height)
		tables = append(tables, m.renderTasksTable(width, height, index, cols, rows))
	}

	return lipgloss.JoinVertical(lipgloss.Top, tables...)
}

func (m dashboardModel) renderTasksTable(width, height, tableIndex int, cols []table.Column, rows []table.Row) string {
	var style lipgloss.Style
	t := m.tasksTables[tableIndex]

	if m.tasksTables[m.currentTasksTable] == t {
		style = dashboardTableFocussedStyle
		t.Focus()
	} else {
		style = dashboardTableBlurredStyle
		t.Blur()
	}

	tableStyles := table.DefaultStyles()
	tableStyles.Header.
		Background(colorBackground).
		Bold(true)

	tableStyles.Selected.
		Foreground(colorFocus).
		Bold(true).
		Italic(true)

	t.SetColumns(cols)
	t.SetRows(rows)
	t.SetWidth(width)
	t.SetHeight(height)
	t.SetStyles(tableStyles)

	return style.Width(width).Render(t.View())
}
