package cmd

import (
	"fmt"
	"github.com/skiba-mateusz/rocket/commandeer"
)

func NewPingCommand() *commandeer.Command {
	cmd := &commandeer.Command{
		Name:        "ping",
		Description: "Just to see if that works",
		Handler: func(cmd *commandeer.Command, args []string) error {
			times, err := cmd.Flags().GetInteger("times")
			if err != nil {
				return err
			}

			name, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}

			for i := 0; i < times; i++ {
				fmt.Printf("%d. %s\n", i, name)
			}

			return nil
		},
	}

	cmd.Flags().Integer("times", "how many times print", 2)
	cmd.Flags().String("name", "name to print", "Mike")

	return cmd
}
