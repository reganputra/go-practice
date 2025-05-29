package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func ReadBalanceFile(filename string) (float64, error) {
	read, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Initializing balance to 0.0")
		return 0.0, errors.New("failed reading balance file")
	}
	valueText := string(read)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		fmt.Println("Initializing balance to 0.0")
		return 0.0, errors.New("error converting balance to float")
	}
	return value, nil
}

func WriteBalanceToFiles(balance float64, filename string) {
	valueText := fmt.Sprint(balance)
	err := os.WriteFile(filename, []byte(valueText), 0644)
	if err != nil {
		return
	}
}
