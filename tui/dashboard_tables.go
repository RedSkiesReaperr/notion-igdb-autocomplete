package tui

import (
	"notion-igdb-autocomplete/core"

	"github.com/charmbracelet/bubbles/table"
)

func (m dashboardModel) getWaitingTasksTableParams(width, height int) (tableIndex int, cols []table.Column, rows []table.Row) {
	tableIndex = 0
	cols = []table.Column{
		{Title: "ID", Width: width / 4},
		{Title: "Search query", Width: width / 4},
		{Title: "Type", Width: width / 11},
		{Title: "Notion Page", Width: width / 4},
		{Title: "Queued at", Width: width / 3},
	}

	rows = []table.Row{}
	for _, task := range m.root.core.FindTasksByStatus(core.WaitingTask) {
		rows = append(rows, table.Row{
			task.Id.String(),
			task.Query,
			task.TypeString(),
			task.NotionId,
			timeToHumanDate(task.QueuedAt),
		})
	}

	return
}

func (m dashboardModel) getRunningTasksTableParams(width, height int) (tableIndex int, cols []table.Column, rows []table.Row) {
	tableIndex = 1
	cols = []table.Column{
		{Title: "ID", Width: width / 4},
		{Title: "Search query", Width: width / 4},
		{Title: "Type", Width: width / 11},
		{Title: "Notion Page", Width: width / 4},
		{Title: "Started at", Width: width / 3},
	}

	rows = []table.Row{}
	for _, task := range m.root.core.FindTasksByStatus(core.RunningTask) {
		rows = append(rows, table.Row{
			task.Id.String(),
			task.Query,
			task.TypeString(),
			task.NotionId,
			timeToHumanDate(task.StartedAt),
		})
	}

	return
}

func (m dashboardModel) getFinishedTasksTableParams(width, height int) (tableIndex int, cols []table.Column, rows []table.Row) {
	tableIndex = 2
	colsCount := 6
	totalRemainingSpace := getProportionOf(width, 0.29)
	spaceBetweenCols := totalRemainingSpace/colsCount - 1

	cols = []table.Column{
		{Title: "ID", Width: getProportionOf(width, 0.15) + spaceBetweenCols},
		{Title: "Search query", Width: getProportionOf(width, 0.25) + spaceBetweenCols},
		{Title: "Type", Width: getProportionOf(width, 0.06) + spaceBetweenCols},
		{Title: "Notion Page", Width: getProportionOf(width, 0.15) + spaceBetweenCols},
		{Title: "Elapsed", Width: getProportionOf(width, 0.06) + spaceBetweenCols},
		{Title: "Succeed?", Width: getProportionOf(width, 0.04)},
	}

	rows = []table.Row{}
	for _, task := range m.root.core.FindTasksByStatus(core.FinishedTask) {
		var succeedValue string

		if task.Succeed() {
			succeedValue = "✓"
		} else {
			succeedValue = "✗"
		}

		rows = append(rows, table.Row{
			task.Id.String(),
			task.Query,
			task.TypeString(),
			task.NotionId,
			task.Elapsed().String(),
			succeedValue,
		})
	}

	return
}
