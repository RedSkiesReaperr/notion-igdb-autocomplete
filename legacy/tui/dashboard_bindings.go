package tui

import "github.com/charmbracelet/bubbles/key"

type dashboardBindings struct {
	Back            key.Binding
	SwitchPanel     key.Binding
	TableLineUp     key.Binding
	TableLineDown   key.Binding
	ShowTaskDetails key.Binding
	RerunTask       key.Binding
}

func newDashboardBindings() dashboardBindings {
	return dashboardBindings{
		Back:            key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "Back to home")),
		SwitchPanel:     key.NewBinding(key.WithKeys("tab"), key.WithHelp("⇆", "Select next task table")),
		TableLineUp:     key.NewBinding(key.WithKeys("up"), key.WithHelp("↑", "Previous task in selected table")),
		TableLineDown:   key.NewBinding(key.WithKeys("down"), key.WithHelp("↓", "Next task in selected table")),
		ShowTaskDetails: key.NewBinding(key.WithKeys("enter"), key.WithHelp("↵", "Show selected task details")),
		RerunTask:       key.NewBinding(key.WithKeys("r"), key.WithHelp("r", "Rerun selected task")),
	}
}

func (d dashboardBindings) ShortHelp() []key.Binding {
	return []key.Binding{d.Back, d.SwitchPanel, d.TableLineUp, d.TableLineDown, d.RerunTask, d.ShowTaskDetails}
}

func (d dashboardBindings) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{d.Back},
		{d.SwitchPanel, d.TableLineUp, d.TableLineDown, d.RerunTask, d.ShowTaskDetails},
	}
}
