package tui

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/lipgloss"
)

const colorBackground = lipgloss.Color("0")
const colorFocus = lipgloss.Color("68")
const colorBlur = lipgloss.Color("246")

var headerStyle = lipgloss.NewStyle().
	Height(1).
	Align(lipgloss.Center).
	Background(lipgloss.Color("8")).
	Bold(true)

var subtitleStyle = lipgloss.NewStyle().
	Height(1).
	Align(lipgloss.Center).
	Background(lipgloss.Color("8")).
	Bold(false).
	Italic(true)

func renderTitle(width int, text string) string {
	return headerStyle.Copy().Width(width).Render(text)
}

func renderSubtitle(width, height int, text string) string {
	return subtitleStyle.Copy().Width(width).Render(text)
}

func renderHelp(help help.Model, binds help.KeyMap) string {
	return help.View(binds)
}
