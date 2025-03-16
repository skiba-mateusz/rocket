package cmd

import "github.com/skiba-mateusz/rocket/commandeer"

func RegisterCommands(cmdr *commandeer.Commandeer) {
	cmdr.RegisterCommand(NewInitCommand())
	cmdr.RegisterCommand(NewAddCommand())
	cmdr.RegisterHelpCommand()
}
