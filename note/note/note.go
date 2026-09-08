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
	title     string
	content   string
	createdAt time.Time
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("Both title and content must be provided.")
	}

	return Note{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil
}

func (note Note) Display() {
	fmt.Printf("Your note titled '%v' has the following content:\n\n'%v'", note.title, note.content)
}

func (note Note) Save() error {
	// Prepare a clean file name
	fileName := strings.ReplaceAll(note.title, " ", "_")
	fileName = strings.ToLower(fileName)

	// Encode data to JSON format
	jsonData, err := json.Marshal(note)
	if err != nil {
		return err
	}

	// Write to the file and return error (if any) returned
	return os.WriteFile(fileName, jsonData, 0644)
}
