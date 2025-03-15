package commandeer

import (
	"fmt"
)

type CommandHandler func(args []string) error

type Command struct {
	Name        string
	Description string
	Handler     CommandHandler
}

type Commandeer struct {
	commands map[string]*Command
}

func NewCommandeer() *Commandeer {
	return &Commandeer{
		commands: map[string]*Command{},
	}
}

func (c *Commandeer) RegisterCommand(command *Command) {
	if _, exists := c.commands[command.Name]; exists {
		fmt.Printf("Command '%s' already exists\n", command.Name)
		return
	}
	c.commands[command.Name] = command
}

func (c *Commandeer) RegisterHelpCommand() {
	c.RegisterCommand(&Command{
		Name:        "help",
		Description: "Show usage information and available commands",
		Handler: func(args []string) error {
			c.showUsage()
			return nil
		},
	})
}

func (c *Commandeer) ExecuteCommand(args []string) error {
	if len(args) < 2 {
		c.showUsage()
		return nil
	}

	name := args[1]
	command, exists := c.commands[name]
	if !exists {
		return fmt.Errorf("command '%s' not found", name)
	}

	if err := command.Handler(args[2:]); err != nil {
		return err
	}

	return nil
}

func (c *Commandeer) showUsage() {
	fmt.Println("🚀 Usage: rocket <command> [arguments]")
	fmt.Println("\nAvailable commands:")
	for _, command := range c.commands {
		fmt.Printf("  %-15s %s\n", command.Name, command.Description)
	}
}
