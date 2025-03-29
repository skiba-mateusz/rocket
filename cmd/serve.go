package cmd

import (
	"github.com/skiba-mateusz/rocket/commandeer"
	"github.com/skiba-mateusz/rocket/logger"
	"github.com/skiba-mateusz/rocket/server"
)

func NewServeCommand() *commandeer.Command {
	cmd := &commandeer.Command{
		Name:        "serve",
		Description: "Serve build files",
		Handler: func(command *commandeer.Command, args []string) error {
			log := logger.NewDefaultLogger(logger.INFO)

			port, err := command.Flags().GetInteger("port")
			if err != nil {
				log.Error(err.Error())
				return nil
			}

			srv, err := server.NewServer(log, port, "public")
			if err != nil {
				log.Error(err.Error())
				return nil
			}

			if err = srv.Run(); err != nil {
				log.Error(err.Error())
			}

			return nil
		},
	}

	cmd.Flags().Integer("port", "Specify port for server", 8080)

	return cmd
}
