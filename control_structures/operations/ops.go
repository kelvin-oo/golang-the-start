package operations

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetValueFromFile(file string) (float64, error) {
	data, err := os.ReadFile(file)

	if err != nil {
		return 1000, errors.New("failed to find balance file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored balance value")
	}

	return balance, nil
}

func WriteBalanceToFile(value float64, file string) {
	balanceText := fmt.Sprint(value)
	os.WriteFile(file, []byte(balanceText), 0644)
}
