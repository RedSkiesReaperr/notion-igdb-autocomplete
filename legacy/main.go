package main

import (
	"flag"
	"log"
	"notion-igdb-autocomplete/config"
	"notion-igdb-autocomplete/core"
	"notion-igdb-autocomplete/tui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	headlessMode := flag.Bool("headless", false, "Set headless mode. Default: false")
	noLogFile := flag.Bool("no-logfile", false, "If provided, no log file will be created. Default: false")
	flag.Parse()

	logger := Logger{headlessMode: *headlessMode, logFileMode: !*noLogFile}
	if err := logger.Setup(); err != nil {
		log.Fatalf("cannot setup logging: %v", err)
	}
	defer logger.Close()

	conf, err := config.Load()
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	core, err := core.New(conf)
	if err != nil {
		log.Fatalf("cannot create core: %v", err)
	}

	if *headlessMode {
		if err := core.Initialize(); err != nil {
			log.Fatalf("cannot initialize core: %v", err)
		}

		core.Launch(true)
		defer core.Stop()
	} else {
		teaProgram := tea.NewProgram(tui.RootScreen(core), tea.WithAltScreen())
		teaProgram.SetWindowTitle("notion-igdb-autocomplete")

		if _, err := teaProgram.Run(); err != nil {
			log.Fatal(err)
		}
	}
}
