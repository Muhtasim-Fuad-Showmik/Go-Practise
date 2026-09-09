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

func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("Todo text must be provided.")
	}

	return Todo{
		Text: text,
	}, nil
}

func (todo Todo) Display() {
	fmt.Printf("Todo: %s", todo.Text)
}

func (todo Todo) Save() error {
	// Prepare a clean file name
	fileName := "todo.json"

	// Encode data to JSON format
	jsonData, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	// Write to the file and return error (if any) returned
	return os.WriteFile(fileName, jsonData, 0644)
}
