package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Greet(name string) (string, error) {
	if name == "" {
		return "", errors.New("invalid name")
	}
	message := fmt.Sprintf(RandomGreetings(), name)
	return message, nil
}

func Greetings(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Greet(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func RandomGreetings() string {
	formatMessage := []string{
		"Hi %v, welcome to go world",
		"Namaste %v, welcome to go world",
		"Ola %v, welcome to go world",
	}
	return formatMessage[rand.Intn(len(formatMessage))]
}
