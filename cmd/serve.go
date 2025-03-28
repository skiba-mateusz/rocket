package cmd

import (
	"github.com/skiba-mateusz/rocket/commandeer"
	"github.com/skiba-mateusz/rocket/logger"
	"github.com/skiba-mateusz/rocket/server"
)

func NewServeCommand() *commandeer.Command {
	cmd := &commandeer.Command{
		Name:        "serve",
		Description: "Serve files from public directory",
		Handler: func(command *commandeer.Command, args []string) error {
			log := logger.NewDefaultLogger(logger.INFO)

			port, err := command.Flags().GetInteger("port")
			if err != nil {
				log.Error("%v", err)
			}

			srv := server.NewServer(log, port)
			if err = srv.Run(); err != nil {
				log.Error("%v", err)
			}

			return nil
		},
	}

	cmd.Flags().Integer("port", "specify port for server", 8080)

	return cmd
}
