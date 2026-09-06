package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(value float64, fileName string) {
	// Write value to file
	valueText := fmt.Sprintf("%.2f", value)
	os.WriteFile(fileName, []byte(valueText), 0644)
	fmt.Printf("Value of %.2f written to file: %s\n", value, fileName)
}

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 0, errors.New("Failed to find file.")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 0, errors.New("Failed to parse stored value.")
	}

	return value, nil
}
