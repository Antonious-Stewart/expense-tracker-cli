package flags

import (
	"errors"
	"flag"
	"log"
	"os"

	"githug.com/Antonious-Stewart/expense-tracker-cli/internal/types"
)

type FlagOperations struct {
	fs map[types.Command]*flag.FlagSet
}

func (f *FlagOperations) New() {
	f.fs[types.Budget] = flag.NewFlagSet("budget", flag.ExitOnError)
	f.fs[types.Add] = flag.NewFlagSet("add", flag.ExitOnError)
}

func Root(args []string) error {
	if len(args) < 2 {
		return errors.New("you must pass a sub-command")
	}

	switch arg := os.Args[1]; arg {
	case types.Budget.String():
		log.Println("Budget")
	case types.Add.String():
		log.Println("Add")
	}
	return nil
}
