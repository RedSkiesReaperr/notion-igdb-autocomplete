package tui

import "github.com/charmbracelet/lipgloss"

var taskMainContentStyle = lipgloss.NewStyle().
	Align(lipgloss.Top)

var taskItemStyle = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder(), false, false, false, true).
	BorderForeground(colorFocus).
	MarginBottom(1).
	PaddingLeft(1)

var taskItemTitleStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(colorFocus)

var taskItemValueStyle = lipgloss.NewStyle().
	Bold(false).
	Italic(true)
