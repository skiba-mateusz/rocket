package commands

import (
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func AddCommand(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("specify content path, e.g. 'rocket add content/posts/my-post.md' (relative path)")
	}

	path := args[0]

	if filepath.Ext(path) == "" {
		path += ".md"
	}

	targetDirPath := filepath.Dir(path)
	if err := os.MkdirAll(targetDirPath, 0755); err != nil {
		return fmt.Errorf("failed to create target directory %s: %v", targetDirPath, err)
	}

	if _, err := os.Stat(path); err == nil {
		return fmt.Errorf("file %s already exists", path)
	}

	title := formatTitle(filepath.Base(path))
	currentTime := time.Now().Format(time.RFC3339)

	defaultContent := generateFrontMatter(title, currentTime)

	if err := os.WriteFile(path, []byte(defaultContent), 0644); err != nil {
		return fmt.Errorf("failed to create file %s: %v", path, err)
	}

	fmt.Printf("Content %s successfully added\n", path)

	return nil
}

func generateFrontMatter(title, date string) string {
	return fmt.Sprintf(`---
title='%s'
date='%s'
---
`, title, date)
}

func formatTitle(filename string) string {
	nameWithoutExt := strings.TrimSuffix(filename, filepath.Ext(filename))
	nameWithSpaces := strings.Replace(nameWithoutExt, "-", " ", -1)
	return cases.Title(language.English).String(nameWithSpaces)
}
