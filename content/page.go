package content

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type FrontMatter struct {
	Title  string `toml:"title"`
	Url    string `toml:"url"`
	Date   string `toml:"date"`
	Layout string `toml:"layout"`
}

func NewPage(contentDir, path string) error {
	fullPath := filepath.Join(contentDir, path)
	pathDir := filepath.Dir(fullPath)

	if err := os.MkdirAll(pathDir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %v", pathDir, err)
	}

	frontMatter, err := generateFrontMatter(contentDir, fullPath)
	if err != nil {
		return err
	}

	if err = os.WriteFile(fullPath, []byte(frontMatter), 0644); err != nil {
		return fmt.Errorf("failed to create file %s: %v", path, err)
	}

	return nil
}

func generateFrontMatter(contentDir, path string) (string, error) {
	title := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
	url := strings.TrimPrefix(strings.TrimSuffix(path, filepath.Ext(path)), contentDir)
	currentDate := time.Now().Format(time.RFC3339)

	fm := FrontMatter{
		Title:  title,
		Url:    url,
		Date:   currentDate,
		Layout: "single.html",
	}

	var builder strings.Builder
	builder.WriteString("---\n")
	if err := toml.NewEncoder(&builder).Encode(fm); err != nil {
		return "", fmt.Errorf("failed to encode front matter: %v", err)
	}
	builder.WriteString("---\n")

	return builder.String(), nil
}
