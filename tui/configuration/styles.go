package configuration

import "github.com/charmbracelet/lipgloss"

const (
	mainBackgroundColor   = lipgloss.Color("234")
	accentBackgroundColor = lipgloss.Color("237")
	focusedColor          = lipgloss.Color("34")
	blurredColor          = accentBackgroundColor
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
var introductionStyle = lipgloss.NewStyle().
	Height(4).
	PaddingTop(1).
	Foreground(lipgloss.Color("255")).
	AlignHorizontal(lipgloss.Left).
	AlignVertical(lipgloss.Left)

var mainStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("1")).
	Background(mainBackgroundColor).
	AlignHorizontal(lipgloss.Center).
	AlignVertical(lipgloss.Center)

var focusedInputStyle = lipgloss.NewStyle().
	Foreground(focusedColor).
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(accentBackgroundColor).
	BorderBackground(mainBackgroundColor).
	BorderTop(true).
	BorderRight(true).
	BorderBottom(true).
	BorderLeft(true)

var focusedChoiceCursorStyle = lipgloss.NewStyle().
	Foreground(focusedColor).
	Blink(true).
	Bold(true)

var focusedChoiceLabelStyle = lipgloss.NewStyle().
	Foreground(focusedColor).
	Bold(true).
	Underline(false).
	Italic(false)

var blurredInputStyle = lipgloss.NewStyle().
	Foreground(blurredColor).
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(accentBackgroundColor).
	BorderBackground(mainBackgroundColor).
	BorderTop(true).
	BorderRight(true).
	BorderBottom(true).
	BorderLeft(true)

var blurredChoiceCursorStyle = lipgloss.NewStyle().
	Foreground(blurredColor)

var blurredChoiceLabelStyle = lipgloss.NewStyle().
	Foreground(blurredColor)

// ////////////////////////
// Help content styles  //
// ////////////////////////
var helpStyle = lipgloss.NewStyle().
	Background(accentBackgroundColor).
	Height(1)
