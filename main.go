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

	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err := teaProgram.Run(); err != nil {
		log.Fatal(err)
	}
}
