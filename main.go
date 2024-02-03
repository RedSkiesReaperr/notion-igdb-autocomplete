package main

import (
	"log"
	"notion-igdb-autocomplete/config"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	conf, err := config.Load()
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	app := NewApp(&conf)
	teaProgram := tea.NewProgram(app)
	teaProgram.SetWindowTitle("notion-igdb-autocomplete")

	if _, err := teaProgram.Run(); err != nil {
		log.Fatal(err)
	}
}
