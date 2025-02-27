package commandeer

import (
	"fmt"
	"os"
)

type Commandeer struct {
	commands map[string]*Command
}

type Command struct {
	Name        string
	Description string
	Handler     HandlerFunc
	flags       *CommandFlagSet
}

type HandlerFunc func(cmd *Command, args []string) error

func NewCommandeer() *Commandeer {
	return &Commandeer{
		commands: map[string]*Command{},
	}
}

func (c *Commandeer) RegisterCommand(command *Command) {
	if _, exists := c.commands[command.Name]; exists {
		fmt.Printf("command %s already exists\n", command.Name)
		return
	}
	c.commands[command.Name] = command
}

func (c *Commandeer) RegisterHelpCommand() {
	c.RegisterCommand(&Command{
		Name:        "help",
		Description: "manual",
		Handler: func(cmd *Command, args []string) error {
			c.showUsage()
			return nil
		},
	})
}

func (c *Commandeer) ExecuteCommand(args []string) error {
	if len(os.Args) < 2 {
		c.showUsage()
		return nil
	}

	name := os.Args[1]

	command, exists := c.commands[name]
	if !exists {
		return fmt.Errorf("command %s not found", name)
	}

	if command.flags != nil {
		if err := command.flags.Parse(args[2:]); err != nil {
			return err
		}
	}

	if err := command.Handler(command, args[2:]); err != nil {
		return err
	}
	return nil
}

func (c *Commandeer) showUsage() {
	fmt.Println("Usage: rocket <command> [arguments]")
	fmt.Println("Available commands:")
	for command := range c.commands {
		fmt.Printf("%s - %s\n", command, c.commands[command].Description)
	}
}
