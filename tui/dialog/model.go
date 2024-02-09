package dialog

import (
	"notion-igdb-autocomplete/tui"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	Width   int
	Height  int
	Title   string
	Message string
	Type    DialogType
	binds   bindings
	help    help.Model
}

func NewModel() Model {
	return Model{
		Width:   0,
		Height:  0,
		Title:   "",
		Message: "",
		Type:    InfoDialog,
		binds:   newBindings(),
		help:    help.New(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.binds.Validate):
			return m, func() tea.Msg { return tui.BackMsg{Width: m.Width, Height: m.Height} }
		}
	}
	return m, nil
}

func (m Model) View() string {
	dialogWidth := m.Width / 2
	dialogHeight := (m.Height - helpStyle.GetHeight()) / 4

	dialogTitle := dialogTitleStyle.Width(dialogWidth).Render(m.Title)
	dialogMessage := dialogMessageStyle.Width(dialogWidth).Height(dialogHeight).Render(m.Message)
	dialogControls := dialogControlsStyle.Width(dialogWidth).Render(dialogButtonStyle.Render("Ok"))

	dialogContent := lipgloss.JoinVertical(lipgloss.Top, dialogTitle, dialogMessage, dialogControls)
	dialogContent = dialogStyle.Copy().Width(dialogWidth).BorderForeground(m.mainColor()).Height(dialogHeight).Render(dialogContent)

	mainHeight := m.Height - helpStyle.GetHeight()
	mainContent := mainStyle.Copy().Width(m.Width).Height(mainHeight).Render(dialogContent)

	helpContent := helpStyle.Copy().Width(m.Width).Render(m.help.View(m.binds))

	return lipgloss.JoinVertical(lipgloss.Top, mainContent, helpContent)
}

func (m Model) mainColor() lipgloss.Color {
	switch m.Type {
	case SuccessDialog:
		return successColor
	case ErrorDialog:
		return errorColor
	default:
		return infoColor
	}
}
