package main

import "github.com/charmbracelet/lipgloss"

var (
	// Overall
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#e67e22"))
	blurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#34495e"))
	errorStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#e74c3c"))
	// UI parts specific
	headerStyle      = lipgloss.NewStyle().Italic(false).Foreground(lipgloss.Color("#16a085")).AlignHorizontal(lipgloss.Center)
	helpPrimaryStyle = lipgloss.NewStyle().Italic(true).Foreground(lipgloss.Color("240"))
	helpAccentStyle  = lipgloss.NewStyle().Italic(true).Foreground(lipgloss.Color("250"))
)
