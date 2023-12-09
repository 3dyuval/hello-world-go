package goErrors

import (
	"errors"
	"fmt"
)

func GoErrors(name string) (string, error) {

	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf("Hi %v\n", name)
	return message, nil
}
