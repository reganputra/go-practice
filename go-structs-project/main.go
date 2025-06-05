package main

import (
	"bufio"
	"fmt"
	"go-structs-project/note"
	"go-structs-project/todo"
	"os"
	"strings"
)

type saver interface {
	Save() error
}

func main() {
	fmt.Println("Welcome to Notes App!")
	title, content := getNoteData()

	todoText := getUserInput("Enter todo item (or press Enter to skip): ")
	newTodo, err := todo.NewTodo(todoText)
	if err != nil {
		return
	}

	userNote, err := note.NewNote(title, content)
	if err != nil {
		fmt.Println("Error creating note:", err)
		return
	}

	newTodo.Display()
	err = newTodo.Save()
	if err != nil {
		fmt.Println("Error saving todo:", err)
		return
	}
	fmt.Println("Todo saved successfully!")

	userNote.Display()
	err = userNote.Save()
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
