package main

import (
	"fmt"
	"log"
	"myprojectgo/goErrors"
	"myprojectgo/hello"
	"myprojectgo/hellos"
	"myprojectgo/randomGreeting"
	"myprojectgo/world"
)

func main() {

	log.SetPrefix("main.go: ")
	log.SetFlags(0) // disable printing the date and time

	hellos.MyUtilityFunction()

	message, err1 := goErrors.GoErrors("Yuval")
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Println(message)

	randomMessage, err2 := randomGreeting.RandomGreeting("Yuval")

	if (err2) != nil {
		log.Fatal(err2)
	}
	fmt.Println(randomMessage)

	hello.Hello()
	world.World()
}
