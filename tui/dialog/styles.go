package dialog

import "github.com/charmbracelet/lipgloss"

const (
	mainBackgroundColor   = lipgloss.Color("234")
	accentBackgroundColor = lipgloss.Color("237")
	successColor          = lipgloss.Color("34")
	errorColor            = lipgloss.Color("124")
	infoColor             = lipgloss.Color("7")
)

// ////////////////////////
// Main content styles  //
// ////////////////////////
var mainStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("1")).
	Background(mainBackgroundColor).
	AlignHorizontal(lipgloss.Center).
	AlignVertical(lipgloss.Center)

var dialogStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderBackground(mainBackgroundColor).
	BorderTop(true).
	BorderLeft(true).
	BorderBottom(true).
	BorderRight(true)

var dialogTitleStyle = lipgloss.NewStyle().
	AlignHorizontal(lipgloss.Center).
	Bold(true)

var dialogMessageStyle = lipgloss.NewStyle().
	AlignHorizontal(lipgloss.Left).
	AlignVertical(lipgloss.Center).
	PaddingLeft(2).
	PaddingRight(2)

var dialogControlsStyle = lipgloss.NewStyle().
	AlignHorizontal(lipgloss.Center)

var dialogButtonStyle = lipgloss.NewStyle().
	Blink(true).
	PaddingLeft(1).
	PaddingRight(1).
	Background(accentBackgroundColor).
	BorderBackground(mainBackgroundColor).
	BorderTop(false).
	BorderLeft(true).
	BorderBottom(false).
	BorderRight(true)

// ////////////////////////
// Help content styles  //
// ////////////////////////
var helpStyle = lipgloss.NewStyle().
	Background(accentBackgroundColor).
	Height(1)
