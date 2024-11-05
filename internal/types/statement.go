package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log"
	"os"
	"path/filepath"
	"time"
)

type Statement struct {
	ID       uuid.UUID
	Period   string
	Expenses []Expense
	Budget   float64
}

func NewStatement(budget float64) {
	currentTime := time.Now()
	statementDate := currentTime.Format("01/06")
	statement := Statement{
		uuid.New(),
		statementDate,
		[]Expense{},
		budget,
	}

	fPath, err := filepath.Abs("../../statements")

	if err != nil {
		log.Panic(err)
	}

	path := fmt.Sprintf("%s/%s.json", fPath, currentTime.Format("2006-01-02"))

	fileData, err := json.Marshal(statement)

	if err != nil {
		log.Panic(err)
	}

	_, err = os.Stat(path)

	if err == nil {
		log.Panic(errors.New("budget has already been set for the month"))
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	if _, err := file.Write(fileData); err != nil {
		log.Fatal(err)
	}
}
