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

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("Both title and content must be provided.")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (note Note) Display() {
	fmt.Printf("Your note titled '%v' has the following content:\n\n'%v'", note.Title, note.Content)
}

func (note Note) Save() error {
	// Prepare a clean file name
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	// Encode data to JSON format
	jsonData, err := json.Marshal(note)
	if err != nil {
		return err
	}

	// Write to the file and return error (if any) returned
	return os.WriteFile(fileName, jsonData, 0644)
}
