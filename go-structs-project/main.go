package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Welcome to Notes App!")
	title, content, err := getNoteData()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func getUserInput(prompt string) (string, error) {
	var input string
	fmt.Println(prompt)
	fmt.Scan(&input)
	if input == "" { // If the input is empty, return an error
		return "", errors.New("input cannot be empty")
	}
	return input, nil
}

func getNoteData() (string, string, error) {
	titleInput, err := getUserInput("Enter title: ")
	if err != nil {
		return "", "", err
	}
	fmt.Println("Content")
	contentInput, err := getUserInput("Enter content: ")
	if err != nil {
		return "", "", err
	}

	return titleInput, contentInput, nil
}
