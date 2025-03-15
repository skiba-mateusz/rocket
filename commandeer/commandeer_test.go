package commandeer

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func TestRegisterCommand(t *testing.T) {
	cmdr := NewCommandeer()

	testCommand := &Command{
		Name:        "test",
		Description: "Just a test command",
		Handler: func(command *Command, args []string) error {
			return nil
		},
	}

	cmdr.RegisterCommand(testCommand)

	assert.Contains(t, cmdr.commands, "test", "The command should be registered")
	assert.Equal(t, testCommand, cmdr.commands["test"], "Registered command should match the input")
}

func TestRegisterDuplicateCommand(t *testing.T) {
	cmdr := NewCommandeer()

	testCommand1 := &Command{
		Name:        "test",
		Description: "Just a test command",
		Handler: func(command *Command, args []string) error {
			return nil
		},
	}
	testCommand2 := &Command{
		Name:        "test",
		Description: "Just a test command",
		Handler: func(command *Command, args []string) error {
			return nil
		},
	}

	cmdr.RegisterCommand(testCommand1)
	cmdr.RegisterCommand(testCommand2)

	assert.Equal(t, testCommand1, cmdr.commands["test"], "The first command should be registered")
}

func TestExecutedValidCommand(t *testing.T) {
	cmdr := NewCommandeer()

	var executed bool

	testCommand := &Command{
		Name:        "test",
		Description: "Just a test command",
		Handler: func(command *Command, args []string) error {
			executed = true
			return nil
		},
	}

	cmdr.RegisterCommand(testCommand)

	err := cmdr.ExecuteCommand([]string{"rocket", "test"})
	assert.NoErrorf(t, err, "Executing a valid command should not return an error")
	assert.True(t, executed, "The command handler should be executed")
}

func TestExecuteInvalidCommand(t *testing.T) {
	cmdr := NewCommandeer()

	err := cmdr.ExecuteCommand([]string{"rocket", "test"})

	assert.Error(t, err, "Executing invalid command should return an error")
	assert.EqualError(t, err, "command 'test' not found", "The error message should indicate the missing command")
}

func TestCommandWithFlags(t *testing.T) {
	cmdr := NewCommandeer()

	var flagValue string

	testCommand := &Command{
		Name:        "greet",
		Description: "Just a command to test flags",
		Handler: func(command *Command, args []string) error {
			name, err := command.Flags().GetString("name")
			if err != nil {
				return err
			}
			if name == "" {
				return fmt.Errorf("name flag is required")
			}

			flagValue = name

			return nil
		},
	}

	testCommand.Flags().String("name", "name flag", "")
	cmdr.RegisterCommand(testCommand)

	err := cmdr.ExecuteCommand([]string{"rocket", "greet", "-name=John"})

	assert.NoError(t, err, "Executing a valid command with flags shoud not return an error")
	assert.Equal(t, "John", flagValue, "The flag value should be parsed correctly")
}

func TestShowUsage(t *testing.T) {
	cmdr := NewCommandeer()

	cmdr.RegisterCommand(&Command{
		Name:        "test",
		Description: "A test command",
	})
	cmdr.RegisterHelpCommand()

	output, err := captureStdout(cmdr.showUsage)
	if err != nil {
		t.Fatal(err)
	}

	assert.Contains(t, output, "🚀 Usage: rocket <command> [arguments] [flags]", "Usage information and should be printed")
	assert.Contains(t, output, "test", "Registered command should be listed in usage information")
}

func captureStdout(fn func()) (string, error) {
	original := os.Stdout

	var buf bytes.Buffer
	r, w, err := os.Pipe()
	if err != nil {
		return "", err
	}
	os.Stdout = w

	fn()

	w.Close()
	os.Stdout = original
	io.Copy(&buf, r)

	return buf.String(), nil
}
