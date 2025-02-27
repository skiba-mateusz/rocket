package main

import (
	"github.com/skiba-mateusz/rocket/cmd"
	"github.com/skiba-mateusz/rocket/commandeer"
	"log"
	"os"
)

func main() {
	cmdr := commandeer.NewCommandeer()

	cmdr.RegisterCommand(cmd.InitCommand())
	cmdr.RegisterCommand(cmd.AddCommand())
	cmdr.RegisterCommand(cmd.ServeCommand())
	cmdr.RegisterCommand(cmd.PingCommand())
	cmdr.RegisterHelpCommand()

	if err := cmdr.ExecuteCommand(os.Args); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}
