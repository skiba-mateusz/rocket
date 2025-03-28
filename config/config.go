package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

type Config struct {
	BaseUrl      string     `toml:"baseUrl"`
	LanguageCode string     `toml:"languageCode"`
	Title        string     `toml:"title"`
	Menu         []MenuItem `toml:"menu"`
}

type MenuItem struct {
	Name string `toml:"name"`
	Url  string `toml:"url"`
}

func LoadConfig() (*Config, error) {
	if _, err := os.Stat("config.toml"); err != nil {
		return nil, fmt.Errorf("config file not found - ensure you are in right directory: %v", err)
	}

	var config Config
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %v", err)
	}

	return &config, nil
}
