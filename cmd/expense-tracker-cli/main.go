package main

import (
	"log"
	"os"

	"githug.com/Antonious-Stewart/expense-tracker-cli/internal/flags"
)

func main() {
	if err := flags.Root(os.Args[1:]); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
