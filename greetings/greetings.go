package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	// Return a greeting that embeds the name in a message.
	// message := fmt.Sprintf("Hi, %v. Welcome!", name)
	// return message, nil
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}
func randomFormat() string {
	formats := []string{
		"Hi %v!",
		"Hello %v!",
		"Great to see you %v!",
	}
	return formats[rand.Intn(len(formats))]
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}
