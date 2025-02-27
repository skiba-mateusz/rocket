package commandeer

import (
	"flag"
	"fmt"
	"strconv"
)

type CommandFlagSet struct {
	*flag.FlagSet
}

type intVal struct {
	val *int
}

type stringVal struct {
	val *string
}

func (c *Command) Flags() *CommandFlagSet {
	if c.flags == nil {
		c.flags = &CommandFlagSet{flag.NewFlagSet(c.Name, flag.ContinueOnError)}
	}
	return c.flags
}

func (c *CommandFlagSet) Integer(name, usage string, defaultVal int) {
	c.FlagSet.Var(intVal{val: &defaultVal}, name, usage)
}

func (c *CommandFlagSet) GetInteger(name string) (int, error) {
	f := c.FlagSet.Lookup(name)
	if f == nil {
		return 0, fmt.Errorf("flag %s not fund", name)
	}
	return strconv.Atoi(f.Value.String())
}

func (c *CommandFlagSet) String(name, usage, defaultVal string) {
	c.FlagSet.Var(stringVal{val: &defaultVal}, name, usage)
}

func (c *CommandFlagSet) GetString(name string) (string, error) {
	f := c.FlagSet.Lookup(name)
	if f == nil {
		return "", fmt.Errorf("flag %s not found", name)
	}
	return f.Value.String(), nil
}

func (v intVal) String() string {
	return strconv.Itoa(*v.val)
}

func (v intVal) Set(s string) error {
	if i, err := strconv.Atoi(s); err != nil {
		return err
	} else {
		*v.val = i
	}
	return nil
}

func (v stringVal) String() string {
	return *v.val
}

func (v stringVal) Set(s string) error {
	*v.val = s
	return nil
}
