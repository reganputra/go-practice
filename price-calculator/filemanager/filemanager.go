package filemanager

import (
	"bufio"
	"errors"
	"os"
)

func ReadLine(path string) ([]string, error) {
	open, err := os.Open(path)
	if err != nil {
		return nil, errors.New("error opening file")
	}
	scanner := bufio.NewScanner(open)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		defer open.Close()
		return nil, errors.New("error reading file")
	}

	defer open.Close()
	return lines, nil
}
