package flags

import (
	"errors"
	"flag"

	"githug.com/Antonious-Stewart/expense-tracker-cli/internal/types"
)

type FlagOperations struct {
	fs map[types.Command]*flag.FlagSet
}

func (f *FlagOperations) Set() {
	f.fs[types.Budget] = flag.NewFlagSet("budget", flag.ExitOnError)
	f.fs[types.Add] = flag.NewFlagSet("add", flag.ExitOnError)
}

func Runner() {
	addCmd := FlagOperations.fs[types.Add]

}

func Root(args []string) error {
	if len(args) < 2 {
		return errors.New("you must pass a sub-command")
	}

	return nil
}
