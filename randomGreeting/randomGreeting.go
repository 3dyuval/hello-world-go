package randomGreeting

import (
	"fmt"
	"math/rand"
)

func RandomGreeting(name string) (string, error) {

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Nice to meet you, %v",
		"Yo %v",
	}
	return formats[rand.Intn(len(formats))]
}
