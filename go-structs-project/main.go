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
	err = saveData(newTodo)
	if err != nil {
		return
	}

	userNote.Display()
	err = saveData(userNote)
	if err != nil {
		return
	}

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

func saveData(data saver) error {
	if data == nil {
		fmt.Println("No data to save.")
		return nil
	}
	err := data.Save()
	if err != nil {
		fmt.Println("Error saving data:", err)
		return err
	}
	fmt.Println("Data saved successfully!")
	return nil
}
