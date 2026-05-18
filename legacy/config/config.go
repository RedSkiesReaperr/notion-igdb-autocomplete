package config

import (
	"fmt"
	"log"
	"reflect"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	NotionAPISecret string `mapstructure:"NOTION_API_SECRET" configKey:"0"`
	NotionPageID    string `mapstructure:"NOTION_PAGE_ID" configKey:"1"`
	IGDBClientID    string `mapstructure:"IGDB_CLIENT_ID" configKey:"2"`
	IGDBSecret      string `mapstructure:"IGDB_SECRET" configKey:"3"`
	RefreshDelay    int    `mapstructure:"REFRESH_DELAY" configKey:"4"`
}

type ConfigKey string

const (
	NotionAPISecret ConfigKey = "0"
	NotionPageID    ConfigKey = "1"
	IGDBClientID    ConfigKey = "2"
	IGDBSecret      ConfigKey = "3"
	RefreshDelay    ConfigKey = "4"
)

func Load() (*Config, error) {
	configFile := ".env"
	config := Config{}

	viper.SetConfigFile(configFile)
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	viper.SetDefault("NOTION_API_SECRET", "")
	viper.SetDefault("NOTION_PAGE_ID", "")
	viper.SetDefault("IGDB_CLIENT_ID", "")
	viper.SetDefault("IGDB_SECRET", "")
	viper.SetDefault("REFRESH_DELAY", 2)
	viper.OnConfigChange(config.onChange)
	viper.WatchConfig()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("can't find config file %s: %v", configFile, err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("can't load config file: %v", err)
	}

	config.Save() // Ensure config file is created

	return &config, nil
}

func (c Config) Get(key ConfigKey) string {
	t := reflect.TypeOf(c)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		col := field.Tag.Get("configKey")

		if col != "" && col == string(key) {
			configKey := field.Tag.Get("mapstructure")

			return viper.GetString(configKey)
		}
	}

	return ""
}

func (c Config) Update(key ConfigKey, value any) {
	t := reflect.TypeOf(c)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		col := field.Tag.Get("configKey")

		if col != "" && col == string(key) {
			configKey := field.Tag.Get("mapstructure")
			viper.Set(configKey, value)
			break
		}
	}

	c.Save()
}

func (c Config) Save() error {
	return viper.WriteConfig()
}

func (c *Config) onChange(e fsnotify.Event) {
	log.Println("config has been updated")
	viper.ReadInConfig()
	viper.Unmarshal(c)
}
