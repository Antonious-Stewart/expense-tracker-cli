package flags

import (
	"errors"
	"flag"
	"fmt"
	"githug.com/Antonious-Stewart/expense-tracker-cli/internal/expense"
	"githug.com/Antonious-Stewart/expense-tracker-cli/internal/statement"
	"githug.com/Antonious-Stewart/expense-tracker-cli/internal/types"
	"log"
	"os"
	"slices"
	"time"
)

type FlagOperations struct {
	FlagSets         map[types.Command]*flag.FlagSet
	StatementHandler statement.Handler
	ExpenseHandler   expense.Handler
}

func containsAll(args []string, flags ...string) bool {
	for _, value := range flags {
		if !slices.Contains(args, value) {
			return false
		}
	}

	return true
}

func NewFlagOperations(statementHandler statement.Handler, expenseHandler expense.Handler) *FlagOperations {
	flagOps := &FlagOperations{
		FlagSets:         make(map[types.Command]*flag.FlagSet),
		StatementHandler: statementHandler,
		ExpenseHandler:   expenseHandler,
	}

	flags := []struct {
		name string
		typ  types.Command
	}{
		{"budget", types.Budget},
		{"add", types.Add},
	}

	for _, flagInfo := range flags {
		flagOps.FlagSets[flagInfo.typ] = flag.NewFlagSet(flagInfo.name, flag.ExitOnError)
	}

	return flagOps
}

func (f *FlagOperations) HandleBudgetCommand(args []string) error {
	setCmd := f.FlagSets[types.Budget].Float64(
		"set",
		0,
		"sets the budget for the current month")

	if err := f.FlagSets[types.Budget].Parse(os.Args[2:]); err != nil {
		return err
	}

	if len(args) == 0 || args[2] != "--set" {
		return errors.New("expected '--set' flag for the budget command")
	}

	return f.StatementHandler.SetBudget(*setCmd)
}

func (f *FlagOperations) HandleAddCommand(args []string) error {

	flagSet := f.FlagSets[types.Add]

	descCmd := flagSet.String(
		"description",
		"",
		"this is the item that will be getting added",
	)

	amtCmd := flagSet.Float64(
		"amount",
		0,
		"used to set the amount of the expense",
	)

	if err := f.FlagSets[types.Add].Parse(os.Args[2:]); err != nil {
		return err
	}

	if !containsAll(args, "--description", "--amount") {
		return errors.New("expected add command to have a description and a amount sub-command")
	}

	if err := flagSet.Parse(args); err != nil {
		return err
	}

	newExpense := types.NewExpense(*amtCmd, *descCmd)
	return f.ExpenseHandler.AddExpense(newExpense)
}

func (f *FlagOperations) runner(args []string) error {
	var err error

	switch arg := args[1]; arg {
	case types.Budget.String():
		err = f.HandleBudgetCommand(args)
	case types.Add.String():
		err = f.HandleAddCommand(args)
	default:
		err = fmt.Errorf("unknown sub-command: %s", args[1])
	}

	return err
}

func Root(args []string) {
	if len(args) < 2 {
		log.Fatal("missing sub-command: expected 'budget' or 'add'")
	}

	if _, err := os.Stat("../../statements"); err != nil {
		if err := os.MkdirAll("../../statements", 0750); err != nil {
			log.Fatal(err)
		}
		log.Printf("%s Directory Created", time.DateTime)
	}

	statementService := statement.NewService()
	expenseService := expense.NewService()
	flags := NewFlagOperations(statementService, expenseService)

	if err := flags.runner(args); err != nil {
		log.Fatal(err)
	}
}
