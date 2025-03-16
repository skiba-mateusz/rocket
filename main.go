package main

import (
	"fmt"
	"github.com/skiba-mateusz/rocket/cmd"
	"github.com/skiba-mateusz/rocket/commandeer"
	"os"
)

func main() {
	cmdr := commandeer.NewCommandeer()
	cmd.RegisterCommands(cmdr)

	if err := cmdr.ExecuteCommand(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
