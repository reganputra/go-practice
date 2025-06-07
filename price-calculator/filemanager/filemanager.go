package filemanager

import (
	"bufio"
	"encoding/json"
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

func WriteJson(path string, data interface{}) error {
	file, err := os.Create(path)
	if err != nil {
		return errors.New("error creating file")
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		defer file.Close()
		return errors.New("error encoding data to JSON")
	}

	return nil
}
