package tui

import "github.com/charmbracelet/lipgloss"

var configurationIntroStyle = lipgloss.NewStyle().
	PaddingTop(1).
	MarginBottom(2).
	AlignHorizontal(lipgloss.Center)

var configurationItemStyle = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder(), false, false, false, true).
	MarginBottom(1)

var configurationInputTitleStyle = lipgloss.NewStyle().
	Bold(true)

var configurationInputDescStyle = lipgloss.NewStyle().
	Italic(true)

var configurationInputStyle = lipgloss.NewStyle().
	Background(colorBackground).
	Border(lipgloss.RoundedBorder(), true, true, true, true).
	BorderBackground(colorBackground).
	MarginBottom(1)

var configurationInputFocusStyle = lipgloss.NewStyle().
	Foreground(colorFocus).
	BorderForeground(colorFocus).
	Background(colorBackground)

var configurationInputBlurStyle = lipgloss.NewStyle().
	Foreground(colorBlur).
	BorderForeground(colorBlur)
