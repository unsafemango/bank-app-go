package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(fileName string, value float64) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644) // 0644 is a way of encoding file permissions (read and write for owner only, other users can only read)
}

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	// code to run when an error occurs
	if err != nil {
		return 1000, errors.New("Failed to find file.")
	}
	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	// code to run when an error occurs
	if err != nil {
		return 1000, errors.New("Failed to parse stored value!")
	}

	// no error then return nil
	return value, nil
}
