package main

import (
	"fmt"
	"log"

	"github.com/jackson-perry-1/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// message, err := greetings.Hello("Jackson")
	message, err := greetings.Hello("")

	// if there was an error, print it and exit
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
