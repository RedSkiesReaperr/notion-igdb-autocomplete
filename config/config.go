package config

import (
	"encoding/json"
	"errors"
	"os"
	"os/user"
	"path"
)

// Config holds all configurations loaded from .env file
type Config struct {
	NotionAPISecret string `json:"NOTION_API_SECRET"`
	NotionPageID    string `json:"NOTION_PAGE_ID"`
	IGDBClientID    string `json:"IGDB_CLIENT_ID"`
	IGDBSecret      string `json:"IGDB_SECRET"`
	RefreshDelay    string `json:"REFRESH_DELAY"`
}

// Load all configs from .env & returns the values (+ error if needed)
func Load() (Config, error) {
	config := Config{}
	user, err := user.Current()
	if err != nil {
		return config, err
	}

	configPath := path.Join(user.HomeDir, "notion-igdb-autocomplete.json")

	if _, err := os.Stat(configPath); errors.Is(err, os.ErrNotExist) {
		if createErr := createFile(configPath); createErr != nil {
			return Config{}, createErr
		}
	}

	if loadErr := config.loadFromFile(configPath); loadErr != nil {
		return Config{}, loadErr
	}

	return config, nil
}

func (c Config) Save() error {
	user, err := user.Current()
	if err != nil {
		return err
	}

	configPath := path.Join(user.HomeDir, "notion-igdb-autocomplete.json")
	marshalled, err := json.Marshal(c)
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, marshalled, 0755)
}

func (c *Config) loadFromFile(filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(content, c)
	if err != nil {
		return err
	}

	return nil
}

func createFile(filePath string) error {
	defaultConfig := Config{
		NotionAPISecret: "",
		NotionPageID:    "",
		IGDBClientID:    "",
		IGDBSecret:      "",
		RefreshDelay:    "5",
	}

	marshalled, err := json.Marshal(defaultConfig)
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, marshalled, 0755)
}
