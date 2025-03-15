package initializer

import (
	"embed"
	"fmt"
	"github.com/skiba-mateusz/rocket/logger"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

//go:embed defaults/*
var defaultsFS embed.FS

type Initializer struct {
	logger logger.Logger
}

func NewInitializer(logger logger.Logger) *Initializer {
	return &Initializer{
		logger: logger,
	}
}

func (i *Initializer) NewSite(title string) error {
	rootDir := filepath.Join(".", title)

	if info, err := os.Stat(rootDir); err == nil {
		if info.IsDir() {
			return fmt.Errorf("site '%s' already exists in current directory", title)
		} else {
			return fmt.Errorf("a file named '%s' alredy exits and is not a directory", title)
		}
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("failed to check site existence: %v", err)
	}

	i.logger.Info("Initializing new site: '%s'", title)

	if err := fs.WalkDir(defaultsFS, "defaults", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("failed to access embedded file %s: %v", path, err)
		}

		return i.handleEmbeddedFile(rootDir, path, d)
	}); err != nil {
		i.logger.Error("failed to initialize site '%s': %v", title, err)
		return err
	}

	i.logger.Success("Site '%s' initialized successfully!", title)

	return nil
}

func (i *Initializer) handleEmbeddedFile(rootDir, path string, d fs.DirEntry) error {
	if d.IsDir() {
		return nil
	}

	content, err := defaultsFS.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read embedded file %s: %v", err)
	}

	outputPath := i.mapEmbeddedPathToOutput(rootDir, path)
	outputPathDir := filepath.Dir(outputPath)

	if err = os.MkdirAll(outputPathDir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %v", outputPath, err)
	}

	if err = os.WriteFile(outputPath, content, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %v", path, err)
	}

	i.logger.Info("Created file: %s\n", outputPath)

	return nil
}

func (i *Initializer) mapEmbeddedPathToOutput(rootDir, path string) string {
	outputPath := strings.TrimPrefix(path, "defaults/")
	return filepath.Join(rootDir, outputPath)
}
