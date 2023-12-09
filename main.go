package main

import (
	"fmt"
	"github.com/3dyuval/go/hello"
	"github.com/3dyuval/go/world"
	"github.com/3dyuval/goErrors"
	"log"
)

func main() {

	log.SetPrefix("main.go: ")
	log.SetFlags(0) // disable printing the date and time

	message, err := goErrors.GoErrors("Yuval")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	hello.Hello()
	world.World()
}
