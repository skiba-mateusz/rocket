package cmd

import (
	"github.com/skiba-mateusz/rocket/builder"
	"github.com/skiba-mateusz/rocket/commandeer"
	"github.com/skiba-mateusz/rocket/config"
	"github.com/skiba-mateusz/rocket/content"
	"github.com/skiba-mateusz/rocket/logger"
	"github.com/skiba-mateusz/rocket/templates"
)

func NewBuildCommand() *commandeer.Command {
	cmd := &commandeer.Command{
		Name:        "build",
		Description: "Build site",
		Handler: func(command *commandeer.Command, args []string) error {
			log := logger.NewDefaultLogger(logger.INFO)

			cfg, err := config.LoadConfig()
			if err != nil {
				return err
			}
			engine, err := templates.NewGoTemplateEngine("templates")
			if err != nil {
				return err
			}

			parser := content.NewMarkdownParser()
			build := builder.NewBuilder(cfg, log, parser, engine, "public", "content")

			return build.Build()
		},
	}

	return cmd
}
