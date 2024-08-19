package tui

import "github.com/charmbracelet/lipgloss"

var menuItemStyle = lipgloss.NewStyle().
	MarginBottom(1).
	MarginTop(1).
	MarginLeft(1).
	Border(lipgloss.NormalBorder(), false, false, false, true)

var menuItemTitleStyle = lipgloss.NewStyle().
	Bold(true)

var menuItemDescStyle = lipgloss.NewStyle().
	Bold(false).
	Italic(true)
