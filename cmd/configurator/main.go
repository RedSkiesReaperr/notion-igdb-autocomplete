package main

import (
	"log"
	"notion-igdb-autocomplete/config"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	config, err := config.Load()
	if err != nil {
		log.Fatalf("Cannot load config: %s\n", err)
	}

	p := tea.NewProgram(newHomeModel(&config))
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
