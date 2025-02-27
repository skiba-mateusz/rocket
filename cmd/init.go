package cmd

import (
	"fmt"
	"github.com/skiba-mateusz/rocket/commandeer"
	"github.com/skiba-mateusz/rocket/skeletons"
	"os"
)

func InitCommand() *commandeer.Command {
	cmd := &commandeer.Command{
		Name:        "init",
		Description: "initialize rocket project",
		Handler: func(cmd *commandeer.Command, args []string) error {
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
		},
	}

	return cmd
}
