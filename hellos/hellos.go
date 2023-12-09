package hellos

import (
	"myprojectgo/randomGreeting"
)

/*
	This is a package without a go.mod.
	To use it within this project, make sure
	this project is inside you go workspace folder
	which is declared $GOPATH
*/

func Hellos(names []string) (map[string]string, error) {

	messages := make(map[string]string)
	for _, name := range names {
		message, err := randomGreeting.RandomGreeting(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}
