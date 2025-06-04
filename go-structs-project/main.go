package main

import (
	"bufio"
	"fmt"
	"go-structs-project/note"
	"os"
	"strings"
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
	err = userNote.SaveNoteToFle()
	if err != nil {
		fmt.Println("Error saving note:", err)
		return
	}
	fmt.Println("Note saved successfully!")
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSpace(text)
	return text
}

func getNoteData() (string, string) {
	titleInput := getUserInput("Enter title: ")
	contentInput := getUserInput("Enter content: ")
	return titleInput, contentInput
}
