package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (note *Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	data, err := json.Marshal(note)
	if err != nil {
		return fmt.Errorf("error marshalling note: %w", err)
	}
	return os.WriteFile(fileName, data, 0644)
}

func NewNote(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return nil, errors.New("title and content cannot be empty")
	}

	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (note *Note) Display() {
	fmt.Printf("Title: %v has following content:\n\n%v\n\n", note.Title, note.Content)
}
