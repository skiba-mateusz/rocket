package cmd

import (
	"github.com/skiba-mateusz/rocket/commandeer"
	"github.com/skiba-mateusz/rocket/initializer"
	"github.com/skiba-mateusz/rocket/logger"
)

func NewInitCommand() *commandeer.Command {
	cmd := &commandeer.Command{
		Name:        "init",
		Description: "Initialize new site",
		Handler: func(command *commandeer.Command, args []string) error {
			log := logger.NewDefaultLogger(logger.INFO)
			init := initializer.NewInitializer(log)

			if len(args) == 0 {
				log.Warn("You must provide a title for the new site, e.g. my-site")
				return nil
			}

			if err := init.NewSite(args[0]); err != nil {
				return err
			}

			return nil
		},
	}

	return cmd
}
