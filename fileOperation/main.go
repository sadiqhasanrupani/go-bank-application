package fileOperation

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 5000.0, errors.New("failed to find the file")
	}

	balanceText := string(data)
	value, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 5000.0, errors.New("failed to parse stored file value")
	}

	return value, nil
}

func WriteFloatToFile(fileName string, balance float64) {
	amount := fmt.Sprint(balance)

	//^ first param => text's name
	//^ second param => array of byte
	//^ third param => file permissions
	os.WriteFile(fileName, []byte(amount), 0644)
}
