package commandeer

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func captureOutput(f func() error) (string, error) {
	orig := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	err := f()
	os.Stdout = orig
	w.Close()
	out, _ := io.ReadAll(r)
	return string(out), err
}

func TestCommandeer(t *testing.T) {
	t.Run("should execute command", func(t *testing.T) {
		cmdr := NewCommandeer()
		executed := false

		testCmd := &Command{
			Name:        "test",
			Description: "test command",
			Handler: func(cmd *Command, args []string) error {
				executed = true
				if len(args) != 2 || args[0] != "arg1" || args[1] != "arg2" {
					return fmt.Errorf("unexpected arguments: %v", args)
				}
				return nil
			},
		}

		cmdr.RegisterCommand(testCmd)

		os.Args = []string{"rocket", "test", "arg1", "arg2"}
		if err := cmdr.ExecuteCommand(os.Args); err != nil {
			t.Fatal(err)
		}

		if !executed {
			t.Fatalf("the command was expected to be executed")
		}
	})

	t.Run("should print help message", func(t *testing.T) {
		cmdr := NewCommandeer()

		test1Cmd := &Command{
			Name:        "test1",
			Description: "test1 command",
			Handler: func(cmd *Command, args []string) error {
				return nil
			},
		}
		test2Cmd := &Command{
			Name:        "test2",
			Description: "test2 command",
			Handler: func(cmd *Command, args []string) error {
				return nil
			},
		}

		cmdr.RegisterCommand(test1Cmd)
		cmdr.RegisterCommand(test2Cmd)
		cmdr.RegisterHelpCommand()

		expectedOutput := `Usage: rocket <command> [arguments]
Available commands:
test1 - test1 command
test2 - test2 command
help - manual`

		os.Args = []string{"rocket", "help"}
		output, err := captureOutput(func() error {
			return cmdr.ExecuteCommand(os.Args)
		})
		if err != nil {
			t.Fatal(err)
		}

		if strings.TrimSpace(expectedOutput) != strings.TrimSpace(output) {
			t.Fatalf("expected output: %s, got: %s", expectedOutput, output)
		}
	})

	t.Run("should handle flag arguments", func(t *testing.T) {
		cmdr := NewCommandeer()

		expectedFlag := "test"

		testCmd := &Command{
			Name:        "test",
			Description: "test command",
			Handler: func(cmd *Command, args []string) error {
				f, err := cmd.Flags().GetString("testFlag")
				if err != nil {
					return err
				}
				if f != "test" {
					return fmt.Errorf("expected flag: %s, got: %s", expectedFlag, f)
				}
				return nil
			},
		}

		testCmd.Flags().String("testFlag", "flag for testing", "")

		cmdr.RegisterCommand(testCmd)

		os.Args = []string{"rocket", "test", "--testFlag=test"}
		if err := cmdr.ExecuteCommand(os.Args); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("should handle unknown command", func(t *testing.T) {
		cmdr := NewCommandeer()
		cmdr.RegisterHelpCommand()

		os.Args = []string{"rocket", "unknown"}
		err := cmdr.ExecuteCommand(os.Args)
		if err.Error() != "command unknown not found" {
			t.Fatal(err)
		}
	})

	t.Run("should handle existing command", func(t *testing.T) {
		cmdr := NewCommandeer()
		cmdr.RegisterHelpCommand()

		testCmd := &Command{
			Name:        "help",
			Description: "help command",
			Handler: func(cmd *Command, args []string) error {
				return nil
			},
		}

		cmdr.RegisterCommand(testCmd)

		output, _ := captureOutput(func() error {
			cmdr.RegisterCommand(testCmd)
			return nil
		})

		expectedOutput := "command help already exists\n"
		if output != expectedOutput {
			t.Fatalf("expected: %s, got: %s", expectedOutput, output)
		}
	})
}
