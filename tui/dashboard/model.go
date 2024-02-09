package dashboard

import (
	"notion-igdb-autocomplete/tui"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	Width  int
	Height int
	binds  bindings
	help   help.Model
	cursor int
	panels []panelState
	pState panelState
}

func NewModel() Model {
	return Model{
		Width:  0,
		Height: 0,
		binds:  newBindings(),
		help:   help.New(),
		cursor: 0,
		panels: []panelState{waitingPanel, runningPanel, finishedPanel},
		pState: waitingPanel,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.binds.SwitchPanel):
			m.cursor += 1
			if m.cursor > len(m.panels)-1 {
				m.cursor = 0
			}
			m.pState = m.panels[m.cursor]
		case key.Matches(msg, m.binds.Back):
			return m, func() tea.Msg { return tui.BackMsg{Width: m.Width, Height: m.Height} }
		}
	}

	return m, nil
}

func (m Model) View() string {
	mainContentWidth := m.Width
	mainContentHeight := m.Height - headerStyle.GetHeight() - statusBarStyle.GetHeight() - helpStyle.GetHeight()

	statusBar := m.viewStatusBar()
	waitingTasks := m.viewWaitingTasks(mainContentWidth, mainContentHeight)
	runningTasks := m.viewRunningTasks(mainContentWidth, mainContentHeight)
	finishedTasks := m.viewFinishedTasks(mainContentWidth, mainContentHeight)

	mainRightPart := lipgloss.JoinVertical(lipgloss.Top, runningTasks, finishedTasks)
	mainPanes := lipgloss.JoinHorizontal(lipgloss.Top, waitingTasks, mainRightPart)

	headerContent := headerStyle.Copy().Width(m.Width).Render("Notion IGDB autocomplete - Dashboard")
	mainContent := mainStyle.Copy().Width(mainContentWidth).Height(mainContentHeight).Render(mainPanes)
	helpContent := helpStyle.Copy().Width(mainContentWidth).Render(m.help.View(m.binds))

	return lipgloss.JoinVertical(lipgloss.Top, headerContent, statusBar, mainContent, helpContent)
}

func (m Model) viewStatusBar() string {
	return statusBarStyle.Copy().Width(m.Width).Render("Watcher status: running | Waiting tasks: 28 | Running tasks: 10 | Finished tasks: 3")
}

