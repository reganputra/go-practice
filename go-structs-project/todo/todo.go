package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (todo *Todo) Save() error {

	fileName := "todo.json"

	data, err := json.Marshal(todo)
	if err != nil {
		return fmt.Errorf("error marshalling note: %w", err)
	}
	return os.WriteFile(fileName, data, 0644)
}

func NewTodo(text string) (*Todo, error) {
	if text == "" {
		return nil, errors.New("title and content cannot be empty")
	}

	return &Todo{
		Text: text,
	}, nil
}

func (todo *Todo) Display() {
	fmt.Println(todo.Text)
}
