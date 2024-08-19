package tui

import (
	"github.com/charmbracelet/lipgloss"
)

var dashboardTableBaseStyle = lipgloss.NewStyle().
	Background(colorBackground).
	BorderStyle(lipgloss.NormalBorder()).
	BorderBackground(colorBackground)

var dashboardTableBlurredStyle = dashboardTableBaseStyle.Copy().
	BorderForeground(lipgloss.Color("240"))

var dashboardTableFocussedStyle = dashboardTableBaseStyle.Copy().
	BorderForeground(colorFocus)
