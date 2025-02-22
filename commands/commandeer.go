package commands

import "fmt"

type CommanderFunc func(args []string) error

type Command struct {
	Name        string
	Description string
	Handler     CommanderFunc
}

type Commandeer struct {
	commands map[string]Command
}

func NewCommandeer() *Commandeer {
	return &Commandeer{
		commands: map[string]Command{},
	}
}

func (c *Commandeer) RegisterCommand(command Command) {
	if _, exists := c.commands[command.Name]; exists {
		fmt.Printf("command %s already exists", command.Name)
	}
	c.commands[command.Name] = command
}

func (c *Commandeer) ExecuteCommand(name string, args []string) error {
	command, exists := c.commands[name]
	if !exists {
		return fmt.Errorf("command %s not found", name)
	}
	if err := command.Handler(args); err != nil {
		return err
	}
	return nil
}

func (c *Commandeer) ShowUsage() {
	fmt.Println("Usage: rocket <command> [arguments]")
	fmt.Println("Available commands:")
	for command := range c.commands {
		fmt.Printf("%s - %s\n", command, c.commands[command].Description)
	}
}

func (c *Commandeer) RegisterHelpCommand() {
	c.RegisterCommand(Command{
		Name:        "help",
		Description: "manual",
		Handler: func(args []string) error {
			c.ShowUsage()
			return nil
		},
	})
}
