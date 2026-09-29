package filemanager

import (
	"bufio"
	"errors"
	"os"
)

func ReadLines(path string) ([]string, error) {
	// Load file
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New("Failed to open file")
	}

	// Prepare scanner for reading the file
	scanner := bufio.NewScanner(file)
	var lines []string

	// Store content of file into lines
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("Failed to read file content")
	}

	file.Close()

	return lines, nil
}
