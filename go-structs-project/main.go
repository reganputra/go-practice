package main

import (
	"fmt"
	"go-structs-project/note"
)

func main() {
	fmt.Println("Welcome to Notes App!")
	title, content := getNoteData()

	userNote, err := note.NewNote(title, content)
	if err != nil {
		fmt.Println("Error creating note:", err)
		return
	}
	userNote.Display()
}

func getUserInput(prompt string) string {
	var input string
	fmt.Println(prompt)
	fmt.Scan(&input)
	return input
}

func getNoteData() (string, string) {
	titleInput := getUserInput("Enter title: ")
	contentInput := getUserInput("Enter content: ")
	return titleInput, contentInput
}
