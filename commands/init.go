package commands

import (
	"fmt"
	"os"
	"path/filepath"
)

func InitCommand(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("specify a directory name, e.g. 'rocket init my-site'")
	}

	rootDir := args[0]

	if _, err := os.Stat(rootDir); err == nil {
		return fmt.Errorf("directory %s already exists", rootDir)
	}

	if err := os.MkdirAll(rootDir, 0755); err != nil {
		return fmt.Errorf("failed to create project: %v", err)
	}

	dirs := []string{"content", "layout", "static"}

	for _, dir := range dirs {
		path := filepath.Join(rootDir, dir)
		if err := os.MkdirAll(path, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %v", path, err)
		}
	}

	configPath := filepath.Join(rootDir, "config.toml")
	defaultConfig := `[site]
title = "My Site"
baseURL = "http://example.com"
	`

	if err := os.WriteFile(configPath, []byte(defaultConfig), 0644); err != nil {
		return fmt.Errorf("failed to write config file %v", err)
	}

	fmt.Printf("Project %s initialized successfully\n", rootDir)

	return nil
}
