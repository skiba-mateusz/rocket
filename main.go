package main

import (
	"github.com/skiba-mateusz/rocket/commands"
	"log"
	"os"
)

func main() {
	commandeer := commands.NewCommandeer()

	commandeer.RegisterCommand(commands.Command{
		Name:        "init",
		Description: "initialize rocket project",
		Handler:     commands.InitCommand,
	})
	commandeer.RegisterCommand(commands.Command{
		Name:        "add",
		Description: "add a new content page",
		Handler:     commands.AddCommand,
	})
	commandeer.RegisterCommand(commands.Command{
		Name:        "ping",
		Description: "See if that works",
		Handler:     commands.PingCommand,
	})
	commandeer.RegisterHelpCommand()

	if len(os.Args) < 2 {
		commandeer.ShowUsage()
		return
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	if err := commandeer.ExecuteCommand(cmd, args); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}
