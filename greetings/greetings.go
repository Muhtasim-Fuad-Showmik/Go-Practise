package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name is given, return an error with a message.
	if name == "" {
		return "", errors.New("Empty Name")
	}

	// Return a greeting that embeds that name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