func (m Model) viewWaitingTasks(containerWidth int, containerHeight int) string {
	width := (containerWidth / 2) - 2 // Width - borders chars
	height := containerHeight - 2     // Height - borders chars

	cols := []table.Column{
		{Title: "ID", Width: width / 11},
		{Title: "Search query", Width: width / 5},
		{Title: "Type", Width: width / 7},
		{Title: "Notion Page ID", Width: width / 3},
		{Title: "Queued at", Width: width / 3},
	}

	rows := []table.Row{
		{"123456789098765432", "Tokyo", "Japan", "37,274,000", "djksqjdlksq"},
		{"123456789098765432", "Tokyo", "Japan", "37,274,000", "djksqjdlksq"},
		{"123456789098765432", "Tokyo", "Japan", "37,274,000", "djksqjdlksq"},
		{"123456789098765432", "Tokyo", "Japan", "37,274,000", "djksqjdlksq"},
		{"123456789098765432", "Tokyo", "Japan", "37,274,000", "djksqjdlksq"},
		{"123456789098765432", "Tokyo", "Japan", "37,274,000", "djksqjdlksq"},
		{"123456789098765432", "Tokyo", "Japan", "37,274,000", "djksqjdlksq"},
		{"123456789098765432", "Tokyo", "Japan", "37,274,000", "djksqjdlksq"},
		{"123456789098765432", "Tokyo", "Japan", "37,274,000", "djksqjdlksq"},
		{"123456789098765432", "Tokyo", "Japan", "37,274,000", "djksqjdlksq"},
		{"123456789098765432", "Tokyo", "Japan", "37,274,000", "djksqjdlksq"},
		{"123456789098765432", "Tokyo", "Japan", "37,274,000", "djksqjdlksq"},
		{"123456789098765432", "Tokyo", "Japan", "37,274,000", "djksqjdlksq"},
		{"123456789098765432", "Tokyo", "Japan", "37,274,000", "djksqjdlksq"},
		{"123456789098765432", "Tokyo", "Japan", "37,274,000", "djksqjdlksq"},
		{"123456789098765432", "Tokyo", "Japan", "37,274,000", "djksqjdlksq"},
		{"123456789098765432", "Tokyo", "Japan", "37,274,000", "djksqjdlksq"},
		{"123456789098765432", "Tokyo", "Japan", "37,274,000", "djksqjdlksq"},
		{"123456789098765432", "Tokyo", "Japan", "37,274,000", "djksqjdlksq"},
		{"123456789098765432", "Tokyo", "Japan", "37,274,000", "djksqjdlksq"},
		{"123456789098765432", "Tokyo", "Japan", "37,274,000", "djksqjdlksq"},
		{"123456789098765432", "Tokyo", "Japan", "37,274,000", "djksqjdlksq"},
		{"123456789098765432", "Tokyo", "Japan", "37,274,000", "djksqjdlksq"},
		{"123456789098765432", "Tokyo", "Japan", "37,274,000", "djksqjdlksq"},
		{"123456789098765432", "Tokyo", "Japan", "37,274,000", "djksqjdlksq"},
		{"123456789098765432", "Tokyo", "Japan", "37,274,000", "djksqjdlksq"},
		{"123456789098765432", "Tokyo", "Japan", "37,274,000", "djksqjdlksq"},
		{"123456789098765432", "Tokyo", "Japan", "37,274,000", "djksqjdlksq"},
		{"123456789098765432", "Tokyo", "Japan", "37,274,000", "djksqjdlksq"},
		{"123456789098765432", "Tokyo", "Japan", "37,274,000", "djksqjdlksq"},
		{"123456789098765432", "Tokyo", "Japan", "37,274,000", "djksqjdlksq"},
		{"123456789098765432", "Tokyo", "Japan", "37,274,000", "djksqjdlksq"},
		{"123456789098765432", "Tokyo", "Japan", "37,274,000", "djksqjdlksq"},
		{"123456789098765432", "Tokyo", "Japan", "37,274,000", "djksqjdlksq"},
		{"123456789098765432", "Tokyo", "Japan", "37,274,000", "djksqjdlksq"},
		{"123456789098765432", "Tokyo", "Japan", "37,274,000", "djksqjdlksq"},
		{"123456789098765432", "Tokyo", "Japan", "37,274,000", "djksqjdlksq"},
		{"123456789098765432", "Tokyo", "Japan", "37,274,000", "djksqjdlksq"},
	}

	t := table.New(
		table.WithColumns(cols),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(height-4), // Substract y borders size (2), title height (1), margin bottom (1)
		table.WithWidth(width-2),
	)

	title := subtitleStyle.Copy().Width(width).Render("Waiting tasks")
	panelContent := lipgloss.JoinVertical(lipgloss.Top, title, t.View())

	return panelContainerStyle.Copy().Width(width).Height(height).BorderForeground(m.getPanelColor(waitingPanel)).Render(panelContent)
}

func (m Model) viewRunningTasks(containerWidth int, containerHeight int) string {
	width := (containerWidth / 2) - 2   // Width - borders chars
	height := (containerHeight / 2) - 2 // Height - borders chars

	return panelContainerStyle.Copy().Width(width).Height(height).BorderForeground(m.getPanelColor(runningPanel)).Render("")
}

func (m Model) viewFinishedTasks(containerWidth int, containerHeight int) string {
	width := (containerWidth / 2) - 2   // Width - borders chars
	height := (containerHeight / 2) - 1 // Height - borders chars

	return panelContainerStyle.Copy().Width(width).Height(height).BorderForeground(m.getPanelColor(finishedPanel)).Render("")
}

func (m Model) getPanelColor(target panelState) lipgloss.Color {
	if target == m.pState {
		return focussedColor
	} else {
		return blurredColor
	}
}
