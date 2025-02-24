package commands

import (
	"fmt"
	"github.com/skiba-mateusz/rocket/skeletons"
	"os"
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

	if err := skeletons.CreateProject(rootDir); err != nil {
		return fmt.Errorf("failed to create project: %v", err)
	}

	fmt.Printf("Project %s initialized successfully\n", rootDir)

	return nil
}
