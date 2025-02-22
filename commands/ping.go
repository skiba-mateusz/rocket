package commands

import "fmt"

func Ping() error {
	fmt.Println("Pong")
	return nil
}
