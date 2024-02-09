package dashboard

import "github.com/charmbracelet/lipgloss"

const (
	mainBackgroundColor   = lipgloss.Color("234")
	accentBackgroundColor = lipgloss.Color("237")
	focussedColor         = lipgloss.Color("92")
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

// //////////////////
// Main styles     //
// //////////////////
var statusBarStyle = lipgloss.NewStyle().
	Background(mainBackgroundColor).
	Height(1).
	Align(lipgloss.Center)

var mainStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("1")).
	Background(mainBackgroundColor).
	AlignHorizontal(lipgloss.Left).
	AlignVertical(lipgloss.Left)

var panelContainerStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderBackground(mainBackgroundColor).
	BorderForeground(accentBackgroundColor).
	BorderTop(true).
	BorderLeft(true).
	BorderBottom(true).
	BorderRight(true)

var subtitleStyle = lipgloss.NewStyle().
	Height(1).
	MarginBottom(1).
	Background(accentBackgroundColor).
	AlignHorizontal(lipgloss.Center).
	Bold(true)

// //////////////////
// Help styles     //
// //////////////////
var helpStyle = lipgloss.NewStyle().
	Background(accentBackgroundColor).
	Height(1)
