package main

import (
	"fmt"
	"log"

	"github.com/jackson-perry-1/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Jackson", "Bob", "Sally"}
	messages, err := greetings.Hellos(names)

	// if there was an error, print it and exit
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
