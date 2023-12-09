package main

import (
	"fmt"
	"log"
	"myprojectgo/hello"
	"myprojectgo/hellos"
	"myprojectgo/world"
)

func main() {

	log.SetPrefix("main.go: ")
	log.SetFlags(0) // disable printing the date and time

	messages, err := hellos.Hellos([]string{
		"Yuval", "Shimon", "Tikva",
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, greet := range messages {
		fmt.Println(greet)
	}

	hello.Hello()
	world.World()
}
