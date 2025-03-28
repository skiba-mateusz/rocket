package builder

import (
	"fmt"
	"github.com/skiba-mateusz/rocket/content"
	"github.com/skiba-mateusz/rocket/logger"
	"github.com/skiba-mateusz/rocket/templates"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	dirPerm  = 0755
	filePerm = 0644
)

type Builder struct {
	parser          content.Parser
	logger          logger.Logger
	engine          templates.Engine
	publicDir       string
	contentDir      string
	totalPages      int
	successfulPages int
}

func NewBuilder(logger logger.Logger, parser content.Parser, engine templates.Engine, publicDir, contentDir string) *Builder {
	return &Builder{
		parser:          parser,
		logger:          logger,
		engine:          engine,
		publicDir:       publicDir,
		contentDir:      contentDir,
		totalPages:      0,
		successfulPages: 0,
	}
}

func (b *Builder) Build() error {
	b.logger.Info("Starting build process")

	if err := os.RemoveAll(b.publicDir); err != nil {
		return fmt.Errorf("failed to remove %s dir: %v", b.publicDir, err)
	}

	startTime := time.Now()
	if err := filepath.WalkDir(b.contentDir, b.processFile); err != nil {
		return err
	}

	b.logger.Success("Building process completed in %s (%d/%d pages)", time.Since(startTime), b.successfulPages, b.totalPages)
	return nil
}

func (b *Builder) processFile(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}

	if d.IsDir() {
		return nil
	}

	b.totalPages++
	b.logger.Info("Building page: %s", path)

	page, err := b.parser.Parse(path)
	if err != nil {
		b.logger.Warn("Failed to parse %s: %v", path, err)
		return nil
	}

	if err = b.buildPage(page); err != nil {
		b.logger.Warn("Failed to build %s: %v", path, err)
	}

	b.successfulPages++
	b.logger.Info("Page built successfully")

	return nil
}

func (b *Builder) buildPage(page *content.Page) error {
	customTemplatesDir := b.determineCustomTemplatesDir(page.Url, page.Filename)

	customTemplates := []string{
		filepath.Join(customTemplatesDir, "item.html"),
		filepath.Join(customTemplatesDir, page.Layout),
	}

	output, err := b.engine.Render(customTemplates, "base.html", page)
	if err != nil {
		return fmt.Errorf("failed to render page: %v", err)
	}

	outputPath := filepath.Join(b.publicDir, filepath.Base(page.Url), "index.html")

	return b.writePageToFile(outputPath, output)
}

func (b *Builder) determineCustomTemplatesDir(url, filename string) string {
	var pageTemplatesDir string

	if strings.TrimSuffix(filename, filepath.Ext(filename)) == "index" {
		pageTemplatesDir = filepath.Base(url)
	} else {
		pageTemplatesDir = filepath.Dir(url)
	}

	if pageTemplatesDir == "/" {
		pageTemplatesDir = "home"
	}

	return pageTemplatesDir
}

func (b *Builder) writePageToFile(path string, data []byte) error {
	if err := os.MkdirAll(filepath.Dir(path), dirPerm); err != nil {
		return fmt.Errorf("failed to create directory %s: %v", path, err)
	}

	if err := os.WriteFile(path, data, filePerm); err != nil {
		return fmt.Errorf("failed to write page file %s: %v", path, err)
	}

	return nil
}
