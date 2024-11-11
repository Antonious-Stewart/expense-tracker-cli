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

func NewStatement(budget float64) error {
	currentTime := time.Now()
	statementDate := currentTime.Format("01/06")
	statement := Statement{
		uuid.New(),
		statementDate,
		[]Expense{},
		budget,
	}

	dirPath, err := filepath.Abs("../../statements")

	if err != nil {
		return err
	}

	path := filepath.Join(dirPath, fmt.Sprintf("%s.json", currentTime.Format("01-2006")))

	fileData, err := json.Marshal(statement)

	if err != nil {
		return err
	}

	_, err = os.Stat(path)

	if err == nil {
		return errors.New("budget has already been set for the month")
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0750)
	if err != nil {
		return err
	}

	defer file.Close()

	if _, err := file.Write(fileData); err != nil {
		return err
	}

	log.Printf("%s File Created", time.DateOnly)

	return nil
}
