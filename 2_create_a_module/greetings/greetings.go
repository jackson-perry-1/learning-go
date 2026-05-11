package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	// return an error if the name string is empty
	if name == "" {
		return "", errors.New("name was empty")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
