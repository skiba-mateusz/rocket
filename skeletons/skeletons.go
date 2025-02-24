package skeletons

import (
	"embed"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

//go:embed all:project/*
var projectFs embed.FS

func CreateProject(rootDir string) error {
	return fs.WalkDir(projectFs, "project", func(path string, d fs.DirEntry, err error) error {
		relPath, _ := filepath.Rel("project", path)
		destPath := filepath.Join(rootDir, relPath)

		if d.IsDir() {
			if err := os.MkdirAll(destPath, 0755); err != nil {
				return err
			}
		} else {
			if err := copyEmbeddedFiles(path, destPath); err != nil {
				return err
			}
		}

		return nil
	})
}

func copyEmbeddedFiles(src, dst string) error {
	srcFile, err := projectFs.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destPath, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destPath.Close()

	if _, err = io.Copy(destPath, srcFile); err != nil {
		return err
	}

	return nil
}
