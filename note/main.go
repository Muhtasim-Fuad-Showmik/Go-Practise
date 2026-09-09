package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

type outputtable interface {
	saver // Embedded interface
	Display()
}

func main() {
	title, content := getNoteData()
	userNote, err := note.New(title, content)
	if err != nil {
		panic(err)
	}

	err = outputData(userNote)
	if err != nil {
		panic("💣 Saving the data failed!" + err.Error())
	}

	todoText := getUserInput("Todo text:")
	todo, err := todo.New(todoText)
	if err != nil {
		panic(err)
	}

	err = outputData(todo)
	if err != nil {
		panic("💣 Saving the data failed!" + err.Error())
	}
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)
	value, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	value = strings.TrimSuffix(value, "\n")
	value = strings.TrimSuffix(value, "\r")

	return value
}

func saveData(s saver) error {
	err := s.Save()
	if err != nil {
		return err
	}

	fmt.Println("💾 Saving the data succeeded!")
	return nil
}

func outputData(o outputtable) error {
	o.Display()
	return saveData(o)
}

func printTypesOfValues(value any) {
	switch value.(type) {
	case string:
		fmt.Println("String: ", value)
	case int, float64:
		fmt.Println("Number: ", value)
	default:
		fmt.Println("Unknown type")
	}
}
