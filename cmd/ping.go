package cmd

import (
	"fmt"
	"github.com/skiba-mateusz/rocket/commandeer"
)

func PingCommand() *commandeer.Command {
	cmd := &commandeer.Command{
		Name:        "ping",
		Description: "pong",
		Handler: func(cmd *commandeer.Command, args []string) error {
			fmt.Println("Pong")
			return nil
		},
	}

	return cmd
}
