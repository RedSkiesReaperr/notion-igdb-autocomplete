package app

import "github.com/charmbracelet/lipgloss"

const (
	mainBackgroundColor   = lipgloss.Color("234")
	accentBackgroundColor = lipgloss.Color("237")
)

// //////////////////
// Header styles  //
// //////////////////
var headerStyle = lipgloss.NewStyle().
	Background(accentBackgroundColor).
	Height(3).
	Align(lipgloss.Center).
	AlignVertical(lipgloss.Center)

// ////////////////////////
// Main content styles  //
// ////////////////////////
var mainStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("1")).
	Background(mainBackgroundColor).
	AlignHorizontal(lipgloss.Center).
	AlignVertical(lipgloss.Center)

var choicePromptStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("34")).
	Blink(true).
	Bold(true)

var choiceLabelStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("34")).
	Bold(true)

var choiceDescStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("250")).
	Italic(true)

// ////////////////////////
// Help content styles  //
// ////////////////////////
var helpStyle = lipgloss.NewStyle().
	Background(accentBackgroundColor).
	Height(1)
