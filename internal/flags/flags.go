package flags

import (
	"errors"
	"flag"
	"githug.com/Antonious-Stewart/expense-tracker-cli/internal/budget"
	"log"
	"os"

	"githug.com/Antonious-Stewart/expense-tracker-cli/internal/types"
)

type FlagOperations struct {
	fs map[types.Command]*flag.FlagSet
}

func NewFlagOperations() *FlagOperations {
	f := &FlagOperations{
		fs: make(map[types.Command]*flag.FlagSet),
	}

	flags := []struct {
		name string
		typ  types.Command
	}{
		{"budget", types.Budget},
		{"add", types.Add},
	}

	for _, flagInfo := range flags {
		f.fs[flagInfo.typ] = flag.NewFlagSet(flagInfo.name, flag.ExitOnError)
	}

	return f
}

func Root(args []string) error {
	if len(args) < 2 {
		return errors.New("you must pass a sub-command")
	}

	flags := NewFlagOperations()
	switch arg := os.Args[1]; arg {
	case types.Budget.String():
		setCmd := flags.fs[types.Budget].Float64(
			"set",
			0,
			"sets the budget for the current month")

		flags.fs[types.Budget].Parse(os.Args[2:])
		budget.Set(*setCmd)
	case types.Add.String():
		log.Println("Add")
	}
	return nil
}
