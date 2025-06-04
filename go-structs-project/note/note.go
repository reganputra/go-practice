package note

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

func (note *Note) SaveNoteToFle() {
	fileName := strings.ReplaceAll(note.title, " ", "_")
	fileName = strings.ToLower(fileName)
	os.WriteFile(fileName, []byte(note.content), 0644)
}

func NewNote(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return nil, errors.New("title and content cannot be empty")
	}

	return &Note{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil
}

func (note *Note) Display() {
	fmt.Printf("Title: %v has following content:\n\n%v\n\n", note.title, note.content)
}
