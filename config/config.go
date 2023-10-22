package config

import (
	"fmt"

	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

// Config holds all configurations loaded from .env file
type Config struct {
	NotionAPISecret  string `env:"NOTION_API_SECRET,required"`
	NotionPageID     string `env:"NOTION_PAGE_ID,required"`
	IGDBClientID     string `env:"IGDB_CLIENT_ID,required"`
	IGDBSecret       string `env:"IGDB_SECRET,required"`
	UpdaterHost      string `env:"UPDATER_HOST,required"`
	UpdaterPort      int    `env:"UPDATER_PORT,required"`
	WatcherTickDelay int    `env:"WATCHER_TICK_DELAY,required"`
}

// Load all configs from .env & returns the values (+ error if needed)
func Load() (Config, error) {
	config := Config{}

	err := godotenv.Load()
	if err != nil {
		return config, err
	}

	err = env.Parse(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func (c *Config) UpdaterURL() string {
	return fmt.Sprintf("http://%s:%d/", c.UpdaterHost, c.UpdaterPort)
}
