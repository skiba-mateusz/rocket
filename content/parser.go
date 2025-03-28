package content

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

type Parser interface {
	Parse(path string) (*Page, error)
}

type MarkdownParser struct{}

func NewMarkdownParser() *MarkdownParser {
	return &MarkdownParser{}
}

func (mp *MarkdownParser) Parse(path string) (*Page, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %v", path, err)
	}

	frontMatter, markdownContent, err := mp.extractContentFile(data)
	if err != nil {
		return nil, err
	}

	page := &Page{
		FrontMatter: frontMatter,
		Content:     template.HTML(markdownContent),
		Filename:    filepath.Base(path),
	}

	return page, nil
}

func (mp *MarkdownParser) extractContentFile(data []byte) (FrontMatter, string, error) {
	parts := strings.SplitN(string(data), "+++", 3)

	if len(parts) < 3 {
		return FrontMatter{}, "", fmt.Errorf("invalid markdown file: missing front matter")
	}

	var frontMatter FrontMatter
	if _, err := toml.Decode(parts[1], &frontMatter); err != nil {
		return FrontMatter{}, "", fmt.Errorf("failed to parse TOML front matter: %v", err)
	}

	// TODO: convert to html
	markdownContent := parts[2]

	return frontMatter, markdownContent, nil
}
