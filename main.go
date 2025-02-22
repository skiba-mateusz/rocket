package main

import (
	"github.com/skiba-mateusz/rocket/commands"
	"log"
	"os"
)

func main() {
	commandeer := commands.NewCommandeer()
	commandeer.RegisterCommand(commands.Command{
		Name:        "ping",
		Description: "See if that works",
		Handler:     commands.Ping,
	})
	commandeer.RegisterHelpCommand()

	if len(os.Args) < 2 {
		commandeer.ShowUsage()
		return
	}

	cmd := os.Args[1]

	if err := commandeer.ExecuteCommand(cmd); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}
