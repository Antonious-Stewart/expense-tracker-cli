package types

import (
	"github.com/google/uuid"
	"time"
)

type Expense struct {
	Date   string
	Amount float64
	ID     uuid.UUID
}

func NewExpense(amount float64) Expense {
	currentTime := time.Now()
	expenseDate := currentTime.Format("01/02/2006")

	return Expense{
		Date:   expenseDate,
		ID:     uuid.New(),
		Amount: amount,
	}
}