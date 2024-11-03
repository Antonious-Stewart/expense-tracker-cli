package types

import (
	"github.com/google/uuid"
	"time"
)

type Command int

const (
	Budget Command = iota
	Add
)

func (c Command) String() string {
	return [...]string{"budget", "add"}[c]
}

type Expense struct {
	Date   time.Time
	Amount float64
	ID     uuid.UUID
}

type Statement struct {
	ID       uuid.UUID
	Period   string
	Expenses []Expense
	Budget   float64
	Balance  float64
}
