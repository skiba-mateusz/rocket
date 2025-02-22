package commands

import "fmt"

func PingCommand(args []string) error {
	fmt.Println("Pong")
	return nil
}
