package budget

import (
	"githug.com/Antonious-Stewart/expense-tracker-cli/internal/types"
	"log"
	"os"
)

func Set(amount float64) {
	if _, err := os.Stat("../../statements"); err != nil {
		if err := os.MkdirAll("../../statements", 0666); err != nil {
			log.Panic(err)
		}
	}

	types.NewStatement(amount)
}
