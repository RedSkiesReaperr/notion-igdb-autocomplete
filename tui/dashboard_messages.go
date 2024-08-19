package tui

import (
	"log"
	"notion-igdb-autocomplete/core"

	tea "github.com/charmbracelet/bubbletea"
)

type taskMsg string // Contains the notification message

func launchCore(core *core.Core) tea.Cmd {
	return func() tea.Msg {
		if err := core.Initialize(); err != nil {
			log.Fatalf("cannot initialize core: %v", err)
		}

		core.Launch(false)
		return nil
	}
}

func waitForTask(sub chan string) tea.Cmd {
	return func() tea.Msg {
		return taskMsg(<-sub)
	}
}
