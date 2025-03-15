package cmd

import (
	"github.com/skiba-mateusz/rocket/commandeer"
	"log"
)

func NewPingCommand() *commandeer.Command {
	return &commandeer.Command{
		Name:        "ping",
		Description: "Just to see if that works",
		Handler: func(args []string) error {
			log.Println("pong")
			return nil
		},
	}
}
