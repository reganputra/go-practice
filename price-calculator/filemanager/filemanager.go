package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type FileManager struct {
	// Store
	InputFilePath  string
	OutputFilePath string
}

func (fm *FileManager) ReadLine() ([]string, error) {
	open, err := os.Open(fm.InputFilePath)
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

func (fm *FileManager) WriteResult(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return errors.New("error creating file")
	}
	defer file.Close()

	time.Sleep(2 * time.Second)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		defer file.Close()
		return errors.New("error encoding data to JSON")
	}

	return nil
}

func New(inputFilePath, outputFilePath string) *FileManager {
	return &FileManager{
		InputFilePath:  inputFilePath,
		OutputFilePath: outputFilePath,
	}
}
