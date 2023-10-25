package config

import (
	"fmt"

	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

// Config holds all configurations loaded from .env file
type Config struct {
	NotionAPISecret  string `env:"NOTION_API_SECRET,required,notEmpty"`
	NotionPageID     string `env:"NOTION_PAGE_ID,required,notEmpty"`
	IGDBClientID     string `env:"IGDB_CLIENT_ID,required,notEmpty"`
	IGDBSecret       string `env:"IGDB_SECRET,required,notEmpty"`
	UpdaterHost      string `env:"UPDATER_HOST,required" envDefault:"127.0.0.1"`
	UpdaterPort      int    `env:"UPDATER_PORT,required" envDefault:"8080"`
	WatcherTickDelay int    `env:"WATCHER_TICK_DELAY,required" envDefault:"5"`
}

// Load all configs from .env & returns the values (+ error if needed)
func Load(file string) (Config, error) {
	config := Config{}

	if file == "" {
		file = ".env"
	}

	err := godotenv.Load(file)
	if err != nil {
		return config, err
	}

	err = env.Parse(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}

// UpdaterURL returns URL to reach the updater process
func (c *Config) UpdaterURL() string {
	return fmt.Sprintf("http://%s:%d/", c.UpdaterHost, c.UpdaterPort)
}
