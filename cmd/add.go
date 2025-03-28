package cmd

import (
	"github.com/skiba-mateusz/rocket/commandeer"
	"github.com/skiba-mateusz/rocket/config"
	"github.com/skiba-mateusz/rocket/content"
	"github.com/skiba-mateusz/rocket/logger"
)

func NewAddCommand() *commandeer.Command {
	cmd := &commandeer.Command{
		Name:        "add",
		Description: "Add new content page",
		Handler: func(command *commandeer.Command, args []string) error {
			log := logger.NewDefaultLogger(logger.INFO)
			_, err := config.LoadConfig()
			if err != nil {
				log.Error("%v", err)
				return nil
			}

			if len(args) == 0 {
				log.Warn("You must provide path, e.g. 'blogs/my-first-blog.md', 'about.md'")
				return nil
			}

			path := args[0]

			log.Info("Adding new content page: '%s'", path)

			if err = content.NewPage("content", path); err != nil {
				log.Error("%v", err)
				return nil
			}

			log.Success("Page '%s' added successfully", path)

			return nil
		},
	}

	return cmd
}
