package main

import (
	"os"

	"githug.com/Antonious-Stewart/expense-tracker-cli/internal/flags"
)

func main() {
	flags.Root(os.Args)
}
