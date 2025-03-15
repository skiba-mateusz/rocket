package commandeer

import (
	"flag"
	"fmt"
	"strconv"
)

type Flags struct {
	*flag.FlagSet
}

type intVal struct {
	val *int
}

type strVal struct {
	val *string
}

func (c *Command) Flags() *Flags {
	if c.flags == nil {
		c.flags = &Flags{flag.NewFlagSet(c.Name, flag.ContinueOnError)}
	}
	return c.flags
}

func (v intVal) String() string {
	if v.val == nil {
		return "0"
	}
	return strconv.Itoa(*v.val)
}

func (v intVal) Set(str string) error {
	if i, err := strconv.Atoi(str); err != nil {
		return err
	} else {
		*v.val = i
	}
	return nil
}

func (f *Flags) Integer(name, usage string, defaultVal int) {
	f.FlagSet.Var(intVal{val: &defaultVal}, name, usage)
}

func (f *Flags) GetInteger(name string) (int, error) {
	lookup := f.FlagSet.Lookup(name)
	if lookup == nil {
		return 0, fmt.Errorf("flag '%s' not found", name)
	}
	return strconv.Atoi(lookup.Value.String())
}

func (v strVal) String() string {
	if v.val == nil {
		return ""
	}
	return *v.val
}

func (v strVal) Set(str string) error {
	*v.val = str
	return nil
}

func (f *Flags) String(name, usage, defaultVal string) {
	f.FlagSet.Var(strVal{val: &defaultVal}, name, usage)
}

func (f *Flags) GetString(name string) (string, error) {
	lookup := f.FlagSet.Lookup(name)
	if lookup == nil {
		return "", fmt.Errorf("flag '%s' not found", name)
	}
	return lookup.Value.String(), nil
}
