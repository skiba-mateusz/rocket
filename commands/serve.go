package commands

import (
	"github.com/skiba-mateusz/rocket/server"
)

func ServeCommand(args []string) error {
	srv := server.NewServer(8080)
	return srv.Run()
}
