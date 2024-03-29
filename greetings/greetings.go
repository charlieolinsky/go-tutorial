package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	//Returns a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomFormat(), name)
	//message := fmt.Sprint(randomFormat()) //BREAK ON PURPOSE
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := map[string]string{}

	for _, name := range names {
		msg, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = msg
	}
	return messages, nil
}

func randomFormat() string {

	//A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	//return a number [0, len(formats))
	return formats[rand.Intn(len(formats))]
}
