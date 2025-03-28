package templates

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

type Engine interface {
	Render(templates []string, baseTemplate string, data interface{}) ([]byte, error)
}

type GoTemplateEngine struct {
	templates    *template.Template
	templatesDir string
}

func NewGoTemplateEngine(templatesDir string) (*GoTemplateEngine, error) {
	defaultTemplates, err := filepath.Glob(filepath.Join(templatesDir, "defaults/*.html"))
	if err != nil {
		return nil, fmt.Errorf("failed to glob default templates: %v", err)
	}

	partialTemplates, err := filepath.Glob(filepath.Join(templatesDir, "partials/*.html"))
	if err != nil {
		return nil, fmt.Errorf("failed to glob partial templates: %v", err)
	}

	allTemplates := append(defaultTemplates, partialTemplates...)

	ts, err := template.ParseFiles(allTemplates...)
	if err != nil {
		return nil, fmt.Errorf("failed to parse global templates: %v", err)
	}

	return &GoTemplateEngine{
		templates:    ts,
		templatesDir: templatesDir,
	}, nil
}

func (gte *GoTemplateEngine) Render(templates []string, baseTemplate string, data interface{}) ([]byte, error) {
	clonedTemplates, err := gte.templates.Clone()
	if err != nil {
		return nil, fmt.Errorf("failed to clone templates: %v", err)
	}

	var fullTemplatesPaths []string

	for _, tmpl := range templates {
		fullPath := filepath.Join(gte.templatesDir, tmpl)

		if _, err = os.Stat(fullPath); err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				return nil, fmt.Errorf("could not check if file exists")
			}
			continue
		}

		fullTemplatesPaths = append(fullTemplatesPaths, fullPath)
	}

	if fullTemplatesPaths != nil {
		clonedTemplates, err = clonedTemplates.ParseFiles(fullTemplatesPaths...)
		if err != nil {
			return nil, fmt.Errorf("failed to parse additional templates: %v", err)
		}
	}

	var buf bytes.Buffer
	if err = clonedTemplates.ExecuteTemplate(&buf, baseTemplate, data); err != nil {
		return nil, fmt.Errorf("failed to execute template: %v", err)
	}

	return buf.Bytes(), nil
}
