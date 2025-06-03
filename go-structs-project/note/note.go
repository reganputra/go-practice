package note

import (
	"errors"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
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
