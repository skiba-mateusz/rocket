package cmd

import (
	"github.com/skiba-mateusz/rocket/commandeer"
	"github.com/skiba-mateusz/rocket/server"
)

func ServeCommand() *commandeer.Command {
	cmd := &commandeer.Command{
		Name:        "serve",
		Description: "run development server",
		Handler: func(cmd *commandeer.Command, args []string) error {
			port, err := cmd.Flags().GetInteger("port")
			if err != nil {
				return err
			}

			srv := server.NewServer(port)
			return srv.Run()
		},
	}

	cmd.Flags().Integer("port", "set port", 8080)

	return cmd
}
