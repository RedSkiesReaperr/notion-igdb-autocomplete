package tui

import (
	"notion-igdb-autocomplete/config"
	"notion-igdb-autocomplete/core"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
)

func RootScreen(core *core.Core) *RootModel {
	root := RootModel{}
	root.Model = homeScreen(&root)
	root.core = core
	root.width = 0
	root.height = 0

	return &root
}

func homeScreen(root *RootModel) *homeModel {
	return &homeModel{
		root:   root,
		binds:  newHomeBindings(),
		help:   help.New(),
		cursor: 0,
		menus: []homeMenuItem{
			{title: "Dashboard", desc: "Launch autocomplete program & show dashboard", targetScreen: dashboard},
			{title: "Configuration", desc: "Display & edit configurations", targetScreen: configuration},
		},
	}
}

func dashboardScreen(root *RootModel) *dashboardModel {
	waitingTable := table.New()
	runningTable := table.New()
	finishedTable := table.New()

	return &dashboardModel{
		root: root,
		tasksTables: []*table.Model{
			&waitingTable,  // Waiting tasks table
			&runningTable,  // Running tasks table
			&finishedTable, // FInished tasks table
		},
		currentTasksTable: 0,
		binds:             newDashboardBindings(),
		help:              help.New(),
	}
}

func taskScreen(root *RootModel, task *core.Task) *taskModel {
	return &taskModel{
		root:  root,
		task:  task,
		binds: newTaskBindings(),
		help:  help.New(),
	}
}

func configurationScreen(root *RootModel) *configurationModel {
	return &configurationModel{
		root:   root,
		binds:  newConfigurationBindings(),
		help:   help.New(),
		cursor: 0,
		items: []configurationItem{
			{title: "Notion API secret", desc: "Your personal Notion API secret (see configuration guide).", target: config.NotionAPISecret, input: textinput.New()},
			{title: "Notion page ID", desc: "The Notion page where you want to autocomplete", target: config.NotionPageID, input: textinput.New()},
			{title: "IGDB client ID", desc: "Your personal IGDB client ID (see configuration guide).", target: config.IGDBClientID, input: textinput.New()},
			{title: "IGDB secret", desc: "Your personal IGDB secret (see configuration guide).", target: config.IGDBSecret, input: textinput.New()},
			{title: "Refresh delay", desc: "How often we will ask Notion for games to update (in seconds).", target: config.RefreshDelay, input: textinput.New()},
		},
	}
}
