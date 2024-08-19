package tui

import (
	"fmt"
	"log"
	"notion-igdb-autocomplete/core"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
)

type dashboardModel struct {
	root              *RootModel
	tasksTables       []*table.Model
	currentTasksTable int
	binds             dashboardBindings
	help              help.Model
}

// Init implements tea.Model interface
func (m dashboardModel) Init() tea.Cmd {
	return tea.Batch(
		launchCore(m.root.core),
		waitForTask(m.root.core.TasksNotificationsQueue))
}

// Update implements tea.Model interface
func (m dashboardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.root.width = msg.Width
		m.root.height = msg.Height
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.binds.SwitchPanel):
			if m.currentTasksTable < len(m.tasksTables)-1 {
				m.currentTasksTable++
			} else {
				m.currentTasksTable = 0
			}
		case key.Matches(msg, m.binds.TableLineUp):
			m.tasksTables[m.currentTasksTable].MoveUp(1)
		case key.Matches(msg, m.binds.TableLineDown):
			m.tasksTables[m.currentTasksTable].MoveDown(1)
		case key.Matches(msg, m.binds.RerunTask):
			task, err := m.getTaskFromTableRow(m.tasksTables[m.currentTasksTable].SelectedRow())
			if err != nil {
				log.Printf("cannot get task to detail: %v", err)
				return m, nil
			}
			m.root.core.RetryTask(task)
		case key.Matches(msg, m.binds.ShowTaskDetails):
			task, err := m.getTaskFromTableRow(m.tasksTables[m.currentTasksTable].SelectedRow())
			if err != nil {
				log.Printf("cannot get task to detail: %v", err)
				return m, nil
			}
			return m.root.switchScreenModel(taskScreen(m.root, task))
		case key.Matches(msg, m.binds.Back):
			m.root.core.Stop()
			return m.root.switchScreen(home)
		}
	case taskMsg:
		return m, waitForTask(m.root.core.TasksNotificationsQueue)
	}

	return m, nil
}

// View implements tea.Model interface
func (m dashboardModel) View() string {
	return m.render()
}

func (m dashboardModel) getTaskFromTableRow(row table.Row) (*core.Task, error) {
	if row == nil {
		return nil, fmt.Errorf("invalid row")
	}

	targetTaskUUID, err := uuid.Parse(row[0])
	if err != nil {
		return nil, fmt.Errorf("invalid row UUID: %v", row[0])
	}

	for _, task := range m.root.core.Tasks {
		if task.Id == targetTaskUUID {
			return task, nil
		}
	}

	return nil, fmt.Errorf("task with UUID %v not found", row[0])
}
